// version 1.0.4
package envstruct

import (
	"errors"
	"os"
	"reflect"
)

type Options struct {
	// Type : Variable Pointer
	// Pointer Struct Variable to config key,type for read env file or os.env variable
	VarPtr interface{} // don't have defualt // nil

	// path to read file
	FileName string // default .env

	// if FileName not empty will change to false
	// if true = not read file
	IgnoreFile bool // default false

	// Incoming in next version
	// if ture = read os.env and put to Variable
	ReadOS bool // default false

	// if true = read all in file and put to os.env ( if PutToOs true )
	// ReadAll bool // default false

	// if true = put variable from file to os.env variable
	// PutToOs bool // default false

	// if true = if already variable in same name in os.env variable will replace
	// OverRide bool // default false

	// if true = if have os.env and env file will choose os.env
	// osFirst bool // default false
}

func Setup(optA ...Options) (err []error) {

	opt := checkOptions(optA)

	fileName := opt.FileName
	var varProp = typeVarProp{}
	var errCheckVar []error

	if opt.VarPtr != nil {
		varProp, errCheckVar = prepareVar(opt.VarPtr)
		err = append(err, errCheckVar...)
	}

	if !opt.IgnoreFile {
		file, errFile := os.Open(fileName)
		if errFile != nil {
			err = append(err, errFile)
		} else {
			defer file.Close()
			var errParser []error
			varProp, errParser = parserFile(file, parserFileOption{
				varProp: varProp,
			})
			err = append(err, errParser...)
		}
	}

	if opt.ReadOS {
		var errParser []error
		varProp, errParser = parserOSEnv(parserOSOption{
			varProp: varProp,
		})
		err = append(err, errParser...)
	}

	if len(errCheckVar) == 0 {
		if errSet := setVar(varProp); len(errSet) != 0 {
			err = append(err, errSet...)
		}
	}
	return
}

func setVar(newVarProp typeVarProp) (err []error) {

	ref := newVarProp.ref
	refType := ref.Type()
	for i := 0; i < refType.NumField(); i++ {
		newProp := newVarProp.prop[i]
		refField := ref.Field(i)
		value := newProp.defaultValue
		fieldee := refField
		if newProp.didRead {
			refTypeField := newProp.refTypeField
			typee := refTypeField.Type
			if typee.Kind() == reflect.Ptr {
				typee = typee.Elem()
				fieldee = refField.Elem()
			}
			value = newProp.readValue
		}
		refValue := reflect.ValueOf(value)
		fieldee.Set(refValue)
		if !newProp.didRead && newProp.required && refValue.IsZero() {
			err = append(err, errors.New("Field "+refField.Type().Name()+" Required is True, But can't get any value."))
		}
	}
	return
}

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
	required     bool
	didRead      bool
	readValue    any
	refTypeField reflect.StructField
}
