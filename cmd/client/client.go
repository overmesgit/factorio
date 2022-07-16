package main

import (
	"context"
	pb "github.com/overmesgit/factorio/grpc"
	"github.com/overmesgit/factorio/nodemap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial(
		nodemap.MapServer+":8080", grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMapClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UpdateMap(
		ctx, &pb.MapRequest{Nodes: []*pb.Node{
			{
				Type:      "IronMine",
				Col:       0,
				Row:       0,
				Ip:        "",
				Items:     nil,
				Direction: "↓",
			},
		}},
	)
	if err != nil {
		log.Fatalf("could not update: %v", err)
	}
	log.Printf("Resp: %s", r.GetNodes())
}
