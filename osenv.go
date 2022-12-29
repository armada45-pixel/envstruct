package envstruct

import (
	"os"
	"strings"
)

type parserOSOption struct {
	varProp typeVarProp
	// envFileMap map[string]string
	opt Options
}

func parserOSEnv(opts ...parserOSOption) (varProp typeVarProp, err []error) {
	varProp = opts[0].varProp
	opt := opts[0].opt
	osenv := os.Environ()

	for _, v := range osenv {
		p := strings.SplitN(v, "=", 2)
		key := p[0]
		value := p[1]

		keyProp, found := varProp.OSname[key]
		if found {
			prop := varProp.prop[keyProp]
			typeVar := prop.refTypeField
			newValue, errParserData := parserData(varProp, typeVar, keyProp, value)
			if len(errParserData) != 0 {
				err = append(err, errParserData...)
			}
			if !prop.didRead || opt.OsFirst {
				varProp.prop[keyProp] = varFieldProp{
					defaultValue: prop.defaultValue,
					required:     prop.required,

					didRead:      true,
					readValue:    newValue,
					refTypeField: typeVar,
				}
			}
		}
	}
	return
}
