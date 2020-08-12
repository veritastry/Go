package login

import (
	"encoding/json"
	config "gin-demo/config"
	"gin-demo/defs"
	"gin-demo/logger"
	"gin-demo/utils"
	"net/url"
)

var (
	qywxConf = config.Config().QyWechat
	code     string
)

// GetSuiteTokenRes get suite token resp
type GetSuiteTokenRes struct {
	// Errcode          int    `json:errcode`
	// Errmsg           string `json:errmsg`
	SuiteAccessToken string `json:"suite_access_token"`
	ExpiresIn        int    `json:"expires_in"`
}

//QywxThirdParty third party login
func QywxThirdParty(queryMap url.Values) string {
	//获取url中的参数
	code = queryMap.Get("code")
	logger.Debug(queryMap.Get("code"))
	//构造post请求参数
	postMap := map[string]string{
		"suite_id":     qywxConf.ThirdParty.SuiteID,
		"suite_secret": qywxConf.ThirdParty.SuitSecret,
		"suite_ticket": "D90wiEwDKQ4_ZiLrt0voBrw5lEWvV22DCgwh_7Xwgv0sUHznOVTBOPZEGk9zLE-R",
	}

	//解析参数为JSON
	paramJSON, err := json.Marshal(postMap)
	if err != nil {
		//handle error
		logger.Error(defs.ParseToJSONErr, err)
		return ""
	}
	logger.Notice("paramJSON: ", string(paramJSON))

	//  获取第三方应用凭证（suite_access_token)
	//  POST
	//  https://qyapi.weixin.qq.com/cgi-bin/service/get_suite_token

	resp, err := utils.HTTPPost(qywxConf.ThirdParty.SuiteTokenURL, string(paramJSON))
	if err != nil {
		logger.Error(defs.CallFuncErr, err)
		return ""
	}
	logger.Debug(resp)
	userInfo, err := getUserInfo(resp)
	if err != nil {
		logger.Error(defs.CallFuncErr, err)
		return ""
	}
	return userInfo
}

// 服务商服务器以code换取 用户唯一标识 userid 、用户所在企业corpid 和 会话密钥 session_key。
// https://qyapi.weixin.qq.com/cgi-bin/service/miniprogram/jscode2session?suite_access_token=SUITE_ACCESS_TOKEN&js_code=CODE&grant_type=authorization_code
func getUserInfo(resp string) (string, error) {
	getSuiteToken := GetSuiteTokenRes{}
	json.Unmarshal([]byte(resp), &getSuiteToken)
	// suite_access_token 第三方应用凭证
	// js_code		登录时获取的 code
	// grant_type	固定为authorization_code
	reqTokenParam := url.Values{
		"suite_access_token": {getSuiteToken.SuiteAccessToken},
		"js_code":            {code},
		"grant_type":         {"authorization_code"},
	}

	encodeURL, err := utils.EncodeURL(qywxConf.ThirdParty.TpURL, reqTokenParam)
	logger.Debug("url: ", encodeURL)
	if err != nil {
		logger.Error(defs.CallFuncErr, err)
		return "", err
	}
	res, err := utils.HTTPGet(encodeURL)
	if err != nil {
		logger.Error(defs.CallFuncErr, err)
		return "", err
	}
	return res, err
}
