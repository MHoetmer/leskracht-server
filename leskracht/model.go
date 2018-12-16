package leskracht

type User struct {
	ID        int       `json:"id","omitempty"`
	FirstName string    `json:"firstName","omitempty"`
	LastName  string    `json:"lastName","omitempty"`
	Email     string    `json:"email","omitempty"`
	BirthDate string    `json:"birtDate","omitempty"`
	Messages  []Message `json:"messages","omitempty"`
}

type Message struct {
	ID      int    `json:"id","omitempty"`
	Author  User   `json:"author","omitempty"`
	Message string `json:"message","omitempty"`
	Date    string `json:"date","omitempty"`
}
