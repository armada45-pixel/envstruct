package envstruct

import (
	"errors"
	"reflect"
)

func parserData(varProp typeVarProp, typeVar reflect.StructField, keyProp int, value string) (newValue any, err []error) {

	typeVarKind := typeVar.Type.Kind()
	parserFunc, foundFunc := defaultBuiltInParsers[typeVarKind]
	if !foundFunc {
		err = append(err, errors.New("Parser Function For Type "+typeVarKind.String()+" In Field "+typeVar.Name+""))
	} else {
		parseValue, errParse := parserFunc(value)
		if errParse != nil {
			err = append(err, errParse)
			// newValue = varProp.prop[keyProp].defaultValue
		} else {
			newValue = parseValue
		}
	}
	return
}
