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
func GetUrlEssayDetail(cid int, eid int) string {
	return "/page/essay/detail?cid=" + strconv.Itoa(cid) + "&eid=" + strconv.Itoa(eid)
}

// 课程详情
func GetUrlCourseDetail(cid int) string {
	return "/page/course/detail?course_id=" + strconv.Itoa(cid)
}

//用户主页
func GetUrlUserProfile(uid int) string {
	return "/page/user/profile?uid=" + strconv.Itoa(uid)
}

//举报详情
func GetUrlAccusationDetail(aid int) string {
	return "/page/accusation/detail?aid=" + strconv.Itoa(aid)
}

//反馈详情
func GetUrlFeedbackDetail(fid int) string {
	return "/page/feedback/detail?fid=" + strconv.Itoa(fid)
}
