package redis

import (
	"bishe/model"
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func GetCircleJoinNumKey() (key string) {
	key = "circle_join_num"
	return
}

func GetCircleJoinNumFreeKey() (key string) {
	key = "circle_join_num_free"
	return
}

func GetCircleJoinNumChargeKey() (key string) {
	key = "circle_join_num_charge"
	return
}

func ZincrCircleJoinNum(ctx context.Context, cid int, isFree bool) (score float64, err error) {
	key := GetCircleJoinNumKey()
	member := &redis.Z{
		Score:  1,
		Member: cid,
	}

	score, err = RedisClient.ZIncr(ctx, key, member).Result()

	if isFree {
		key = GetCircleJoinNumFreeKey()
		member = &redis.Z{
			Score:  1,
			Member: cid,
		}

		score, err = RedisClient.ZIncr(ctx, key, member).Result()
	} else {
		key = GetCircleJoinNumChargeKey()
		member = &redis.Z{
			Score:  1,
			Member: cid,
		}

		score, err = RedisClient.ZIncr(ctx, key, member).Result()
	}

	return
}

func ZrevrangeCircleJoinNum(ctx context.Context) (circleSlice []model.Circle, err error) {
	key := GetCircleJoinNumKey()

	start := int64(0)
	stop := int64(model.CIRCLE_RANK_LEN)

	zslice, err := RedisClient.ZRevRangeWithScores(ctx, key, start, stop).Result()
	if err != nil {
		return
	}

	for _, v := range zslice {
		cidStr, _ := v.Member.(string)
		cid, _ := strconv.Atoi(cidStr)
		if cid == 0 {
			continue
		}
		circleJoinNum := model.Circle{
			Id:      cid,
			JoinNum: int(v.Score),
		}
		circleSlice = append(circleSlice, circleJoinNum)
	}

	return
}

func ZrevrangeCircleJoinNumFree(ctx context.Context) (circleSlice []model.Circle, err error) {
	key := GetCircleJoinNumFreeKey()

	start := int64(0)
	stop := int64(model.CIRCLE_RANK_LEN)

	zslice, err := RedisClient.ZRevRangeWithScores(ctx, key, start, stop).Result()
	if err != nil {
		return
	}

	for _, v := range zslice {
		cidStr, _ := v.Member.(string)
		cid, _ := strconv.Atoi(cidStr)
		if cid == 0 {
			continue
		}
		circleJoinNum := model.Circle{
			Id:      cid,
			JoinNum: int(v.Score),
		}
		circleSlice = append(circleSlice, circleJoinNum)
	}

	return
}

func ZrevrangeCircleJoinNumCharge(ctx context.Context) (circleSlice []model.Circle, err error) {
	key := GetCircleJoinNumChargeKey()

	start := int64(0)
	stop := int64(model.CIRCLE_RANK_LEN - 1)

	zslice, err := RedisClient.ZRevRangeWithScores(ctx, key, start, stop).Result()
	if err != nil {
		return
	}

	for _, v := range zslice {
		cidStr, _ := v.Member.(string)
		cid, _ := strconv.Atoi(cidStr)
		if cid == 0 {
			continue
		}

		circleJoinNum := model.Circle{
			Id:      cid,
			JoinNum: int(v.Score),
		}
		circleSlice = append(circleSlice, circleJoinNum)
	}

	return
}
