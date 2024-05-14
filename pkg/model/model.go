package model

import (
	"context"

	pb "github.com/erteldg/grpcnamespaceservice/pkg/proto"
	v1c "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Server struct {
	pb.UnimplementedNamespaceServiceServer
	Clientset *kubernetes.Clientset
}

func (s *Server) CreateNamespace(ctx context.Context, in *pb.NamespaceRequest) (*pb.NamespaceReply, error) {
	namespace, err := s.Clientset.CoreV1().Namespaces().Create(ctx, &v1c.Namespace{
		ObjectMeta: v1.ObjectMeta{
			Name: in.GetName(),
		},
	}, v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return &pb.NamespaceReply{Message: "Namespace " + namespace.Name + " created."}, nil
}

func (s *Server) DeleteNamespace(ctx context.Context, in *pb.NamespaceRequest) (*pb.NamespaceReply, error) {
	err := s.Clientset.CoreV1().Namespaces().Delete(ctx, in.GetName(), v1.DeleteOptions{})
	if err != nil {
		return nil, err
	}
	return &pb.NamespaceReply{Message: "Namespace " + in.GetName() + " deleted."}, nil
}

func (s *Server) ListNamespaces(ctx context.Context, in *pb.NamespaceRequest) (*pb.NamespaceListReply, error) {
	namespaces, err := s.Clientset.CoreV1().Namespaces().List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var names []string
	for _, namespace := range namespaces.Items {
		names = append(names, namespace.Name)
	}
	return &pb.NamespaceListReply{Namespaces: names}, nil
}
