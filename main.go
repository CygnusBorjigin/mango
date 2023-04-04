package mango

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type AtlasConnection struct {
	connectionUrl    string
	connectContext   context.Context
	connectionClient *mongo.Client
}

func NewAtlasConnection(connectionString string) AtlasConnection {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	newConnection := establishConnection(connectionString, ctx)
	return AtlasConnection{
		connectionUrl:    connectionString,
		connectContext:   ctx,
		connectionClient: newConnection,
	}
}

func establishConnection(connectionUrl string, ctx context.Context) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionUrl))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return client
}

func (a *AtlasConnection) TerminateConnection() {
	err := a.connectionClient.Disconnect(a.connectContext)
	if err != nil {
		fmt.Println(err)
	}
}

func (a *AtlasConnection) ListDatabase() []string {
	databasesList, err := a.connectionClient.ListDatabaseNames(a.connectContext, bson.M{})
	if err != nil {
		fmt.Println("error when trying to list database")
		fmt.Println(err)
		return nil
	}
	return databasesList
}

func (a *AtlasConnection) GetClient() *mongo.Client {
	return a.connectionClient
}
