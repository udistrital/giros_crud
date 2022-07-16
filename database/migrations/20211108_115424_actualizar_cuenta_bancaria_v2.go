package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ActualizarCuentaBancariaV2_20211108_115424 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ActualizarCuentaBancariaV2_20211108_115424{}
	m.Created = "20211108_115424"

	migration.Register("ActualizarCuentaBancariaV2_20211108_115424", m)
}

// Run the migrations
func (m *ActualizarCuentaBancariaV2_20211108_115424) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20211108_115424_actualizar_cuenta_bancaria_v2_up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}

// Reverse the migrations
func (m *ActualizarCuentaBancariaV2_20211108_115424) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20211108_115424_actualizar_cuenta_bancaria_v2_down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}
