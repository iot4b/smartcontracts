package proxy

import (
	"net"
	"net/url"

	"github.com/coalalib/coalago"
	log "github.com/ndmsystems/golog"
	"github.com/patrickmn/go-cache"
)

type proxiedMessage struct {
	msg   *coalago.CoAPMessage
	saddr *net.UDPAddr
}

func senderBackWorker(conn *net.UDPConn, msgs <-chan proxiedMessage) {
	for {
		m := <-msgs
		buf, err := coalago.Serialize(m.msg)
		if err != nil {
			continue
		}
		conn.WriteToUDP(buf, m.saddr)
	}

}

type forwardMessage struct {
	msg      *coalago.CoAPMessage
	proxyURI string
}

func senderForwardWorker(conn *net.UDPConn, forwardList *cache.Cache, msgs <-chan forwardMessage) {
	for {
		msg := <-msgs
		parsedURL, err := url.Parse(msg.proxyURI)
		if err != nil {
			log.Debugw("senderForwardWorker err", "error", err, "msg", msg.msg.ToReadableString())
			continue
		}

		addr, err := net.ResolveUDPAddr("udp", parsedURL.Host)
		if err != nil {
			log.Debugw("senderForwardWorker err", "error", err, "msg", msg.msg.ToReadableString())
			continue
		}

		deleteProxyOptions(msg.msg)
		optClientHello := msg.msg.GetOption(coalago.CoapHandshakeTypeClientHello)
		if optClientHello != nil {
			log.Debugw("Debug proxy sessions client hello", "len", msg.msg.Payload.Length(), "sender", msg.msg.Sender.String())
		} else {
			optPeerHello := msg.msg.GetOption(coalago.CoapHandshakeTypePeerHello)
			if optPeerHello != nil {
				log.Debugw("Debug proxy sessions peer hello", "len", msg.msg.Payload.Length(), "sender", msg.msg.Sender.String())
			}
		}
		buf, err := coalago.Serialize(msg.msg)
		if err != nil {
			log.Debugw("senderForwardWorker err", "error", err, "msg", msg.msg.ToReadableString())
			continue
		}

		forwardList.SetDefault(msg.msg.GetTokenString()+addr.String(), msg.msg.Sender)
		Rate.Inc()
		conn.WriteTo(buf, addr)
	}
}

func serverWorker(server *coalago.Server, msgs <-chan *coalago.CoAPMessage) {
	for {
		msg := <-msgs
		server.ServeMessage(msg)
	}
}
