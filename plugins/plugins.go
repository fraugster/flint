package plugins

import (
	"context"
	"log"
	"sync"

	"github.com/pkg/errors"
)

var (
	plugs = make(map[string]Interface)
	lock  = &sync.RWMutex{}
)

// Interface is the single file checker
type Interface interface {
	Process(context.Context, []byte) error
}

// Register a new plugin
func Register(name string, p Interface) {
	lock.Lock()
	defer lock.Unlock()

	if _, ok := plugs[name]; ok {
		log.Panic("plugin with the same name is already registered")
	}

	plugs[name] = p
}

// Call the plugins on a single file
func Call(ctx context.Context, src []byte) error {
	lock.RLock()
	defer lock.RUnlock()

	for i := range plugs {
		if err := plugs[i].Process(ctx, src); err != nil {
			// TODO: Multi err
			return errors.Wrapf(err, "plugin '%s' failed", i)
		}
	}

	return nil
}
