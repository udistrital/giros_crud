package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ActualizarCuentaBancaria_20211013_155612 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ActualizarCuentaBancaria_20211013_155612{}
	m.Created = "20211013_155612"

	migration.Register("ActualizarCuentaBancaria_20211013_155612", m)
}

// Run the migrations
func (m *ActualizarCuentaBancaria_20211013_155612) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/giros_update_cuenta_up.sql")

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
func (m *ActualizarCuentaBancaria_20211013_155612) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/giros_update_cuenta_down.sql")

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
