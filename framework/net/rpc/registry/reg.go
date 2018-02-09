package registry

import (
	"os"

	"github.com/abronan/valkeyrie/store"
	"github.com/cuigh/auxo/config"
	"github.com/cuigh/auxo/log"
	"github.com/cuigh/auxo/net/rpc/registry"
	"github.com/cuigh/auxo/net/rpc/registry/valkeyrie"
)

func init() {
	registerBuild()
}

func registerBuild() {
	backend := config.GetString("global.registry.store.backend")
	var _backend store.Backend
	switch backend {
	case "etcd":
		_backend = store.ETCD
	case "etcdv3":
		_backend = store.ETCDV3
	case "consul":
		_backend = store.CONSUL
	default:
		log.Get("runtime").Error("registry store backend is not supported!")
		os.Exit(-1)
	}
	ret := &valkeyrie.Builder{Backend: _backend}
	registry.Register(ret)
}
