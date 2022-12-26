package envstruct

import (
	"errors"
	"os"
	"strings"
)

type parserOSOption struct {
	varProp    typeVarProp
	envFileMap map[string]string
}

func parserEnv(opts ...parserOSOption) (varProp typeVarProp, err []error) {
	varProp = opts[0].varProp
	osenv := os.Environ()
	err = []error{}
	// r := map[string]string{}
	for _, v := range osenv {
		p := strings.SplitN(v, "=", 2)
		key := p[0]
		value := p[1]

		keyProp, found := varProp.OSname[key]
		if found {
			typeVar := varProp.prop[keyProp].refTypeField
			typeVarKind := typeVar.Type.Kind()
			parserFunc, foundFunc := defaultBuiltInParsers[typeVarKind]
			if !foundFunc {
				err = append(err, errors.New("Parser Function For Type "+typeVarKind.String()+" In Field "+typeVar.Name+""))
			} else {
				parseValue, errParse := parserFunc(value)
				newValue := varProp.prop[keyProp].defaultValue
				if errParse != nil {
					err = append(err, errParse)
				} else {
					newValue = parseValue
				}
				varProp.prop[keyProp] = varFieldProp{
					defaultValue: varProp.prop[keyProp].defaultValue,
					required:     varProp.prop[keyProp].required,

					didRead:      true,
					readValue:    newValue,
					refTypeField: typeVar,
				}
			}
		}
	}
	return
}
