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
			err = client.Write_blocking(to, int32(x), 2)
			if err != nil {
				log.Fatal(err)
			}
		}
		for x := 0; x < k; x++ {
			log.Println("Read response from", y, x)
			v, err := client.Read_blocking(from, 2)
			if err != nil {
				log.Fatal(err)
			}
			if v != int32(x*3) {
				log.Fatal("mismatch!", x, x*3, v)
			}
		}
		log.Println("Wait 5 seconds...", y)
		time.Sleep(5 * time.Second)
	}
}
