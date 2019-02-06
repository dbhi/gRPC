package main

import (
	"context"
	"errors"
	"log"
	"regexp"
	"time"

	"github.com/umarcor/dbhi/gRPC/lib"
	grpc "google.golang.org/grpc"
)

var client lib.ChansClient
var ctx context.Context

// connect establishes a connection with the server at 'addr', and gets identifiers for both the client and the context.
// For now, it is insecure (see WithInsecure).
func connect(addr string) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to backend:", err)
	}
	client = lib.NewChansClient(conn)
	ctx = context.Background()
}

// register checks if each of the elements in a slice of IDs is found in the list retrieved from the server.
// If it is not present, it is registered (added).
func register(ids []string) (err error) {
	l, err := list()
	if err != nil {
		return
	}
	for _, id := range ids {
		if !check_id(l, id) {
			_, err = client.Reg(ctx, &lib.Register{Id: id, Length: 8192})
			if err != nil {
				return
			}
		}
	}
	return
}

// list retrieves a list of channel identifiers from the server.
func list() (l []string, err error) {
	u := &lib.ChanList{}
	u, err = client.List(ctx, &lib.Void{})
	if (err == nil) && len(u.Chans) != 0 {
		for _, v := range u.Chans {
			l = append(l, v.Id)
		}
	}
	return
}

// check_id checks if a channel, 'id', is found in a list of identifiers, 'l'.
func check_id(l []string, id string) bool {
	if len(l) != 0 {
		for _, v := range l {
			if v == id {
				return true
			}
		}
	}
	return false
}

// read pops a value from channel 'id'.
// Returns the value and an error. If the channel is empty, the content of the error is "empty".
func read(id string) (v int32, err error) {
	d := &lib.Value{}
	d, err = client.Rd(ctx, &lib.Id{Id: id})
	if err != nil {
		if m, _ := regexp.Match("desc = empty", []byte(err.Error())); m {
			err = errors.New("empty")
		}
		return
	}
	v = d.Val
	return
}

// read_blocking tries to pop a value from channel 'id' with intervals of 't' seconds, until a successful read is achieved.
// Returns the value and any error which is not "empty".
func read_blocking(id string, t int) (v int32, err error) {
	for {
		v, err = read(id)
		if err == nil {
			break
		}
		if err.Error() == "empty" {
			log.Println("Empty stream. Retry in", t, "seconds...")
			time.Sleep(time.Duration(t) * time.Second)
		} else {
			return
		}
	}
	return
}

// write pushes a value, 'v', to channel 'id'.
// Returns an error. If the channel is full, the content of the error is 'full'.
func write(id string, v int32) (err error) {
	_, err = client.Wr(ctx, &lib.Write{
		Id:  id,
		Val: v,
	})
	if err != nil {
		if m, _ := regexp.Match("desc = full", []byte(err.Error())); m {
			err = errors.New("full")
		}
	}
	return
}

// write_blocking tries to push a value 'v' to channel 'id' with intervals of 't' seconds, until a successful write is achieved.
// Returns any wrror which is not 'full'.
func write_blocking(id string, v int32, t int) (err error) {
	for {
		err = write(id, v)
		if err == nil {
			break
		}
		if err.Error() == "full" {
			log.Println("Full stream. Retry in", t, "seconds...")
			time.Sleep(time.Duration(t) * time.Second)
		} else {
			return
		}
	}
	return
}
