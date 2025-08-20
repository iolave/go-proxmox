package api

import (
	"fmt"
	"reflect"

	"github.com/ggicci/httpin/core"
	"github.com/iolave/go-errors"
)

type directiveQueryNFormatter struct{}

func (f *directiveQueryNFormatter) Decode(rtm *core.DirectiveRuntime) error {
	return errors.NewWithName("http_in_error", "query_n decoding is not supported")
}

func (f *directiveQueryNFormatter) Encode(rtm *core.DirectiveRuntime) error {
	if len(rtm.Directive.Argv) != 1 {
		return errors.NewWithName("http_in_error", "query_n requires exactly one argument")
	}

	prefix := rtm.Directive.Argv[0]

	if rtm.Value.Type().Kind() != reflect.Slice {
		return errors.NewWithName("http_in_error", "not a slice")
	}

	len := rtm.Value.Len()
	currentValue := rtm.Value.Slice(0, len)

	for i := range len {
		rtm.GetRequestBuilder().Query.Set(
			fmt.Sprintf("%s[%d]", prefix, i),
			currentValue.Index(i).String(),
		)
	}

	return nil
}

func httpinInit() {
	core.RegisterDirective("query_n", &directiveQueryNFormatter{})
}
