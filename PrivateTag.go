package goi18n

import (
	"strings"
)

type PrivateTag interface {
	BuildPrivateTagString(bool, bool, *strings.Builder) error
}

type SimplePrivateTag struct {
	Rendition string
}

func(tag *SimplePrivateTag) BuildPrivateTagString(
	initalHyphen bool,
	withX bool,
	builder *strings.Builder,
) (err error) {
	//TODO
	return
}

var _ PrivateTag = &SimplePrivateTag{}
