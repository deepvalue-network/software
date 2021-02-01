package hydro

import (
	"errors"
	"fmt"
	"reflect"
)

type manager struct {
	mp map[string]Bridge
}

func createManager(
	mp map[string]Bridge,
) Manager {
	out := manager{
		mp: mp,
	}

	return &out
}

// Fetch fetches a bridge by interface or struct name
func (app *manager) Fetch(pkg string, name string) (Bridge, error) {
	path := fmt.Sprintf("%s/%s", pkg, name)
	if bridge, ok := app.mp[path]; ok {
		return bridge, nil
	}

	str := fmt.Sprintf("the interface or struct '%s' is not registered", path)
	return nil, errors.New(str)
}

// Register registers a bridge
func (app *manager) Register(bridge Bridge) {
	// hydrated:
	hydrated := bridge.Hydrated()

	// hydrated pointer:
	hydratedPtrType := reflect.Indirect(reflect.ValueOf(hydrated.Pointer())).Type()
	hydratedPtrName := fmt.Sprintf(doubleStringPattern, hydratedPtrType.PkgPath(), hydratedPtrType.Name())
	app.mp[hydratedPtrName] = bridge

	// dehydrated:
	dehyrated := bridge.Dehydrated()

	// dehydrated interface:
	dehyratedInterfaceType := reflect.TypeOf(dehyrated.Interface())
	dehyratedInterfaceName := fmt.Sprintf(doubleStringPattern, dehyratedInterfaceType.PkgPath(), dehyratedInterfaceType.Name())
	app.mp[dehyratedInterfaceName] = bridge

	// dehydrated pointer:
	dehydratedPtrType := reflect.Indirect(reflect.ValueOf(dehyrated.Pointer())).Type()
	dehydratedPtrName := fmt.Sprintf(doubleStringPattern, dehydratedPtrType.PkgPath(), dehydratedPtrType.Name())
	app.mp[dehydratedPtrName] = bridge
}
