package requeststruct

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type nestedStruct struct {
	Field1 string
	Field2 int
}

func TestResolveParameterValue(t *testing.T) {
	query := "example"
	limit := 10
	active := true
	date := time.Now()
	uintField := uint(100)
	floatField := 3.14
	nested := nestedStruct{Field1: "nested", Field2: 2}

	tests := []struct {
		value    interface{}
		expected string
	}{
		{&query, "example"},
		{&limit, "10"},
		{&active, "true"},
		{&date, date.Format(time.RFC3339)},
		{&uintField, strconv.FormatUint(uint64(uintField), 10)},
		{&floatField, strconv.FormatFloat(floatField, 'f', -1, 64)},
		{nil, ""},
		{&nested, ""}, // nested structs are not supported yet
	}

	for _, tt := range tests {
		val := reflect.ValueOf(tt.value)
		result := ResolveParameterValue(val, nil)
		assert.Equal(t, tt.expected, result)
	}
}
