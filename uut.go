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
