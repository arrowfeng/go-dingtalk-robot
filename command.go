/*
 * File : command.go
 * Author : arrowfeng
 * Date : 2019/11/13
 */
package main

import (
	"bufio"
	"errors"
	"github.com/urfave/cli"
	"os"
	"strings"
)

var text Text
var link Link
var md MarkDown
var actionCard ActionCard
var feedCard FeedCard

var textCommand = cli.Command{
	Name:      "text",
	Usage:     "to send context of text type",
	Action:    textAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "content, c",
			Usage:       "to set content of sending",
			EnvVar:      "DINGTALK_CONTENT",
			Destination: &text.content,
			Required:    false,
			Value:       "",
		},
		cli.StringFlag{
			Name:        "file, f",
			Usage:       "to set content of sending from file",
			EnvVar:      "DINGTALK_FILEPATH",
			Required:    false,
			Hidden:      false,
		},
		cli.StringSliceFlag{
			Name:     "at, a",
			Usage:    "to @ someone",
			EnvVar:   "DINGTALK_ATMOBILES",
			Required: false,
		},
		cli.BoolFlag{
			Name:        "all",
			Usage:       "@ all people",
			EnvVar:      "DINGTALK_ISATALL",
			Required:    false,
			Destination: &text.isAtAll,
		},
	},
}

var linkCommand = cli.Command{
	Name:      "link",
	Usage:     "to send context of link type",
	Action:    linkAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "title, t",
			Usage:       "to set title of sending",
			EnvVar:      "DINGTALK_TITLE",
			Required:    true,
			Value:       "",
			Destination: &link.title,
		},
		cli.StringFlag{
			Name:        "content, c",
			Usage:       "to set content of sending",
			EnvVar:      "DINGTALK_CONTENT",
			Destination: &link.text,
			Required:    false,
			Value:       "",
		},
		cli.StringFlag{
			Name:        "file, f",
			Usage:       "to set content of sending from file",
			EnvVar:      "DINGTALK_FILEPATH",
			Required:    false,
		},
		cli.StringFlag{
			Name:     "purl, p",
			Usage:    "to set picture of sending",
			EnvVar:   "DINGTALK_PICURL",
			Required: false,
			Destination: &link.picUrl,
		},
		cli.StringFlag{
			Name:        "murl, m",
			Usage:       "to set link content of sending",
			EnvVar:      "DINGTALK_MESSAGEURL",
			Required:    true,
			Destination: &link.messageUrl,
		},
	},
}

var markDownCommand = cli.Command{
	Name:      "md",
	Usage:     "to send context of markdown type",
	Action:    markDownAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "content, c",
			Usage:       "to set content of sending",
			EnvVar:      "DINGTALK_CONTENT",
			Destination: &md.text,
			Required:    false,
			Value:       "",
		},
		cli.StringFlag{
			Name:        "file, f",
			Usage:       "to set content of sending from file",
			EnvVar:      "DINGTALK_FILEPATH",
			Required:    false,
			Hidden:      false,
		},
		cli.StringFlag{
			Name:        "title, t",
			Usage:       "to set title of sending context",
			EnvVar:      "DINGTALK_TITLE",
			Required:    true,
			Destination: &md.title,
		},
		cli.StringSliceFlag{
			Name:     "at, a",
			Usage:    "to @ someone",
			EnvVar:   "DINGTALK_ATMOBILES",
			Required: false,
		},
		cli.BoolFlag{
			Name:        "all",
			Usage:       "@ all people",
			EnvVar:      "DINGTALK_ISATALL",
			Required:    false,
			Destination: &md.isAtAll,
		},
	},
}

var actionCardCommand = cli.Command{
	Name:      "ac",
	Usage:     "to send context of actionCard type",
	Action:    actionCardAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "content, c",
			Usage:       "to set content of sending",
			EnvVar:      "DINGTALK_CONTENT",
			Destination: &actionCard.text,
			Required:    false,
			Value:       "",
		},
		cli.StringFlag{
			Name:        "file, f",
			Usage:       "to set content of sending from file",
			EnvVar:      "DINGTALK_FILEPATH",
			Required:    false,
			Hidden:      false,
		},
		cli.StringFlag{
			Name:        "title, t",
			Usage:       "to set title of sending context",
			EnvVar:      "DINGTALK_TITLE",
			Required:    true,
			Destination: &actionCard.title,
		},
		cli.BoolFlag{
			Name:        "avatar, a",
			Usage:       "whether hide avatar or not",
			EnvVar:      "DINGTALK_ISHIDEAVATER",
		},
		cli.BoolFlag{
			Name:        "orientation",
			Usage:       "to set button orientation",
			EnvVar:      "DINGTALK_BTORIENTATION",
		},
		cli.BoolFlag{
			Name:        "m",
			Usage:       "whether enable independent jump or not",
			EnvVar:      "DINGTALK_ISINDEPENDENT",
			Required:    false,
			Destination: &actionCard.independent,
		},
		cli.StringSliceFlag{
			Name:      "stitle",
			Usage:     "to set button title",
			EnvVar:    "DINGTALK_BTTITLE",
			Required:  true,
		},
		cli.StringSliceFlag{
			Name:      "surl",
			Usage:     "jump to specific url when you click button",
			EnvVar:    "DINGTALK_BTURL",
			Required:  true,
		},
	},
}

var feedCardCommand = cli.Command{
	Name:      "fc",
	Usage:     "to send context of feedCard type",
	Action:    feedCardAction,
	Flags: []cli.Flag{
		cli.StringSliceFlag{
			Name:        "title",
			Usage:       "to set content title",
			EnvVar:      "DINGTALK_TITLE",
			Required:    false,
		},
		cli.StringSliceFlag{
			Name:        "murl",
			Usage:       "to set message url",
			EnvVar:      "DINGTALK_MESSAGEURL",
			Required:    false,
		},
		cli.StringSliceFlag{
			Name:     "purl",
			Usage:    "to set picture url",
			EnvVar:   "DINGTALK_PICTUREURL",
			Required: false,
		},
		cli.StringFlag{
			Name:        "file, f",
			Usage:       "load content from specific file",
			EnvVar:      "DINGTALK_FILEPATH",
			Required:    false,
		},
	},
}

func feedCardAction(c *cli.Context) error {

	title := c.StringSlice("title")
	murl := c.StringSlice("murl")
	purl := c.StringSlice("purl")

	if len(title) != len(murl) || len(title) != len(purl) || len(murl) != len(purl) {
		return errors.New("[error] the number of title option must be consistent with murl as well as purl")
	}

	if len(title) == 0 {
		return errors.New("null")
	}

	feedCard.msgtype = "feedCard"
	feedCard.title = title
	feedCard.messageURL = murl
	feedCard.picURL = purl

	return send(feedCard.Package())
}

func actionCardAction(c *cli.Context) error {

	filepath := c.String("file")

	if actionCard.text == "" && filepath == "" {
		return errors.New("[error] -c or -f should be provided but not")
	}

	if actionCard.text == "" {
		sb, err := readLine(filepath)
		if err != nil {
			return err
		}
		actionCard.text = sb.String()
	}

	actionCard.msgtype = "actionCard"

	if c.Bool("avatar") {
		actionCard.hideAvatar = "1"
	} else {
		actionCard.hideAvatar = "0"
	}

	if c.Bool("orientation") {
		actionCard.btnOrientation = "1"
	} else {
		actionCard.btnOrientation = "0"
	}

	actionCard.sTitle = c.StringSlice("stitle")
	actionCard.sURL =  c.StringSlice("surl")

	if len(actionCard.sTitle) != len(actionCard.sURL) {
		return errors.New("[error] the number of stitle option must be consistent with surl")
	}

	return send(actionCard.Package())
}

func markDownAction(c *cli.Context) error {

	filepath := c.String("file")

	if md.text == "" && filepath == "" {
		return errors.New("[error] -c or -f should be provided but not")
	}

	if md.text == "" {
		sb, err := readLine(filepath)
		if err != nil {
			return err
		}
		md.text = sb.String()
	}

	md.msgtype = "markdown"
	md.atMobiles = c.StringSlice("at")

	return send(md.Package())
}

func linkAction(c *cli.Context) error {

	filepath := c.String("file")

	if link.text == "" {
		sb, err := readLine(filepath)
		if err != nil {
			return err
		}
		link.text = sb.String()
	}

	link.msgtype = "link"

	return send(link.Package())
}

func textAction(c *cli.Context) error {

	filepath := c.String("file")

	if text.content == "" && filepath == "" {
		return errors.New("[error] -c or -f should be provided but not")
	}

	if text.content == "" {
		sb, err := readLine(filepath)
		if err != nil {
			return err
		}
		text.content = sb.String()
	}

	text.msgtype = "text"
	text.atMobiles = c.StringSlice("at")

	return send(text.Package())
}

func readLine(filepath string) (strings.Builder, error) {

	var sb strings.Builder

	f, err := os.Open(filepath)

	if err != nil {
		return sb, err
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		sb.WriteString(s.Text())
		sb.WriteString("\n")
	}

	err = s.Err()

	if err != nil {
		return sb, err
	}

	return sb, nil
}
