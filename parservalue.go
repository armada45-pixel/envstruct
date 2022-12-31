// version 1.0.7
package envstruct

import (
	"errors"
	"reflect"
)

func parserData(typeKind reflect.Type, typeVar reflect.StructField, value string, allParserFunc map[reflect.Type]TypeDefaultBy) (newValue any, err []error) {

	typeVarKind := typeVar.Type.Kind()
	searchDefault, found := allParserFunc[typeKind]
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
