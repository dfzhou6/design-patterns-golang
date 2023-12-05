package singleton

import (
	"fmt"
	"sync"
)

type Application struct {
}

var app *Application
var lock sync.Mutex
var once sync.Once

func GetApp(index int) *Application {
	if app == nil {
		lock.Lock()
		defer lock.Unlock()

		if app == nil {
			app = &Application{}
			fmt.Printf("new app success, index: %v\n", index)
		}
	}

	return app
}

func GetAppBySyncOnce(index int) *Application {
	once.Do(func() {
		if app == nil {
			app = &Application{}
			fmt.Printf("new app by sync.once success, index: %v\n", index)
		}
	})

	return app
}
