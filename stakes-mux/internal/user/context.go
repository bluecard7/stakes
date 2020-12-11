package user

import (
	"context"
)

type key struct{}

var ctxKey = key{}

// NewContext returns a new Context carrying email.
func NewContext(ctx context.Context, email string) context.Context {
	return context.WithValue(ctx, ctxKey, email)
}

// FromContext extracts the user email from ctx, if present.
func FromContext(ctx context.Context) (string, bool) {
	email, ok := ctx.Value(ctxKey).(string)
	return email, ok
}
