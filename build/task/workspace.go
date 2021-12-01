package task

import (
	"context"
	"time"

	. "github.com/hyroge/pluginbot/utils/prelude"
	"github.com/zyxar/argo/rpc"
)

func Init() {
	ctx, cancal := context.WithCancel(context.Background())
	rpc.New(ctx, "localhost:", "hello", time.ParseDuration("180s"), 
}
