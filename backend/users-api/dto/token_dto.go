package dto

type TokenDto struct {
	Token   string `json:"token"`
	Id_user int    `json:"idUser"`
	Role    int    `json:"role"`
}
