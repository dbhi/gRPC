package main

import "C"

import (
	"log"
)

//export go_connect
func go_connect(addr string) {
	log.Println(addr)
	connect(addr)
}

//export go_register
func go_register(s []string) {
	err := register(s)
	if err != nil {
		log.Fatal(err)
	}
}

//export go_read
func go_read(id string) (v int32, e string) {
	e = ""
	var err error
	v, err = read(id)
	if err != nil {
		e = err.Error()
	}
	return
}

//export go_read_blocking
func go_read_blocking(id string, t int) (v int32) {
	v, err := read_blocking(id, t)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

//export go_write
func go_write(id string, v int32) (e string) {
	e = ""
	err := write(id, v)
	if err != nil {
		e = err.Error()
	}
	return
}

//export go_write_blocking
func go_write_blocking(id string, v int32, t int) {
	err := write_blocking(id, v, t)
	if err != nil {
		log.Fatal(err)
	}
	return
}
