package re2postfix

import (
	"testing"

	"github.com/celicoo/re2postfix/errors"
	"github.com/celicoo/re2postfix/internal/base"
	"github.com/celicoo/re2postfix/internal/container"

	"github.com/celicoo/ttest"
)

func TestRun(t *testing.T) {
	ttest.
		NewTest(t).
		SetSubject(Run).
		Run([]ttest.TestCase{
			// no loop
			{
				Give: []any{
					&container.Slice[base.Character]{},
				},
				Want: []any{
					&container.Slice[base.Character]{},
					nil,
				},
			},
			// default
			{
				Give: []any{
					&container.Slice[base.Character]{
						base.Character('a'),
					},
				},
				Want: []any{
					&container.Slice[base.Character]{
						base.Character('a'),
					},
					nil,
				},
			},
			// default and .
			{
				Give: []any{
					&container.Slice[base.Character]{
						base.Character('a'),
						base.Character('b'),
					},
				},
				Want: []any{
					&container.Slice[base.Character]{
						base.Character('a'),
						base.Character('b'),
						base.Character('.'),
					},
					nil,
				},
			},
			// |
			{
				Give: []any{
					&container.Slice[base.Character]{
						base.Character('a'),
						base.Character('|'),
						base.Character('b'),
					},
				},
				Want: []any{
					&container.Slice[base.Character]{
						base.Character('a'),
						base.Character('b'),
						base.Character('|'),
					},
					nil,
				},
			},
			// *
			{
				Give: []any{
					&container.Slice[base.Character]{
						base.Character('*'),
					},
				},
				Want: []any{
					&container.Slice[base.Character]{},
					errors.Syntax("nothing to repeat"),
				},
			},
			{
				Give: []any{
					&container.Slice[base.Character]{
						base.Character('a'),
						base.Character('*'),
					},
				},
				Want: []any{
					&container.Slice[base.Character]{
						base.Character('a'),
						base.Character('*'),
					},
					nil,
				},
			},
			// +
			{
				Give: []any{
					&container.Slice[base.Character]{
						base.Character('*'),
					},
				},
				Want: []any{
					&container.Slice[base.Character]{},
					errors.Syntax("nothing to repeat"),
				},
			},
			// ?
			{
				Give: []any{
					&container.Slice[base.Character]{
						base.Character('*'),
					},
				},
				Want: []any{
					&container.Slice[base.Character]{},
					errors.Syntax("nothing to repeat"),
				},
			},
			// (
			{
				Give: []any{
					&container.Slice[base.Character]{
						base.Character('('),
					},
				},
				Want: []any{
					&container.Slice[base.Character]{},
					errors.Syntax("unmatched opening parenthesis"),
				},
			},
			{
				Give: []any{
					&container.Slice[base.Character]{
						base.Character('('),
						base.Character('a'),
						base.Character(')'),
					},
				},
			},

			// )
			{
				Give: []any{
					&container.Slice[base.Character]{
						base.Character(')'),
					},
				},
				Want: []any{
					&container.Slice[base.Character]{},
					errors.Syntax("unmatched closing parenthesis"),
				},
			},
		})
}
