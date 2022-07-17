package nodemap

import (
	"encoding/json"
	"fmt"
	pb "github.com/overmesgit/factorio/grpc"
	"html/template"
	"io"
	"io/ioutil"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"time"
)

func RunHttpServer() {
	http.HandleFunc(
		"/update", func(w http.ResponseWriter, r *http.Request) {
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
			sugar.Infof("Got data %v", nodes)
			updatedNodes(nodes)

			var reply []*pb.NodeState
			for _, n := range mapItems.nodes {
				reply = append(reply, n)
			}

			mapItems, err := json.Marshal(reply)
			if err != nil {
				sugar.Error(err)
				return
			}
			_, err = w.Write(mapItems)
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

		},
	)

	http.Handle(
		"/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./nodemap/static"))),
	)

	http.HandleFunc(
		"/", func(w http.ResponseWriter, r *http.Request) {
			indexTemplate, parseError := template.ParseFiles("nodemap/index.html")
			if parseError != nil {
				sugar.Errorw("error reading file", parseError)
				return
			}

			type vars struct {
				Data string
			}

			var arr []*pb.Node
			for _, node := range mapNodes.nodes {
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
		},
	)

	http.HandleFunc(
		"/logs/", getPodLogs,
	)

	if err := http.ListenAndServe(":8081", nil); err != nil {
		sugar.Error(err)
	}
}

func getPodLogs(w http.ResponseWriter, r *http.Request) {
	nodeName := r.URL.Query().Get("node")
	if nodeName == "" {
		w.Write([]byte("node not specified"))
		return
	}
	podLogOpts := apiv1.PodLogOptions{
		Follow:    true,
		SinceTime: &metav1.Time{time.Now().Add(-10 * time.Second)},
	}

	label := fmt.Sprintf("app=%v", nodeName)
	nodes, err := clientset.CoreV1().Pods("").List(
		r.Context(), metav1.ListOptions{LabelSelector: label},
	)
	if err != nil {
		w.Write([]byte("node not found" + err.Error()))
		return
	}
	if len(nodes.Items) == 0 {
		w.Write([]byte("node not found"))
		return
	}
	node := nodes.Items[0]
	req := clientset.CoreV1().Pods("default").GetLogs(node.Name, &podLogOpts)
	podLogs, err := req.Stream(r.Context())
	if err != nil {
		w.Write([]byte("error in opening stream" + err.Error()))
		return

	}
	defer podLogs.Close()

	for {
		io.Copy(w, podLogs)
	}
}

func updatedNodes(nodes []*pb.Node) {
	seenKey := make(map[Key]struct{})
	for _, node := range nodes {
		key := Key{
			row: node.Row,
			col: node.Col,
		}

		if exist, ok := mapNodes.nodes[key]; !ok {
			go createPod(node)
		} else {
			if exist.Type != node.Type || exist.Direction != node.Direction {
				go recreatePod(node)
			}
		}

		seenKey[key] = struct{}{}
		mapNodes.nodes[key] = node
	}

	for k := range mapNodes.nodes {
		if _, ok := seenKey[k]; !ok {
			delete(mapNodes.nodes, k)
			go deletePod(k.row, k.col)
		}
	}
}
