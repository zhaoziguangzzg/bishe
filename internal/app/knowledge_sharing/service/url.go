package service

import "strconv"

const (
	HOST string = "http://localhost:8080"
)

func GetUrlIndex() string {
	return "/"
}

func GetUrlLogin() string {
	return "/page/user/login"
}

// 圈子首页
func GetUrlCircleIndex(cid int) string {
	return "/page/circle/index?cid=" + strconv.Itoa(cid)
}

//文章详情
func GetUrlEssayDetail(eid int) string {
	return "/page/essay/detail?eid=" + strconv.Itoa(eid)
}

// 课程详情
func GetUrlCourseDetail(cid int) string {
	return "/page/course/detail?course_id=" + strconv.Itoa(cid)
}
