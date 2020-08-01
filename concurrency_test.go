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

func TestConcurrentTreeOperations(t *testing.T) {

	initialMap := map[string]interface{}{
		"abc":         "a",
		"foobar":      "b",
		"foo/bar/baz": "c",
		"foo/baz/bar": "d",
		"foo/zip/zap": "e",
		"zipzap":      "zipzap",
	}
	addedMap := map[string]interface{}{
		"vanilla":                    "a",
		"vanilla-icecream":           "b",
		"vanilla-icecream-milkshake": "c",
		"vanilla-icecream-cake":      "d",
		"blackforest":                "e",
		"blackforest-cake":           "blackforest-cake",
	}
	removedKeys := []string{
		"vanilla-icecream",
		"vanilla-icecream-milkshake",
		"vanilla-icecream-cake",
	}

	r := NewFromMap(initialMap)

	type exp struct {
		inp string
		out []string
	}

	wg := new(sync.WaitGroup)
	wg.Add(3)

	go func() {
		t.Log("Add keys")
		defer wg.Done()
		r.InsertFromMap(addedMap)
	}()

	go func() {
		t.Log("Get Longest Prefix")
		defer wg.Done()
		out, found := r.LongestPrefix(Key("a"))
		if out != "a" {
			t.Fatalf(" failed to Longest get prefix, expected %v, got %v", "a", out)
		}
		if !found {
			t.Fatalf(" failed to find Longest get prefix for %v, expected true", "a")
		}
	}()

	go func() {
		defer wg.Done()
		t.Log("Get Max/Min/Len")
		max, _ := r.Maximum()
		min, _ := r.Minimum()
		if min != "a" {
			t.Fatalf(" failed to Longest get prefix, expected min  %v, got min %v", "a", min)
		}
		if max != "zipzap" {
			t.Fatalf(" failed to Longest get prefix, expected max  %v, got max %v", "zipzap", max)
		}
	}()

	wg.Wait()

	wg.Add(2)

	go func() {
		t.Log("Delete keys")
		defer wg.Done()
		for _, k := range removedKeys {
			r.Delete(Key(k))
		}
	}()

	go func() {
		t.Log("Get Longest Prefix")
		defer wg.Done()
		out, found := r.LongestPrefix(Key("blackforest-"))
		if out != "blackforest-cake" {
			t.Fatalf(" failed to Longest get prefix, expected %v, got %v", "blackforest-cake", out)
		}
		if !found {
			t.Fatalf(" failed to find Longest get prefix for %v, expected true", "blackforest-cake")
		}
	}()

	wg.Wait()

	if r.Size() != len(initialMap)+len(addedMap)-len(removedKeys) {
		t.Fatalf("bad len: %v %v", r.Size(), len(initialMap)+len(addedMap)-len(removedKeys))
	}
}
