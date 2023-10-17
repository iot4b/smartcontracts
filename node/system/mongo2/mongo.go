package mongoDB

import (
	"context"
	"github.com/ever-iot/node/utils"

	log "github.com/ndmsystems/golog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Init(Hosts []string, Name, Username, Password string, useCerts bool) (*mongo.Database, context.Context) {
	log.Info(Hosts, Name)
	mongoctx := context.Background()

	opts := []*options.ClientOptions{
		options.Client().SetHosts(Hosts),
		options.Client().SetAppName(utils.Hostname()),                     //задаем имя приложения для логов монги
		options.Client().SetReadPreference(readpref.SecondaryPreferred()), //задаем режим чтения со слейвов
	}
	if useCerts {
		opts = append(opts, options.Client().SetAuth(options.Credential{Username: Username, Password: Password}))
	}

	client, err := mongo.Connect(mongoctx, opts...)
	if err != nil {
		log.Fatal(err)
	}
	return client.Database(Name), mongoctx
}
