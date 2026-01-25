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

	CODE_USER_NAME_EXIST       int = 2001
	CODE_USER_NAME_NOT_EXIST   int = 2002
	CODE_USER_PASSWORD_INVALID int = 2003
	CODE_USER_NOT_LOGIN        int = 2004
	CODE_USER_NAME_LEN_INVALID int = 2005
	CODE_USER_EMAIL_INVALID    int = 2006
	CODE_USER_PHONE_INVALID    int = 2007

	CODE_CHECKIN_REPEAT   int = 3001
	CODE_COMMENT_TOO_LONG int = 3002

	CODE_CIRCLE_EXIST                     int = 4001
	CODE_CIRCLE_TITLE_LEN_INVASLID        int = 4002
	CODE_CIRCLE_INTRODUCTION_LEN_INVASLID int = 4003
	CODE_CIRCLE_PRICE_INVASLID            int = 4004
	CODE_CIRCLE_NOT_EXIST                 int = 4005
	CODE_USER_NOT_CREATE_CIRCLE           int = 4006
)

var CodeMsgMap map[int]string = map[int]string{
	CODE_SUCCESS:      "成功",
	CODE_SYS_ERROR:    "服务出错，稍后再试",
	CODE_PARAMS_ERROR: "参数错误",

	CODE_USER_NAME_EXIST:       "用户名已存在",
	CODE_USER_NAME_NOT_EXIST:   "用户名不存在",
	CODE_USER_PASSWORD_INVALID: "用户密码错误",
	CODE_USER_NOT_LOGIN:        "用户未登录",
	CODE_USER_NAME_LEN_INVALID: "用户名长度错误",
	CODE_USER_EMAIL_INVALID:    "用户邮箱错误",
	CODE_USER_PHONE_INVALID:    "用户手机号错误",

	CODE_CHECKIN_REPEAT:   "重复打卡",
	CODE_COMMENT_TOO_LONG: "评论最多200字",

	CODE_CIRCLE_EXIST:                     "圈子已存在",
	CODE_CIRCLE_TITLE_LEN_INVASLID:        "圈子标题长度错误",
	CODE_CIRCLE_INTRODUCTION_LEN_INVASLID: "圈子简介长度错误",
	CODE_CIRCLE_PRICE_INVASLID:            "圈子价格超过1w",
	CODE_CIRCLE_NOT_EXIST:                 "圈子不存在",
	CODE_USER_NOT_CREATE_CIRCLE:           "用户未加入圈子",
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
