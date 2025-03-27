package response

type BaseResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}
