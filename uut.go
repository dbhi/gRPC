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

	"github.com/dbhi/gRPC/client"
)

func main() {
	client.Connect(":8888")
	from := "uut/axis/s"
	to := "uut/axis/m"
	err := client.Register([]string{from, to})
	if err != nil {
		log.Fatal(err)
	}

	for {
		adr, dat, err := client.Read_blocking(from, 2)
		if err != nil {
			log.Fatal(err)
		}
		err = client.Write_blocking(to, adr, uut(dat), 2)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func uut(i int32) int32 {
	return 3 * i
}
