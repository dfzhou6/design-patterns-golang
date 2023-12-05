package flyweight

import "fmt"

type User struct {
	Name string
}

func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}

type UserCache struct {
	userMap map[string]*User
}

func NewUserCache() *UserCache {
	return &UserCache{
		userMap: make(map[string]*User, 0),
	}
}

func (impl *UserCache) Get(name string) *User {
	user, ok := impl.userMap[name]
	if !ok {
		user = NewUser(name)
		impl.userMap[name] = user
	}
	return user
}

func (impl *UserCache) PrintCache() {
	for _, item := range impl.userMap {
		fmt.Printf("user: %+v\n", *item)
	}
}
