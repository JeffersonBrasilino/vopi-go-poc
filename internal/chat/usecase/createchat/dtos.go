package createchat

type CreateChatInput struct {
	ChannelId    string
	Participants []*Person
	Messages     []*Message
	BotName      string
}

type Message struct {
	Content string
	Status  string
	Sender  string
}

type Person struct {
	Name     string `json:"name"`
	Document string `json:"document"`
	Contact string `json:"contact"`
	//PersonType int      `json:"personType"`
}

//application -> input -> infra
