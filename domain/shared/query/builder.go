package query

import (
	"eko-car/domain/shared/constant"
	"eko-car/domain/shared/model"
	"fmt"
	"strings"
)

func SearchQueryBuilder(conditions string) string {
	var (
		tempConditions []string
	)

	conds := strings.Split(conditions, "|")
	for _, cond := range conds {
		temps := strings.Split(cond, "=")
		for j, temp := range temps {
			key := strings.ReplaceAll(temps[0], " ", "")
			if j == 1 {
				temp := strings.ReplaceAll(temp, " ", "")
				arrs := strings.Split(temp, ",")
				if len(arrs) > 1 {
					tempConditions = append(tempConditions, key+" IN ('"+strings.Join(arrs, "','")+"')")
				} else {
					tempConditions = append(tempConditions, key+" = '"+temp+"'")
				}
			}
		}
	}

	return " AND " + strings.Join(tempConditions, " AND ")
}

func ConditionsBuilder(condition *model.Filter) (query string) {

	var (
		sortedFields          = SortByFieldName(condition.Filters)
		counterFields         = FindCounterField(sortedFields)
		conditionLen          = len(sortedFields)
		beforeField, newQuery string
		listsValue            []string
	)

	for i, field := range sortedFields {
		switch field.Option {
		case constant.EQUAL:
			if counterFields[field.FieldName] > 1 {
				newQuery, beforeField, listsValue = FieldINValueBuilder(field, beforeField, i, conditionLen, counterFields[field.FieldName], listsValue)
				query += newQuery
			} else {
				query += FieldValueBuilder(field, "=", i, conditionLen)
			}
		case constant.BETWEEN:
			query += FieldValueBuilder(field, "BETWEEN", i, conditionLen)
		case constant.NOT_BETWEEN:
			query += FieldValueBuilder(field, "NOT BETWEEN", i, conditionLen)
		case constant.LESS_THAN:
			query += FieldValueBuilder(field, "<", i, conditionLen)
		case constant.GREATER_THAN:
			query += FieldValueBuilder(field, ">", i, conditionLen)
		case constant.LESS_THAN_EQUAL:
			query += FieldValueBuilder(field, "<=", i, conditionLen)
		case constant.GREATER_THAN_EQUAL:
			query += FieldValueBuilder(field, ">=", i, conditionLen)
		case constant.NOT_EQUAL:
			query += FieldValueBuilder(field, "<>", i, conditionLen)
		case constant.CONTAINS_PATTERN:
			query += FieldValueBuilder(field, "LIKE", i, conditionLen)
		case constant.CONTAINS_NO_PATTERN:
			query += FieldValueBuilder(field, "NOT LIKE", i, conditionLen)
		}

		// if field.Option == constant.EQUAL {
		// 	if counterFields[field.FieldName] > 1 {
		// 		newQuery, beforeField, listsValue = FieldINValueBuilder(field, beforeField, i, conditionLen, counterFields[field.FieldName], listsValue)
		// 		query += newQuery
		// 	} else {
		// 		query += FieldValueBuilder(field, "=", i, conditionLen)
		// 	}
		// } else if field.Option == constant.BETWEEN {
		// 	query += FieldValueBuilder(field, "BETWEEN", i, conditionLen)
		// } else if field.Option == constant.NOT_BETWEEN {
		// 	query += FieldValueBuilder(field, "NOT BETWEEN", i, conditionLen)
		// } else if field.Option == constant.LESS_THAN {
		// 	query += FieldValueBuilder(field, "<", i, conditionLen)
		// } else if field.Option == constant.GREATER_THAN {
		// 	query += FieldValueBuilder(field, ">", i, conditionLen)
		// } else if field.Option == constant.LESS_THAN_EQUAL {
		// 	query += FieldValueBuilder(field, "<=", i, conditionLen)
		// } else if field.Option == constant.GREATER_THAN_EQUAL {
		// 	query += FieldValueBuilder(field, ">=", i, conditionLen)
		// } else if field.Option == constant.NOT_EQUAL {
		// 	query += FieldValueBuilder(field, "<>", i, conditionLen)
		// } else if field.Option == constant.CONTAINS_PATTERN {
		// 	query += FieldValueBuilder(field, "LIKE", i, conditionLen)
		// } else if field.Option == constant.CONTAINS_NO_PATTERN {
		// 	query += FieldValueBuilder(field, "NOT LIKE", i, conditionLen)
		// }
	}

	return
}

func FieldValueBuilder(field model.Fields, operations string, idx, fieldsLen int) (query string) {
	if field.ToValue != nil {
		if field.DataType == constant.NUMBER {
			query += fmt.Sprintf("(%s %s %s AND %s)", field.FieldName, operations, CastToString(field.FromValue), CastToString(field.ToValue))
		} else if field.DataType == constant.STRING || field.DataType == constant.TIME || field.DataType == constant.DATE || field.DataType == constant.DECIMAL {
			query += fmt.Sprintf("(%s %s %s AND %s)", field.FieldName, operations, ValueStrBuilder(CastToString(field.FromValue)), ValueStrBuilder(CastToString(field.ToValue)))
		}
	} else {
		if field.DataType == constant.NUMBER {
			query += fmt.Sprintf("%s %s %s", field.FieldName, operations, CastToString(field.FromValue))
		} else if field.DataType == constant.STRING || field.DataType == constant.TIME || field.DataType == constant.DATE || field.DataType == constant.DECIMAL {
			query += fmt.Sprintf("%s %s %s", field.FieldName, operations, ValueStrBuilder(CastToString(field.FromValue)))
		}
	}

	if idx < fieldsLen-1 {
		query += " AND "
	} else {
		query += " "
	}

	return
}

func FieldINValueBuilder(field model.Fields, beforeField string, idx, fieldsLen, counterFields int, listsValue []string) (query, afterField string, afterList []string) {

	if beforeField == "" || beforeField != field.FieldName {
		beforeField = field.FieldName
		listsValue = []string{}
	}

	if beforeField == field.FieldName {
		if field.DataType == constant.NUMBER {
			listsValue = append(listsValue, CastToString(field.FromValue))
		} else if field.DataType == constant.STRING || field.DataType == constant.TIME || field.DataType == constant.DATE || field.DataType == constant.DECIMAL {
			listsValue = append(listsValue, ValueStrBuilder(CastToString(field.FromValue)))
		}

		if len(listsValue) == counterFields && idx < fieldsLen-1 {
			query += fmt.Sprintf(" (%s IN(%s)) AND ", field.FieldName, strings.Join(listsValue, ", "))
		} else if len(listsValue) == counterFields {
			query += fmt.Sprintf(" (%s IN(%s)) ", field.FieldName, strings.Join(listsValue, ", "))
		}
	}

	afterField = beforeField
	afterList = listsValue

	return
}
