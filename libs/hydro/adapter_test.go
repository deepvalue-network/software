package hydro

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/steve-care-software/products/libs/hash"
	"github.com/steve-care-software/products/libs/hydro/internals"
)

func TestHydrate_Success(t *testing.T) {
	first := "firstValue"
	second := "secondValue"
	another := uint(567)
	simple, _ := internals.NewSimple(first, second)

	hsh, _ := hash.NewAdapter().FromBytes([]byte("this is an hash"))
	complex, _ := internals.NewComplex(simple, another, *hsh)

	// create simple bridge:
	simpleBridge, err := NewBridgeBuilder().Create().
		WithConstructor(internals.NewSimple).
		WithPointer(new(internals.HydrateSimpleStruct)).
		WithInterfaceName("github.com/steve-care-software/products/libs/hydro/internals/SimpleInterface").
		WithStructName("github.com/steve-care-software/products/libs/hydro/internals/dehydrateSimpleStruct").
		Now()

	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// create complex bridge:
	complexBridge, err := NewBridgeBuilder().Create().
		WithConstructor(internals.NewComplex).
		WithPointer(new(internals.HydrateComplexStruct)).
		WithInterfaceName("github.com/steve-care-software/products/libs/hydro/internals/ComplexInterface").
		WithStructName("github.com/steve-care-software/products/libs/hydro/internals/dehydrateComplexStruct").
		OnHydrate(internals.ComplexStructOnHydrateEventFn).
		OnDehydrate(internals.ComplexStructOnDehydrateEventFn).
		Now()

	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// build the manager:
	manager := NewManagerFactory().Create()

	// register the bridges:
	manager.Register(simpleBridge)
	manager.Register(complexBridge)

	// create the adapter:
	adapter, err := NewAdapterBuilder().Create().WithManager(manager).Now()
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// hydrates:
	hydrate, err := adapter.Hydrate(complex)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// convert to json:
	js, err := json.Marshal(hydrate)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// dehydrates:
	dehydrate, err := adapter.Dehydrate(hydrate)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// hydrates again:
	reHydrate, err := adapter.Hydrate(dehydrate)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	reJS, err := json.Marshal(reHydrate)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if bytes.Compare(js, reJS) != 0 {
		t.Errorf("the conversions failed: \n1): %s\n2): %s\n", js, reJS)
		return
	}
}
