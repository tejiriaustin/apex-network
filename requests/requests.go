package requests

type CreatePlayerRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
