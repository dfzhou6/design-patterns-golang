package flyweight

import "testing"

func TestFlyweight(t *testing.T) {
	cache := NewUserCache()
	_ = cache.Get("user1")
	_ = cache.Get("user2")
	_ = cache.Get("user1")
	_ = cache.Get("user2")
	cache.PrintCache()
}
