package createchat

type CreateChatInput struct {
	ChannelId    string
	Participants []*Person
	Messages     []*Message
	BotName string
}

type Message struct {
	Content string
	Status  string
	Sender  *Person
}

type Person struct {
	Contact string
	Name    string
}


//application -> input -> infra