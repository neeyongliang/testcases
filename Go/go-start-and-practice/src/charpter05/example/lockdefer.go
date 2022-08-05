package main

import "sync"

var (
	valueByKey      = make(map[string]int)
	valueByKeyGuard sync.Mutex
)

// 普通的操作
func readValue(key string) int {
	valueByKeyGuard.Lock()
	v := valueByKey[key]
	valueByKeyGuard.Unlock()
	return v
}

func readValue2(key string) int {
	valueByKeyGuard.Lock()
	defer valueByKeyGuard.Unlock()
	return valueByKey[key]
}

func main() {
}
