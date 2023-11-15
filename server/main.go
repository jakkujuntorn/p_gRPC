package main

import (
	"fmt"
	"log"
	"net"
	"server/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// option เอาไว้ทำอะไร ทำตาม vdo
	option := []grpc.ServerOption{}
	// ฝั่ง server จะใช้ NewServer
	s := grpc.NewServer(option...)

	//สร้าง net.Listen ก่อน ทำแค่ฝั่ง server
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	// มาจาก gRPC ที่ build  มา
	// เอามาบริการงาน gRPC ที่ทำไว้
	// รับ พารามิเตอร์ 2 ตัว ไปดูว่ามาจากอะไรบ้าง *******
	// ตรงนี้สำคัญต้องใส เพื่อมาบริการงาน grpc มันจะลิ้งเข้าไปหา fun ที่ทำงานจากตรงนี้ ******
	// ชื่อมันอิงมาจาก proto
	// parameter  s = grpc, services.NewCalculatorServer() = func ที่ส้รางเองในไฟล์ service ที่ conform ตาม interface CalculatorServer ในไฟล์ proto ที่  build
	services.RegisterCalculatorServer(s, services.NewServer_Service())

	// ทำให้ evans เห็น service ทั้งหมด
	// ต้องการ reflection.GRPCServer
	reflection.Register(s)

	fmt.Println("gRPC Start Server .........")

	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

	// stop gRPC
	// s.Stop()
}
