package nrcontext_test

import (
	"context"
	"errors"
	"net/http/httptest"

	"github.com/TV4/nrcontext"
	"github.com/TV4/nrcontext/nrmock"
)

func ExampleApplication() {
	// Get a context containing a newrelic.Application (a mock in this example)
	ctx := nrcontext.WithApplication(context.Background(), nrmock.NewApplication())

	// Now you can retrieve the newrelic.Application from
	// ctx using the Application function
	app := nrcontext.Application(ctx)

	// Setup a http.ResponseWriter and a *http.Request for use in this example
	w, r := httptest.NewRecorder(), httptest.NewRequest("GET", "/foo", nil)

	// Start a transaction
	txn := app.StartTransaction("name", w, r)

	// End the transaction
	txn.End()
}

func ExampleTransaction() {
	// Get a context containing a newrelic.Transaction (a mock in this example)
	ctx := nrcontext.WithTransaction(context.Background(), nrmock.NewTransaction("name"))

	// Now you can retrieve the newrelic.Transaction from the context
	txn := nrcontext.Transaction(ctx)

	// You can now use the transaction to notice errors
	txn.NoticeError(errors.New("some error you want to notice"))

	// And to add attributes
	txn.AddAttribute("foo", "bar")
}

func ExampleWithApplication() {
	// Get a context containing a newrelic.Application (a mock in this example)
	ctx := nrcontext.WithApplication(context.Background(), nrmock.NewApplication())

	ctx.Done()
}

func ExampleWithTransaction() {
	// Get a context containing a newrelic.Transaction (a mock in this example)
	ctx := nrcontext.WithTransaction(context.Background(), nrmock.NewTransaction("name"))

	ctx.Done()
}
