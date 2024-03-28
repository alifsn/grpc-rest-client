package request

type CheckInRequest struct {
	Username string `json:"username"`
	Datetime string `json:"datetime"`
}
