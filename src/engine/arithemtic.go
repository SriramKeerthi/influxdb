package engine

import (
	"common"
	"fmt"
	"parser"
	"protocol"
	"strconv"
)

type ArithmeticOperator func(elems []*parser.Value, fields []string, point *protocol.Point) (*protocol.FieldValue, error)

var registeredArithmeticOperator map[string]ArithmeticOperator

func init() {
	registeredArithmeticOperator = make(map[string]ArithmeticOperator)
	registeredArithmeticOperator["+"] = PlusOperator
	registeredArithmeticOperator["-"] = MinusOperator
	registeredArithmeticOperator["*"] = MultiplyOperator
	registeredArithmeticOperator["/"] = DivideOperator
}

func getValue(value *parser.Value, fields []string, point *protocol.Point) (*protocol.FieldValue, error) {
	switch value.Type {
	case parser.ValueSimpleName:
		for idx, f := range fields {
			if f == value.Name {
				return point.Values[idx], nil
			}
		}
		return nil, fmt.Errorf("Invalid column name %s", value.Name)
	case parser.ValueExpression:
		operator := registeredArithmeticOperator[value.Name]
		return operator(value.Elems, fields, point)
	case parser.ValueInt:
		v, _ := strconv.ParseInt(value.Name, 10, 64)
		return &protocol.FieldValue{Int64Value: &v}, nil
	case parser.ValueFloat:
		v, _ := strconv.ParseFloat(value.Name, 64)
		return &protocol.FieldValue{DoubleValue: &v}, nil
	}

	return nil, fmt.Errorf("Value cannot be evaluated for type %v", value.Type)
}

func PlusOperator(elems []*parser.Value, fields []string, point *protocol.Point) (*protocol.FieldValue, error) {
	leftValue, err := getValue(elems[0], fields, point)
	if err != nil {
		return nil, err
	}
	rightValues, err := getValue(elems[1], fields, point)
	if err != nil {
		return nil, err
	}
	left, right, valueType := common.CoerceValues(leftValue, rightValues)
	switch valueType {
	case common.TYPE_DOUBLE:
		value := left.(float64) + right.(float64)
		return &protocol.FieldValue{DoubleValue: &value}, nil
	case common.TYPE_INT:
		value := left.(int64) + right.(int64)
		return &protocol.FieldValue{Int64Value: &value}, nil
	}
	return nil, fmt.Errorf("+ operator doesn't work with %v types", valueType)
}

func MinusOperator(elems []*parser.Value, fields []string, point *protocol.Point) (*protocol.FieldValue, error) {
	leftValue, err := getValue(elems[0], fields, point)
	if err != nil {
		return nil, err
	}
	rightValues, err := getValue(elems[1], fields, point)
	if err != nil {
		return nil, err
	}
	left, right, valueType := common.CoerceValues(leftValue, rightValues)
	switch valueType {
	case common.TYPE_DOUBLE:
		value := left.(float64) - right.(float64)
		return &protocol.FieldValue{DoubleValue: &value}, nil
	case common.TYPE_INT:
		value := left.(int64) - right.(int64)
		return &protocol.FieldValue{Int64Value: &value}, nil
	}
	return nil, fmt.Errorf("- operator doesn't work with %v types", valueType)
}

func MultiplyOperator(elems []*parser.Value, fields []string, point *protocol.Point) (*protocol.FieldValue, error) {
	leftValue, err := getValue(elems[0], fields, point)
	if err != nil {
		return nil, err
	}
	rightValues, err := getValue(elems[1], fields, point)
	if err != nil {
		return nil, err
	}
	left, right, valueType := common.CoerceValues(leftValue, rightValues)
	switch valueType {
	case common.TYPE_DOUBLE:
		value := left.(float64) * right.(float64)
		return &protocol.FieldValue{DoubleValue: &value}, nil
	case common.TYPE_INT:
		value := left.(int64) * right.(int64)
		return &protocol.FieldValue{Int64Value: &value}, nil
	}
	return nil, fmt.Errorf("* operator doesn't work with %v types", valueType)
}

func DivideOperator(elems []*parser.Value, fields []string, point *protocol.Point) (*protocol.FieldValue, error) {
	leftValue, err := getValue(elems[0], fields, point)
	if err != nil {
		return nil, err
	}
	rightValues, err := getValue(elems[1], fields, point)
	if err != nil {
		return nil, err
	}
	left, right, valueType := common.CoerceValues(leftValue, rightValues)
	switch valueType {
	case common.TYPE_DOUBLE:
		value := left.(float64) / right.(float64)
		return &protocol.FieldValue{DoubleValue: &value}, nil
	case common.TYPE_INT:
		value := left.(int64) / right.(int64)
		return &protocol.FieldValue{Int64Value: &value}, nil
	}
	return nil, fmt.Errorf("/ operator doesn't work with %v types", valueType)
}
