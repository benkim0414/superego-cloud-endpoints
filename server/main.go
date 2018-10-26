//go:generate protoc -I ../people/v1alpha1 -I $GOOGLEAPIS_DIR --go_out=plugins=grpc:$GOPATH/src ../people/v1alpha1/profile_service.proto
package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"

	"cloud.google.com/go/datastore"
	peoplepb "github.com/benkim0414/superego-cloud-endpoints/people/v1alpha1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	addr = flag.String("addr", ":50051", "Network host:port to listen on for gRPC connections.")
)

const (
	profileKind = "Profile"
)

type server struct {
	client *datastore.Client
}

// Creates a profile, consisting of the user profile information of Firebase.
func (s *server) CreateProfile(ctx context.Context, req *peoplepb.CreateProfileRequest) (*peoplepb.Profile, error) {
	key := datastore.NameKey(profileKind, req.Name, nil)
	req.Profile.Name = req.Name
	_, err := s.client.Put(ctx, key, req.Profile)
	if err != nil {
		return nil, err
	}
	return req.Profile, nil
}

// Gets the profile of a specific user.
func (s *server) GetProfile(ctx context.Context, req *peoplepb.GetProfileRequest) (*peoplepb.Profile, error) {
	key := datastore.NameKey(profileKind, req.Name, nil)
	p := &peoplepb.Profile{}
	err := s.client.Get(ctx, key, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	projectID := os.Getenv("DATASTORE_PROJECT_ID")
	client, err := datastore.NewClient(context.Background(), projectID)
	if err != nil {
		log.Fatalf("failed to create a datastore client: %v", err)
	}
	defer client.Close()

	s := grpc.NewServer()
	peoplepb.RegisterProfileManagerServer(s, &server{
		client: client,
	})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
