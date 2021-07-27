package resources_test

import (
	"context"
	"fmt"
	"net"
	"testing"

	"google.golang.org/grpc"

	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/faker/v3"
	"github.com/hashicorp/go-hclog"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/cq-provider-yandex/resources"
	compute1 "github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	"github.com/yandex-cloud/go-sdk/gen/compute"
)

func TestComputeDisks(t *testing.T) {
	resource := providertest.ResourceTestData{
		Table: resources.ComputeDisks(),
		Config: client.Config{
			FolderIDs: []string{"testFolder"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			computeSvc, err := createDiskServer()
			if err != nil {
				return nil, err
			}
			c := client.NewYandexClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testFolder"}, &client.Services{
				Compute: computeSvc,
			})
			return c, nil
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
}

type FakeDiskServiceServer struct {
	compute1.UnimplementedDiskServiceServer
	Disk *compute1.Disk
}

func NewFakeDiskServiceServer() (*FakeDiskServiceServer, error) {
	var disk compute1.Disk
	faker.SetIgnoreInterface(true)
	err := faker.FakeData(&disk)
	if err != nil {
		return nil, err
	}
	return &FakeDiskServiceServer{Disk: &disk}, nil
}

func (s *FakeDiskServiceServer) List(context.Context, *compute1.ListDisksRequest) (*compute1.ListDisksResponse, error) {
	return &compute1.ListDisksResponse{Disks: []*compute1.Disk{s.Disk}}, nil
}

func createDiskServer() (*compute.Compute, error) {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		return nil, err
	}

	serv := grpc.NewServer()
	fakeDiskServiceServer, err := NewFakeDiskServiceServer()

	if err != nil {
		return nil, err
	}

	compute1.RegisterDiskServiceServer(serv, fakeDiskServiceServer)

	go func() {
		err := serv.Serve(lis)
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	return compute.NewCompute(
		func(ctx context.Context) (*grpc.ClientConn, error) {
			return conn, nil
		},
	), nil
}
