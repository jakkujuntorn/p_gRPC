package services

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ไม่ต้องใช้ interface CalculatorClient ใน grpc.go เพราะ layer นี้แค่แสดงผล
// เลยไม่ต้อง return ค่าตาม interface ที่ทาง proto ให้มา
type I_CalculatorService interface {
	Hello_Client(name string, time_out time.Duration) error
	Fibonacci_Client(n uint32) error

	//  ที่ใช้  float64 เพราะ ทางฝั่ง proto ใช้ double
	Average_Client(number ...float64) error

	Sum_Client(number ...int32) error
}

type calculatorService struct {
	// เรียกใช้ interface CalculatorClient ใน _grpc.go มาช้วยเหลือ
	// หลักการ port anf adapter
	// ใน CalculatorClient มี servie ต่างๆที่ build มาจาก grpc
	calculatorClient CalculatorClient
}

// รับ CalculatorClient เข้ามาเพราะต้องเอาใช้
// ถ้าไม่รับเข้ามาจะใช้ CalculatorClient  ที่ struct จะใช้ได้ไหม
func NewClient_Service(calcultor_Client CalculatorClient) I_CalculatorService {
	return calculatorService{calculatorClient: calcultor_Client}
}

// Sum
func (base calculatorService) Sum_Client(n1 ...int32) error {

	// มาจาก grpc ที่ build ขึ้นมา
	stream, err := base.calculatorClient.Sum(context.Background())
	if err != nil {
		return err
	}

	// request จะทำหน้าที่ ส่งอย่างเดียว
	// ส่ง request stream ไปหา server
	// send *************
	go func() {
		// ใช้ for ไม่รู้จบ เพราะ เราไม่รู้ว่าจะจบเมื่อไหร
		for _, n := range n1 {
			// ปั้น request
			req := SumRequest{
				Number: n,
			}
			time.Sleep(time.Second)
			// ส่ง  stream ไปหา server
			stream.Send(&req)
			fmt.Println("Request:", req.Number)
		}
		//จบการส่ง
		stream.CloseSend()
	}()

	// response จะทำหน้าที่ รับอย่างเดียว
	// รับ response stream จาก server
	done := make(chan bool)
	errs := make(chan error)

	// ทำตาม vdo เอาไว้ทำงานอะไร
	wait := make(chan struct{})

	// recive ************
	go func() {
		// ใช้ for ไม่รู้จบเพราะ เราไม่รู็ว่า จะส่งเลขไป บวก กี่ตัว
		for {
			//รับ stream จาก server
			res, err := stream.Recv()

			//เช็คการรับว่าจบหรือยัง
			// ถ้าจบ จะมาที่ done <- true
			if err == io.EOF {
				break
			}
			if err != nil {
				errs <- err
				// ดูใน vdo มันคืออะไร *****
				// close(<-wait)
			}

			// ทำงานยังไง ดูจาก vdo
			close(wait)

			// แสดงผล
			fmt.Println("Response:", res.Result)
			fmt.Println("====================")
		}

		// block unit everything is done
		<-wait

		// ถ้าจบแบบสมบูรณ์ done จะมีค่า
		done <- true

	}()

	// select เอามาดักตอนจบว่าจะมีเหตุการณ์อะไรเกิดขึ้น
	select {
	case <-done:
		return nil
	case err := <-errs:
		return err
	}

}

// Aerage
func (base calculatorService) Average_Client(number ...float64) error {

	// เอาค่า stream  จาก Func Average จาก proto ใน _grpc.go
	// เราจะใส time out ก็ได้
	stream, err := base.calculatorClient.Average(context.Background())
	if err != nil {
		return err
	}

	// number เป็น array เลยตัองวนเพื่อส่งค่าที่ละค่า
	for _, n := range number {
		// client ปั้น Request เพื่อส่งไปหา Server ******
		req := AverageRequest{
			Number: n,
		}

		// ใช้ stream ส่งข้อมูลไปหา server
		stream.Send(&req)
		time.Sleep(time.Second)
	}

	// ใช้ stream รับข้อมูลจาก Server แล้วปิด *******
	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	fmt.Println("=================")
	fmt.Println("Input :", number)
	fmt.Println("Len :", len(number))
	fmt.Println("Average : ", res.Result)
	fmt.Println("=================")

	return nil
}

// ทำ reciver Func ให้ตรงกับ I_CalculatorService
func (base calculatorService) Fibonacci_Client(n uint32) error {
	req := FibonacciRequest{
		N: n,
	}

	// ตั้งเวลา time out ได้ ให้ตัด connection  กรณีไม่สามารถติดต่อ service
	// ตัวอย่างนี้ หน่วงเวลาที่ฝั่ง server ไว้ 1 วิ 
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()

	// เรียก คำสั่งไปหา server แล้วรอรับค่ากลับ *****
	stream, err := base.calculatorClient.Fibonacci(ctx, &req)
	if err != nil {
		return err
	}

	// ตอน reponse มันไม่รู้จบ เลยต้อง for
	// fmt.Println("Service : Fibonacci")
	// fmt.Printf("Request : %v\n", req.N)
	// fmt.Printf("Request : %v\n", n)
	// fmt.Println("==================")

	// loop ไม่รู้จบ เพราะ client ไม่รู้ว่า server จะส่งกลับมากี่จำนวน *******
	for {
		// การรับค่า จาก Server
		res, err := stream.Recv()

		// ถ้า stream หมดแล้ว err จะได้ค่า io.EOF
		if err == io.EOF {
			fmt.Println("End of Stream :", err)
			break
		}
		if err != nil {
			return err
		}
		// ผลลัพท์ของ stream **
		fmt.Println("Response:", res.Result)
	}

	return nil

}

func (base calculatorService) Hello_Client(name string, time_out time.Duration) error {
	//ปั้น request เตรียมส่งไปหา server
	req := HelloRequest{
		Name:        name, // string
		CreatedDate: timestamppb.Now(), // time
	}

	// ทำ Time out ด้วย context
	ctx, cancel := context.WithTimeout(context.Background(), time_out)
	defer cancel()

	//****************** ส่งไปให้ server **************
	// มี func hello ให้ใช้ มาจากการ build proto
	res, err := base.calculatorClient.Hello(ctx, &req)

	if err != nil {

		resError, ok := status.FromError(err)
		if ok {
			if resError.Code() == codes.InvalidArgument {
				return errors.New(resError.Message())
			} else {
				return errors.New(resError.Message())
			}
		}
		return err

	}

	fmt.Println("Service Hello")
	fmt.Printf("Request : %v\n", req.Name) // HelloRequest
	fmt.Printf("Response : %v\n", res.Result) // HellpResponse

	return nil
}
