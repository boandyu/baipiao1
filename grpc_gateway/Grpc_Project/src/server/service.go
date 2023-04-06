package main

import (
	"GrpcDemo/pb/person"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"sync"
)

// 取出server
type server struct {
	person.UnimplementedSearchServiceServer
}

func (s *server) Search(ctx context.Context, p *person.Person) (*person.Person, error) {
	name := p.Name
	age := p.Age
	stu := p.Stu
	res := &person.Person{Name: "接受信息：" + name, Age: age, Stu: stu}
	return res, nil
}

// 注册服务
func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go registerGateWay(&wg)
	go registerGrpc(&wg)
	wg.Wait()

}
func registerGateWay(wg *sync.WaitGroup) {
	conn, _ := grpc.DialContext(context.Background(), "127.0.0.1:8888", grpc.WithInsecure(), grpc.WithBlock())
	mux := runtime.NewServeMux() // 一个对外开放的mux
	s := &http.Server{
		Handler: mux,
		Addr:    ":8089",
	}
	err := person.RegisterSearchServiceHandler(context.Background(), mux, conn)
	if err != nil {
		fmt.Println(err)
	}
	s.ListenAndServe()
	wg.Done()
}
func registerGrpc(wg *sync.WaitGroup) {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
	}
	newServer := grpc.NewServer()
	person.RegisterSearchServiceServer(newServer, &server{})
	newServer.Serve(listen)
	wg.Done()
}
