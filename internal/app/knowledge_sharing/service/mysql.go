package service

import "bishe/internal/app/knowledge_sharing/dao/mysql"

// 初始化数据库连接
func ServiceInitDB(dsn string) (err error) {
	err = mysql.DaoInitDB(dsn)
	return
}
