package mediator

import "testing"

func TestMediator(t *testing.T) {
	chatRoom := NewChatRoomImpl()
	alice := NewUserImpl("Alice", chatRoom)
	mike := NewUserImpl("Mike", chatRoom)

	chatRoom.RegisterUser(alice)
	chatRoom.RegisterUser(mike)

	alice.SendMsg(mike.GetName(), "hello")
	mike.SendMsg(alice.GetName(), "hi")
}
