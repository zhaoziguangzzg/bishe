package model

import (
	"time"
)

/*
-- knowledge_sharing.level_score definition

CREATE TABLE `level_score` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `cid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '圈子id',
  `score` int unsigned NOT NULL DEFAULT '0' COMMENT '等级分数',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '1删除，0未删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_uid_cid` (`uid`,`cid`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='等级表';
*/

// LevelScoreRecord 定义等级结构体
type LevelScoreRecord struct {
	Id int `json:"id" gorm:"column:id" mapstructure:"id"`
	//用户id
	Uid int `json:"uid" gorm:"column:uid" mapstructure:"uid"`
	//圈子id
	Cid int `json:"cid" gorm:"column:cid" mapstructure:"cid"`
	//分数
	Score int `json:"score" gorm:"column:score" mapstructture:"score"`
	//分数相关id
	RelateId int `json:"relateId" gorm:"column:relate_id" mapstructture:"relateId"`
	//分数增加类型
	Typei int `json:"type" gorm:"column:type" mapstructture:"type"`

	CreateAt    *time.Time `json:"createAt" gorm:"column:create_at" mapstructure:"-"`
	CreateAtStr string     `json:"-" gorm:"-" mapstructure:"createAt"`
	UpdateAt    *time.Time `json:"updateAt" gorm:"column:update_at" mapstructure:"-"`
	UpdateAtStr string     `json:"-" gorm:"-" mapstructure:"updateAt"`
	IsDeleted   int        `json:"isDeleted" gorm:"column:is_deleted" mapstructture:"isDeleted"`
}

const (
	LEVEL_SCORE_RECORD_TYPE_ESSAY int = 1 //发文章

	LEVEL_SCORE_RECORD_TYPE_LIKE    int = 2 //点赞
	LEVEL_SCORE_RECORD_TYPE_COLLECT int = 3 //收藏
	LEVEL_SCORE_RECORD_TYPE_COMMENT int = 4 //评论

	LEVEL_SCORE_RECORD_TYPE_LIKED     int = 5 //被点赞
	LEVEL_SCORE_RECORD_TYPE_COLLECTED int = 6 //被收藏
	LEVEL_SCORE_RECORD_TYPE_COMMENTED int = 7 //被评论

	LEVEL_SCORE_RECORD_TYPE_ESSENCE int = 8 //加精
	LEVEL_SCORE_RECORD_TYPE_WEEKLY  int = 9 //文章周刊
)

var LevelTypeScoreMap = map[int]int{
	LEVEL_SCORE_RECORD_TYPE_ESSAY:     100,
	LEVEL_SCORE_RECORD_TYPE_LIKE:      2,
	LEVEL_SCORE_RECORD_TYPE_COLLECT:   5,
	LEVEL_SCORE_RECORD_TYPE_COMMENT:   10,
	LEVEL_SCORE_RECORD_TYPE_LIKED:     3,
	LEVEL_SCORE_RECORD_TYPE_COLLECTED: 3,
	LEVEL_SCORE_RECORD_TYPE_COMMENTED: 3,
	LEVEL_SCORE_RECORD_TYPE_ESSENCE:   100,
	LEVEL_SCORE_RECORD_TYPE_WEEKLY:    100,
}

func GetScoreByType(typei int) int {
	return LevelTypeScoreMap[typei]
}

// 指定LevelScoreRecord对应的表名
func (LevelScoreRecord) TableName() string {
	return "level_score_record"
}
