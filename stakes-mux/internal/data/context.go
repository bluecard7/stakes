package data

import (
	"context"
)

type key struct{}

var ctxKey = key{}

// NewContext returns a new Context carrying recordTable.
func NewContext(ctx context.Context, recordTable RecordTable) context.Context {
	return context.WithValue(ctx, ctxKey, recordTable)
}

// FromContext extracts the record table from ctx, if present.
func FromContext(ctx context.Context) (RecordTable, bool) {
	recordTable, ok := ctx.Value(ctxKey).(RecordTable)
	return recordTable, ok
}
