package hydro

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type adapter struct {
	manager Manager
}

func createAdapter(
	manager Manager,
) Adapter {
	out := adapter{
		manager: manager,
	}

	return &out
}

// Hydrate takes an interface instance and hydrate it to a struct instance, in the given ptr
func (app *adapter) Hydrate(dehydrate interface{}) (interface{}, error) {
	// value:
	val := reflect.ValueOf(dehydrate)
	indVal := reflect.Indirect(val)

	// type:
	dehydrateType := reflect.Indirect(val).Type()

	// pointer:
	bridge, err := app.manager.Fetch(dehydrateType.PkgPath(), dehydrateType.Name())
	if err != nil {
		return nil, err
	}

	ptr := bridge.Pointer()
	ptrVal := reflect.ValueOf(ptr)
	indPtrVal := reflect.Indirect(ptrVal)

	amount := dehydrateType.NumField()
	for i := 0; i < amount; i++ {
		field := dehydrateType.Field(i)
		tagName, ok := field.Tag.Lookup("hydro")
		if !ok {
			str := fmt.Sprintf("there is no tag 'hydro' on field (name: %s) of struct (name: %s)", field.Name, dehydrateType.Name())
			return nil, errors.New(str)
		}

		tagProperties := strings.Split(tagName, ",")
		if len(tagProperties) != 2 {
			str := fmt.Sprintf("the hydro tag was expected 2 properties on field (name: %s) of struct (name: %s), %d given", field.Name, dehydrateType.Name(), len(tagProperties))
			return nil, errors.New(str)
		}

		fetchFnName := strings.TrimSpace(tagProperties[0])
		setFieldName := strings.TrimSpace(tagProperties[1])
		callResults := val.MethodByName(fetchFnName).Call([]reflect.Value{})
		if len(callResults) != 1 {
			str := fmt.Sprintf(
				"the method (name: %s) declared in the hydro tag of field (name: %s) of struct (name: %s) was expected to return 1 value, %d returned",
				fetchFnName,
				field.Name,
				dehydrateType.Name(),
				len(callResults),
			)

			return nil, errors.New(str)
		}

		fieldVal := indVal.Field(i)
		indFieldVal := reflect.Indirect(fieldVal)
		fieldTypeKind := indFieldVal.Type().Kind()
		switch fieldTypeKind {
		case reflect.Interface:
			if !callResults[0].CanInterface() {
				str := fmt.Sprintf(
					"the method (name: %s) declared in the hydro tag of field (name: %s) of struct (name: %s) was expected to be used without panicking",
					fetchFnName,
					field.Name,
					dehydrateType.Name(),
				)

				return nil, errors.New(str)
			}

			fieldPtr, err := app.Hydrate(callResults[0].Interface())
			if err != nil {
				return nil, err
			}

			app.setField(dehydrateType, indPtrVal, setFieldName, fieldPtr, bridge)
		default:
			app.setField(dehydrateType, indPtrVal, setFieldName, callResults[0].Interface(), bridge)
			break
		}
	}

	return ptr, nil
}

// Dehydrate takes a struct instance and dehydrate it to an interface instance, then returns it
func (app *adapter) Dehydrate(hydrate interface{}) (interface{}, error) {
	// value:
	val := reflect.ValueOf(hydrate)
	indVal := reflect.Indirect(val)

	// type:
	deHydrateType := reflect.Indirect(val).Type()

	// bridge:
	bridge, err := app.manager.Fetch(deHydrateType.PkgPath(), deHydrateType.Name())
	if err != nil {
		return nil, err
	}

	paramsIns := []interface{}{}
	amount := deHydrateType.NumField()
	for i := 0; i < amount; i++ {
		paramsIns = append(paramsIns, nil)
	}

	for i := 0; i < amount; i++ {
		field := deHydrateType.Field(i)
		tagName, ok := field.Tag.Lookup("hydro")
		if !ok {
			str := fmt.Sprintf("there is no tag 'hydro' on field (name: %s) of struct (name: %s)", field.Name, deHydrateType.Name())
			return nil, errors.New(str)
		}

		index, err := strconv.Atoi(strings.TrimSpace(tagName))
		if err != nil {
			return nil, err
		}

		fieldVal := indVal.Field(i)
		fieldValIns := fieldVal.Interface()
		indFieldVal := reflect.Indirect(fieldVal)
		indFieldType := indFieldVal.Type()
		fieldTypeKind := indFieldType.Kind()
		switch fieldTypeKind {
		case reflect.Struct:
			dehydrate, err := app.Dehydrate(fieldValIns)
			if err != nil {
				return nil, err
			}

			ptrVal, err := app.executeOnDehydrateEvent(bridge, dehydrate, field.Name, deHydrateType.Name())
			if err != nil {
				return nil, err
			}

			paramsIns[index] = ptrVal
			break
		default:
			ptrVal, err := app.executeOnDehydrateEvent(bridge, fieldValIns, field.Name, deHydrateType.Name())
			if err != nil {
				return nil, err
			}

			paramsIns[index] = ptrVal
			break
		}
	}

	params := []reflect.Value{}
	for _, oneParamIns := range paramsIns {
		params = append(params, reflect.ValueOf(oneParamIns))
	}

	constructorFnIns := bridge.ConstructorFn()
	constructorFn := reflect.Indirect(reflect.ValueOf(constructorFnIns))
	results := constructorFn.Call(params)
	if len(results) != 2 {
		str := fmt.Sprintf(
			"the constructor (bridge interface: %s, struct: %s) results wered expected to contain 2 elements, %d returned",
			bridge.Interface(),
			bridge.Struct(),
			len(results),
		)

		return nil, errors.New(str)
	}

	if !results[1].IsNil() {
		return nil, results[1].Interface().(error)
	}

	return results[0].Interface(), nil
}

func (app *adapter) setField(strctType reflect.Type, ptr reflect.Value, fieldName string, ins interface{}, bridge Bridge) error {
	processedIns, err := app.executeOnHydrateEvent(bridge, ins, fieldName, strctType.Name())
	if err != nil {
		return err
	}

	val := reflect.ValueOf(processedIns)
	fieldToSet := ptr.FieldByName(fieldName)
	if !fieldToSet.CanSet() {
		str := fmt.Sprintf("the field (name: %s) of struct (name: %s) is not exported", fieldName, strctType.Name())
		return errors.New(str)
	}

	fieldToSet.Set(val)
	return nil
}

func (app *adapter) executeOnHydrateEvent(bridge Bridge, ins interface{}, fieldName string, structTypeName string) (interface{}, error) {
	if bridge.HasEvents() {
		evts := bridge.Events()
		if evts.HasOnHydrate() {
			onHydrateFn := evts.OnHydrate()
			processedIns, err := onHydrateFn(ins, fieldName, structTypeName)
			if err != nil {
				return nil, err
			}

			if processedIns != nil {
				return processedIns, nil
			}
		}
	}

	return ins, nil
}

func (app *adapter) executeOnDehydrateEvent(bridge Bridge, ins interface{}, fieldName string, structTypeName string) (interface{}, error) {
	if bridge.HasEvents() {
		evts := bridge.Events()
		if evts.HasOnDehydrate() {
			onDehydrateFn := evts.OnDehydrate()
			processedIns, err := onDehydrateFn(ins, fieldName, structTypeName)
			if err != nil {
				return nil, err
			}

			if processedIns != nil {
				return processedIns, nil
			}
		}
	}

	return ins, nil
}
