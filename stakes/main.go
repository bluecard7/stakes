package main

import (
	"stakes/handler"
	pb "stakes/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("stakes"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterStakesHandler(srv.Server(), new(handler.Stakes))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
