package model

type APIResponse struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// 榜单中的圈子
type RankCircle struct {
	Id            int    `json:"id"`
	Rank          int    `json:"rank"`
	Title         string `json:"title"`
	Price         int    `json:"price"`
	PriceText     string `json:"priceText"`
	CircleOwnerId int    `json:"circleOwnerId"`
	OwnerName     string `json:"ownerName"`
	JoinNum       int    `json:"joinNum"`
}

// 榜单
type TypeList struct {
	TypeName    string       `json:"typeName"`
	ListType    int          `json:"listType"`
	RankCircles []RankCircle `json:"rankCircles"`
}

type CommentEssay struct {
	Comment UserEssayComment `json:"comment"`
	Essay   Essay            `json:"essay"`
}

type PurchaseCourse struct {
	Purchase Purchase `json:"purchase"`
	Course   Course   `json:"course"`
}

type UserChatContact struct {
	ChatUser    User        `json:"chatUser"`
	ChatContact ChatContact `json:"chatContact"`
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
	//作者积分
	Score int `json:"score"`
	//文章
	Essay Essay `json:"essay"`
	//当前用户是否喜欢
	IsLike bool `json:"isLike"`
	//当前用户是否收藏
	IsCollect bool `json:"isCollect"`
}

type UserCourse struct {
	//作者
	Author User `json:"author"`
	//课程
	Course Course `json:"course"`
}

type UserCircle struct {
	User   User   `json:"user"`
	Circle Circle `json:"circle"`
}

type UserComment struct {
	User    User             `json:"user"`
	Comment UserEssayComment `json:"comment"`
}

type NoticeUrl struct {
	UserUrl string `json:"userUrl"`
	Url     string `json:"circleUrl"`
	Notice  Notice `json:"notice"`
}
