package tracer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTracer(t *testing.T) {
	jaeger, err := JaegerTraceProvider("test-app", "test", "http://localhost:14268/api/traces")
	assert.NoError(t, err)
	assert.NotNil(t, jaeger)

	RegisterTracer(jaeger)

	ctx := context.Background()
	ctx, span := SpanStart(ctx, "test-span")
	assert.NotNil(t, ctx)
	assert.NotNil(t, span)
	defer span.Finish()

	t.Run("Trace ID", func(t *testing.T) {
		assert.NotEmpty(t, span.TraceId())
	})

	t.Run("Trace Events", func(t *testing.T) {
		span.AddEvents("test-event")
	})

	t.Run("Trace Tags String", func(t *testing.T) {
		span.AddTags(SpanTagString("test-key-string", "test-value"))
	})

	t.Run("Trace Tags Int64", func(t *testing.T) {
		span.AddTags(SpanTagInt64("test-key-int", 100))
	})

	t.Run("Trace Tags Bool", func(t *testing.T) {
		span.AddTags(SpanTagBool("test-key-bool", true))
	})

	t.Run("Trace Tags Float64", func(t *testing.T) {
		span.AddTags(SpanTagFloat64("test-key-float", 100.0))
	})

	t.Run("Trace Tags String Slice", func(t *testing.T) {
		span.AddTags(SpanTagStringSlice("test-key-string-slice", []string{"test-value"}))
	})

	t.Run("Trace Tags Int", func(t *testing.T) {
		span.AddTags(SpanTagInt("test-key-int", 100))
	})

	t.Run("Trace Tags Any", func(t *testing.T) {
		span.AddAnyTags(map[string]interface{}{
			"test-key-string": "test-value",
			"test-key-int":    100,
			"test-key-bool":   true,
			"test-key-float":  100.0,
		})
	})

	t.Run("Trace Error", func(t *testing.T) {
		span.AddError(assert.AnError)
	})
}
