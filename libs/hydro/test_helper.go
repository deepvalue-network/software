package hydro

import (
	"bytes"
	"encoding/json"
	"testing"
)

// VerifyAdapterUsingJSForTests verifies the adapter for tests, using JSON conversions
func VerifyAdapterUsingJSForTests(adapter Adapter, dehydrated interface{}, t *testing.T) {
	// hydrates:
	hydrated, err := adapter.Hydrate(dehydrated)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// convert to json:
	js, err := json.Marshal(hydrated)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// dehydrate:
	dehydrated, err = adapter.Dehydrate(hydrated)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// hydrates again:
	reHydrated, err := adapter.Hydrate(dehydrated)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	reJS, err := json.Marshal(reHydrated)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if bytes.Compare(js, reJS) != 0 {
		t.Errorf("the conversions failed: \n1): %s\n2): %s\n", js, reJS)
		return
	}
}
