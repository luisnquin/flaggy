package flaggy

import (
	"fmt"
	"strconv"
	"strings"
	"text/template"
)

type attribute int

// The possible state of text inside the application, either Bold, faint, italic or underline.
//
// These constants are called through the use of the Styler function.
const (
	reset attribute = iota

	FGBold
	FGFaint
	FGItalic
	FGUnderline
)

// The possible colors of text inside the application.
//
// These constants are called through the use of the Styler function.
const (
	FGBlack attribute = iota + 30
	FGRed
	FGGreen
	FGYellow
	FGBlue
	FGMagenta
	FGCyan
	FGWhite
)

// The possible background colors of text inside the application.
//
// These constants are called through the use of the Styler function.
const (
	BGBlack attribute = iota + 40
	BGRed
	BGGreen
	BGYellow
	BGBlue
	BGMagenta
	BGCyan
	BGWhite
)

const esc = "\033["

var ResetCode = fmt.Sprintf("%s%dm", esc, reset)

// Styler is a function that accepts multiple possible styling transforms from the state,
// color and background colors constants and transforms them into a templated string
// to apply those styles in the CLI.
//
// The returned styling function accepts a string that will be extended with
// the wrapping function's styling attributes.
func Styler(attrs ...attribute) func(interface{}) string {
	attrStrs := make([]string, len(attrs))
	for i, v := range attrs {
		attrStrs[i] = strconv.Itoa(int(v))
	}

	seq := strings.Join(attrStrs, ";")

	return func(v interface{}) string {
		end := ""
		s, ok := v.(string)
		if !ok || !strings.HasSuffix(s, ResetCode) {
			end = ResetCode
		}
		return fmt.Sprintf("%s%sm%v%s", esc, seq, v, end)
	}
}

func getTemplateStyleFuncs() template.FuncMap {
	return template.FuncMap{
		"black":     Styler(FGBlack),
		"red":       Styler(FGRed),
		"green":     Styler(FGGreen),
		"yellow":    Styler(FGYellow),
		"blue":      Styler(FGBlue),
		"magenta":   Styler(FGMagenta),
		"cyan":      Styler(FGCyan),
		"white":     Styler(FGWhite),
		"bgBlack":   Styler(BGBlack),
		"bgRed":     Styler(BGRed),
		"bgGreen":   Styler(BGGreen),
		"bgYellow":  Styler(BGYellow),
		"bgBlue":    Styler(BGBlue),
		"bgMagenta": Styler(BGMagenta),
		"bgCyan":    Styler(BGCyan),
		"bgWhite":   Styler(BGWhite),
		"bold":      Styler(FGBold),
		"faint":     Styler(FGFaint),
		"italic":    Styler(FGItalic),
		"underline": Styler(FGUnderline),
	}
}
