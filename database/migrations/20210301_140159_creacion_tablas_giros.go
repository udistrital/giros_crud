package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreacionTablasGiros_20210301_140159 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionTablasGiros_20210301_140159{}
	m.Created = "20210301_140159"

	migration.Register("CreacionTablasGiros_20210301_140159", m)
}

// Run the migrations
func (m *CreacionTablasGiros_20210301_140159) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/giros_create_up.sql")

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

// Reverse the migrations
func (m *CreacionTablasGiros_20210301_140159) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/giros_create_down.sql")

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
