package localmap

import (
	"context"
	"fmt"
	pb "github.com/overmesgit/factorio/grpc"
	"google.golang.org/grpc"
	"log"
	"os/exec"
	"time"
)

var Nodes []*pb.Node

func createInstance(row, col int32, nodeType string) error {
	nodeName := fmt.Sprintf("r%vc%v", row, col)
	command := fmt.Sprintf("gcloud compute instances create %v --image-family cos-stable --image-project cos-cloud --metadata-from-file user-data=infra/init --metadata=cos-metrics-enabled=true --zone=asia-northeast1-a --machine-type=e2-micro --project=factorio2022", nodeName)
	log.Println(command)
	runner := exec.Command("bash", "-c", command)
	cmd, err := runner.Output()
	log.Println("cmd ========>", string(cmd))
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			log.Println(string(ee.Stderr))
		}
	}

	Nodes = append(Nodes, &pb.Node{
		Type: nodeType,
		Col:  col,
		Row:  row,
	})

	return nil
}

func UpdateMap(conn *grpc.ClientConn) {
	for {
		time.Sleep(time.Second)
		Update(conn)
	}
}

func Update(conn *grpc.ClientConn) bool {
	c := pb.NewMapClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UpdateMap(ctx, &pb.MapRequest{
		Nodes: Nodes,
	})
	if err != nil {
		log.Printf("could not update nodes: %v\n", err)
		return true
	}
	log.Printf("response: %s\n", r.String())
	return false
}
