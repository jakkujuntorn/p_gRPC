package main

import (
	"client/services"
	_ "context"
	"fmt"
	_ "fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

func main() {

	// สร้าง insecure
	creds := insecure.NewCredentials()

	// ฝั่ง client จะใช้ Dial
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	// NewCalculatorClient มาจาก  build proto
	calculator_Client := services.NewCalculatorClient(cc)

	// **********เอา calculator_Client ไปใช้เลยไม่ได้เหรอ ******
	// ********* เขียนตรงๆก็ใช้ได้ ************
	// req := services.HelloRequest{
	// 	Name:        "Jack Richer",
	// 	CreatedDate: timestamppb.Now(),
	// }
	// res, err := calculator_Client.Hello(context.Background(), &req)
	// if err != nil {
	// 	// return err
	// 	log.Fatal(err)
	// }
	// fmt.Println(res.Result)
	// ********************************************


	//********** NewCalculstorService สร้าง เอง เอามาครอบ grpc เอาไว้อีกที ******
	// port and adaptor
	calculator_Service := services.NewClient_Service(calculator_Client)
	_ = calculator_Service

	// เรียกใช้ func ที่เราสร้างขึ้นมาเอง ใน func พวกนี้จะเข้าไปติดต่อกับ func ของ grpc
	err = calculator_Service.Hello_Client("Russy", 5*time.Second)
	// err = calculator_Service.Fibonacci_Client(uint32(3))
	//  err = calculator_Service.Average_Client(10,10,2,10,5)
	// err = calculator_Service.Sum_Client(10, 10, 2, 10, 5)

	if err != nil {

		// เช็คว่า error มาจาก gRPC รึ ป่าว
		if grpc_Error, ok := status.FromError(err); ok {
			// ถ้ามาจาก grpc error okจะเป็น true
			log.Printf("Code:%v ; Message: %v", grpc_Error.Code(), grpc_Error.Message())
		} else {
			log.Println(err)
		}


		// อีก แบบ ทำเอง *********
		grpc_Err, ok := status.FromError(err)
		if ok {
			if grpc_Err.Code() == codes.InvalidArgument {
				fmt.Println("Error for gRPC")
			}
		}

	}

}

func do_gRPC() {

}
