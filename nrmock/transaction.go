package nrmock

import (
	"net/http"

	newrelic "github.com/newrelic/go-agent"
)

var _ newrelic.Transaction = &Transaction{}

// Transaction is a mock implementation of newrelic.Transaction
type Transaction struct {
	Name       string
	Ended      bool
	Ignored    bool
	Errors     int
	Code       int
	Attributes map[string]interface{}
}

// NewTransaction creates a new *Transaction
func NewTransaction(name string, options ...func(*Transaction)) *Transaction {
	txn := &Transaction{Name: name}

	for _, option := range options {
		option(txn)
	}

	return txn
}

// End marks the transaction as ended
func (t *Transaction) End() error {
	t.Ended = true

	return nil
}

// Ignore marks the transaction as ignored
func (t *Transaction) Ignore() error {
	t.Ignored = true

	return nil
}

// SetName sets the transaction name
func (t *Transaction) SetName(name string) error {
	t.Name = name

	return nil
}

// NoticeError increments the errors counter
func (t *Transaction) NoticeError(err error) error {
	t.Errors++

	return nil
}

// AddAttribute adds the provided key and value to the transaction attributes for later inspection
func (t *Transaction) AddAttribute(key string, value interface{}) error {
	if t.Attributes == nil {
		t.Attributes = map[string]interface{}{}
	}

	t.Attributes[key] = value

	return nil
}

// StartSegmentNow returns a blank newrelic.SegmentStartTime
func (t *Transaction) StartSegmentNow() newrelic.SegmentStartTime {
	return newrelic.SegmentStartTime{}
}

// Mock implementation of the ResponseWriter interface

// Header returns and empty header map
func (t *Transaction) Header() http.Header {
	return http.Header{}
}

// Write writes nothing
func (t *Transaction) Write(b []byte) (int, error) {
	return len(b), nil
}

// WriteHeader stores the provided status code for later inspection
func (t *Transaction) WriteHeader(code int) {
	t.Code = code
}
