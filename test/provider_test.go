package test

import (
	_ "github.com/whileW/enze-wx-api/test/app"
	"github.com/whileW/enze-global/initialize"
	"github.com/whileW/enze-wx-api/service/qiye_wx/provider"
	"testing"
)

func init()  {
	initialize.Db()
}

//获取服务商凭证
// go test provider_test.go -v -test.run TestGetProviderToken
func TestGetProviderToken(t *testing.T)  {
	resp,err := provider.GetProviderTokenById(1)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(resp)
}

// 获取第三方应用凭证
// go test provider_test.go -v -test.run TestGetSuiteToken
func TestGetSuiteToken(t *testing.T)  {
	resp,err := provider.GetSuiteTokenById(1)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(resp)
}