package hydro

import (
	"testing"

	"github.com/deepvalue-network/software/libs/hash"
	"github.com/deepvalue-network/software/libs/hydro/internals"
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
		WithDehydratedInterface((*internals.SimpleInterface)(nil)).
		WithDehydratedConstructor(internals.NewSimple).
		WithDehydratedPointer(new(internals.DehydrateSimpleStruct)).
		WithHydratedPointer(new(internals.HydrateSimpleStruct)).
		Now()

	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// create complex bridge:
	complexBridge, err := NewBridgeBuilder().Create().
		WithDehydratedInterface((*internals.ComplexInterface)(nil)).
		WithDehydratedConstructor(internals.NewComplex).
		WithDehydratedPointer(new(internals.DehydrateComplexStruct)).
		WithHydratedPointer(new(internals.HydrateComplexStruct)).
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

	// execute:
	VerifyAdapterUsingJSForTests(adapter, complex, t)
}
