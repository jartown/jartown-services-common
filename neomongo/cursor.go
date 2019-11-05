package neomongo

import "context"

type Cursor interface {
	All(ctx context.Context, results interface{}) error
	Close(ctx context.Context) error
	Decode(val interface{}) error
	Err() error
	ID() int64
	Next(ctx context.Context) bool
}
