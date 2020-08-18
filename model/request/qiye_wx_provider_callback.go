package request

type QiYeWxProviderCallBack struct {
	// 企业微信加密签名，msg_signature计算结合了企业填写的token、请求中的timestamp、nonce、加密的消息体。
	// 签名计算方法参考 https://work.weixin.qq.com/api/doc/90000/90139/90968#%E6%B6%88%E6%81%AF%E4%BD%93%E7%AD%BE%E5%90%8D%E6%A0%A1%E9%AA%8C
	MsgSignature 		string			`json:"msg_signature"`
	// 时间戳。与nonce结合使用，用于防止请求重放攻击。
	Timestamp			string			`json:"timestamp"`
	// 随机数。与timestamp结合使用，用于防止请求重放攻击。
	Nonce				string			`json:"nonce"`
	// 加密的字符串。需要解密得到消息内容明文，解密后有random、msg_len、msg、receiveid四个字段，
	// 其中msg即为消息内容明文
	// https://work.weixin.qq.com/api/doc/90000/90139/90968#%E5%AF%86%E6%96%87%E8%A7%A3%E5%AF%86%E5%BE%97%E5%88%B0msg%E7%9A%84%E8%BF%87%E7%A8%8B
	Echostr				string			`json:"echostr"`
	// 企业微信的CorpID，当为第三方应用回调事件时，CorpID的内容为suiteid
	ToUserName			string			`json:"to_user_name"`
	// 接收的应用id，可在应用的设置页面获取。仅应用相关的回调会带该字段。
	AgentId				string			`json:"agent_id"`
}

