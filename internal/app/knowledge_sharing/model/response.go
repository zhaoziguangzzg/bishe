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
	//作者
	Author User `json:"author"`
	//作者等级
	Level int `json:"level"`
	//文章
	Essay Essay `json:"essay"`
	//当前用户是否喜欢
	IsLike bool `json:"isLike"`
	//当前用户是否收藏
	IsCollect bool `json:"isCollect"`
}

type UserCircle struct {
	Uid    int    `json:"uid"`
	Name   string `json:"userName"`
	Circle Circle `json:"circle"`
}
