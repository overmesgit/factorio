package mine

import (
	"context"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/network"
	pb "github.com/overmesgit/factorio/grpc"
	"github.com/overmesgit/factorio/mine/sugar"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"runtime"
	"time"
)

func RunMapper() {
	go func() {
		url := "map:8080"
		if os.Getenv("local") != "" {
			url = "host.minikube.internal:8080"
		}
		sugar.Sugar.Infof("Map url %v", url)
		conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			sugar.Sugar.Errorw("failed to connect: %v", err)
			return
		}
		defer conn.Close()

		c := pb.NewMapClient(conn)
		for {
			time.Sleep(300 * time.Millisecond)
			UpdateMapState(c)
		}
	}()

	go func() {
		for {
			err := UpdateStats()
			if err != nil {
				sugar.Sugar.Errorw("failed to update stats: %v", err)
			}
		}
	}()
}

var nodeStats pb.Stats

func UpdateStats() error {
	before, err := cpu.Get()
	if err != nil {
		return err
	}
	beforeNet, err := network.Get()
	if err != nil {
		return err
	}

	time.Sleep(time.Duration(1) * time.Second)

	after, err := cpu.Get()
	if err != nil {
		return err
	}
	afterNet, err := network.Get()
	if err != nil {
		return err
	}

	total := float32(after.Total - before.Total)
	nodeStats.CpuLoad = float32(after.User-before.User) / total * 100

	var totalRx, totalTx uint64
	for i := range afterNet {
		totalRx = afterNet[i].RxBytes - beforeNet[i].RxBytes
		totalTx = afterNet[i].TxBytes - beforeNet[i].TxBytes
	}
	nodeStats.NetworkRx = int32(totalRx)
	nodeStats.NetworkTx = int32(totalTx)

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	nodeStats.MemoryUsage = int32(m.Alloc)

	return nil

}

func UpdateMapState(client pb.MapClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	counter := MyWorker.GetItemCount()
	itemsCounter := make([]*pb.ItemCounter, 0, len(counter))
	for _, c := range counter {
		itemsCounter = append(
			itemsCounter, &pb.ItemCounter{
				Type:  c.Type,
				Count: c.Count,
			},
		)
	}
	state := pb.NodeState{
		Node: &pb.Node{
			Col:       MyNode.Col,
			Row:       MyNode.Row,
			Type:      string(MyNode.NodeType),
			Direction: string(MyNode.Direction),
		},
		Items:     itemsCounter,
		NodeStats: &nodeStats,
	}
	sugar.Sugar.Infof("Update state with: %v\n", &state)
	r, err := client.UpdateNodeState(
		ctx, &state,
	)
	if err != nil {
		sugar.Sugar.Errorf("Could not update my status: %v\n", err)
		return
	}
	sugar.Sugar.Infof("Response: %v\n", r)
}
