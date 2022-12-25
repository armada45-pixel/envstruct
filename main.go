// version 1.0.1
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
	// ReadOS bool // default false

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
	err = []error{}
	opt := checkOptions(optA)

	fileName := opt.FileName
	if !opt.IgnoreFile {
		file, errFile := os.Open(fileName)
		if errFile != nil {
			err = append(err, errFile)
		} else {
			defer file.Close()
			var varProp = typeVarProp{}
			var errCheckVar []error
			if opt.VarPtr != nil {
				varProp, errCheckVar = prepareVar(opt.VarPtr)
				err = append(err, errCheckVar...)
			}
			newVarProp, errParser := parserFile(file, parserOption{
				varProp: varProp,
			})
			err = append(err, errParser...)
			// if ((len(errCheckVar) == 0) && (errSet := set(newVarProp); len(errSet) != 0)) {
			if len(errCheckVar) == 0 {
				if errSet := set(newVarProp); len(errSet) != 0 {
					err = append(err, errSet...)
				}
			}
		}
	}
	return
}

func set(newVarProp typeVarProp) (err []error) {
	err = []error{}
	ref := newVarProp.ref
	refType := ref.Type()
	for i := 0; i < refType.NumField(); i++ {
		newProp := newVarProp.prop[i]
		refTypeField := newProp.refTypeField
		newValue := newProp.readValue
		typee := refTypeField.Type
		refField := ref.Field(i)
		fieldee := refField
		if typee.Kind() == reflect.Ptr {
			typee = typee.Elem()
			fieldee = refField.Elem()
		}
		if !newProp.didRead {
			newValue = newProp.defaultValue
		}
		if !reflect.ValueOf(newValue).IsZero() {
			fieldee.Set(reflect.ValueOf(newValue))
		} else if newProp.required {
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
