package dishook

import (
	"errors"
	"fmt"
	"net/url"
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
	Username  string  `json:"username,omitempty"`
	AvatarUrl Url     `json:"avatar_url,omitempty"`
	Content   string  `json:"content"`
	Tts       bool    `json:"tts,omitempty"`
	Embeds    []Embed `json:"embeds,omitempty"`
}

const (
	ColorTrace = 3092790
	ColorDebug = 10170623
	ColorInfo  = 3581519
	ColorWarn  = 14327864
	ColorError = 13631488
	ColorPanic = 13631488
	ColorFatal = 13631488
)

func (u Url) MarshalJSON() ([]byte, error) {
	if err := u.validate(); err != nil {
		return nil, err
	}
	s := fmt.Sprintf("\"%s\"", string(u))
	return []byte(s), nil
}

func (u Url) validate() error {
	resURL, err := url.ParseRequestURI(string(u))
	if err != nil {
		return err
	} else if resURL.Scheme != "https" {
		return errors.New("scheme must be: https")
	}

	return nil
}
