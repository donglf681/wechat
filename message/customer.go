package message

import (
	"fmt"

	"github.com/trrtly/wechat/context"
	"github.com/trrtly/wechat/util"
)

const (
	customerMsgURL = "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s"
)

//CustomerMsg 客服消息
type CustomerMsg struct {
	*context.Context
}

//reqCustomerText 发送文本消息数据
type reqCustomerTextMsg struct {
	Touser  string          `json:"touser"`
	Msgtype string          `json:"msgtype"`
	Text    reqCustomerText `json:"text"`
}

//reqCustomerText 文本消息数据
type reqCustomerText struct {
	Content string `json:"content"`
}

//reqCustomerImage 发送图片消息数据
type reqCustomerImageMsg struct {
	Touser  string           `json:"touser"`
	Msgtype string           `json:"msgtype"`
	Image   reqCustomerImage `json:"image"`
}

//reqCustomerImage 图片消息数据
type reqCustomerImage struct {
	MediaID string `json:"media_id"`
}

//NewCustomerMsg 实例化
func NewCustomerMsg(context *context.Context) *CustomerMsg {
	return &CustomerMsg{Context: context}
}

//SendImageMsg 发送图片消息
func (cus *CustomerMsg) SendImageMsg(touser string, mediaID string) error {
	accessToken, err := cus.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf(customerMsgURL, accessToken)
	reqImage := &reqCustomerImageMsg{
		Touser:  touser,
		Msgtype: "image",
		Image: reqCustomerImage{
			MediaID: mediaID,
		},
	}

	response, err := util.PostJSON(uri, reqImage)
	if err != nil {
		return err
	}

	return util.DecodeWithCommonError(response, "SetMenu")
}

//SendTextMsg 发送文本消息
func (cus *CustomerMsg) SendTextMsg(touser string, context string) error {
	accessToken, err := cus.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf(customerMsgURL, accessToken)
	reqImage := &reqCustomerTextMsg{
		Touser:  touser,
		Msgtype: "text",
		Text: reqCustomerText{
			Content: context,
		},
	}

	response, err := util.PostJSON(uri, reqImage)
	if err != nil {
		return err
	}

	return util.DecodeWithCommonError(response, "SetMenu")
}
