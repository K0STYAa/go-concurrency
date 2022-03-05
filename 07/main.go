/*
Реализовать конкурентную запись данных в map.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Стандартная библиотека предоставляет два вида мьютексов в пакете sync: sync.Mutex и sync.RWMutex;
последний оптимизирован для случаев,когда ваша программа имеет дело с множеством читателей
и очень небольшим количеством записывателей.
*/

type muMap struct {
	sync.RWMutex
	data map[string]int
}

func create() *muMap {
	return &muMap{
		data: make(map[string]int),
	}
}

func (m *muMap) get(value string) int {
	m.RLock()
	defer m.RUnlock()
	v := m.data[value]
	return v
}

func (m *muMap) set(key string, value int) {
	m.Lock()
	defer m.Unlock()
	m.data[key] = value
}

// inc increments the counter for the given key.
func (m *muMap) inc(key string) {
	m.Lock()
	defer m.Unlock()
	m.data[key]++
}

func (m *muMap) delete(key string) {
	m.Lock()
	defer m.Unlock()
	delete(m.data, key)
}

func main() {
	m := create()
	for i := 0; i < 1000; i++ {
		go m.inc("hello world")
	}

	time.Sleep(time.Millisecond)
	fmt.Println(m.get("hello world"))

}
