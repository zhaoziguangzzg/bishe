package controller

import (
	"bishe/internal/app/knowledge_sharing/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	CODE_SUCCESS      int = 0
	CODE_SYS_ERROR    int = 1
	CODE_PARAMS_ERROR int = 1001

	CODE_USER_NAME_EXIST              int = 2001
	CODE_USER_NAME_NOT_EXIST          int = 2002
	CODE_USER_PASSWORD_INVALID        int = 2003
	CODE_USER_NOT_LOGIN               int = 2004
	CODE_USER_NAME_LEN_INVALID        int = 2005
	CODE_USER_EMAIL_INVALID           int = 2006
	CODE_USER_PHONE_INVALID           int = 2007
	CODE_USER_REPLACE                 int = 2008
	CODE_USER_AGE_INVALID             int = 2009
	CODE_PASSWORD_WRONG               int = 2010
	CODE_USER_AVATAR_NAME_LEN_INVALID int = 2011
	CODE_USER_NOT_EXIST               int = 2012
	CODE_ADMIN_USER_NOT_LOGIN         int = 2013

	CODE_CHECKIN_REPEAT   int = 3001
	CODE_COMMENT_TOO_LONG int = 3002
	CODE_TITLE_REPLACE    int = 3003

	CODE_ANNOUNCE_NOT_EXIST            int = 3006
	CODE_ANNOUNCE_TITLE_LEN_INVASLID   int = 3007
	CODE_ANNOUNCE_CONTENT_LEN_INVASLID int = 3008

	CODE_COURSE_TITLE_LEN_INVASLID   int = 1011
	CODE_COURSE_CONTENT_LEN_INVASLID int = 1012
	CODE_COURSE_PRICE_MAX_INVASLID   int = 1013
	CODE_USER_NOT_COURSE_OWNER       int = 1014
	CODE_COURSE_NOT_EXIST            int = 1015

	CODE_LESSON_TITLE_LEN_INVASLID   int = 1016
	CODE_LESSON_CONTENT_LEN_INVASLID int = 1017

	CODE_MENU_NAME_EXIST int = 1021
	CODE_MENU_NOT_EXIST  int = 1022

	CODE_ROLE_NAME_EXIST int = 1031
	CODE_ROLE_NOT_EXIST  int = 1032

	CODE_CIRCLE_EXIST                     int = 4001
	CODE_CIRCLE_TITLE_LEN_INVASLID        int = 4002
	CODE_CIRCLE_INTRODUCTION_LEN_INVASLID int = 4003
	CODE_CIRCLE_PRICE_INVASLID            int = 4004
	CODE_CIRCLE_NOT_EXIST                 int = 4005
	CODE_USER_NOT_JOIN_CIRCLE             int = 4006
	CODE_USER_JOIN_CIRCLE                 int = 4007
	CODE_CIRCLE_FREE                      int = 4008
	CODE_USER_CIRCLE_RENEW                int = 4009
	CODE_USER_NOT_CIRCLE_OWNER            int = 4010

	CODE_ESSAY_EXIST                int = 5001
	CODE_ESSAY_TITLE_LEN_INVASLID   int = 5002
	CODE_ESSAY_CONTENT_LEN_INVASLID int = 5003
	CODE_ESSAY_NOT_EXIST            int = 5004
	CODE_USER_NOT_ESSAY_OWNER       int = 5005

	CODE_ADVERT_NOT_EXIST             int = 5006
	CODE_ADVERT_POSITION_LEN_INVASLID int = 5007
	CODE_ADVERT_CONTENT_LEN_INVASLID  int = 5008
	CODE_ADVERT_ADDR_LEN_INVASLID     int = 5009

	CODE_COMMENT_CONTENT_LEN_INVASLID int = 6003

	CODE_CONTACT_NOT_EXIST int = 6006

	CODE_FAVORITE_NOT_EXIST             int = 7001
	CODE_FAVORITE_EXIST                 int = 7002
	CODE_INTERACT_FAVORITE_LEN_INVASLID int = 7003
	CODE_LIKE_NOT_EXIST                 int = 7004
	CODE_COLLECT_NOT_EXIST              int = 7005
	CODE_COLLECT_EXIST                  int = 7006
	CODE_LIKE_EXIST                     int = 7007
	CODE_COLLECT_DELETED                int = 7008
	CODE_LIKE_DELETED                   int = 7009

	CODE_CHAT_CONTENT_LEN_INVASLID int = 8003
	CODE_INFORMATION_NOT_EXIST     int = 8004

	CODE_FEEDBACK_CONTENT_LEN_INVASLID int = 8006
	CODE_FEEDBACK_EXIST                int = 8007
	CODE_FEEDBACK_NOT_EXIST            int = 8008

	CODE_USER_NOT_FOLLOW int = 9001

	CODE_ACCUSATION_CONTENT_LEN_INVASLID int = 9006
	CODE_ACCUSATION_EXIST                int = 9007
	CODE_ACCUSATION_NOT_EXIST            int = 9008

	CODE_ORDERS_NOT_EXIST int = 9011
	CODE_USER_PURCHASED   int = 9012

	CODE_LOCKED int = 9021
)

var CodeMsgMap map[int]string = map[int]string{
	CODE_SUCCESS:      "成功",
	CODE_SYS_ERROR:    "服务出错，稍后再试",
	CODE_PARAMS_ERROR: "参数错误",

	CODE_USER_NAME_EXIST:              "用户名已存在",
	CODE_USER_NAME_NOT_EXIST:          "用户名不存在",
	CODE_USER_PASSWORD_INVALID:        "用户密码错误",
	CODE_PASSWORD_WRONG:               "密码错误",
	CODE_USER_NOT_LOGIN:               "用户未登录",
	CODE_USER_NAME_LEN_INVALID:        "用户名长度错误",
	CODE_USER_EMAIL_INVALID:           "用户邮箱错误",
	CODE_USER_PHONE_INVALID:           "用户手机号错误",
	CODE_USER_REPLACE:                 "请更换用户名",
	CODE_USER_AGE_INVALID:             "用户年龄错误",
	CODE_USER_AVATAR_NAME_LEN_INVALID: "用户头像名长度错误",
	CODE_ADMIN_USER_NOT_LOGIN:         "管理员用户未登录",

	CODE_CHECKIN_REPEAT:   "重复打卡",
	CODE_COMMENT_TOO_LONG: "评论最多200字",
	CODE_TITLE_REPLACE:    "请更换标题",

	CODE_ANNOUNCE_NOT_EXIST:            "公告不存在",
	CODE_ANNOUNCE_TITLE_LEN_INVASLID:   "公告标题长度错误",
	CODE_ANNOUNCE_CONTENT_LEN_INVASLID: "公告内容长度错误",

	CODE_COURSE_TITLE_LEN_INVASLID:   "课程标题长度错误",
	CODE_COURSE_CONTENT_LEN_INVASLID: "课程内容简介长度错误",
	CODE_COURSE_PRICE_MAX_INVASLID:   "课程价格超过1w",
	CODE_USER_NOT_COURSE_OWNER:       "用户不是课程作者",
	CODE_COURSE_NOT_EXIST:            "课程不存在",

	CODE_LESSON_TITLE_LEN_INVASLID:   "课时标题长度错误",
	CODE_LESSON_CONTENT_LEN_INVASLID: "课时内容长度错误",

	CODE_MENU_NAME_EXIST: "菜单名已存在",
	CODE_MENU_NOT_EXIST:  "菜单不存在",

	CODE_ROLE_NAME_EXIST: "角色名已存在",
	CODE_ROLE_NOT_EXIST:  "角色不存在",

	CODE_CIRCLE_EXIST:                     "圈子已存在",
	CODE_CIRCLE_TITLE_LEN_INVASLID:        "圈子标题长度错误",
	CODE_CIRCLE_INTRODUCTION_LEN_INVASLID: "圈子简介长度错误",
	CODE_CIRCLE_PRICE_INVASLID:            "圈子价格超过1w",
	CODE_CIRCLE_NOT_EXIST:                 "圈子不存在",
	CODE_USER_NOT_JOIN_CIRCLE:             "用户未加入圈子",
	CODE_USER_JOIN_CIRCLE:                 "用户已加入圈子",
	CODE_CIRCLE_FREE:                      "免费圈子",
	CODE_USER_CIRCLE_RENEW:                "用户圈子需要付费",
	CODE_USER_NOT_CIRCLE_OWNER:            "用户不是圈主",

	CODE_ESSAY_EXIST:                "文章已存在",
	CODE_ESSAY_TITLE_LEN_INVASLID:   "文章标题长度错误",
	CODE_ESSAY_CONTENT_LEN_INVASLID: "文章内容长度错误",
	CODE_ESSAY_NOT_EXIST:            "文章不存在",
	CODE_USER_NOT_ESSAY_OWNER:       "用户不是文章作者",

	CODE_ADVERT_NOT_EXIST:             "广告不存在",
	CODE_ADVERT_POSITION_LEN_INVASLID: "广告位置长度错误",
	CODE_ADVERT_CONTENT_LEN_INVASLID:  "广告内容长度错误",
	CODE_ADVERT_ADDR_LEN_INVASLID:     "广告地址长度错误",

	CODE_COMMENT_CONTENT_LEN_INVASLID: "文章评论内容长度错误",

	CODE_CONTACT_NOT_EXIST: "联系人不存在",

	CODE_FAVORITE_NOT_EXIST:             "收藏夹不存在",
	CODE_FAVORITE_EXIST:                 "收藏夹已存在",
	CODE_INTERACT_FAVORITE_LEN_INVASLID: "文章收藏夹名长度错误",
	CODE_LIKE_NOT_EXIST:                 "用户文章喜欢不存在",
	CODE_COLLECT_NOT_EXIST:              "用户文章收藏不存在",
	CODE_COLLECT_EXIST:                  "用户文章已收藏",
	CODE_LIKE_EXIST:                     "用户文章已喜欢",
	CODE_COLLECT_DELETED:                "用户文章收藏已删除",
	CODE_LIKE_DELETED:                   "用户文章喜欢已删除",

	CODE_CHAT_CONTENT_LEN_INVASLID: "消息内容长度错误",
	CODE_INFORMATION_NOT_EXIST:     "消息不存在",

	CODE_FEEDBACK_CONTENT_LEN_INVASLID: "反馈内容长度错误",
	CODE_FEEDBACK_EXIST:                "反馈已存在",
	CODE_FEEDBACK_NOT_EXIST:            "反馈不存在",

	CODE_USER_NOT_FOLLOW: "该用户未关注",

	CODE_ACCUSATION_CONTENT_LEN_INVASLID: "举报内容长度错误",
	CODE_ACCUSATION_EXIST:                "举报已存在",
	CODE_ACCUSATION_NOT_EXIST:            "举报不存在",

	CODE_ORDERS_NOT_EXIST: "有未支付订单",
	CODE_USER_PURCHASED:   "用户已购买",

	CODE_LOCKED: "已锁定，请稍后再试",
}

func GetMsgByCode(code int) (msg string) {
	msg, ok := CodeMsgMap[code]
	if ok {
		return
	} else {
		msg = CodeMsgMap[CODE_SYS_ERROR]
		return
	}
}

func MakeApiResponse(c *gin.Context, code int, data interface{}) {
	msg := GetMsgByCode(code)
	if data == nil {
		data = make(map[int]string)
	}

	c.JSON(http.StatusOK, model.APIResponse{
		Code:    code,
		Data:    data,
		Message: msg,
	})
}

func MakeApiResponseSuccess(c *gin.Context, data interface{}) {
	MakeApiResponse(c, CODE_SUCCESS, data)
}

func MakeApiResponseSuccessDefault(c *gin.Context) {
	MakeApiResponse(c, CODE_SUCCESS, nil)
}

func MakeApiResponseError(c *gin.Context, code int) {
	MakeApiResponse(c, code, nil)
}

// 参数错误
func MakeApiResponseErrorParams(c *gin.Context) {
	MakeApiResponse(c, CODE_PARAMS_ERROR, nil)
}

// 系统错误
func MakeApiResponseErrorDefault(c *gin.Context) {
	MakeApiResponse(c, CODE_SYS_ERROR, nil)
}
