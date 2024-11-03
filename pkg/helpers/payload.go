package helpers

import (
	"fmt"
	"net/url"
)

func AddPayloadValue(p *url.Values, key string, value any) {
	switch t := value.(type) {
	case bool:
		p.Set(key, string(BoolToInt(t)))
		return
	case *bool:
		if t == nil {
			return
		}
		v := BoolToInt(*t)
		p.Set(key, fmt.Sprintf("%d", v))
		return
	case int:
		p.Set(key, fmt.Sprintf("%d", t))
		return
	case *int:
		if t == nil {
			return
		}
		p.Set(key, fmt.Sprintf("%d", *t))
		return
	case string:
		p.Set(key, t)
		return
	case *string:
		if t == nil {
			return
		}
		p.Set(key, *t)
		return
	default:
		panic("'AddPayloadValue' value parameter type not supported")
	}
}
