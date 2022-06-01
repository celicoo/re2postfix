package container

import (
	"testing"

	"github.com/celicoo/re2postfix/internal/base"

	"github.com/celicoo/ttest"
)

func TestSlice_Pop(t *testing.T) {
	ttest.NewTest(t).
		SetSubject((*Slice[base.Character]).Pop).
		Run([]ttest.TestCase{
			{
				Give: []any{
					&Slice[base.Character]{},
				},
				Want: []any{
					base.Character(0),
				},
			},
			{
				Give: []any{
					&Slice[base.Character]{
						'a',
						'b',
					},
				},
				Want: []any{
					base.Character('b'),
				},
			},
		})
}

func TestSlice_Push(t *testing.T) {
	ttest.NewTest(t).
		SetSubject((*Slice[base.Character]).Push).
		Run([]ttest.TestCase{
			{
				Give: []any{
					&Slice[base.Character]{},
				},
				Want: []any{
					0,
				},
			},
			{
				Give: []any{
					&Slice[base.Character]{},
					base.Character('a'),
				},
				Want: []any{
					1,
				},
			},
			{
				Give: []any{
					&Slice[base.Character]{},
					base.Character('a'),
					base.Character('b'),
				},
				Want: []any{
					2,
				},
			},
		})
}

func TestSlice_Shift(t *testing.T) {
	ttest.NewTest(t).
		SetSubject((*Slice[base.Character]).Shift).
		Run([]ttest.TestCase{
			{
				Give: []any{
					&Slice[base.Character]{},
				},
				Want: []any{
					base.Character(0),
				},
			},
			{
				Give: []any{
					&Slice[base.Character]{
						'a',
						'b',
					},
				},
				Want: []any{
					base.Character('a'),
				},
			},
		})
}

func TestSlice_Unshift(t *testing.T) {
	ttest.NewTest(t).
		SetSubject((*Slice[base.Character]).Unshift).
		Run([]ttest.TestCase{
			{
				Give: []any{
					&Slice[base.Character]{},
				},
				Want: []any{
					0,
				},
			},
			{
				Give: []any{
					&Slice[base.Character]{},
					base.Character(97),
				},
				Want: []any{
					1,
				},
			},
			{
				Give: []any{
					&Slice[base.Character]{},
					base.Character('a'),
					base.Character('b'),
				},
				Want: []any{
					2,
				},
			},
		})
}
