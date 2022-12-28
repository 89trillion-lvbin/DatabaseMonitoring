package gmongoclient

import (
	setting2 "dm/dm/setting"
	"dm/pkg/gmongo"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	MgoCore *mongo.Database
)

func Setup() {
	MgoCore = gmongo.NewClient(setting2.MongoCoreSetting.ApplyURI, setting2.MongoBaseSetting)
}
