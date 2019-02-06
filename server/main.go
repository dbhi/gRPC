package main

import (
	"errors"
	"log"
	"net"

	"github.com/umarcor/dbhi/gRPC/lib"
	"golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

func init() {
	db = make(map[string]*stream)
}

func main() {
	srv := grpc.NewServer()
	var chans chanServer
	lib.RegisterChansServer(srv, chans)
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("could not listen to :8888: %v", err)
	}
	log.Fatal(srv.Serve(l))
}

type chanServer struct{}

var chans lib.ChanList

var db map[string]*stream

type stream chan int32

func (s *stream) read() (i int32, err error) {
	valid := false
	select {
	case i, valid = <-*s:
		if !valid {
			err = errors.New("closed")
		}
	default:
		err = errors.New("empty")
	}
	return
}

func (s *stream) write(i int32) (err error) {
	select {
	case *s <- i:
	default:
		err = errors.New("full")
	}
	return
}

func (chanServer) List(ctx context.Context, void *lib.Void) (*lib.ChanList, error) {
	return &chans, nil
}

func (chanServer) Wr(ctx context.Context, args *lib.Write) (v *lib.Void, err error) {
	v = &lib.Void{}
	id := args.Id
	if s, ok := db[id]; ok {
		err = s.write(args.Val)
	}
	return v, errors.New("chan does not exist")
}

func (chanServer) Rd(ctx context.Context, i *lib.Id) (v *lib.Value, err error) {
	v = &lib.Value{}
	id := i.Id
	if s, ok := db[id]; ok {
		v.Val, err = s.read()
		return
	}
	return v, errors.New("chan does not exist")
}

func (chanServer) Reg(ctx context.Context, r *lib.Register) (v *lib.Void, err error) {
	log.Println("REG", r)
	v = &lib.Void{}
	id := r.Id
	if _, ok := db[id]; !ok {
		s := make(stream, r.Length)
		db[id] = &s
		chans.Chans = append(chans.Chans, &lib.Id{Id: id})
		return
	}
	return v, errors.New("chan exists")
}
