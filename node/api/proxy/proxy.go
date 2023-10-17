package proxy

import (
	"net"
	"runtime"
	"time"

	"github.com/coalalib/coalago"
	"github.com/patrickmn/go-cache"
)

type Proxy struct {
	conn        *net.UDPConn
	forwardList *cache.Cache
}

func New() *Proxy {
	return &Proxy{
		forwardList: cache.New(time.Minute, time.Second),
	}
}

func (p *Proxy) RunWithServer(addr string, server *coalago.Server) error {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}
	p.conn, err = net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}

	server.Serve(p.conn)

	proxiedMessages := make(chan proxiedMessage, 1024)
	forwardMessages := make(chan forwardMessage, 1024)
	serveMessages := make(chan *coalago.CoAPMessage, 1024)

	for i := 0; i < runtime.NumCPU()*4; i++ {
		go senderBackWorker(p.conn, proxiedMessages)
		go senderForwardWorker(p.conn, p.forwardList, forwardMessages)
		go serverWorker(server, serveMessages)
	}

	for {
		message, err := p.receiveMessage()
		if err != nil {
			return err
		}

		proxyURI := message.GetOptionProxyURIasString()
		if len(proxyURI) != 0 {
			forwardMessages <- forwardMessage{
				msg:      message,
				proxyURI: proxyURI,
			}
			continue
		}

		saddr, ok := p.forwardList.Get(message.GetTokenString() + message.Sender.String())
		if ok {
			proxiedMessages <- proxiedMessage{
				msg:   message,
				saddr: saddr.(*net.UDPAddr),
			}
			p.forwardList.SetDefault(message.GetTokenString()+message.Sender.String(), saddr)

			continue
		}

		serveMessages <- message
	}
}

func (p *Proxy) receiveMessage() (*coalago.CoAPMessage, error) {
	for {
		buff := make([]byte, 1500)
		n, sender, err := p.conn.ReadFromUDP(buff)
		if err != nil {
			return nil, err
		}

		message, err := coalago.Deserialize(buff[:n])
		if err != nil {
			continue
		}

		message.Sender = sender
		return message, nil
	}
}

func deleteProxyOptions(message *coalago.CoAPMessage) {
	message.RemoveOptions(coalago.OptionProxyScheme)
	message.RemoveOptions(coalago.OptionProxyURI)
}
