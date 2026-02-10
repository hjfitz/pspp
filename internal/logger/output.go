package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

func NewLogger() zerolog.Logger {
	output := zerolog.ConsoleWriter{
		Out:           os.Stdout,
		NoColor:       true,
		PartsOrder:    []string{"level", "one", "two", "three", "message"},
		FieldsExclude: []string{"one", "two", "three"},
	}

	output.FormatLevel = func(i interface{}) string { return strings.ToUpper(fmt.Sprintf("%-6s", i)) }
	output.FormatFieldName = func(i interface{}) string { return fmt.Sprintf("%s:", i) }
	output.FormatPartValueByName = func(i interface{}, s string) string {
		var ret string
		switch s {
		case "one":
			ret = strings.ToUpper(fmt.Sprintf("%s", i))
		case "two":
			ret = strings.ToLower(fmt.Sprintf("%s", i))
		case "three":
			ret = strings.ToLower(fmt.Sprintf("(%s)", i))
		}
		return ret
	}
	log := zerolog.New(output)
	return log
}
