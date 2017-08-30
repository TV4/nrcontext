package nrmock

import (
	"errors"
	"testing"
)

func TestNewTransaction(t *testing.T) {
	txn := NewTransaction("", func(txn *Transaction) {
		txn.Code = 123
	})

	if txn == nil {
		t.Fatalf("NewTransaction() = nil, want *Transaction")
	}

	if got, want := txn.Code, 123; got != want {
		t.Fatalf("txn.Code = %d, want %d", got, want)
	}
}

func TestEnd(t *testing.T) {
	txn := &Transaction{}

	if got, want := txn.Ended, false; got != want {
		t.Fatalf("txn.Ended = %v, want %v", got, want)
	}

	txn.End()

	if got, want := txn.Ended, true; got != want {
		t.Fatalf("txn.Ended = %v, want %v", got, want)
	}
}

func TestIgnore(t *testing.T) {
	txn := &Transaction{}

	if got, want := txn.Ignored, false; got != want {
		t.Fatalf("txn.Ignored = %v, want %v", got, want)
	}

	txn.Ignore()

	if got, want := txn.Ignored, true; got != want {
		t.Fatalf("txn.Ignored = %v, want %v", got, want)
	}
}

func TestSetName(t *testing.T) {
	txn := &Transaction{}

	if got, want := txn.Name, ""; got != want {
		t.Fatalf("txn.Name = %q, want %q", got, want)
	}

	txn.SetName("bar")

	if got, want := txn.Name, "bar"; got != want {
		t.Fatalf("txn.Name = %q, want %q", got, want)
	}
}

func TestNoticeError(t *testing.T) {
	txn := &Transaction{}

	for i := 0; i < 25; i++ {
		if got, want := txn.Errors, i; got != want {
			t.Fatalf("txn.Errors = %d, want %d", got, want)
		}

		if err := txn.NoticeError(errors.New("test error")); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	if got, want := txn.Errors, 25; got != want {
		t.Fatalf("txn.Errors = %d, want %d", got, want)
	}
}

func TestAddAttribute(t *testing.T) {
	txn := &Transaction{}

	txn.AddAttribute("foo", 123)
	txn.AddAttribute("bar", true)
	txn.AddAttribute("baz", "qux")

	if _, ok := txn.Attributes["foo"].(int); !ok {
		t.Fatalf("expected foo attribute to be int")
	}

	if _, ok := txn.Attributes["bar"].(bool); !ok {
		t.Fatalf("expected bar attribute to be bool")
	}

	if _, ok := txn.Attributes["baz"].(string); !ok {
		t.Fatalf("expected baz attribute to be string")
	}
}

func TestStartSegmentNow(t *testing.T) {
	txn := &Transaction{}

	txn.StartSegmentNow()

	txn.End()
}

func TestHeader(t *testing.T) {
	txn := &Transaction{}

	if got, want := len(txn.Header()), 0; got != want {
		t.Fatalf("len(txn.Header()) = %d, want %d", got, want)
	}
}

func TestWrite(t *testing.T) {
	txn := &Transaction{}

	n, err := txn.Write([]byte("foo"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got, want := n, 3; got != want {
		t.Fatalf("n = %d, want %d", got, want)
	}
}

func TestWriteHeader(t *testing.T) {
	txn := &Transaction{}

	txn.WriteHeader(768)

	if got, want := txn.Code, 768; got != want {
		t.Fatalf("txn.Code = %d, want %d", got, want)
	}
}
