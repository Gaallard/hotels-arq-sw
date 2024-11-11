package domain

type UserData struct {
	Id       int    `json: "id"`
	User     string `json: "user"`
	Password string `json: "password"`
	Admin    bool   `json: "admin"`
}

type LoginData struct {
	Token  string `json: "token"`
	IdU    int    `json: "idu"`
	AdminU bool   `json:"adminu"`
}

type UsersData []UserData
