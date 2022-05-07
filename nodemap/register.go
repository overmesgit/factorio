package nodemap

import (
	"encoding/json"
	pb "github.com/overmesgit/factorio/grpc"
	"os"
	"sync"
	"time"
)

type Key struct {
	row, col int32
}

type Map struct {
	nodes map[Key]*pb.Node
	sync.Mutex
}

type MapItems struct {
	nodes      map[Key]*pb.NodeState
	lastUpdate map[Key]int64
	sync.Mutex
}

var mapNodes = Map{nodes: make(map[Key]*pb.Node, 0)}
var mapItems = MapItems{nodes: make(map[Key]*pb.NodeState, 0), lastUpdate: make(map[Key]int64, 0)}

func init() {
	data, err := os.ReadFile("/mnt/data/db.json")
	if err != nil {
		sugar.Error(err)
		return
	}

	var nodes []*pb.Node
	err = json.Unmarshal(data, &nodes)
	if err != nil {
		sugar.Error(err)
		return
	}
	sugar.Infof("Loaded nodes from db %v", nodes)
	updatedNodes(nodes)

}

func CleanItems() {
	for {
		time.Sleep(time.Second)
		DoCleanItems()
	}
}

func DoCleanItems() {
	mapItems.Lock()
	defer mapItems.Unlock()

	var toDelete []Key
	now := time.Now().Unix()
	for key, val := range mapItems.lastUpdate {
		if now-val > 5 {
			toDelete = append(toDelete, key)
		}
	}

	sugar.Infof("Delete outdated item statuses: %v", toDelete)
	for _, key := range toDelete {
		delete(mapItems.nodes, key)
		delete(mapItems.lastUpdate, key)
	}

}

//
//func (s *server) GetAdjustedNodes(in *pb.IpRequest) []*pb.Node {
//	currentKey := Key{
//		row: in.Row,
//		col: in.Col,
//	}
//	resp := []*pb.Node{mapNodes.nodes[currentKey]}
//	for _, offset := range [][]int32{
//		{1, 0},
//		{-1, 0},
//		{0, 1},
//		{0, -1},
//	} {
//		k := Key{
//			row: in.Row + offset[0],
//			col: in.Col + offset[1],
//		}
//		val, ok := mapNodes.nodes[k]
//		if ok {
//			resp = append(resp, val)
//		}
//	}
//	return resp
//}
