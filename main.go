package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	snapshotsapi "github.com/containerd/containerd/api/services/snapshots/v1"
	"github.com/containerd/containerd/contrib/snapshotservice"
	"github.com/rumpl/squashfs-snapshotter/squashfs"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	rpc := grpc.NewServer()
	sn, err := squashfs.NewSnapshotter()
	if err != nil {
		return err
	}

	service := snapshotservice.FromSnapshotter(sn)
	snapshotsapi.RegisterSnapshotsServer(rpc, service)

	l, err := net.Listen("unix", "/var/squashfs-snapshotter.sock")
	if err != nil {
		return err
	}

	return rpc.Serve(l)
}
