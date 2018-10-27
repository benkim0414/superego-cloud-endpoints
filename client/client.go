package client

import (
	"context"

	peoplepb "github.com/benkim0414/superego-cloud-endpoints/people/v1alpha1"
	"google.golang.org/grpc"
)

// ProfileManagerClient is a client for interacting with People API.
type ProfileManagerClient struct {
	// The connection to ther service.
	conn *grpc.ClientConn

	// The gRPC API client.
	profileManagerClient peoplepb.ProfileManagerClient
}

// NewProfileManagerClient creates a new profile manager client.
func NewProfileManagerClient(addr string) (*ProfileManagerClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := &ProfileManagerClient{
		conn:                 conn,
		profileManagerClient: peoplepb.NewProfileManagerClient(conn),
	}
	return c, nil
}

// Connection returns the client's connection to the API service.
func (c *ProfileManagerClient) Connection() *grpc.ClientConn {
	return c.conn
}

// Close closes the connection to the API service. The user should invoke this
// when the client is no longer required.
func (c *ProfileManagerClient) Close() error {
	return c.conn.Close()
}

// CreateProfile creates a profile, consisting of the user profile information
// of Firebase.
func (c *ProfileManagerClient) CreateProfile(ctx context.Context, req *peoplepb.CreateProfileRequest) (*peoplepb.Profile, error) {
	resp, err := c.profileManagerClient.CreateProfile(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Gets the profile of a specific user.
func (c *ProfileManagerClient) GetProfile(ctx context.Context, req *peoplepb.GetProfileRequest) (*peoplepb.Profile, error) {
	resp, err := c.profileManagerClient.GetProfile(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
