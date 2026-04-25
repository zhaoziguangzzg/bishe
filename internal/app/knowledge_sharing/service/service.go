package service

import "go.uber.org/zap"

func ServiceInit() (err error) {
	err = LoggerInit()
	if err != nil {
		return
	}
	//Logger, err = zap.NewDevelopment()

	err = LoadConfig()
	if err != nil {
		Logger.Error("LoadConfig err", zap.Error(err))
		return
	}

	// 初始化数据库
	err = ServiceInitDB(Cfg.Database.Dsn)
	if err != nil {
		Logger.Error("InitDB err", zap.Error(err))
		return
	}

	//初始化 Redis
	ServiceInitRedis(Cfg.Redis.Addr, Cfg.Redis.Password, Cfg.Redis.DB)

	// //初始化 kafka
	// err = ServiceInitKafka()
	// if err != nil {
	// 	Logger.Error("InitKafka err", zap.Error(err))
	// 	return
	// }
	// defer Closekafka()

	return
}
