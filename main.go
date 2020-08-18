package main

import (
	"github.com/gin-gonic/gin"
	"github.com/whileW/enze-global"
	"github.com/whileW/enze-global/initialize"
	"github.com/whileW/enze-global/utils/resp"
	"github.com/whileW/enze-wx-api/model"
	"github.com/whileW/enze-wx-api/router"
)

func main() {
	init_db_tables()
	init_server()
}

// 注册数据库表
func init_db_tables() {
	initialize.Db()
	db := global.GVA_DB.Get("wxapi")
	db.AutoMigrate(model.QyWxProvider{})
	global.GVA_LOG.Info("register table success")
}
//加载http监听
func init_server() {
	r := gin.Default()
	// 跨域
	r.Use(resp.Cors())
	//捕获异常
	r.Use(gin.Recovery())
	//挂载路由
	ApiGroup := r.Group("")
	router.InitQyWxProviderRouter(ApiGroup)				//企业微信服务商

	global.GVA_LOG.Error(r.Run(":"+global.GVA_CONFIG.SysSetting.HttpAddr))
}