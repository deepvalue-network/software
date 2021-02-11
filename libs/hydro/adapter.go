package hydro

import (
	"errors"
	"fmt"
	"log"
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
	if dehydrate == nil {
		return nil, nil
	}

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

	hydratedBridge := bridge.Hydrated()
	ptr := hydratedBridge.Pointer()
	ptrVal := reflect.ValueOf(ptr)
	indPtrVal := reflect.Indirect(ptrVal)

	amount := dehydrateType.NumField()
	for i := 0; i < amount; i++ {
		field := dehydrateType.Field(i)
		tagName, ok := field.Tag.Lookup("hydro")
		if !ok {
			// there is no tag log:
			str := fmt.Sprintf("there is no tag 'hydro' on field (name: %s) of struct (name: %s), skip it.", field.Name, dehydrateType.Name())
			log.Println(str)

			// skip:
			continue
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
		if fieldVal.IsZero() {
			return nil, nil
		}

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
				// log:
				str := fmt.Sprintf("there was a property that was probably not registered. so trying the instance directly: %s", err.Error())
				log.Println(str)

				// we try to use the instance directly, since it was not registered, we will deal with it in the event:
				fieldPtr = callResults[0].Interface()
			}

			app.setHydratedField(dehydrateType, indPtrVal, setFieldName, fieldPtr, hydratedBridge)
		case reflect.Map:
			app.setHydratedMapField(dehydrateType, indPtrVal, setFieldName, callResults[0].Interface(), hydratedBridge)
			break
		case reflect.Slice:
			app.setHydratedSliceField(dehydrateType, indPtrVal, setFieldName, callResults[0].Interface(), hydratedBridge)
		default:
			app.setHydratedField(dehydrateType, indPtrVal, setFieldName, callResults[0].Interface(), hydratedBridge)
			break
		}
	}

	return ptr, nil
}

// Dehydrate takes a struct instance and dehydrate it to an interface instance, then returns it
func (app *adapter) Dehydrate(hydrate interface{}) (interface{}, error) {
	if hydrate == nil {
		return nil, nil
	}

	// value:
	val := reflect.ValueOf(hydrate)
	indVal := reflect.Indirect(val)

	// type:
	hydrateType := reflect.Indirect(val).Type()

	// bridge:
	bridge, err := app.manager.Fetch(hydrateType.PkgPath(), hydrateType.Name())
	if err != nil {
		return nil, err
	}

	dehydratedBridge := bridge.Dehydrated()
	paramsIns := []interface{}{}
	amount := hydrateType.NumField()
	for i := 0; i < amount; i++ {
		paramsIns = append(paramsIns, nil)
	}

	for i := 0; i < amount; i++ {
		field := hydrateType.Field(i)
		tagName, ok := field.Tag.Lookup("hydro")
		if !ok {
			// log:
			str := fmt.Sprintf("there is no tag 'hydro' on field (name: %s) of struct (name: %s), skip", field.Name, hydrateType.Name())
			log.Println(str)

			// remove the params:
			paramsIns = append(paramsIns[:i], paramsIns[i+1:]...)

			// skip:
			continue
		}

		index, err := strconv.Atoi(strings.TrimSpace(tagName))
		if err != nil {
			return nil, err
		}

		fieldVal := indVal.Field(i)
		indFieldVal := reflect.Indirect(fieldVal)
		indFieldKind := indFieldVal.Kind()
		if fieldVal.IsZero() {
			if indFieldKind == reflect.Ptr {
				paramsIns[index] = fieldVal
				continue
			}

		}

		fieldValIns := fieldVal.Interface()
		switch indFieldKind {
		case reflect.Struct:
			dehydrate, err := app.Dehydrate(fieldValIns)
			if err != nil {
				// log:
				str := fmt.Sprintf("there was a property that was probably not registered. so trying the instance directly: %s", err.Error())
				log.Println(str)

				// we try to use the instance directly, since it was not registered, we will deal with it in the event:
				dehydrate = fieldValIns
			}

			if err == nil && dehydrate == nil {
				paramsIns[index] = nil
				break
			}

			ptrVal, err := app.executeOnDehydrateEvent(dehydratedBridge, dehydrate, field.Name, hydrateType.Name())
			if err != nil {
				return nil, err
			}

			paramsIns[index] = ptrVal
			break
		case reflect.Map:
			var results reflect.Value
			mapKeys := fieldVal.MapKeys()
			for index, keyname := range mapKeys {

				el := fieldVal.MapIndex(keyname)
				dehydrated, err := app.Dehydrate(el.Interface())
				if err != nil {
					return nil, err
				}

				if index <= 0 {
					dehydratedVal := reflect.ValueOf(dehydrated)
					dehydrateType := reflect.Indirect(dehydratedVal).Type()
					elBridge, err := app.manager.Fetch(dehydrateType.PkgPath(), dehydrateType.Name())
					if err != nil {
						return nil, err
					}

					outMapValue := reflect.TypeOf(elBridge.Dehydrated().Interface()).Elem()
					mapType := reflect.MapOf(keyname.Type(), outMapValue)
					results = reflect.MakeMapWithSize(mapType, len(mapKeys))
				}

				elem := reflect.ValueOf(dehydrated)
				results.SetMapIndex(keyname, elem)
			}

			ptrVal, err := app.executeOnDehydrateEvent(dehydratedBridge, results.Interface(), field.Name, hydrateType.Name())
			if err != nil {
				return nil, err
			}

			paramsIns[index] = ptrVal
			break
		case reflect.Slice:
			var results reflect.Value
			sliceLength := fieldVal.Len()
			for i := 0; i < sliceLength; i++ {
				el := fieldVal.Index(i)
				dehydrated, err := app.Dehydrate(el.Interface())
				if err != nil {
					return nil, err
				}

				if i <= 0 {
					dehydratedVal := reflect.ValueOf(dehydrated)
					dehydrateType := reflect.Indirect(dehydratedVal).Type()
					elBridge, err := app.manager.Fetch(dehydrateType.PkgPath(), dehydrateType.Name())
					if err != nil {
						return nil, err
					}

					ourSliceValue := reflect.TypeOf(elBridge.Dehydrated().Interface()).Elem()
					sliceType := reflect.SliceOf(ourSliceValue)
					results = reflect.MakeSlice(sliceType, sliceLength, fieldVal.Cap())
				}

				elem := reflect.ValueOf(dehydrated)
				results.Index(i).Set(elem)
			}

			ptrVal, err := app.executeOnDehydrateEvent(dehydratedBridge, results.Interface(), field.Name, hydrateType.Name())
			if err != nil {
				return nil, err
			}

			paramsIns[index] = ptrVal
			break
		default:
			ptrVal, err := app.executeOnDehydrateEvent(dehydratedBridge, fieldValIns, field.Name, hydrateType.Name())
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

	constructorFnIns := dehydratedBridge.ConstructorFn()
	constructorFn := reflect.Indirect(reflect.ValueOf(constructorFnIns))
	results := constructorFn.Call(params)
	if len(results) != 2 {
		dehydratedBridgePointerType := reflect.TypeOf(dehydratedBridge.Pointer())
		str := fmt.Sprintf(
			"the dehydrated bridge's constructor (pointer struct: %s) results wered expected to contain 2 elements, %d returned",
			dehydratedBridgePointerType.Name(),
			len(results),
		)

		return nil, errors.New(str)
	}

	if !results[1].IsNil() {
		return nil, results[1].Interface().(error)
	}

	return results[0].Interface(), nil
}

func (app *adapter) setHydratedMapField(strctType reflect.Type, ptr reflect.Value, fieldName string, ins interface{}, bridge Hydrated) error {
	val := reflect.ValueOf(ins)
	indVal := reflect.Indirect(val)
	mapKeys := indVal.MapKeys()

	var results reflect.Value

	for index, keyname := range mapKeys {
		if index <= 0 {
			outMapValue := ptr.FieldByName(fieldName)
			mapType := reflect.MapOf(keyname.Type(), outMapValue.Type().Elem())
			results = reflect.MakeMapWithSize(mapType, len(mapKeys))
		}

		el := indVal.MapIndex(keyname)
		hydrated, err := app.Hydrate(el.Interface())
		if hydrated == nil && err != nil {
			return err
		}

		elem := reflect.ValueOf(hydrated)
		results.SetMapIndex(keyname, elem)
	}

	return app.setHydratedField(strctType, ptr, fieldName, results.Interface(), bridge)
}

func (app *adapter) setHydratedSliceField(strctType reflect.Type, ptr reflect.Value, fieldName string, ins interface{}, bridge Hydrated) error {
	val := reflect.ValueOf(ins)
	indVal := reflect.Indirect(val)
	sliceLength := indVal.Len()

	var results reflect.Value
	for i := 0; i < sliceLength; i++ {
		if i <= 0 {
			if ptr.IsZero() {
				str := fmt.Sprintf("the pointer (name: %s) contains a field (name: %s) that is nil", ptr.Type().Name(), fieldName)
				return errors.New(str)
			}

			ptrField := ptr.FieldByName(fieldName)
			if !ptrField.IsValid() {
				str := fmt.Sprintf("the field (type: %s, field: %s) is invalid", ptr.Type().Name(), fieldName)
				return errors.New(str)
			}

			if ptrField.IsZero() {
				str := fmt.Sprintf("the slice (type: %s, field: %s) is empty", ptr.Type().Name(), fieldName)
				return errors.New(str)
			}

			outSliceValue := ptr.FieldByName(fieldName).Index(i)
			sliceType := reflect.SliceOf(outSliceValue.Type())
			results = reflect.MakeSlice(sliceType, sliceLength, indVal.Cap())
		}

		el := indVal.Index(i)
		hydrated, err := app.Hydrate(el.Interface())
		if hydrated == nil && err != nil {
			return err
		}

		elem := reflect.ValueOf(hydrated)
		results.Index(i).Set(elem)
	}

	return app.setHydratedField(strctType, ptr, fieldName, results.Interface(), bridge)
}

func (app *adapter) setHydratedField(strctType reflect.Type, ptr reflect.Value, fieldName string, ins interface{}, bridge Hydrated) error {
	processedIns, err := app.executeOnHydrateEvent(bridge, ins, fieldName, strctType.Name())
	if err != nil {
		return err
	}

	if processedIns == nil {
		return nil
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

func (app *adapter) executeOnHydrateEvent(bridge Hydrated, ins interface{}, fieldName string, structTypeName string) (interface{}, error) {
	if bridge.HasEvent() {
		onHydrateFn := bridge.Event()
		processedIns, err := onHydrateFn(ins, fieldName, structTypeName)
		if err != nil {
			return nil, err
		}

		if processedIns != nil {
			return processedIns, nil
		}
	}

	return ins, nil
}

func (app *adapter) executeOnDehydrateEvent(bridge Dehydrated, ins interface{}, fieldName string, structTypeName string) (interface{}, error) {
	if bridge.HasEvent() {
		onDehydrateFn := bridge.Event()
		processedIns, err := onDehydrateFn(ins, fieldName, structTypeName)
		if err != nil {
			return nil, err
		}

		if processedIns != nil {
			return processedIns, nil
		}
	}

	return ins, nil
}
