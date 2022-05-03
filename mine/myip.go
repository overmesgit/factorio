package mine

import (
	"context"
	pb "github.com/overmesgit/factorio/grpc"
	"github.com/overmesgit/factorio/localmap"
	"google.golang.org/grpc"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func RunMapper(conn *grpc.ClientConn) {
	myIp, err := GetIp()
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	name, err := GetHostName()
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		for {
			time.Sleep(time.Second)
			RegisterInMapServer(conn, myIp, name)
		}
	}()
}

func GetIp() (string, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "http://metadata/computeMetadata/v1/instance/network-interfaces/0/access-configs/0/external-ip", nil)
	if err != nil {
		return "", err
	}
	request.Header.Set("Metadata-Flavor", "Google")
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(ip), err
}

func GetHostName() (string, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "http://metadata/computeMetadata/v1/instance/hostname", nil)
	if err != nil {
		return "", err
	}
	request.Header.Set("Metadata-Flavor", "Google")
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	//r1c1.asia-northeast1-a.c.factorio2022.internal
	split := strings.Split(string(ip), ".")
	return split[0], err
}

func GetRowCol(name string) (int32, int32, error) {
	// r0c0
	res := strings.Split(name[1:], "c")
	row, err := strconv.Atoi(res[0])
	if err != nil {
		return 0, 0, err
	}

	col, err := strconv.Atoi(res[1])
	if err != nil {
		return 0, 0, err
	}
	return int32(row), int32(col), nil
}

func RegisterInMapServer(conn *grpc.ClientConn, ip string, name string) {

	row, col, err := GetRowCol(name)
	if err != nil {
		log.Printf("could not update ip: %v\n", err)
		return
	}

	c := pb.NewMapClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.NotifyIp(ctx, &pb.IpRequest{
		Ip:    ip,
		Col:   col,
		Row:   row,
		Items: nil,
	})
	if err != nil {
		log.Printf("could not update ip: %v\n", err)
		return
	}
	log.Printf("Response: %s\n", r.String())

	AdjustedNodes := r.GetAdjustedNodes()
	for _, node := range AdjustedNodes {
		if node.Row == row && node.Col == col {
			MyType = localmap.Type(node.Type)
			log.Printf("Set my type: %s\n", MyType)
		}
	}
}
