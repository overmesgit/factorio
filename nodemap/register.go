package nodemap

import (
	"errors"
	"fmt"
	"github.com/overmesgit/factorio/grpc"
	pb "github.com/overmesgit/factorio/grpc"
	"sync"
)

type Key struct {
	row, col int32
}

type Map struct {
	nodes map[Key]*pb.Node
	sync.Mutex
}

var nodeMap = Map{nodes: make(map[Key]*pb.Node, 0)}

func (s *server) RegisterServer(in *grpc.IpRequest) error {
	nodeMap.Lock()
	defer nodeMap.Unlock()
	k := Key{
		row: in.Row,
		col: in.Col,
	}
	node, ok := nodeMap.nodes[k]
	if !ok {
		err := errors.New(fmt.Sprintf("Trying to map unregistered node %v", in))
		s.logger.Errorw(err.Error())
		return err
	}

	s.logger.Infof("Registering node %v %v", in.GetIp(), in.GetItems())
	node.Ip = in.GetIp()
	node.Items = in.GetItems()
	return nil
}

func (s *server) GetAdjustedNodes(in *pb.IpRequest) []*pb.Node {
	currentKey := Key{
		row: in.Row,
		col: in.Col,
	}
	resp := []*pb.Node{nodeMap.nodes[currentKey]}
	for _, offset := range [][]int32{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	} {
		k := Key{
			row: in.Row + offset[0],
			col: in.Col + offset[1],
		}
		val, ok := nodeMap.nodes[k]
		if ok {
			resp = append(resp, val)
		}
	}
	return resp
}

func (s *server) RunUpdateMap(in *pb.MapRequest) []*pb.Node {
	nodeMap.Lock()
	defer nodeMap.Unlock()
	for _, node := range in.GetNodes() {
		k := Key{
			row: node.GetRow(),
			col: node.GetCol(),
		}
		if val, ok := nodeMap.nodes[k]; ok {
			val.Type = node.Type
			val.Direction = node.Direction
		} else {
			nodeMap.nodes[k] = node
		}
	}
	s.logger.Infof("UpdatedNodes: %v", nodeMap.nodes)

	nodesList := make([]*pb.Node, 0)
	for _, node := range nodeMap.nodes {
		nodesList = append(nodesList, node)
	}
	return nodesList
}
