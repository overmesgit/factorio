package localmap

import (
	"encoding/json"
	pb "github.com/overmesgit/factorio/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"net/http"
)

type Type string
type Direction string

const (
	MapServer = "34.84.65.135"

	Up    Direction = "A"
	Down  Direction = "V"
	Left  Direction = "<"
	Right Direction = ">"

	IronMine Type = "MI"
	Belt     Type = "BE"
)

type ItemType string

const (
	Iron ItemType = "IRON"
	Coal ItemType = "COAL"
)

func RunServer() {
	conn, err := grpc.Dial(MapServer+":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer conn.Close()
	go UpdateMap(conn)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var nodes []*pb.Node
		data, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}

		err = json.Unmarshal(data, &nodes)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Got data", nodes)

		syncInstances(nodes)

		resp, err := json.Marshal(DateFromMap)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = w.Write(resp)
		if err != nil {
			log.Println(err)
		}
	})

	http.HandleFunc("/map", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for index.")
		http.ServeFile(w, r, "localmap/index.html")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
