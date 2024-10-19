package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2"
)

// Connect using mgo
func connectWithMgoDial() {
	// Method 1: Using mgo.Dial
	session, err := mgo.Dial("mongodb://root:example@mongo:27017/admin")
	if err != nil {
		log.Fatal("mgo.Dial error:", err)
	}
	defer session.Close()
	
	log.Println("Connected with mgo.Dial")
}

// Connect using mgo with DialInfo
func connectWithMgoDialInfo() {
	// Method 2: Using mgo.DialInfo
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{"mongo:27017"},
		Database: "admin",
		Username: "root",
		Password: "example",
		Timeout:  60 * time.Second,
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatal("mgo.DialWithInfo error:", err)
	}
	defer session.Close()
	log.Println("Connected with mgo.DialInfo")
}

// Connect using mgo with ParseURL
func connectWithMgoParseURL() {
	// Method 3: Using mgo.ParseURL
	uri := "mongodb://root:example@mongo:27017/admin"
	info, err := mgo.ParseURL(uri)
	if err != nil {
		log.Fatal("mgo.ParseURL error:", err)
	}
	info.Timeout = 60 * time.Second

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		log.Fatal("mgo.DialWithInfo error:", err)
	}
	defer session.Close()
	log.Println("Connected with mgo.ParseURL")
}

// Connect using official MongoDB driver
func connectWithMongoDriver() {
	// Method 4: Using mongo-go-driver
	uri := "mongodb://root:example@mongo:27017/admin"
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("mongo.Connect error:", err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("mongo.Ping error:", err)
	}
	defer client.Disconnect(ctx)

	log.Println("Connected with mongo-go-driver")
}

func main() {
	// Uncomment one method at a time to test
	connectWithMongoDriver()
	connectWithMgoDial()
	// connectWithMgoDialInfo()
	// connectWithMgoParseURL()

}
