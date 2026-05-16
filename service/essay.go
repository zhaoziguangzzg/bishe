package service

import (
	"bishe/dao/mysql"
	"bishe/model"

	"github.com/gin-gonic/gin"
)

func SetEidToContext(c *gin.Context, eid int) {
	c.Set("eid", eid)
}

func GetEidFromContext(c *gin.Context) (eid int) {
	eid = c.GetInt("eid")
	return
}

func SetEssayToContext(c *gin.Context, essay *model.Essay) {
	c.Set("essay", essay)
}

// 从context获取文章
func GetEssayFromContext(c *gin.Context) (essay *model.Essay, ok bool) {
	essayAny, ok := c.Get("essay")
	if !ok {
		return
	}

	essay, ok = essayAny.(*model.Essay)
	if !ok {
		return
	}
	return
}

// create文章
func CreateEssay(newEssay *model.Essay) (err error) {
	return mysql.CreateEssay(newEssay)
}

// get用户全部文章
func GetAllEssayByUid(uid int, page int, pagesize int) (essays []model.Essay, err error) {
	return mysql.GetAllEssayByUid(uid, page, pagesize)
}

// get圈子中的文章
func GetAllEssayByCid(cid int, page int, pagesize int) (essays []model.Essay, err error) {
	return mysql.GetAllEssayByCid(cid, page, pagesize)
}

// 根据eid获取文章
func GetEssayByEid(eid int) (essay *model.Essay, err error) {
	return mysql.GetEssayByEid(eid)
}

// 根据title获取文章
func GetEssayByTitle(title string, cid int) (essay *model.Essay, err error) {
	return mysql.GetEssayByTitle(title, cid)
}

// 根据title关键词like获取文章
func GetEssayByLikeTitle(title string, cid int, page int, pagesize int) (essays []model.Essay, err error) {
	return mysql.GetEssayByLikeTitle(title, cid, page, pagesize)
}

// 根据eid更新文章信息
func UpdateEssayByEid(eid int, updateMap map[string]interface{}) (int64, error) {
	return mysql.UpdateEssayByEid(eid, updateMap)
}

// 更新IsDeleted删除essay
func UpdateEssayIsDeleted(eid int) (int64, error) {
	return mysql.UpdateEssayIsDeleted(eid)
}

// 更新文章周刊
func UpdateEssayWeekly(eid int, isWeekly int) (int64, error) {
	return mysql.UpdateEssayWeekly(eid, isWeekly)
}

// 获取文章周刊
func GetEssayWeeklyList(cid int, page int, pagesize int) (essays []model.Essay, err error) {
	return mysql.GetEssayWeeklyList(cid, page, pagesize)
}

// update essay essence
func UpdateEssayEssence(eid int, isEssence int) (int64, error) {
	return mysql.UpdateEssayEssence(eid, isEssence)
}

// 获取文章精粹
func GetEssayEssenceList(cid int, page int, pagesize int) (essays []model.Essay, err error) {
	return mysql.GetEssayEssenceList(cid, page, pagesize)
}

// 更新文章点赞数
func UpdateEssayLikeNum(eid int, num int) (int64, error) {
	return mysql.UpdateEssayLikeNum(eid, num)
}

// 更新文章评论数
func UpdateEssayCommentNum(eid int, num int) (int64, error) {
	return mysql.UpdateEssayCommentNum(eid, num)
}

// 更新文章收藏数
func UpdateEssayCollectNum(eid int, num int) (int64, error) {
	return mysql.UpdateEssayCollectNum(eid, num)
}
