package demo

import (
	"context"
	"go-template/pkg/log"
)

func Test(ctx context.Context, query string) (map[string]string, error) {
	log.WithContext(ctx).Debug("test request")

	return map[string]string{
		"q": query,
	}, nil
}
