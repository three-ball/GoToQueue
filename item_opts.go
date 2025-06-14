// Package gotoqueue provides enqueue options using the functional options pattern.
package gotoqueue

import (
	"context"
	"time"
)

// EnqueueOption represents an option for enqueuing items
type EnqueueOption func(*EnqueueOptions)

// EnqueueOptions contains all the options for enqueuing an item
type EnqueueOptions struct {
	itemCtx    context.Context // Context for the enqueue operation
	metadata   map[string]interface{}
	expireTime time.Time
}

// WithContext sets the context for the enqueue operation
func WithContext(ctx context.Context) EnqueueOption {
	return func(opts *EnqueueOptions) {
		opts.itemCtx = ctx
	}
}

// WithExpiration sets an expiration time for the queue item
func WithExpiration(expireTime time.Time) EnqueueOption {
	return func(opts *EnqueueOptions) {
		opts.expireTime = expireTime
	}
}

// WithExpirationDuration sets an expiration duration from now for the queue item
func WithExpirationDuration(duration time.Duration) EnqueueOption {
	return func(opts *EnqueueOptions) {
		opts.expireTime = time.Now().Add(duration)
	}
}

// WithMetadata sets metadata for the queue item
func WithMetadata(metadata map[string]interface{}) EnqueueOption {
	return func(opts *EnqueueOptions) {
		if opts.metadata == nil {
			opts.metadata = make(map[string]interface{})
		}
		for k, v := range metadata {
			opts.metadata[k] = v
		}
	}
}

// defaultEnqueueOptions returns default options
func defaultEnqueueOptions() *EnqueueOptions {
	return &EnqueueOptions{
		itemCtx:  context.Background(),
		metadata: nil,
	}
}

// applyEnqueueOptions applies all options to the default options
func applyEnqueueOptions(opts ...EnqueueOption) *EnqueueOptions {
	options := defaultEnqueueOptions()
	for _, opt := range opts {
		opt(options)
	}
	return options
}
