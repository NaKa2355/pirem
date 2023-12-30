package domain

type ButtonTag string

func NewButtonTag(tag string) ButtonTag {
	return ButtonTag(tag)
}
