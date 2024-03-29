package requeststruct

import (
	"net/url"
	"reflect"
)

type RequestParams struct {
	HeaderParams map[string]string
	CookieParams map[string]string
	PathParams   map[string]string
	QueryParams  url.Values
	BodyParam    interface{}
}

func ResolveRequestParams(requestStruct any) (RequestParams, error) {
	result := RequestParams{
		HeaderParams: map[string]string{},
		CookieParams: map[string]string{},
		PathParams:   map[string]string{},
		QueryParams:  url.Values{},
		BodyParam:    nil,
	}
	structType := reflect.TypeOf(requestStruct)
	structVal := reflect.ValueOf(requestStruct)

	for i := 0; i < structType.NumField(); i++ {
		fieldType := structType.Field(i)
		valType := structVal.Field(i)
		tagValue := fieldType.Tag

		// headerParam
		if val, ok := tagValue.Lookup("headerParam"); ok {
			// parse
			data := parseKVTags(val)
			name := data["name"]
			style := data["style"]
			if style != "simple" {
				style = "simple"
			}
			explode := data["explode"]
			if explode != "true" {
				explode = "false"
			}

			// add to result
			result.HeaderParams[name] = ResolveParameterValue(valType, nil)
		}

		// cookieParams
		if val, ok := tagValue.Lookup("cookieParam"); ok {
			// parse
			data := parseKVTags(val)
			name := data["name"]
			style := data["style"]
			if style != "form" && style != "spaceDelimited" && style != "pipeDelimited" && style != "deepObject" {
				style = "form"
			}

			// add to result
			result.CookieParams[name] = ResolveParameterValue(valType, nil)
		}

		// pathParam
		if val, ok := tagValue.Lookup("pathParam"); ok {
			// parse
			data := parseKVTags(val)
			name := data["name"]
			style := data["style"]
			if style != "simple" && style != "label" && style != "matrix" {
				style = "simple"
			}
			explode := data["explode"]
			if explode != "true" {
				explode = "false"
			}

			// add to result
			result.PathParams[name] = ResolveParameterValue(valType, nil)
		}

		// queryParam
		if val, ok := tagValue.Lookup("queryParam"); ok {
			// parse
			data := parseKVTags(val)
			name := data["name"]
			style := data["style"]
			if style != "form" && style != "spaceDelimited" && style != "pipeDelimited" && style != "deepObject" {
				style = "form"
			}

			// add to result
			result.QueryParams.Add(name, ResolveParameterValue(valType, nil))
		}

		// bodyParam
		if _, ok := tagValue.Lookup("bodyParam"); ok {
			// add to result
			result.BodyParam = valType.Interface()
		}
	}

	return result, nil
}
