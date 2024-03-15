package goi18n

import (
	"fmt"
	"errors"
	"strings"
	"github.com/UncleSniper/goutil"
)

// https://datatracker.ietf.org/doc/html/rfc5646.html#section-2.2.6

type Extension interface {
	fmt.Stringer
	Singleton() byte
	Subtags() []ExtSubtag
	BuildExtString(bool, *strings.Builder) error
}

func ExtString(ext Extension) string {
	var builder strings.Builder
	if ext != nil {
		ext.BuildExtString(false, &builder)
	}
	return builder.String()
}

type ExtSubtag interface {
	BuildExtSubtagString(bool, *strings.Builder) error
}

func IsValidExtensionSingleton(b byte) bool {
	if b >= 'A' && b <= 'Z' {
		return b != 'X'
	} else if b >= 'a' && b <= 'z' {
		return b != 'x'
	} else {
		return b >= '0' && b <= '9'
	}
}

type SimpleExtension struct {
	SingletonByte byte
	SubtagsSlice []ExtSubtag
}

func(ext *SimpleExtension) String() string {
	return ExtString(ext)
}

func(ext *SimpleExtension) Singleton() byte {
	if ext == nil {
		return 0
	} else {
		return ext.SingletonByte
	}
}

func(ext *SimpleExtension) Subtags() []ExtSubtag {
	if ext == nil || len(ext.SubtagsSlice) == 0 {
		return nil
	}
	return append([]ExtSubtag(nil), ext.SubtagsSlice...)
}

func(ext *SimpleExtension) BuildExtString(initalHyphen bool, builder *strings.Builder) (err error) {
	if ext == nil {
		if !initalHyphen {
			err = goutil.NewNilTargetError(&SimpleExtension{}, "BuildExtString")
		}
		return
	}
	if !IsValidExtensionSingleton(ext.SingletonByte) {
		err = errors.New(fmt.Sprintf("Extension subtag singleton byte %d is invalid", ext.SingletonByte))
		return
	}
	if initalHyphen {
		builder.WriteRune('-')
	}
	builder.WriteRune(rune(ext.SingletonByte))
	hadSub := false
	for _, sub := range ext.SubtagsSlice {
		if sub != nil {
			before := builder.Len()
			err = sub.BuildExtSubtagString(true, builder)
			if err != nil {
				return
			}
			if builder.Len() > before {
				hadSub = true
			}
		}
	}
	if !hadSub {
		err = errors.New(fmt.Sprintf("Extension with singleton byte %d has no subtags", ext.SingletonByte))
	}
	return
}

type SimpleExtSubtag struct {
	Rendition string
}

func(tag *SimpleExtSubtag) BuildExtSubtagString(initalHyphen bool, builder *strings.Builder) (err error) {
	if tag == nil {
		if !initalHyphen {
			err = goutil.NewNilTargetError(&SimpleExtSubtag{}, "BuildExtSubtagString")
		}
		return
	}
	if len(tag.Rendition) == 0 {
		if !initalHyphen {
			err = errors.New("Extension subtag has empty rendition")
		}
		return
	}
	if initalHyphen {
		builder.WriteRune('-')
	}
	if len(tag.Rendition) > 1 && (tag.Rendition[0] == 'x' || tag.Rendition[0] == 'X') && tag.Rendition[1] == '-' {
		err = errors.New(fmt.Sprintf("Extension subtag '%s' starts with 'x-'", tag.Rendition))
		return
	}
	builder.WriteString(tag.Rendition)
	return
}

var _ ExtSubtag = &SimpleExtSubtag{}
