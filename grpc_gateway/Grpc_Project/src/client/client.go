package main

import "context"

func main() {
	// 定义一个连接
	//conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure(), grpc.WithBlock())
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	// 创建路由
	//mux := runtime.NewServeMux()

	//注册服务服务端的方法
	//service.Regis
	//
	//defer conn.Close()
	//client := person.NewSearchServiceClient(conn)
	//io, err := client.Search(context.Background(), &person.Person{Name: "zhangSan"})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(io)

	ctx := context.Background()
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()



}
