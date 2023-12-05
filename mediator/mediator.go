package mediator

import "fmt"

type ChatRoom interface {
	RegisterUser(user User)
	NotifyMsg(from string, to string, msg string)
}

type ChatRoomImpl struct {
	userMap map[string]User
}

func NewChatRoomImpl() *ChatRoomImpl {
	return &ChatRoomImpl{
		userMap: make(map[string]User, 0),
	}
}

func (impl *ChatRoomImpl) RegisterUser(user User) {
	impl.userMap[user.GetName()] = user
}

func (impl *ChatRoomImpl) NotifyMsg(from string, to string, msg string) {
	toUser, ok := impl.userMap[to]
	if !ok {
		fmt.Printf("to user %v doesn't exist\n", to)
		return
	}
	toUser.ReceiveMsg(from, msg)
}

type User interface {
	SendMsg(to string, msg string)
	ReceiveMsg(from string, msg string)
	GetName() string
}

type UserImpl struct {
	name     string
	chatRoom ChatRoom
}

func NewUserImpl(name string, chatRoom ChatRoom) *UserImpl {
	return &UserImpl{
		name:     name,
		chatRoom: chatRoom,
	}
}

func (impl *UserImpl) GetName() string {
	return impl.name
}

func (impl *UserImpl) SendMsg(to string, msg string) {
	impl.chatRoom.NotifyMsg(impl.name, to, msg)
}

func (impl *UserImpl) ReceiveMsg(from string, msg string) {
	fmt.Printf("user %v receive msg from user %v, msg is %v\n", impl.name, from, msg)
}
