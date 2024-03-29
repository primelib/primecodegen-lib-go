package requeststruct

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Sample struct to test ResolveRequestParams
type TestRequest struct {
	StringField  *string     `headerParam:"name=string_field,style=simple,explode=false"`
	IntField     *int        `cookieParam:"name=int_field,style=form"`
	BoolField    *bool       `pathParam:"name=bool_field,style=label"`
	TimeField    *time.Time  `queryParam:"name=time_field,style=form"`
	MissingField *float64    // No tag
	SomeStruct   interface{} `bodyParam:""`
}

func TestResolveRequestParams(t *testing.T) {
	stringField := "string_value"
	intField := 42
	boolField := true
	timeField := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)

	req := TestRequest{
		StringField:  &stringField,
		IntField:     &intField,
		BoolField:    &boolField,
		TimeField:    &timeField,
		MissingField: nil,
		SomeStruct:   5,
	}

	expected := RequestParams{
		HeaderParams: map[string]string{"string_field": "string_value"},
		CookieParams: map[string]string{"int_field": "42"},
		PathParams:   map[string]string{"bool_field": "true"},
		QueryParams:  url.Values{"time_field": {timeField.Format(time.RFC3339)}},
		BodyParam:    5,
	}

	result, err := ResolveRequestParams(req)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
