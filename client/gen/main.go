package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/umarcor/dbhi/router/lib"
	grpc "google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to backend: %v\n", err)
		os.Exit(1)
	}
	c := lib.NewChansClient(conn)

	ctx := context.Background()

	to := "uut/axis/s"
	from := "uut/axis/m"

	for _, id := range []string{to, from} {
		err = isChanRegistered(ctx, c, id)
		if err != nil {
			log.Fatal("Check for chan", id, "failed:", err)
		}
	}

	k := 8192

	for y := 0; y < 3; y++ {
		x := 0
		for x < k {
			_, err = c.Wr(ctx, &lib.Write{
				Id:  to,
				Val: int32(x),
			})
			if err != nil {
				if m, _ := regexp.Match("desc = full", []byte(err.Error())); m {
					log.Println("Full stream. Retry in 2 seconds...")
					time.Sleep(2 * time.Second)
				} else {
					log.Fatal(err)
				}
			} else {
				x++
			}
		}

		x = 0
		for x < k {
			v, err := c.Rd(ctx, &lib.Id{Id: from})
			if err != nil {
				if m, _ := regexp.Match("desc = empty", []byte(err.Error())); m {
					log.Println("Empty stream. Retry in 2 seconds...")
					time.Sleep(2 * time.Second)
				} else {
					log.Fatal(err)
				}
			} else {
				if v.Val != int32(x*3) {
					log.Fatal("mismatch!", x, x*3, v.Val)
				}
				x++
			}
		}
		log.Println("Wait 5 seconds...", y)
		time.Sleep(5 * time.Second)
		y++
	}
}

func isChanRegistered(ctx context.Context, client lib.ChansClient, id string) (err error) {
	reg := false
	for !reg {
		u, err := client.List(ctx, &lib.Void{})
		if err != nil {
			return err
		}
		for _, t := range u.Chans {
			if id == t.Id {
				reg = true
				break
			}
		}
		if !reg {
			log.Println("Chan", id, "not registered yet. Retry in 2 seconds...")
			time.Sleep(2 * time.Second)
		}
	}
	return
}
