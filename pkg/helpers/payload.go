package helpers

import (
	"fmt"
	"net/url"
)

func AddPayloadValue[T string | bool | int](p *url.Values, key string, value *T, defaultValue *T) {
	switch t := any(value).(type) {
	case *bool:
		if t == nil {
			if defaultValue == nil {
				return
			}
			v := BoolToInt(*any(defaultValue).(*bool))
			p.Set(key, fmt.Sprintf("%d", v))
			return
		}
		v := BoolToInt(*t)
		p.Set(key, fmt.Sprintf("%d", v))
		return
	case *int:
		if t == nil {
			if defaultValue == nil {
				return
			}
			p.Set(key, fmt.Sprintf("%d", *any(defaultValue).(*int)))
			return
		}
		p.Set(key, fmt.Sprintf("%d", *t))
		return
	case *string:
		if t == nil {
			if defaultValue == nil {
				return
			}
			p.Set(key, *any(defaultValue).(*string))
			return
		}
		p.Set(key, *t)
		return
	default:
		panic("'AddPayloadValue' value parameter type not supported")
	}
}
