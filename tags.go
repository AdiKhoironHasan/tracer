package tracer

import "go.opentelemetry.io/otel/attribute"

// SpanTagString creates a string attribute key value pair
func SpanTagString(key string, value string) attribute.KeyValue {
	return attribute.Key(key).String(value)
}

// SpanTagInt64 creates an int64 attribute key value pair
func SpanTagInt64(key string, value int64) attribute.KeyValue {
	return attribute.Key(key).Int64(value)
}

// SpanTagBool creates a bool attribute key value pair
func SpanTagBool(key string, value bool) attribute.KeyValue {
	return attribute.Key(key).Bool(value)
}

// SpanTagFloat64 creates a float64 attribute key value pair
func SpanTagFloat64(key string, value float64) attribute.KeyValue {
	return attribute.Key(key).Float64(value)
}

// SpanTagStringSlice creates a string slice attribute key value pair
func SpanTagStringSlice(key string, value []string) attribute.KeyValue {
	return attribute.Key(key).StringSlice(value)
}

// SpanTagInt creates an int attribute key value pair
func SpanTagInt(key string, value int) attribute.KeyValue {
	return attribute.Key(key).Int(value)
}
