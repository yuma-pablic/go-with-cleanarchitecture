package user

type userResponseModel struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Type  string `json:"type"`
}
