package strings

import "sync"

var (
	strMap sync.Map
	mutex  sync.RWMutex
)

func InternString(str string) string {

	interned, _ := strMap.LoadOrStore(str, str)

	return interned.(string)
}

// TODO: cleanup
