package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/whileW/enze-global"
	"github.com/whileW/enze-wx-api/model/request"
	"github.com/whileW/enze-wx-api/service/qiye_wx/provider"
	"net/http"
)

//企业微信服务商回调
func CallBack(c *gin.Context)  {
	req := &request.QiYeWxProviderCallBack{}
	c.ShouldBindJSON(req)
	result,err := provider.CallBack(req)
	global.GVA_LOG.Infow("企业微信服务商回调处理",
		"req",req,
		"resp",result,
		"err",err)
	if err != nil {
		c.String(http.StatusInternalServerError,"",err)
		return
	}
	c.String(http.StatusOK,result)
	return
}
