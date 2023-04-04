package mango

import (
	movingOn "github.com/cygnusborjigin/movingOn"
	"go.mongodb.org/mongo-driver/bson"
)

func MapStringInterfaceToBson(target map[string]interface{}) bson.M {
	res := bson.M{}
	for key, value := range target {
		res[key] = value
	}
	return res
}

func ParseQuery(target []map[string]interface{}) [][]bson.M {
	var res [][]bson.M
	for _, eachQuery := range target {
		dbQuery := bson.M{}
		returnField := bson.M{}

		searchCondition := movingOn.StringInterfaceSubMapByExcluding(eachQuery, []string{"returnField"})
		attributeRequired := movingOn.InterfaceSliceToStringSlice(eachQuery["returnField"].([]interface{}))

		for eachKey, eachCondition := range searchCondition {
			dbQuery[eachKey] = eachCondition
		}

		for _, attribute := range attributeRequired {
			returnField[attribute] = 1
		}
		res = append(res, []bson.M{dbQuery, returnField})
	}
	return res
}
