package envstruct

import (
	"errors"
	"reflect"
	"strconv"
)

type typeVarProp struct {
	check   bool
	prop    map[int]varFieldProp
	OSname  map[string]int
	ENVname map[string]int
	varPtr  interface{}
	ref     reflect.Value
	// refType reflect.Type
}

type varFieldProp struct {
	defaultValue any
	defaultIsSet bool
	required     bool
	didRead      bool
	readValue    any
	refTypeField reflect.StructField
	typeProp     reflect.Type
}

func checkOptions(optArray []Options) Options {
	// Default Option
	defOpt := Options{
		FileName: ".env",
		// IgnoreFile: false,
		// ReadAll:    false,
		// PutToOs:    false,
		// OverRide:   false,
		// osFirst:    false,
		// ReadOS:     false,
	}
	if len(optArray) == 0 {
		return defOpt
	}
	opt := optArray[0]

	if !opt.IgnoreFile && opt.FileName == "" {
		opt.FileName = defOpt.FileName
	}

	return opt
}

func checkVarType(VarPtr interface{}) (reflect.Value, error) {
	ptrRef := reflect.ValueOf(VarPtr)
	if ptrRef.Kind() != reflect.Ptr {
		return reflect.ValueOf(nil), ErrNotAStructPtr
	}
	ref := ptrRef.Elem()
	if ref.Kind() != reflect.Struct {
		return reflect.ValueOf(nil), ErrNotAStructPtr
	}
	return ref, nil
}

func prepareVar(VarPtr interface{}, ref reflect.Value, allParserFunc map[reflect.Type]TypeDefaultBy) (ls typeVarProp, err []error) {

	ls = typeVarProp{
		check:   false,
		prop:    make(map[int]varFieldProp),
		OSname:  make(map[string]int),
		ENVname: make(map[string]int),
	}
	refType := ref.Type()
	ls.ref = ref

	for i := 0; i < refType.NumField(); i++ {
		ls.check = true
		refField := refType.Field(i)
		envName := refField.Tag.Get("env")
		if len(envName) != 0 {
			ls.ENVname[envName] = i
		}

		osName := refField.Tag.Get("os")
		if len(osName) != 0 {
			ls.OSname[osName] = i
		}

		var required bool = false
		requiredString := refField.Tag.Get("required")
		if len(requiredString) != 0 {
			requiredBool, errorParse := strconv.ParseBool(requiredString)
			if errorParse != nil {
				err = append(err, errorParse)
			} else {
				required = requiredBool
			}
		}

		var defaultValue any
		var defaultValueField_i = ref.Field(i)
		var defaultIsSet bool
		sometype := defaultValueField_i.Type()
		if defaultValueField_i.IsZero() && (reflect.TypeOf(reflect.Bool) != sometype) {
			defaultString := refField.Tag.Get("default")
			if len(defaultString) != 0 {
				searchDefault, found := allParserFunc[sometype]
				if !found {
					err = append(err, errors.New("Parser Function For Type "+sometype.String()+" In Field "+refField.Name+""))
				} else {
					parseValue, errParse := searchDefault.ParserFunc(defaultString)
					if errParse != nil {
						err = append(err, errParse)
					} else {
						defaultValue = parseValue
						defaultIsSet = true
					}
				}
			}
		}

		ls.prop[i] = varFieldProp{
			defaultIsSet: defaultIsSet,
			defaultValue: defaultValue,
			required:     required,
			refTypeField: refField,
			typeProp:     defaultValueField_i.Type(),
		}
	}
	return
}
