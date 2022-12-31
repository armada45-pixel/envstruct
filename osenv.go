// version 1.0.7
package envstruct

import (
	"os"
	"reflect"
	"strings"
)

type parserOSOption struct {
	varProp       typeVarProp
	envMap        map[string]string
	opt           Options
	allParserFunc map[reflect.Kind]TypeDefaultBy
}

func parserOSEnv(opts ...parserOSOption) (varProp typeVarProp, err []error) {
	varProp = opts[0].varProp
	opt := opts[0].opt
	envMap := opts[0].envMap
	osenv := os.Environ()

	for _, v := range osenv {
		p := strings.SplitN(v, "=", 2)
		key := p[0]
		value := p[1]

		keyProp, foundVarMap := varProp.OSname[key]
		fileValue, foundEnv := envMap[key]
		if foundVarMap {
			if opt.ReadOS {
				prop := varProp.prop[keyProp]
				typeVar := prop.refTypeField
				newValue, errParserData := parserData(prop.typeProp, typeVar, value, opt.CustomType)
				if len(errParserData) != 0 {
					err = append(err, errParserData...)
				} else {
					if !prop.didRead || opt.OsFirst {
						varProp.prop[keyProp] = varFieldProp{
							defaultIsSet: prop.defaultIsSet,
							defaultValue: prop.defaultValue,
							required:     prop.required,
							typeProp:     prop.typeProp,

							didRead:      true,
							readValue:    newValue,
							refTypeField: typeVar,
						}
					}
				}
			}
		} else if foundEnv {
			os.Setenv(key, fileValue)
			delete(envMap, key)
		}
	}
	return
}
