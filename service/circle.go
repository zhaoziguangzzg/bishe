package service

import (
	"bishe/dao/mysql"
	"bishe/dao/redis"
	"bishe/model"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	LIST_TYPE_ALL    int = 0
	LIST_TYPE_FREE   int = 1
	LIST_TYPE_CHARGE int = 2
)

func SetCidToContext(c *gin.Context, cid int) {
	c.Set("cid", cid)
}

func GetCidFromContext(c *gin.Context) (cid int) {
	cid = c.GetInt("cid")
	return
}

func SetCircleToContext(c *gin.Context, circle *model.Circle) {
	c.Set("circle", circle)
}

// 从context获取圈子
func GetCircleFromContext(c *gin.Context) (circle *model.Circle, ok bool) {
	circleAny, ok := c.Get("circle")
	if !ok {
		return
	}

	circle, ok = circleAny.(*model.Circle)
	if !ok {
		return
	}
	return
}

// create圈子
func CreateCircle(newCircle *model.Circle) (err error) {
	return mysql.CreateCircle(newCircle)
}

// 根据cid获取圈子
func GetCircleByCid(cid int) (circle *model.Circle, err error) {
	return mysql.GetCircleByCid(cid)
}

func GetCircleByCidList(cidList []int) (circleMap map[int]model.Circle, err error) {
	return mysql.GetCircleByCidList(cidList)
}

// 根据title获取圈子
func GetCircleByTitle(title string) (circle *model.Circle, err error) {
	return mysql.GetCircleByTitle(title)
}

// 通过like title关键词获取圈子
func GetCircleByLikeTitle(title string, page int, pagesize int) (circles []model.Circle, err error) {
	return mysql.GetCircleByLikeTitle(title, page, pagesize)
}

// get 付费圈子
func GetCircleAllChargeOrderByJoinNum(page int, pagesize int) (circles []model.Circle, err error) {
	return mysql.GetCircleAllChargeOrderByJoinNum(page, pagesize)
}

// get 免费圈子
func GetCricleAllFreeOrderByJoinNum(page int, pagesize int) (circles []model.Circle, err error) {
	return mysql.GetCricleAllFreeOrderByJoinNum(page, pagesize)
}

// get all圈子
func GetCircleRank(ctx context.Context) (circles []model.Circle, err error) {
	return redis.ZrevrangeCircleJoinNum(ctx)
}

func GetCircleRankFree(ctx context.Context) (circles []model.Circle, err error) {
	return redis.ZrevrangeCircleJoinNumFree(ctx)
}

func GetCircleRankCharge(ctx context.Context) (circles []model.Circle, err error) {
	return redis.ZrevrangeCircleJoinNumCharge(ctx)
}

// get 用户创建的圈子
func GetUserCreateCircleByUid(uid int, page int, pagesize int) (circles []model.Circle, err error) {
	return mysql.GetUserCreateCircleByUid(uid, page, pagesize)
}

// get 用户加入的圈子
func GetUserJoinCircleListByUid(uid int, page int, pagesize int) (circles []model.Circle, err error) {
	return mysql.GetUserJoinCircleListByUid(uid, page, pagesize)
}

// 更新圈子join num 增加
func IncrUpdateCircleJoinNumByCid(ctx context.Context, cid int, isFree bool) (affectRows int64, score float64, err error) {
	affectRows, err = mysql.IncrUpdateCircleJoinNumByCid(cid)
	if err != nil {
		return
	}

	//key "circle_join_num",score join_num,member cid
	score, err = redis.ZincrCircleJoinNum(ctx, cid, isFree)

	return
}

// 更新圈子join num 减少
func DecrrUpdateCircleJoinNumByCid(cid int) (int64, error) {
	return mysql.DecrrUpdateCircleJoinNumByCid(cid)
}

// 更新圈子信息
func UpdateCircleByCid(cid int, updateMap map[string]interface{}) (int64, error) {
	return mysql.UpdateCircleByCid(cid, updateMap)
}

// 更新圈子状态
func UpdateCircleByTitle(title string) (int64, error) {
	return mysql.UpdateCircleByTitle(title)

}

func GetCircleRankByType(ctx context.Context, isFree, isCharge bool) (circleList []model.RankCircle, listType int, err error) {
	var circleRank []model.Circle
	if isFree {
		circleRank, err = GetCircleRankFree(ctx)
		listType = LIST_TYPE_FREE
	} else if isCharge {
		circleRank, err = GetCircleRankCharge(ctx)
		listType = LIST_TYPE_CHARGE
	} else {
		circleRank, err = GetCircleRank(ctx)
		listType = LIST_TYPE_ALL
	}
	if err != nil {
		return
	}

	// 只取前50名
	if len(circleRank) > model.CIRCLE_RANK_LEN {
		circleRank = circleRank[:model.CIRCLE_RANK_LEN]
	}

	circleList = make([]model.RankCircle, 0)
	if len(circleRank) == 0 {
		return
	}

	//根据cids获取圈子信息
	cidList := make([]int, 0)
	for _, v := range circleRank {
		cidList = append(cidList, v.Id)
	}
	circleMap, err := GetCircleByCidList(cidList)
	if err != nil {
		return
	}

	uidList := make([]int, 0)
	for _, v := range circleMap {
		uidList = append(uidList, v.CircleOwnerId)
	}

	userMap, err := GetUserByUidList(uidList)
	if err != nil {
		return
	}

	for k, v := range circleRank {
		cid := v.Id
		joinNum := v.JoinNum

		circle, ok := circleMap[cid]
		if !ok {
			continue
		}

		uid := circle.CircleOwnerId
		user, ok := userMap[uid]
		if !ok {
			continue
		}

		priceText := ""
		if circle.Price > 0 {
			priceText = fmt.Sprintf("%d元", circle.Price)
		} else {
			priceText = "免费"
		}

		item := model.RankCircle{
			Id:            cid,
			Rank:          k + 1,
			Title:         circle.Title,
			Price:         circle.Price,
			PriceText:     priceText,
			CircleOwnerId: uid,
			OwnerName:     user.Name,
			JoinNum:       joinNum,
		}

		circleList = append(circleList, item)
	}

	return
}
