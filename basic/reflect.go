package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	p := reflect.ValueOf(&x)
	value := p.Elem()
	value.SetFloat(33.3)
	fmt.Println("修改后的值：x = ", x)
}
