package services

import (
	context "context"
	"errors"
	"fmt"
	"io"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// conform ตาม interface CalculatorServer ที่มาจากการ build gRPC
type calculator_Server struct {
	// ใน struct ไม่ต้องใช้งานอะไร
}

// CalculatorServer มาจาก _grpc.go    ส่วนของ server
// เพราะเราทำ go ในส่วน server
// อันนี้น่าจะเป็น port
func NewServer_Service() CalculatorServer {
	return calculator_Server{}
}

func (calculator_Server) mustEmbedUnimplementedCalculatorServer() {}

// HelloResponse มาจาก  pb.go
func (calculator_Server) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {

	//***** ทำ delay ให้นานๆ เพื่อเช็คการ error time out **********

	//handler error
	if req.Name == "" {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Name is Request",
		)
	}

	if req.Name == "admin" {
		return nil, status.Errorf(
			codes.Unauthenticated,
			"Name is not Unauthenticated",
		)
	}

	// รับ request *****
	result := fmt.Sprintf("Hello,Hi: %v, At:%v", req.Name, req.CreatedDate.AsTime().Local())

	// ****** ปั้น response ตอบกลับ  *****
	res := HelloResponse{
		Result: result,
	}

	// _=res

	return &res, nil

	// return &HelloResponse{
	// 	Result: result,
	// }, nil
}

// Func นี้เอามาจากตรงไหนในฝั่ง server
// Calculator_FibonacciServer เป็น interface จาก gRPC
func (calculator_Server) Fibonacci(req *FibonacciRequest, streaM Calculator_FibonacciServer) error {

	// req เป็น uint32
	// loop  จนกว่า i <= req.N
	for i := uint32(0); i <= req.N; i++ {

		// ส่งเข้าไปคำนวณ Fib
		result := fib(i)

		res := FibonacciResponse{
			Result: result,
		}

		// ตอนส่ง ต้องใช้ stream.Send() *****
		time.Sleep(time.Second)
		// จะส่งเลขที่ละตัวไปเรื้อยๆ
		streaM.Send(&res)

		// แบบนี้  error เพราะ มันจะจบการทำงานเลย
		// return streaM.Send(&res)
	}

	// type uint32 for range ได้ไหม
	// for n:= range uint32(req.N) {

	// }

	return nil
}

func fib(n uint32) uint32 {

	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fib(n-1) + fib(n-2)
	}

}

func (calculator_Server) Average(stream Calculator_AverageServer) error {
	sum := 0.0
	count := 0.0

	// for ไม่รู้จบเพราะ เราไม่รู้ว่า client จะ stream ถึงเมื่อไหร
	for {
		// รับค่า stream จาก client
		req, err := stream.Recv()

		// client stream หมดแล้ว err == io.EOF
		if err == io.EOF {
			break
		}
		if err != nil {
			// return err
			return errors.New(err.Error())
		}

		// คำนวณผลค่าเฉลี่ย
		sum += req.Number
		count++
	}

	// ปั้น reponse ไปให้ client
	res := AverageResponse{
		Result: sum / count,
	}

	// ส่ง response กลับไปให้ client
	return stream.SendAndClose(&res)
}

func (calculator_Server) Sum(stream Calculator_SumServer) error {

	var sum int32

	// for เพราะเราไม่รู้ว่า client จะส่งถึงเมื่อไหร *****
	for {
		// รับ request จาก Client
		req, err := stream.Recv()
		// เมื่อ client ส่งจบ err = io.EOF
		if err == io.EOF {
			break
		}
		if err != nil {
			// return err
			return status.Error(codes.InvalidArgument, err.Error())
		}
		// logic
		sum += req.Number

		// ปั้น Resoponse ส่งไปหา Client
		// และส่งกลับทันที
		res := SumResponse{
			Result: sum,
		}

		// ส่ง Response ไปหา Client
		err = stream.Send(&res)
		if err != nil {
			return err
		}
	}

	return nil
}
