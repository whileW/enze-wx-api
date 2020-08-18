package model

import (
	"github.com/jinzhu/gorm"
	"github.com/whileW/enze-global"
)

//企业微信服务商
type QyWxProvider struct {
	gorm.Model
	// 每个服务商同时也是一个企业微信的企业，都有唯一的corpid。
	// 获取此信息可在服务商管理后台“应用开发”－“通用开发参数”可查看
	Corpid 					string 			`json:"corpid" gorm:"type:varchar(128)"`
	// 作为服务商身份的调用凭证，应妥善保管好该密钥，务必不能泄漏。
	ProviderSecret			string			`json:"provider_secret" gorm:"type:varchar(128)"`
	// 应用的唯一身份标识
	Suiteid 				string			`json:"suiteid" gorm:"type:varchar(128)"`
	// 调用身份密钥
	SuiteSecret				string			`json:"suite_secret" gorm:"type:varchar(128)"`
	// suite_ticket与suite_secret配套使用，用于获取suite_access_token。
	// suite_ticket由企业微信后台向登记的应用指令回调地址定期推送（10分钟），
	// 用于加强调用者身份确认（即使suite_secret泄露，也无法获取到suite_access_token）。
	// 若开发者丢失suite_ticket，除了等待定时推送的事件外，开发者也可以在管理端手动触发推送。
	SuiteTicket				string			`json:"suite_ticket" gorm:"type:varchar(128)"`

	//回调配置
	// Token用于计算签名
	// 由英文或数字组成且长度不超过32位的自定义字符串
	Token					string 			`json:"token" gorm:"type:varchar(32)"`
	// EncodingAESKey用于消息内容加密，
	// 由英文或数字组成且长度为43位的自定义字符串
	EncodingAESKey			string			`json:"encoding_aes_key" gorm:"type:varchar(43)"`
}

func (p *QyWxProvider)Save() error {
	db := global.GVA_DB.Get(WxApi)
	if err := db.Save(p).Error;err != nil{
		return err
	}
	return nil
}