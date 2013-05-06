package main 

type Payload struct {
	Before     string     `json:"before"`
	After      string     `json:"after"`
	Ref        string     `json:"ref"`
	Commits    []Commit   `json:"commits"`
	Repository Repository `json:"repository"`
}

type Commit struct {
	Id        string   `json:"id"`
	Message   string   `json:"message"`
	Timestamp string   `json:"timestamp"`
	Url       string   `json:"url"`
	Added     []string `json:"added"`
	Removed   []string `json:"removed"`
	Modified  []string `json:"modified"`
	Author    Author   `json:"author"`
}

type Repository struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	Pledgie     string `json:"pledgie"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Watchers    int    `json:"watchers"`
	Forks       int    `json:"forks"`
	Private     bool   `json:"private"`
	Owner       Author `json:"owner"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
