package nrmock

import (
	"net/http"
	"time"

	newrelic "github.com/newrelic/go-agent"
)

var _ newrelic.Application = &Application{}

// Application is a mock implementation of newrelic.Application
type Application struct {
	AllTransactions   []*Transaction
	LatestTransaction *Transaction

	CustomEvents map[string][]map[string]interface{}
}

// NewApplication creates a new *Application
func NewApplication(options ...func(*Application)) *Application {
	app := &Application{}

	for _, option := range options {
		option(app)
	}

	return app
}

// StartTransaction creates a mocked transaction, stores it for later inspection, and returns it
func (a *Application) StartTransaction(name string, w http.ResponseWriter, r *http.Request) newrelic.Transaction {
	txn := &Transaction{Name: name}

	a.AllTransactions = append(a.AllTransactions, txn)
	a.LatestTransaction = txn

	return txn
}

// RecordCustomEvent stores the provided eventType and params for later inspection
func (a *Application) RecordCustomEvent(eventType string, params map[string]interface{}) error {
	if a.CustomEvents == nil {
		a.CustomEvents = map[string][]map[string]interface{}{}
	}

	if a.CustomEvents[eventType] == nil {
		a.CustomEvents[eventType] = []map[string]interface{}{}
	}

	a.CustomEvents[eventType] = append(a.CustomEvents[eventType], params)

	return nil
}

// WaitForConnection does nothing
func (a *Application) WaitForConnection(timeout time.Duration) error {
	return nil
}

// Shutdown does nothing
func (a *Application) Shutdown(timeout time.Duration) {}
