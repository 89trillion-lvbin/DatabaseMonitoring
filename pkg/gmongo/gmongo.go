package gmongo

import (
	"context"
	"time"

	"dm/pkg/commonmodel"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(applyUri string, baseSetting *commonmodel.MongoBase) (cli *mongo.Database) {
	/*op:=options.ClientOptions{uri: setting.MongoSetting.ApplyURI,
		MaxPoolSize: &(setting.MongoSetting.MaxPoolSize),
	}*/

	baseOpt := options.ClientOptions{
		ConnectTimeout:  &baseSetting.ConnTimeout,
		SocketTimeout:   &baseSetting.SocketTimeout,
		MaxPoolSize:     &baseSetting.MaxPoolSize,
		MaxConnIdleTime: &baseSetting.MaxConnIdleTime,
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(applyUri), &baseOpt)
	if err != nil {
		return nil
	}
	// ping5次，都不通认为失败
	for i := 0; i < 5; i++ {
		err := client.Ping(context.TODO(), nil)
		if err == nil {
			return client.Database("iw")
		}
		time.Sleep(100 * time.Microsecond)
	}
	return nil
}
