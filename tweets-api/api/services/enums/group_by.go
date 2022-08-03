package enums

import (
	"fmt"
	"strings"
)

type GroupByEnum string

var (
	Year  GroupByEnum = "year"
	Month GroupByEnum = "month"
	Hour  GroupByEnum = "hour"
)

var groupByMap = map[string]GroupByEnum{
	"year":  Year,
	"month": Month,
	"hour":  Hour,
}

// ConvertToGroupByArray converts array of strings to GroupByEnum
func ConvertToGroupByArray(groupByList []string) ([]GroupByEnum, error) {
	result := make([]GroupByEnum, 0)

	for _, value := range groupByList {
		groupByEnum, ok := groupByMap[strings.ToLower(value)]

		if ok {
			result = append(result, groupByEnum)
		} else {
			return nil, fmt.Errorf("%s is not valid value", value)
		}
	}

	return result, nil
}

// ConvertToStringsGroupBy converts array of GroupByEnum to strings
func ConvertToStringsGroupBy(groupByList []GroupByEnum) ([]string, error) {
	result := make([]string, 0)

	for _, value := range groupByList {
		_, ok := groupByMap[string(value)]

		if ok {
			result = append(result, string(value))
		} else {
			return nil, fmt.Errorf("%s is not valid value", value)
		}
	}

	return result, nil
}
