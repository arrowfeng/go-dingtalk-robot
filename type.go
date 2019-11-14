/*
 * File : type.go
 * Author : arrowfeng
 * Date : 2019/11/13
 */
package main

type PackageData interface {
	Package() map[string]interface{}
}

type Text struct {
	msgtype   string
	content   string
	atMobiles []string
	isAtAll   bool
}

var data map[string]interface{}

func init() {
	data = make(map[string]interface{}, 0)
}

func (this *Text) Package() map[string]interface{} {

	data["msgtype"] = this.msgtype
	data["text"] = map[string]string{"content": this.content}
	data["at"] = map[string]interface{}{"atMobiles": this.atMobiles, "isAtAll": this.isAtAll}

	return data

}

type Link struct {
	msgtype    string
	title      string
	text       string
	messageUrl string
	picUrl     string
}

func (this *Link) Package() map[string]interface{} {

	data["msgtype"] = this.msgtype
	data["link"] = map[string]string{
		"text":       this.text,
		"title":      this.title,
		"picUrl":     this.picUrl,
		"messageUrl": this.messageUrl,
	}

	return data
}

type MarkDown struct {
	msgtype   string
	title     string
	text      string
	atMobiles []string
	isAtAll   bool
}

func (this *MarkDown) Package() map[string]interface{} {

	data["msgtype"] = this.msgtype
	data["markdown"] = map[string]string{
		"title": this.title,
		"text":  this.text,
	}
	data["at"] = map[string]interface{}{"atMobiles": this.atMobiles, "isAtAll": this.isAtAll}

	return data
}

type ActionCard struct {
	msgtype        string
	title          string
	text           string
	independent    bool
	sTitle         []string
	sURL           []string
	btnOrientation string
	hideAvatar     string
}

func (this *ActionCard) Package() map[string]interface{} {

	data["msgtype"] = this.msgtype
	tmp := map[string]interface{}{
		"title": this.title,
		"text": this.text,
		"hideAvatar": this.hideAvatar,
		"btnOrientation": this.btnOrientation,
	}

	if this.independent {
		btns := make([]map[string]string, 0)
		for i, v := range this.sTitle {
			btns = append(btns, map[string]string{"title": v, "actionURL": this.sURL[i]})
		}
		tmp["btns"] = btns
	} else {
		tmp["singleTitle"] = this.sTitle[0]
		tmp["singleURL"] = this.sURL[0]
	}

	data["actionCard"] = tmp

	return data
}

type FeedCard struct {
	msgtype    string
	title      []string
	messageURL []string
	picURL     []string
}

func (this *FeedCard) Package() map[string]interface{} {

	data["msgtype"] = this.msgtype


	links := make([]map[string]string, 0)

	for i, v := range this.title {
		links = append(links, map[string]string{"title": v, "messageURL": this.messageURL[i], "picURL": this.picURL[i]})
	}
	data["feedCard"] = map[string]interface{}{"links": links}

	return data
}

type Response struct {
	errcode int32
	errmsg  string
}
