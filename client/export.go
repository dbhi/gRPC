package main

import "C"

import (
	"log"
)

//export goConnect
func goConnect(addr string) {
	log.Println(addr)
	connect(addr)
}

/*
//export goList
func goList() (l []string, e string) {
	var err error
	l, err = list()
	if err != nil {
		return l, err.Error()
	}
	return
}
*/

//export goCheckIds
func goCheckIds(ids []string) bool {
	l, err := list()
	if err != nil {
		return false
	}
	for _, id := range ids {
		if !check_id(l, id) {
			return false
		}
	}
	return true
}

//export goRegister
func goRegister(s []string) (e string) {
	err := register(s)
	if err != nil {
		return err.Error()
	}
	return
}

//export goRead
func goRead(id string) (v int32, e string) {
	var err error
	v, err = read(id)
	if err != nil {
		return v, err.Error()
	}
	return
}

//export goReadBlocking
func goReadBlocking(id string, t int) (v int32) {
	v, err := read_blocking(id, t)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

//export goWrite
func goWrite(id string, v int32) (e string) {
	err := write(id, v)
	if err != nil {
		return err.Error()
	}
	return
}

//export goWriteBlocking
func goWriteBlocking(id string, v int32, t int) {
	err := write_blocking(id, v, t)
	if err != nil {
		log.Fatal(err)
	}
	return
}
