package tools

import (
	"errors"
	"sync"
)

var mutex sync.Mutex

func SafeCall(f func() error) error {
	mutex.Lock()
	defer mutex.Unlock()

	if nil != f {
		return f()
	}
	return errors.New(`缺乏回调`)
}
