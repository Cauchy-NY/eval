package main

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var data string
	echo := func(s string, xs ...interface{}) {
		data += fmt.Sprintf(s, xs...) + "\n"
	}

	echo(`// Code generated by parser/generate/main.go. DO NOT EDIT.`)
	echo(``)
	echo(`package parser`)
	echo(`import (`)
	echo(`"fmt"`)
	echo(`"reflect"`)
	echo(`)`)

	types := []string{
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"float32",
		"float64",
	}

	echo(`
		func or(a, b interface{}) interface{} {
			switch x := a.(type) {
			case bool:
				switch y := b.(type) {
				case bool:
					return x || y
				}
			}
			panic(fmt.Sprintf("invalid operation: %!T(MISSING) %!v(MISSING) %!T(MISSING)", a, "&&", b))
		}
		
		func and(a, b interface{}) interface{} {
			switch x := a.(type) {
			case bool:
				switch y := b.(type) {
				case bool:
					return x && y
				}
			}
			panic(fmt.Sprintf("invalid operation: %!T(MISSING) %!v(MISSING) %!T(MISSING)", a, "&&", b))
		}
		
		func not(a interface{}) interface{} {
			switch x := a.(type) {
			case bool:
				return !x
			}
			panic(fmt.Sprintf("invalid operation: %!v(MISSING) %!T(MISSING)", "!", a))
		}
		
		func ne(a, b interface{}) interface{} {
			return reflect.DeepEqual(false, eq(a, b))
		}
	`)

	helpers := []struct {
		name, op        string
		noFloat, string bool
	}{
		{
			name:   "eq",
			op:     "==",
			string: true,
		},
		{
			name:   "lt",
			op:     "<",
			string: true,
		},
		{
			name:   "gt",
			op:     ">",
			string: true,
		},
		{
			name:   "le",
			op:     "<=",
			string: true,
		},
		{
			name:   "ge",
			op:     ">=",
			string: true,
		},
		{
			name:   "add",
			op:     "+",
			string: true,
		},
		{
			name: "sub",
			op:   "-",
		},
		{
			name: "mul",
			op:   "*",
		},
		{
			name: "div",
			op:   "/",
		},
		{
			name:    "mod",
			op:      "%",
			noFloat: true,
		},
	}

	for _, helper := range helpers {
		name := helper.name
		op := helper.op
		echo(`func %v(a, b interface{}) interface{} {`, name)
		echo(`switch x := a.(type) {`)
		for i, a := range types {
			if helper.noFloat && strings.HasPrefix(a, "float") {
				continue
			}
			echo(`case %v:`, a)
			echo(`switch y := b.(type) {`)
			for j, b := range types {
				if helper.noFloat && strings.HasPrefix(b, "float") {
					continue
				}
				echo(`case %v:`, b)
				if i == j {
					echo(`return x %v y`, op)
				}
				if i < j {
					echo(`return %v(x) %v y`, b, op)
				}
				if i > j {
					echo(`return x %v %v(y)`, op, a)
				}
			}
			echo(`}`)
		}
		if helper.string {
			echo(`case string:`)
			echo(`switch y := b.(type) {`)
			echo(`case string: return x %v y`, op)
			echo(`}`)
		}
		echo(`}`)
		if name == "eq" {
			echo(`if isNil(a) && isNil(b) { return true }`)
			echo(`return reflect.DeepEqual(a, b)`)
		} else {
			echo(`panic(fmt.Sprintf("invalid operation: %%T %%v %%T", a, "%v", b))`, op)
		}
		echo(`}`)
		echo(``)
	}

	echo(`
		func isNil(v interface{}) bool {
			if v == nil {
				return true
			}
			r := reflect.ValueOf(v)
			switch r.Kind() {
			case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
				return r.IsNil()
			default:
				return false
			}
		}
	`)

	b, err := format.Source([]byte(data))
	check(err)
	err = ioutil.WriteFile("./parser/helpers.go", b, 0644)
	check(err)
}
