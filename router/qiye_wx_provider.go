package router

import (
	"github.com/gin-gonic/gin"
	"github.com/whileW/enze-wx-api/api/v1"
)

func InitQyWxProviderRouter(Router *gin.RouterGroup) {
	QyWxProviderRouter := Router.Group("qiyewx/provider")
	{
		QyWxProviderRouter.POST("callback", v1.CallBack)    //企业微信服务商回调
		QyWxProviderRouter.GET("callback", v1.CallBack)     //企业微信服务商回调
	}
}
