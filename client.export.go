package main

import (
	"C"

	"log"

	"github.com/dbhi/gRPC/client"
)

func main() {}

//export goConnect
func goConnect(addr string) {
	log.Println(addr)
	client.Connect(addr)
}

//export goRegister
func goRegister(s []string) (e string) {
	err := client.Register(s)
	if err != nil {
		return err.Error()
	}
	return
}

//export goReadBlocking
func goReadBlocking(id string, t int) (v int32) {
	v, err := client.Read_blocking(id, t)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

//export goWriteBlocking
func goWriteBlocking(id string, v int32, t int) {
	err := client.Write_blocking(id, v, t)
	if err != nil {
		log.Fatal(err)
	}
	return
}
