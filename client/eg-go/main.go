package main

import "C"
import "log"

func main() {
	connect(":8888")
	ids := []string{"uut/axis/s", "uut/axis/m"}
	err := register(ids)
	if err != nil {
		log.Fatal(err)
	}
	for {
		v, err := read_blocking(ids[0], 2)
		if err != nil {
			log.Fatal(err)
		}
		err = write_blocking(ids[1], uut(v), 2)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func uut(i int32) int32 {
	return 3 * i
}
