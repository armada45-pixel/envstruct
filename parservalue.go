package envstruct

import (
	"errors"
	"reflect"
)

func parserData(typeKind reflect.Type, typeVar reflect.StructField, value string) (newValue any, err []error) {

	typeVarKind := typeVar.Type.Kind()
	searchDefault, found := defaultByType[typeKind]
	if !found {
		err = append(err, errors.New("Parser Function For Type "+typeVarKind.String()+" In Field "+typeVar.Name+""))
	} else {
		parseValue, errParse := searchDefault.ParserFunc(value)
		if errParse != nil {
			err = append(err, errParse)
		} else {
			newValue = parseValue
		}
	}
	return
}
