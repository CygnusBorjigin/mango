package mango

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type AtlasCollection struct {
	collectionClient  *mongo.Collection
	collectionContext context.Context
}

func NewAtlasCollection(targetClient *mongo.Client, targetDatabase string, targetCollection string) AtlasCollection {
	databaseConnection := targetClient.Database(targetDatabase)
	collectionConnection := databaseConnection.Collection(targetCollection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return AtlasCollection{
		collectionClient:  collectionConnection,
		collectionContext: ctx,
	}
}

func (a *AtlasCollection) ListAllObject() {
	// Get the result
	res, err := a.collectionClient.Find(a.collectionContext, bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	// parse the result
	var parsedRes []bson.M
	if err = res.All(a.collectionContext, &parsedRes); err != nil {
		fmt.Println(err)
	}
	fmt.Println(parsedRes)
}

func (a *AtlasCollection) queryCollection(target []map[string]interface{}) ([][]bson.M, *string) {
	var resData [][]bson.M
	parsedQuery := ParseQuery(target)
	for _, eachQuery := range parsedQuery {
		queryOptions := options.Find().SetProjection(eachQuery[1])
		res, err := a.collectionClient.Find(a.collectionContext, eachQuery[0], queryOptions)
		if err != nil {
			errorMessage := "query failed"
			return resData, &errorMessage
		}

		var parsedReturnValue []bson.M
		if err = res.All(a.collectionContext, &parsedReturnValue); err != nil {
			errorMessage := "parsed return value failed"
			return resData, &errorMessage
		}
		resData = append(resData, parsedReturnValue)
	}
	return resData, nil
}

func (a *AtlasCollection) parseDBReturn(target [][]bson.M) [][]map[string]interface{} {
	var res [][]map[string]interface{}
	for _, eachResult := range target {
		var parsedRes []map[string]interface{}
		for _, eachObject := range eachResult {
			parsedObject := make(map[string]interface{})
			parsedMap, err := bson.MarshalExtJSON(eachObject, true, true)
			if err != nil {
				fmt.Println("parse error")
			}
			err = json.Unmarshal(parsedMap, &parsedObject)
			parsedRes = append(parsedRes, parsedObject)
		}
		res = append(res, parsedRes)
	}
	return res
}

func (a *AtlasCollection) GetObject(target []map[string]interface{}) [][]map[string]interface{} {
	rawRes, err := a.queryCollection(target)
	if err != nil {
		fmt.Println(err)
	}
	parsedRes := a.parseDBReturn(rawRes)
	return parsedRes
}
