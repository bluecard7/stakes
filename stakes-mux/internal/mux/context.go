package mux

import (
	"context"
)

type userIDKey struct{}

var userIDCtxKey = userIDKey{}

func newContextWithUserID(ctx context.Context, email string) context.Context {
	return context.WithValue(ctx, userIDCtxKey, email)
}

func userIDFromContext(ctx context.Context) (string, bool) {
	email, ok := ctx.Value(userIDCtxKey).(string)
	return email, ok
}
