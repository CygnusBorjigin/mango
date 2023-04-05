package mango

import (
	"fmt"
	mangoDataType "github.com/cygnusborjigin/mango/dataTypes"
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

func handelNumericCondition(target string) bson.M {
	comparisonType, numericalValue := movingOn.NumericInequality(target)
	switch comparisonType {
	case movingOn.EqualTo:
		return bson.M{"$eq": numericalValue}
	case movingOn.GreaterThan:
		return bson.M{"$gt": numericalValue}
	case movingOn.GreaterThanOrEqualTo:
		return bson.M{"$gte": numericalValue}
	case movingOn.LesserThan:
		return bson.M{"$lt": numericalValue}
	case movingOn.LesserThanOrEqualTo:
		return bson.M{"$lte": numericalValue}
	case movingOn.NotNumericComparison:
		return nil
	}
	return nil
}

func checkConditionType(target interface{}) mangoDataType.ConditionType {
	parsedCondition, parseSuccess := target.(string)
	if !parseSuccess {
		return mangoDataType.Undefined
	}
	comparison, _ := movingOn.NumericInequality(parsedCondition)
	if comparison == movingOn.NotNumericComparison {
		return mangoDataType.StringComparison
	}
	return mangoDataType.NumericComparison
}

func ParseQuery(target []map[string]interface{}) [][]bson.M {
	var res [][]bson.M
	for _, eachQuery := range target {
		dbQuery := bson.M{}
		returnField := bson.M{}

		searchCondition := movingOn.StringInterfaceSubMapByExcluding(eachQuery, []string{"returnField"})
		attributeRequired := movingOn.InterfaceSliceToStringSlice(eachQuery["returnField"].([]interface{}))

		for eachKey, eachCondition := range searchCondition {
			switch checkConditionType(eachCondition) {
			case mangoDataType.StringComparison:
				dbQuery[eachKey] = eachCondition
			case mangoDataType.NumericComparison:
				parsedNumericalQuery := handelNumericCondition(eachCondition.(string))
				if parsedNumericalQuery != nil {
					dbQuery[eachKey] = handelNumericCondition(eachCondition.(string))
				} else {
					fmt.Println("Numerical parse error")
				}
			}
		}

		for _, attribute := range attributeRequired {
			returnField[attribute] = 1
		}
		res = append(res, []bson.M{dbQuery, returnField})
	}
	return res
}
