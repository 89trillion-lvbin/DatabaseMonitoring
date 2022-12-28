package service

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"strings"
	"sync"
	"time"

	"dm/dm/client/gmongoclient"
	"dm/dm/internal/config"
	"dm/pkg/eventCenter"
	"dm/pkg/myerr"
	"dm/pkg/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StreamObject struct {
	Id                *WatchId `bson:"_id"`
	OperationType     string
	FullDocument      map[string]interface{}
	Ns                NS
	UpdateDescription map[string]interface{}
	DocumentKey       map[string]interface{}
}

type CollectionStruct struct {
	CollectionName string
	ExitChan       chan struct{}
	Client         *mongo.Database
	Pipeline       []bson.D
	Opt            *options.ChangeStreamOptions
}

type NS struct {
	Database   string `bson:"db"`
	Collection string `bson:"coll"`
}

type WatchId struct {
	Data string `bson:"_data"`
}

var (
	resumeToken  bson.Raw
	watchChannel = make(chan string)
)

const (
	ArenaRank  = "arena_rank_room"
	ArenaStage = "arena_stage_room"
	Pe         = "pe_user_room"
	Pm         = "pm_user_room"
	Inf        = "inf_user_room"
	Stc        = "stc_user_room"
	SmallRank  = "pvp_sr_room"
)

var GameModeNowSeasonMap = map[string]func() (int, int, int){
	ArenaRank:  GetArenaSeasonInfo,
	ArenaStage: GetArenaSeasonInfo,
	Pe:         GetPeSeasonInfo,
	Pm:         GetPmSeasonInfo,
	Inf:        GetInfSeasonInfo,
	Stc:        GetSuperSeasonInfo,
	SmallRank:  GetSmallRankSeasonInfo,
}

var GameModeNextSeasonMap = map[string]func() (int, int, int){
	ArenaRank:  GetNextArenaSeasonInfo,
	ArenaStage: GetNextArenaSeasonInfo,
	Pe:         GetNextPeSeasonInfo,
	Pm:         GetNextPmSeasonInfo,
	Inf:        GetNextInfSeasonInfo,
	Stc:        GetNextSuperSeasonInfo,
	SmallRank:  GetNextSmallRankSeasonInfo,
}

func Sync() {
	client := gmongoclient.MgoCore
	var wg sync.WaitGroup
	wg.Add(1)
	go producer()
	go consumer(client)
	wg.Wait()
}

var m = make(map[string]int, 0)

func producer() {
	for {
		nTime := util.Ntime()
		for k, v := range GameModeNowSeasonMap {
			nowSid, startTime, endTime := v()
			if nowSid == 0 {
				continue
			}
			_, ok := m[fmt.Sprintf(fmt.Sprintf("%s_%d", k, nowSid))]
			if !ok && startTime <= nTime && endTime >= nTime {
				collectionName := fmt.Sprintf("%s_%d", k, nowSid)
				m[collectionName] = 1
				watchChannel <- collectionName
			}
		}
		for k, v := range GameModeNextSeasonMap {
			nextSid, startTime, endTime := v()
			if nextSid == 0 {
				continue
			}
			_, ok := m[fmt.Sprintf("%s_%d", k, nextSid)]
			if !ok && startTime <= nTime+86400*2 && endTime >= nTime+86400*2 {
				nextCollectionName := fmt.Sprintf("%s_%d", k, nextSid)
				m[nextCollectionName] = 1
				watchChannel <- nextCollectionName
			}
		}

		time.Sleep(time.Hour)
	}
}

func consumer(client *mongo.Database) {
	for {
		name := <-watchChannel
		consume(name, client)
	}
}

func (c *CollectionStruct) watch() {
	//获得watch监听
	fmt.Println(c)
	watch, err := c.Client.Collection(c.CollectionName).Watch(context.TODO(), c.Pipeline, c.Opt)
	if err != nil {
		log.Fatal("watch监听失败：", err)
	}
	for {
		select {
		case <-c.ExitChan:
			runtime.Goexit()
		default:
			watchListen(watch)
		}
	}
}

func (c *CollectionStruct) check() {
	for {
		//校验赛季
		splitNames := strings.Split(c.CollectionName, "_")
		sid := splitNames[len(splitNames)-1]
		name := strings.Join(splitNames[:len(splitNames)-1], "_")
		f := GameModeNowSeasonMap[name]
		seasonId, _, _ := f()
		if util.Atoi(sid) < seasonId {
			c.quit()
			break
		}
		time.Sleep(time.Hour)
	}
}

func (c *CollectionStruct) quit() {
	fmt.Println("quit:", c.CollectionName)
	c.ExitChan <- struct{}{}
}

func CreateNewCollection(name string, client *mongo.Database) *CollectionStruct {
	pipeline := mongo.Pipeline{
		bson.D{{"$match",
			bson.M{"operationType": bson.M{"$in": bson.A{"insert", "update"}}},
		}},
	}
	//当前时间前10分钟
	now := time.Now()
	times, _ := time.ParseDuration("-10m")
	now = now.Add(times)
	timestamp := &primitive.Timestamp{
		T: uint32(now.Unix()),
		I: 0,
	}

	//设置监听option
	opt := options.ChangeStream().SetFullDocument(options.UpdateLookup).SetStartAtOperationTime(timestamp)
	if resumeToken != nil {
		opt.SetResumeAfter(resumeToken)
		opt.SetStartAtOperationTime(nil)
	}
	return &CollectionStruct{
		CollectionName: name,
		ExitChan:       make(chan struct{}),
		Client:         client,
		Pipeline:       pipeline,
		Opt:            opt,
	}
}

func consume(name string, client *mongo.Database) {
	s := CreateNewCollection(name, client)
	go s.watch()
	go s.check()
}

func watchListen(watch *mongo.ChangeStream) {
	for watch.TryNext(context.TODO()) {
		var stream StreamObject
		err := watch.Decode(&stream)
		if err != nil {
			log.Println("watch数据失败：", err)
		}
		log.Println("=============", stream.FullDocument["_id"])
		operate := stream.OperationType
		collections := stream.Ns.Collection
		roomId := stream.FullDocument["roomId"]
		userId := util.Itoa(stream.FullDocument["userId"])
		fmt.Println(roomId, userId, collections, operate)
		eventCenter.PublishEvent(userId, "changeData", "change.stream", map[string]string{
			"roomId":         util.Itoa(roomId),
			"collectionName": util.Itoa(collections),
			"operate":        util.Itoa(operate),
		})
	}
}

func GetSmallRankSeasonInfo() (int, int, int) {
	return config.GetSmallRankSid()
}

func GetSuperSeasonInfo() (int, int, int) {
	sid, startTime, endTime, myErr := config.SuperTroopSeasonSid()
	if myErr != myerr.SUCCESS {
		return 0, 0, 0
	}
	return sid, startTime, endTime
}

func GetInfSeasonInfo() (int, int, int) {
	sid, _ := config.GetInfWarSeasonId(0)
	startTime, endTime := config.GetInfWarSeasonTime(sid)
	return sid, startTime, endTime
}

func GetPeSeasonInfo() (int, int, int) {
	sid := config.GetPointRaceSid(true)
	startTime, _, endTime := config.GetPointRaceStartEndTime(sid, true)
	return sid, startTime, endTime
}

func GetPmSeasonInfo() (int, int, int) {
	sid := config.GetPointRaceSid(false)
	startTime, _, endTime := config.GetPointRaceStartEndTime(sid, false)
	return sid, startTime, endTime
}

func GetArenaSeasonInfo() (int, int, int) {
	nowSid := config.GetBattlePassCurSeasonId()
	startTime, endTime := config.GetBattlePassSeasonInfo(nowSid)
	return nowSid, startTime, endTime
}

func GetNextSmallRankSeasonInfo() (int, int, int) {
	sid, startTime, endTime := config.GetSmallRankSid()
	if util.Ntime()-startTime <= 86400*2 {
		return sid, startTime, endTime
	}
	return 0, 0, 0
}

func GetNextSuperSeasonInfo() (int, int, int) {
	sid, startTime, endTime, myErr := config.SuperTroopSeasonSid()
	if myErr != myerr.SUCCESS {
		return 0, 0, 0
	}
	if util.Ntime()-startTime <= 86400*2 {
		return sid, startTime, endTime
	}
	return 0, 0, 0
}

func GetNextInfSeasonInfo() (int, int, int) {
	sid, _ := config.GetInfWarSeasonId(0)
	startTime, endTime := config.GetInfWarSeasonTime(sid + 1)
	return sid + 1, startTime, endTime
}

func GetNextPeSeasonInfo() (int, int, int) {
	sid := config.GetPointRaceSid(true)
	startTime, _, endTime := config.GetPointRaceStartEndTime(sid+1, true)
	return sid + 1, startTime, endTime
}

func GetNextPmSeasonInfo() (int, int, int) {
	sid := config.GetPointRaceSid(false)
	startTime, _, endTime := config.GetPointRaceStartEndTime(sid+1, false)
	return sid + 1, startTime, endTime
}

func GetNextArenaSeasonInfo() (int, int, int) {
	nowSid := config.GetBattlePassCurSeasonId()
	startTime, endTime := config.GetBattlePassSeasonInfo(nowSid + 1)
	return nowSid + 1, startTime, endTime
}
