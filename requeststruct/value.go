package requeststruct

import (
	"reflect"
	"strconv"
	"time"
)

type ValueToStringConfig struct {
	TimeFormat string // TimeFormat is the format to use for time.Time values.
}

var defaultValToStringConfig = &ValueToStringConfig{
	TimeFormat: time.RFC3339,
}

// ResolveParameterValue resolves the value of a field in a struct into a string to pass as request parameter.
func ResolveParameterValue(value reflect.Value, cfg *ValueToStringConfig) string {
	if cfg == nil {
		cfg = defaultValToStringConfig
	}

	// ptr
	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return ""
		}
		value = value.Elem()
	}

	// value
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(value.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(value.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(value.Float(), 'f', -1, 64)
	case reflect.String:
		return value.String()
	case reflect.Bool:
		return strconv.FormatBool(value.Bool())
	case reflect.Struct:
		if value.Type().PkgPath() == "time" && value.Type().Name() == "Time" {
			return value.Interface().(time.Time).Format(cfg.TimeFormat)
		}
	default:
		return ""
	}

	return ""
}
