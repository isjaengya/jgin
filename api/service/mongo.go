package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"jgin/api/config"
	"log"
	"time"
)

var mongodb *mongo.Client
var mongodbDefaultDatabase string

func MongoInit() {
	var err error
	//uri := "mongodb://@localhost:27017"
	//mongoUser := config.Conf.GetString("mongo.user")
	//mongoPassword := config.Conf.GetString("mongo.password")
	mongoHost := config.Conf.GetString("mongo.host")
	mongoPort := config.Conf.GetInt("mongo.port")
	mongoMaxOpenConns := uint16(config.Conf.GetUint("mongo.maxOpenConns"))
	mongodbDefaultDatabase = config.Conf.GetString("mongo.defaultDatabase")

	uri := fmt.Sprintf("mongodb://@%s:%d", mongoHost, mongoPort)
	mongodb, err = mongo.Connect(getContext(), options.Client().ApplyURI(uri).SetMaxPoolSize(mongoMaxOpenConns)) // 连接池
	if err != nil {
		log.Fatal("open mongodb error: ", err.Error())
	}

	err = mongodb.Ping(nil, nil)
	if err != nil {
		log.Fatal("mongodb 初始化失败, ", err.Error())
	}

	FindDate()
}

func getContext() (ctx context.Context) {
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	return
}

func GetMongoClient() *mongo.Client {
	return mongodb
}

func GetMongoDatabase(s string) (db *mongo.Database) {
	db = mongodb.Database(s)
	return
}

func GetMongoDefaultDatabase() (db *mongo.Database) {
	db = mongodb.Database(mongodbDefaultDatabase)
	return
}

// 这个应该是用的最多的，上面的可以忽略
func GetMongoCollection(collectionName string, databaseName ...string) (c *mongo.Collection) {
	var s string
	if databaseName != nil {
		s = databaseName[0]
	} else {
		s = mongodbDefaultDatabase
	}
	db := GetMongoDatabase(s)
	c = db.Collection(collectionName)
	return
}

type test1 struct {
	TestId primitive.ObjectID `bson:"_id"`
	Name   string             `bson:"name"`
	Age    int                `bson:"age"`
}

func FindDate() {
	mongoCli := GetMongoCollection("test")
	cursor, err := mongoCli.Find(getContext(), bson.M{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if err = cursor.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
	defer cursor.Close(nil)
	for cursor.Next(nil) {
		var t test1
		if err = cursor.Decode(&t); err != nil {
			fmt.Println(t, "tttttttttttttttttttttttttttttt")
		}
		fmt.Println(t)
	}
}

func checkErr(err error) {
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("没有查到数据")
			//os.Exit(0)
		} else {
			fmt.Println(err, "yyyyyyyyyyyyyyyyyyyyyyyyyyy")
			//os.Exit(0)
		}

	}
}
