package dishook

import (
	"errors"
	"fmt"
	"regexp"
)

type Url string

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

type Image struct {
	ImageUrl Url `json:"url,omitempty"`
}

type Icon struct {
	IconUrl Url `json:"icon_url,omitempty"`
}

type Footer struct {
	Text string `json:"text"`
	Icon
}

type Author struct {
	Name string `json:"name"`
	Image
	Icon
}

type Embed struct {
	Author      Author  `json:"author,omitempty"`
	Title       string  `json:"title,omitempty"`
	Url         Url     `json:"url,omitempty"`
	Description string  `json:"description,omitempty"`
	Color       uint    `json:"color,omitempty"`
	Fields      []Field `json:"fields,omitempty"`
	Thumbnail   Image   `json:"thumbnail,omitempty"`
	Image       Image   `json:"image,omitempty"`
	Footer      Footer  `json:"footer,omitempty"`
}

type Payload struct {
	webhookUrl Url     `json:"-"`
	Username   string  `json:"username,omitempty"`
	AvatarUrl  Url     `json:"avatar_url,omitempty"`
	Content    string  `json:"content"`
	Tts        bool    `json:"tts,omitempty"`
	Embeds     []Embed `json:"embeds,omitempty"`
}

func (u Url) MarshalJSON() ([]byte, error) {
	if err := u.validate(); err != nil {
		return nil, err
	}
	s := fmt.Sprintf("\"%s\"", string(u))
	return []byte(s), nil
}

func (u Url) validate() error {
	rgxUrl := regexp.MustCompile("^https?://.*")
	if !rgxUrl.MatchString(string(u)) {
		return errors.New("invalid url")
	}
	return nil
}

func (p Payload) Send() ([]byte, error) {
	return Send(string(p.webhookUrl), p)
}

func (p *Payload) SetWebhookUrl(url string) error {
	u := Url(url)
	if err := u.validate(); err != nil {
		return err
	}
	p.webhookUrl = u
	return nil
}
