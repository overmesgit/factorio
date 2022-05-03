package nodemap

import (
	"errors"
	"fmt"
	"github.com/overmesgit/factorio/grpc"
	pb "github.com/overmesgit/factorio/grpc"
	"log"
	"sync"
)

type Key struct {
	row, col int32
}

type Map struct {
	nodes map[Key]*pb.Node
	sync.Mutex
}

var nodeMap = Map{}

func RegisterServer(in *grpc.IpRequest) error {
	nodeMap.Lock()
	defer nodeMap.Unlock()
	k := Key{
		row: in.Row,
		col: in.Col,
	}
	val, ok := nodeMap.nodes[k]
	if !ok {
		err := errors.New(fmt.Sprintf("unknown node %v", in))
		log.Println(err)
		return err
	}
	val.Ip = in.GetIp()
	val.Items = in.Items
	return nil
}

func UpdateMap(in *pb.MapRequest) []*pb.Node {
	nodeMap.Lock()
	defer nodeMap.Unlock()
	for _, node := range in.GetNodes() {
		k := Key{
			row: node.GetRow(),
			col: node.GetCol(),
		}
		if val, ok := nodeMap.nodes[k]; ok {
			val.Type = node.Type
		} else {
			nodeMap.nodes[k] = node
		}
	}
	log.Printf("UpdatedNodes: %v", nodeMap.nodes)

	nodesList := make([]*pb.Node, 0)
	for _, node := range nodeMap.nodes {
		nodesList = append(nodesList, node)
	}
	return nodesList
}
