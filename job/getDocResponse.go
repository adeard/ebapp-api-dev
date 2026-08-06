package job

type GetDocResponse struct {
	Result  bool      `json:"result"`
	Objek   []DocItem `json:"objek"`
	Message string    `json:"message"`
}

type DocItem struct {
	DocumentID             string `json:"DocumentId"`
	Description            string `json:"Description"`
	PotensialActivityOwner string `json:"PotensialActivityOwner"`
	Activity               string `json:"Activity"`
	Status                 string `json:"Status"`
	ActivityOwner          string `json:"ActivityOwner"`
}

type UserEmailResponse struct {
	Result  bool       `json:"result"`
	Objek   []UserMail `json:"objek"`
	Message string     `json:"message"`
}

type UserMail struct {
	Email string `json:"Email"`
}

type LoginResponse struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
	Datas   string `json:"datas"`
}

type UserEmailByUsernameResponse struct {
	Result  bool   `json:"result"`
	Objek   string `json:"objek"`
	Message string `json:"message"`
}
