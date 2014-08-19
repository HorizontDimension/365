package server

import (
	//"pretty"
	"github.com/kr/pretty"
	"reflect"
)

func GenerateForms(any interface{}) {
	typ := reflect.TypeOf(any)
	// if a pointer to a struct is passed, get the type of the dereferenced object
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	// create an attribute data structure as a map of types keyed by a string.
	attrs := make(map[string]reflect.Type)
	// Only structs are supported so return an empty result if the passed object
	// isn't a struct
	if typ.Kind() != reflect.Struct {
		//pretty.Printf("%v type can't have attributes inspected\n", typ.Kind())
		//	pretty.Println(attrs)
	}

	fieldsinstruct(typ, attrs)

	pretty.Println(attrs)
}

func fieldsinstruct(typ reflect.Type, attrs map[string]reflect.Type) {

	// if a pointer to a struct is passed, get the type of the dereferenced object
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		//	pretty.Printf("%v type can't have attributes inspected\n", typ.Kind())
		//pretty.Println(attrs)
	}

	// loop through the struct's fields and set the map
	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i)
		if p.Type.Kind() == reflect.Struct {
			fieldsinstruct(p.Type, attrs)
		}
		if !p.Anonymous {
			attrs[p.Name] = p.Type
		}

	}

}
