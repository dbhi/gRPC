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
	"context"
	"errors"
	"log"
	"net"

	"github.com/dbhi/gRPC/lib"
	grpc "google.golang.org/grpc"
)

func main() {
	log.Println("DBHI gRPC server")
	srv := grpc.NewServer()
	var chans chanServer
	chans.db = make(map[string]*stream)
	lib.RegisterChansServer(srv, chans)
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("could not listen to :8888: %v", err)
	}
	log.Fatal(srv.Serve(l))
}

type chanServer struct {
	db map[string]*stream
}

type message struct {
	adr int32
	dat int32
}

type stream chan message

func (s *stream) read() (msg message, err error) {
	valid := false
	select {
	case msg, valid = <-*s:
		if !valid {
			err = errors.New("closed")
		}
	default:
		err = errors.New("empty")
	}
	return
}

func (s *stream) write(msg message) (err error) {
	select {
	case *s <- msg:
	default:
		err = errors.New("full")
	}
	return
}

// API

func (c chanServer) List(ctx context.Context, void *lib.Void) (*lib.ChanList, error) {
	var list lib.ChanList
	for k := range c.db {
		list.Ids = append(list.Ids, &lib.Id{Id: k})
	}
	return &list, nil
}

func (c chanServer) Reg(ctx context.Context, args *lib.Register) (*lib.Void, error) {
	log.Println("REG", args)
	v := &lib.Void{}
	if _, ok := c.db[args.Id]; ok {
		return v, errors.New("chan exists")
	}
	s := make(stream, args.Length)
	c.db[args.Id] = &s
	log.Println("Registered channels:")
	for k := range c.db {
		log.Println("-", k)
	}
	return v, nil
}

func (c chanServer) Wr(ctx context.Context, args *lib.Write) (*lib.Void, error) {
	v := &lib.Void{}
	if s, ok := c.db[args.Id]; ok {
		return v, s.write(message{adr: args.Msg.Adr, dat: args.Msg.Dat})
	}
	return v, errors.New("chan does not exist")
}

func (c chanServer) Rd(ctx context.Context, args *lib.Id) (*lib.Message, error) {
	if s, ok := c.db[args.Id]; ok {
		item, err := s.read()
		return &lib.Message{Adr: item.adr, Dat: item.dat}, err
	}
	return &lib.Message{}, errors.New("chan does not exist")
}
