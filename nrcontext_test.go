package nrcontext

import (
	"context"
	"testing"

	"github.com/TV4/nrcontext/nrmock"
)

func TestWithApplication(t *testing.T) {
	app := nrmock.NewApplication()

	if got := WithApplication(context.Background(), app).Value(appKey); got != app {
		t.Fatalf("expected context value for appKey to be app, got %v", got)
	}

	if got := WithApplication(context.Background(), nil).Value(appKey); got != nil {
		t.Fatalf("expected context value for appKey to be nil, got %v", got)
	}
}

func TestWithTransaction(t *testing.T) {
	txn := nrmock.NewTransaction("foo")

	if got := WithTransaction(context.Background(), txn).Value(txnKey); got != txn {
		t.Fatalf("expected context value for txnKey to be txn, got %v", got)
	}

	if got := WithTransaction(context.Background(), nil).Value(txnKey); got != nil {
		t.Fatalf("expected context value for txnKey to be nil, got %v", got)
	}
}

func TestApplication(t *testing.T) {
	app := nrmock.NewApplication()

	if got := Application(WithApplication(context.Background(), app)); got != app {
		t.Fatalf("expected app, got %v", got)
	}

	if got := Application(context.Background()); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}

func TestTransaction(t *testing.T) {
	txn := nrmock.NewTransaction("foo")

	if got := Transaction(WithTransaction(context.Background(), txn)); got != txn {
		t.Fatalf("expected app, got %v", got)
	}

	if got := Transaction(context.Background()); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}
