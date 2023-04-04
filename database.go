package mango

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type AtlasDatabase struct {
	clientConnection  *mongo.Client
	connectionContext context.Context
	databaseName      string
}

func NewAtlasDatabase(targetClient *mongo.Client, targetDatabase string) AtlasDatabase {
	defaultContext, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return AtlasDatabase{
		clientConnection:  targetClient,
		connectionContext: defaultContext,
		databaseName:      targetDatabase,
	}
}

func (a *AtlasDatabase) ListAllCollections() []string {
	targetDatabase := a.clientConnection.Database(a.databaseName)
	res, err := targetDatabase.ListCollectionNames(a.connectionContext, bson.M{})

	if err != nil {
		fmt.Println(err)
	}

	return res
}
