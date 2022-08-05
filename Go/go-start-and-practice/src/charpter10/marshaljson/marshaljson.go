package main

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	type Skill struct {
		Name  string
		Level int
	}

	type Actor struct {
		Name   string
		Age    int
		Skills []Skill
	}

	a := Actor{
		Name: "cow boy",
		Age:  37,
		Skills: []Skill{
			{Name: "Roll and roll", Level: 1},
			{Name: "Flash your dog eye", Level: 2},
			{Name: "Time to have Lunch", Level: 3},
		},
	}
	if result, err := MarshalJson(a); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}

func MarshalJson(v interface{}) (string, error) {
	var b bytes.Buffer

	if err := writeAny(&b, reflect.ValueOf(v)); err == nil {
		return b.String(), nil
	} else {
		return "", err
	}
}

func writeAny(buff *bytes.Buffer, value reflect.Value) error {
	switch value.Kind() {
	case reflect.String:
		// 带双引号的字符串
		buff.WriteString(strconv.Quote(value.String()))
	case reflect.Int:
		buff.WriteString(strconv.FormatInt(value.Int(), 10))
	case reflect.Slice:
		return WriteSlice(buff, value)
	case reflect.Struct:
		return WriteStruct(buff, value)
	default:
		return errors.New("unsupport kind: " + value.Kind().String())
	}

	return nil
}

func WriteSlice(buff *bytes.Buffer, value reflect.Value) error {
	buff.WriteString("[")
	for s := 0; s < value.Len(); s++ {
		sliceValue := value.Index(s)
		writeAny(buff, sliceValue)
		if s < value.Len()-1 {
			// 每个切片后面写入逗号，最后一个不写
			buff.WriteString(",")
		}
	}
	buff.WriteString("]")
	return nil
}

func WriteStruct(buf *bytes.Buffer, value reflect.Value) error {
	valueType := value.Type()
	buf.WriteString("{")
	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)
		fieldType := valueType.Field(i)
		buf.WriteString("\":")
		buf.WriteString(fieldType.Name)
		buf.WriteString("\"")
		writeAny(buf, fieldValue)
		if i < value.NumField()-1 {
			buf.WriteString(",")
		}
	}
	buf.WriteString("}")
	return nil
}
