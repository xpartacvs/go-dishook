package dishook

import (
	"errors"
	"regexp"
)

type Url string
type Color uint

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
	Color       Color   `json:"color,omitempty"`
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
	ColorTrace Color = 3092790
	ColorDebug Color = 10170623
	ColorInfo  Color = 3581519
	ColorWarn  Color = 14327864
	ColorError Color = 13631488
	ColorPanic Color = 13631488
	ColorFatal Color = 13631488
)

func (u Url) MarshalJSON() ([]byte, error) {
	if err := u.validate(); err != nil {
		return nil, err
	}
	s := "\"" + string(u) + "\""
	return []byte(s), nil
}

func (u Url) validate() error {
	rgxUrl := regexp.MustCompile("(?i)^https?://.*")
	if !rgxUrl.MatchString(string(u)) {
		return errors.New("invalid url")
	}
	return nil
}
