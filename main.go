package main

import (
	"grpc-client/handler"
	"grpc-client/pb"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func main() {
	opts := grpc.WithInsecure()
	grpcUser, err := grpc.Dial("localhost:8080", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer grpcUser.Close()

	client := pb.NewAttendanceServiceClient(grpcUser)
	// request := &pb.CheckInRequest{Username: "Bejo", Datetime: "2024-03-12 07:00"}
	// now := time.Now().Format("2006-01-02 15:04")
	// request := &pb.CheckInRequest{Username: "Bejo", Datetime: now}

	// resp, err := client.CheckIn(context.Background(), request)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Receive response => %s ", resp)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	handler.InitHttpHandler(e, client)
	e.Logger.Fatal(e.Start(":9090"))

}
