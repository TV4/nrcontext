package nrcontext

import (
	"context"

	newrelic "github.com/newrelic/go-agent"
)

type key int

const (
	appKey key = iota
	txnKey
)

// WithApplication returns a copy of parent in which the value associated with appKey is app
func WithApplication(parent context.Context, app newrelic.Application) context.Context {
	return context.WithValue(parent, appKey, app)
}

// Application returns newrelic.Application stored in ctx, returns nil if missing
func Application(ctx context.Context) newrelic.Application {
	if app, ok := ctx.Value(appKey).(newrelic.Application); ok {
		return app
	}

	return nil
}

// WithTransaction returns a copy of parent in which the value associated with txnKey is txn
func WithTransaction(parent context.Context, txn newrelic.Transaction) context.Context {
	return context.WithValue(parent, txnKey, txn)
}

// Transaction returns newrelic.Transaction stored in ctx, returns nil if missing
func Transaction(ctx context.Context) newrelic.Transaction {
	if txn, ok := ctx.Value(txnKey).(newrelic.Transaction); ok {
		return txn
	}

	return nil
}
