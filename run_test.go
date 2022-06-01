package re2postfix

import (
	"testing"

	"github.com/celicoo/ttest"
)

func TestRun(t *testing.T) {
	ttest.NewTest(t).
		SetSubject(Run).
		Run([]ttest.TestCase{
			{
				Give: []any{
					". . ..",
				},
				Want: []any{
					"",
					nil,
				},
			},
			//{
			//	Give: []any{
			//		"(",
			//	},
			//	Want: []any{
			//		"",
			//		errors.MissingClosingParenthesis,
			//	},
			//},
			//{
			//	Give: []any{
			//		")",
			//	},
			//	Want: []any{
			//		"",
			//		errors.MissingOpeningParenthesis,
			//	},
			//},
			//{
			//	Give: []any{
			//		"ab",
			//	},
			//	Want: []any{
			//		"ab.",
			//		nil,
			//	},
			//},
		})
}
