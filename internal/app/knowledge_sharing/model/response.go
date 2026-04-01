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

type UserChatContact struct {
	Uid      int    `json:"uid"`
	Name     string `json:"userName"`
	Content  string `json:"content"`
	UpdateAt string `json:"updateAt"`
}

type UserNotice struct {
	Uid      int    `json:"uid"`
	Name     string `json:"userName"`
	Content  string `json:"content"`
	Type     int    `json:"type"`
	UpdateAt string `json:"updateAt"`
}

type UserEssay struct {
	Uid   int    `json:"uid"`
	Name  string `json:"userName"`
	Level int    `json:"level"`
	Essay Essay  `json:"essay"`
}
