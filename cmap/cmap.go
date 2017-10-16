package main

import (
	"crypto/md5"
	"fmt"
	"sync"
)

var inputs = []string{
	"Go is awesome",
	"Go is great",
	"Go is fun",
}

func main() {
	sums := make([]string, len(inputs))
	var wg sync.WaitGroup
	for i, input := range inputs {
		wg.Add(1)
		go func(i int, input string) {
			defer wg.Done()
			hash := md5.Sum([]byte(input))
			sums[i] = fmt.Sprintf("%x", hash)
		}(i, input)
	}

	wg.Wait()
	for i, in := range inputs {
		sum := sums[i]
		fmt.Printf("%-20s %s\n", in, sum)
	}
}

package main

import (
	"crypto/md5"
	"fmt"
	"sync"
)

var inputs = []string{
	"Go is awesome",
	"Go is great",
	"Go is fun",
}

func main() {
	var outputs sync.Map
	var wg sync.WaitGroup
	for _, input := range inputs {
		wg.Add(1)
		go func(input string) {
			hash := md5.Sum([]byte(input))
			outputs.Store(input, fmt.Sprintf("%x", hash))
			wg.Done()
		}(input)
	}
	wg.Wait()
	
	for _, k := range inputs {
		v, _ := outputs.Load(k)
		fmt.Printf("%-20s %s\n", k, v)
	}
}

package main

import (
	"crypto/md5"
	"fmt"
)

var inputs = []string{
	"Go is awesome",
	"Go is great",
	"Go is fun",
}

type Handler struct {
	d    map[string]string
	ch   chan string
	done chan struct{}
}

func New() Handler {
	h := Handler{
		d:    map[string]string{},
		ch:   make(chan string),
		done: make(chan struct{}),
	}
	go func() {
		for v := range h.ch {
			h.d[v] = fmt.Sprintf("%x",md5.Sum([]byte(v)))
		}
		close(h.done)
	}()
	return h
}

func (h Handler) Close() map[string]string{
	close(h.ch)
	<-h.done
	return h.d
}

func (h Handler) Add(s string) {
	h.ch <- s
}

func main() {
	h := New()
	for _, input := range inputs {
		h.Add(input)
	}
	outputs := h.Close()
	for _, k := range inputs {
		fmt.Printf("%-20s %s\n", k, outputs[k])
	}
}

