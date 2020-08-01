package art

import (
	"sync"
	"testing"
)

func TestCC(t *testing.T) {
	var radix Tree
	wg := &sync.WaitGroup{}

	for z := 0; z < 20; z++ {
		radix = New()

		for f := 0; f < 20; f++ {
			wg.Add(1)
			go func(r Tree, w *sync.WaitGroup) {
				r.Insert([]byte("node"), 1)
				w.Done()
			}(radix, wg)

			wg.Add(1)
			go func(r Tree, w *sync.WaitGroup) {
				r.Insert([]byte("node"), 1)
				w.Done()
			}(radix, wg)

			wg.Add(1)
			go func(r Tree, w *sync.WaitGroup) {
				r.Search([]byte("node"))
				w.Done()
			}(radix, wg)

			wg.Add(1)
			go func(r Tree, w *sync.WaitGroup) {
				r.Search([]byte("node"))
				w.Done()
			}(radix, wg)

			wg.Add(1)
			go func(r Tree, w *sync.WaitGroup) {
				r.Delete([]byte("node"))
				w.Done()
			}(radix, wg)

			wg.Add(1)
			go func(r Tree, w *sync.WaitGroup) {
				r.Delete([]byte("node"))
				w.Done()
			}(radix, wg)
		}
		wg.Wait()
	}
}
