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

package client

import (
	"context"
	"errors"
	"log"
	"regexp"
	"time"

	"github.com/dbhi/gRPC/lib"
	grpc "google.golang.org/grpc"
)

var client lib.ChansClient
var ctx context.Context

// Connect establishes a connection with the server at 'addr', and gets identifiers for both the client and the context.
// For now, it is insecure (see WithInsecure).
func Connect(addr string) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to backend:", err)
	}
	client = lib.NewChansClient(conn)
	ctx = context.Background()
}

// Register checks if each of the elements in a slice of IDs is found in the list retrieved from the server.
// If it is not present, it is registered (added).
func Register(ids []string) (err error) {
	l, err := list()
	if err != nil {
		return
	}
	for _, id := range ids {
		if !check_id(l, id) {
			_, err = client.Reg(ctx, &lib.Register{Id: id, Length: 1024})
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
	if (err == nil) && len(u.Ids) != 0 {
		for _, v := range u.Ids {
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

// Read pops a message (adr and dat) from channel 'id'.
// Returns the message and any error.
func Read(id string) (int32, int32, error) {
	msg, err := client.Rd(ctx, &lib.Id{Id: id})
	if err == nil {
		return msg.Adr, msg.Dat, nil
	}
	if m, _ := regexp.Match("desc = empty", []byte(err.Error())); m {
		err = errors.New("EMPTY")
	}
	return 0, 0, err
}

// Read_blocking tries to pop a message from channel 'id' with intervals of 't' seconds, until a successful read is achieved.
// Returns the message (adr and dat) and any error which is not "empty".
func Read_blocking(id string, t int) (adr int32, dat int32, err error) {
	for {
		adr, dat, err = Read(id)
		if err == nil || (err.Error() != "EMPTY") {
			return
		}
		log.Println("Empty stream", id, "| Retry in", t, "seconds...")
		time.Sleep(time.Duration(t) * time.Second)
	}
}

// Write pushes a message (adr and dat) to channel 'id'.
// Returns any error.
func Write(id string, adr int32, dat int32) error {
	_, err := client.Wr(ctx, &lib.Write{Id: id, Msg: &lib.Message{Adr: adr, Dat: dat}})
	if err == nil {
		return nil
	}
	if m, _ := regexp.Match("desc = full", []byte(err.Error())); err == nil && m {
		err = errors.New("FULL")
	}
	return err
}

// Write_blocking tries to push a message (adr and dat) to channel 'id' with intervals of 't' seconds, until a successful write is achieved.
// Returns any error which is not 'full'.
func Write_blocking(id string, adr int32, dat int32, t int) (err error) {
	for {
		err = Write(id, adr, dat)
		if err == nil || (err.Error() != "FULL") {
			return
		}
		log.Println("Full stream", id, "| Retry in", t, "seconds...")
		time.Sleep(time.Duration(t) * time.Second)
	}
}
