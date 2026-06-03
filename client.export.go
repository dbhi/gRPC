/*
Copyright 2018-2026
Unai Martinez-Corral <unai.martinezcorral@ehu.eus>
University of the Basque Country (UPV/EHU) <ehu.eus>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"C"

	"log"

	"github.com/dbhi/gRPC/client"
)

func main() {}

//export goConnect
func goConnect(addr string) {
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

//export goRead
func goRead(id string) (adr int32, dat int32, e string) {
	adr, dat, err := client.Read(id)
	if err == nil {
		return adr, dat, ""
	}
	e = err.Error()
	if e != "EMPTY" {
		log.Fatal(err)
	}
	return 0, 0, e
}

//export goReadBlocking
func goReadBlocking(id string, t int) (adr int32, dat int32) {
	adr, dat, err := client.Read_blocking(id, t)
	if err != nil {
		log.Fatal(err)
	}
	return adr, dat
}

//export goWrite
func goWrite(id string, adr int32, dat int32) (e string) {
	err := client.Write(id, adr, dat)
	if err == nil {
		return ""
	}
	e = err.Error()
	if e != "FULL" {
		log.Fatal(err)
	}
	return e
}

//export goWriteBlocking
func goWriteBlocking(id string, adr int32, dat int32, t int) {
	err := client.Write_blocking(id, adr, dat, t)
	if err != nil {
		log.Fatal(err)
	}
	return
}
