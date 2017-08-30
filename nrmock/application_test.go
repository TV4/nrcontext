package nrmock

import "testing"

func TestNewApplication(t *testing.T) {
	app := NewApplication(func(*Application) {})

	if app == nil {
		t.Fatalf("NewApplication() = nil, want *Application")
	}
}

func TestStartTransaction(t *testing.T) {
	name := "foo"

	app := &Application{}

	txn, ok := app.StartTransaction(name, nil, nil).(*Transaction)
	if !ok {
		t.Fatalf("expected app.StartTransaction to return a *Transaction")
	}

	if got, want := txn.Name, name; got != want {
		t.Fatalf("txn.Name = %q, want %q", got, want)
	}
}

func TestRecordCustomEvent(t *testing.T) {
	app := &Application{}

	app.RecordCustomEvent("foo", map[string]interface{}{"bar": 123})
	app.RecordCustomEvent("foo", map[string]interface{}{"baz": 456})

	foo := app.CustomEvents["foo"]

	if got, want := len(foo), 2; got != want {
		t.Fatalf("len(foo) = %d, want %d", got, want)
	}
}

func TestWaitForConnection(t *testing.T) {
	app := &Application{}

	if app.WaitForConnection(0) != nil {
		t.Fatalf("expected WaitForConnection to return nil")
	}
}

func TestShutdown(t *testing.T) {
	NewApplication().Shutdown(0)
}
