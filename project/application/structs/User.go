package structs

type User struct {
	ID       uint   `json:"user_id,integer,omitempty"`
	Nickname string `json:"nickname,string,omitempty"`
	Email    string `json:"email,string,omitempty"`
	Password string `json:"password,string,omitempty"`
	Avatar   string `json:"avatar,string,omitempty"`
}
