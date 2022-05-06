package nodemap

import (
	"encoding/json"
	"github.com/overmesgit/factorio/grpc"
	pb "github.com/overmesgit/factorio/grpc"

	"html/template"
	"io"
	"log"
	"net/http"
)

func RunHttpServer() {
	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		var nodes []*grpc.Node
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

		for _, node := range nodes {
			key := Key{
				row: node.Row,
				col: node.Col,
			}
			nodeMap.nodes[key] = node
		}

		log.Println("Got data", nodes)
		_, err = w.Write([]byte("{}"))
		if err != nil {
			log.Println(err)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		indexTemplate, parseError := template.ParseFiles("nodemap/index.html")
		if parseError != nil {
			log.Println("error reading file", parseError)
			return
		}

		type vars struct {
			Data string
		}

		var arr []*pb.Node
		for _, node := range nodeMap.nodes {
			arr = append(arr, node)
		}

		data, err := json.Marshal(arr)
		if err != nil {
			log.Println("marshal error", data)
			return
		}

		if err := indexTemplate.Execute(w, vars{string(data)}); err != nil {
			log.Println("indexTemplate error", err)
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
