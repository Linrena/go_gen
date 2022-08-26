package util

import (
	"bytes"
	"strings"

	"github.com/Linrena/go_gen/internal/util/consts"
	"github.com/fatih/camelcase"
)

func InitialismsWithCamel(camelStr string) string {
	splits := camelcase.Split(camelStr)
	buf := bytes.NewBuffer(nil)
	for _, word := range splits {
		upper := strings.ToUpper(word)
		if consts.CommonInitialisms[upper] {
			buf.WriteString(upper)
		} else {
			buf.WriteString(word)
		}
	}
	return buf.String()
}
