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
	"log"
	"time"

	"github.com/dbhi/gRPC/client"
)

func main() {
	client.Connect(":8888")
	to := "uut/axis/s"
	from := "uut/axis/m"
	err := client.Register([]string{to, from})
	if err != nil {
		log.Fatal(err)
	}

	k := 10

	for y := 0; y < 3; y++ {
		for x := 0; x < k; x++ {
			log.Println("Write request to", y, x)
			err = client.Write_blocking(to, int32(1000+x), int32(x), 2)
			if err != nil {
				log.Fatal(err)
			}
		}
		for x := 0; x < k; x++ {
			log.Println("Read response from", y, x)
			adr, dat, err := client.Read_blocking(from, 2)
			if err != nil {
				log.Fatal(err)
			}
			if adr != int32(1000+x) {
				log.Fatal("mismatch!", x, 1000+x, adr)
			}
			if dat != int32(x*3) {
				log.Fatal("mismatch!", x, x*3, dat)
			}
		}
		log.Println("Wait 5 seconds...", y)
		time.Sleep(5 * time.Second)
	}
}
