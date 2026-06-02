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
func goReadBlocking(id string, t int) (adr int32, dat int32) {
	adr, dat, err := client.Read_blocking(id, t)
	if err != nil {
		log.Fatal(err)
	}
	return adr, dat
}

//export goWriteBlocking
func goWriteBlocking(id string, adr int32, dat int32, t int) {
	err := client.Write_blocking(id, adr, dat, t)
	if err != nil {
		log.Fatal(err)
	}
	return
}
