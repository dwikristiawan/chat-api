package service

import "context"

type service struct{}

type Service interface {
	NewChat(context.Context, *uint)
}
