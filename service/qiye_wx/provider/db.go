package provider

import (
	"github.com/whileW/enze-global"
	"github.com/whileW/enze-wx-api/model"
)

func GetInfoById(id int) (*model.QyWxProvider,error) {
	db := global.GVA_DB.Get(model.WxApi)
	provider := &model.QyWxProvider{}
	if err := db.Model(provider).First(provider,"id = ?",id).Error;err != nil{
		return nil,err
	}
	return provider,nil
}

func GetInfoBySuiteId(suite_id string) (*model.QyWxProvider,error) {
	db := global.GVA_DB.Get(model.WxApi)
	provider := &model.QyWxProvider{}
	if err := db.Model(provider).First(provider,"suiteid = ?",suite_id).Error;err != nil{
		return nil,err
	}
	return provider,nil
}

func Update()  {
	db := global.GVA_DB.Get(model.WxApi)
	
}