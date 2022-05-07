package nodemap

import (
	"context"
	"fmt"
	pb "github.com/overmesgit/factorio/grpc"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
)

var clientset *kubernetes.Clientset

func init() {
	var kubeconfig string
	var config *rest.Config
	var err error

	if os.Getenv("KUBERNETES_SERVICE_HOST") != "" {
		config, err = rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
	} else {
		home := homedir.HomeDir()
		kubeconfig = filepath.Join(home, ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			sugar.Error(err)
			return
		}
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		sugar.Error(err)
		return
	}
}

func createPod(node *pb.Node) {
	createDeployment(node)
	createService(node)
}

func createDeployment(node *pb.Node) {
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	name := getName(node.Row, node.Col)
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: appsv1.DeploymentSpec{
			Strategy: appsv1.DeploymentStrategy{Type: "Recreate"},
			Replicas: Ptr(int32(1)),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": name,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": name,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  name,
							Image: "gcr.io/factorio2022-349510/mine:latest",
							Env: []apiv1.EnvVar{
								{Name: "ROW", Value: fmt.Sprint(node.Row)},
								{Name: "COL", Value: fmt.Sprint(node.Col)},
								{Name: "TYPE", Value: fmt.Sprint(node.Type)},
								{Name: "DIRECTION", Value: fmt.Sprint(node.Direction)}},
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 8080,
								},
							},
							Resources: apiv1.ResourceRequirements{
								Limits: apiv1.ResourceList{
									"memory":            resource.MustParse("64Mi"),
									"cpu":               resource.MustParse("250m"),
									"ephemeral-storage": resource.MustParse("100Mi"),
								},
								Requests: apiv1.ResourceList{
									"memory":            resource.MustParse("128Mi"),
									"cpu":               resource.MustParse("250m"),
									"ephemeral-storage": resource.MustParse("100Mi"),
								},
							},
						},
					},
				},
			},
		},
	}

	sugar.Infof("Creating deployment %v", name)
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		sugar.Error(err)
	}
	sugar.Infof("Node was created %v %v", name, result)
}

func createService(node *pb.Node) {
	name := getName(node.Row, node.Col)
	_, err := clientset.CoreV1().Services(apiv1.NamespaceDefault).Create(context.TODO(), &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: apiv1.ServiceSpec{
			Selector: map[string]string{
				"app": name,
			},
			Ports: []apiv1.ServicePort{
				{
					Protocol: "TCP",
					Port:     8080,
				},
			},
		},
	}, metav1.CreateOptions{})
	if err != nil {
		sugar.Errorf("could not create service %v", err)
	}
}

func getName(row int32, col int32) string {
	return fmt.Sprintf("r%vc%v", row, col)
}

func deletePod(row, col int32) {
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	name := getName(row, col)
	sugar.Infof("Node was created %v", name)

	err := deploymentsClient.Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		sugar.Errorf("could not delete pod %v", err)
	}
}

func recreatePod(node *pb.Node) {
	deletePod(node.Row, node.Col)
	createPod(node)
}

func Ptr[T any](v T) *T {
	return &v
}
