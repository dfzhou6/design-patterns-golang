package singleton

import (
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(ii int) {
			defer wg.Done()
			GetApp(ii)
		}(i)
	}

	wg.Wait()
}

func TestSingletonBySyncOnce(t *testing.T) {
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(ii int) {
			defer wg.Done()
			GetAppBySyncOnce(ii)
		}(i)
	}

	wg.Wait()
}
