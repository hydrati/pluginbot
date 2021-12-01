package aria2

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	// . "github.com/hyroge/pluginbot/utils/prelude"
	"github.com/zyxar/argo/rpc"
)

type NotifierFunc func([]rpc.Notifier)
type DefaultNotifier struct {
	OnDownloadStart      NotifierFunc
	OnDownloadPause      NotifierFunc
	OnDownloadStop       NotifierFunc
	OnDownloadError      NotifierFunc
	OnDownloadComplete   NotifierFunc
	OnBtDownloadComplete NotifierFunc
}

var (
	ERR_RPC_CLOSED      = errors.New("rpc closed")
	ERR_CLIENT_UNLOCKED = errors.New("client unlocked")
)

type ClientLockGuard struct {
	client   rpc.Client
	unlocker Unlocker
	unlocked bool
}

func (g *ClientLockGuard) Get() rpc.Client {
	if g.IsUnlocked() {
		panic(ERR_CLIENT_UNLOCKED)
	}
	return g.client
}

func (g *ClientLockGuard) IsUnlocked() bool {
	return g.unlocked
}

// alias of `Unlock()`
func (g *ClientLockGuard) Close() {
	g.Unlock()
}

func (g *ClientLockGuard) Unlock() {
	if !g.IsUnlocked() {
		g.unlocked = true
		g.unlocker()
	}
}

type RpcOptions struct {
	Host      string
	Port      uint
	Secret    string
	Transport string
	Timeout   string
	Notifier  rpc.Notifier
}

type RpcGuard struct {
	client  rpc.Client
	ctx     context.Context
	cancel  context.CancelFunc
	closed  bool
	timeout time.Duration
	mutex   *sync.Mutex
}

func NewClient(opts RpcOptions) (*RpcGuard, error) {
	ctx, cancel := context.WithCancel(context.Background())
	uri := fmt.Sprintf("%s://%s:%d/jsonrpc", opts.Transport, opts.Host, opts.Port)
	timeout, err := time.ParseDuration(opts.Timeout)
	if err != nil {
		return nil, err
	}

	client, err := rpc.New(ctx, uri, opts.Secret, timeout, opts.Notifier)
	if err != nil {
		return nil, err
	}

	mutex := &sync.Mutex{}

	return &RpcGuard{
		client:  client,
		timeout: timeout,
		ctx:     ctx,
		closed:  false,
		cancel:  cancel,
		mutex:   mutex,
	}, nil
}

func (r *RpcGuard) Close() error {
	r.mutex.Lock()
	if r.closed {
		return nil
	}
	r.closed = true
	r.cancel()
	return r.client.Close()
}

func (r *RpcGuard) IsClosed() bool {
	return r.closed
}

type Unlocker func()

func (r *RpcGuard) GetClient() (*ClientLockGuard, error) {
	if r.IsClosed() {
		return nil, ERR_RPC_CLOSED
	}
	r.mutex.Lock()
	return &ClientLockGuard{client: r.client, unlocker: r.mutex.Unlock, unlocked: false}, nil
}
