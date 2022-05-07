package nodemap

import (
	"encoding/json"
	pb "github.com/overmesgit/factorio/grpc"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
)

func RunHttpServer() {
	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		var nodes []*pb.Node
		data, err := io.ReadAll(r.Body)
		if err != nil {
			sugar.Error(err)
			return
		}
		err = json.Unmarshal(data, &nodes)
		if err != nil {
			sugar.Error(err)
			return
		}

		updatedNodes(nodes)

		sugar.Infof("Got data %v", nodes)
		_, err = w.Write([]byte("{}"))
		if err != nil {
			sugar.Error(err)
		}

		file, err := json.Marshal(nodes)
		if err != nil {
			sugar.Error(err)
			return
		}

		err = ioutil.WriteFile("/mnt/data/db.json", file, 0644)
		if err != nil {
			sugar.Error(err)
		}

	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		indexTemplate, parseError := template.ParseFiles("nodemap/index.html")
		if parseError != nil {
			sugar.Errorw("error reading file", parseError)
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
			sugar.Errorw("marshal error", data)
			return
		}

		if err := indexTemplate.Execute(w, vars{string(data)}); err != nil {
			sugar.Errorw("indexTemplate error", err)
		}
	})

	if err := http.ListenAndServe(":8081", nil); err != nil {
		sugar.Error(err)
	}
}

func updatedNodes(nodes []*pb.Node) {
	seenKey := make(map[Key]struct{})
	for _, node := range nodes {
		key := Key{
			row: node.Row,
			col: node.Col,
		}

		if exist, ok := nodeMap.nodes[key]; !ok {
			createPod(node)
		} else {
			if exist.Type != node.Type || exist.Direction != node.Direction {
				go recreatePod(node)
			}
		}

		seenKey[key] = struct{}{}
		nodeMap.nodes[key] = node
	}

	for k := range nodeMap.nodes {
		if _, ok := seenKey[k]; !ok {
			delete(nodeMap.nodes, k)
			go deletePod(k.row, k.col)
		}
	}
}
