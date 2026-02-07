package model

type APIResponse struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type CommentEssay struct {
	Comment UserEssayComment `json:"comment"`
	Essay   Essay            `json:"essay"`
}
