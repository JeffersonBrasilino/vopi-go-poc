package entity

type Contact struct {
	contact string
}

func NewContact(contact string) (*Contact, error) {
	return &Contact{
		contact: contact,
	}, nil
}

func (c *Contact) Contact() string {
	return c.contact
}
