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

const (
	MapServer = "35.221.65.163"

	MINE Type = "MINE"
)

func RunServer() {
	conn, err := grpc.Dial(MapServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer conn.Close()
	go UpdateMap(conn)

	http.HandleFunc("/add_node", func(w http.ResponseWriter, r *http.Request) {
		node := pb.Node{}
		data, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}

		err = json.Unmarshal(data, &node)
		if err != nil {
			log.Println(err)
			return
		}

		err = createInstance(node.Row, node.Col, node.Type)
		resp := ""
		if err != nil {
			log.Println(err)
			resp = err.Error()
		} else {
			resp = "OK"
		}
		_, err = w.Write([]byte(resp))
		if err != nil {
			log.Println(err)
		}
	})

	http.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
