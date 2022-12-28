package eventCenter

import (
	"context"
	"time"

	"dm/pkg/commonmodel"

	"github.com/89trillion/CommonECProducerSDK/producer"
	"github.com/sirupsen/logrus"
)

var eventProducer = producer.Producer{}

type KeyString interface{}

func SetUp(eventKafkaSetting *commonmodel.EventKafka) {
	producerConfig := producer.Config{
		Ip:                     eventKafkaSetting.BootStrapServer,
		Ack:                    eventKafkaSetting.Ack,
		BatchSize:              eventKafkaSetting.BatchSize,
		AllowAutoTopicCreation: eventKafkaSetting.AllowAutoTopicCreation,
		OutTime:                eventKafkaSetting.OutTime * time.Second,
		FlushTime:              eventKafkaSetting.FlushTime * time.Millisecond,
	}
	eventProducer.NewProducer(producerConfig)
}

func PublishEvent(userId, topic, event string, paramsMap map[string]string) {
	ctx := context.WithValue(context.Background(), KeyString(producer.UserId), userId)
	ctx = context.WithValue(ctx, KeyString(producer.Version), int32(1))
	ctx = context.WithValue(ctx, KeyString(producer.DeviceId), "")
	ctx = context.WithValue(ctx, KeyString(producer.ClientIP), "")
	ctx = context.WithValue(ctx, KeyString(producer.Platform), int32(1))
	ctx = context.WithValue(ctx, KeyString(producer.GameLanguage), "")
	ctx = context.WithValue(ctx, KeyString(producer.SysVer), int32(1))
	err := eventProducer.Publish(ctx, topic, event, paramsMap)
	if err != nil {
		logrus.Error(err, userId, topic, paramsMap)
	}
}
