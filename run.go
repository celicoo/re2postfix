package re2postfix

import (
	"strings"

	"github.com/celicoo/re2postfix/internal/base"
	"github.com/celicoo/re2postfix/internal/container"
	"github.com/celicoo/re2postfix/internal/re2postfix"
)

// Run returns infix in postfix notation. All concatenation symbols (dots) are
// stripped out from infix prior to transformation.
// Run returns "" and an error if infix cannot be parsed.
func Run(infix string) (string, error) {
	infix = strings.ReplaceAll(infix, ".", "")
	{
		var (
			infix      = container.Slice[base.Character](infix)
			postfix, e = re2postfix.Run(&infix)
		)
		if e != nil {
			return "", e
		}
		return string(*postfix), e
	}
}
