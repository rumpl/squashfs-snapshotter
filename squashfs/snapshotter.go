package squashfs

import (
	"context"

	"github.com/containerd/containerd/mount"
	"github.com/containerd/containerd/snapshots"
)

type snapshotter struct {
}

func NewSnapshotter() (snapshots.Snapshotter, error) {
	return &snapshotter{}, nil
}

func (*snapshotter) Close() error {
	panic("unimplemented")
}

func (*snapshotter) Commit(ctx context.Context, name string, key string, opts ...snapshots.Opt) error {
	panic("unimplemented")
}

func (*snapshotter) Mounts(ctx context.Context, key string) ([]mount.Mount, error) {
	panic("unimplemented")
}

func (*snapshotter) Prepare(ctx context.Context, key string, parent string, opts ...snapshots.Opt) ([]mount.Mount, error) {
	panic("unimplemented")
}

func (*snapshotter) Remove(ctx context.Context, key string) error {
	panic("unimplemented")
}

func (*snapshotter) Stat(ctx context.Context, key string) (snapshots.Info, error) {
	panic("unimplemented")
}

func (*snapshotter) Update(ctx context.Context, info snapshots.Info, fieldpaths ...string) (snapshots.Info, error) {
	panic("unimplemented")
}

func (*snapshotter) Usage(ctx context.Context, key string) (snapshots.Usage, error) {
	panic("unimplemented")
}

func (*snapshotter) View(ctx context.Context, key string, parent string, opts ...snapshots.Opt) ([]mount.Mount, error) {
	panic("unimplemented")
}

func (*snapshotter) Walk(ctx context.Context, fn snapshots.WalkFunc, filters ...string) error {
	panic("unimplemented")
}
