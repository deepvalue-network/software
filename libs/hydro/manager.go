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
	in := bridge.Interface()
	app.mp[in] = bridge

	strctName := bridge.Struct()
	app.mp[strctName] = bridge

	ptr := bridge.Pointer()
	ptrType := reflect.Indirect(reflect.ValueOf(ptr)).Type()
	ptrName := fmt.Sprintf("%s/%s", ptrType.PkgPath(), ptrType.Name())
	app.mp[ptrName] = bridge
}
