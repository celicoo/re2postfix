package re2postfix

import (
	"github.com/celicoo/re2postfix/errors"
	"github.com/celicoo/re2postfix/internal/base"
	"github.com/celicoo/re2postfix/internal/container"
)

func Run(infix *container.Slice[base.Character]) (postfix *container.Slice[base.Character], e error) {
	{
		var (
			alternative bool
			postfix     = &container.Slice[base.Character]{}
		)
		for len(*infix) > 0 {
			c := infix.Shift()
			switch c {
			default:
				postfix.Push(c)
				if len(*postfix) > 1 != alternative {
					infix.Unshift('.')
				}
				alternative = false
			case '|':
				alternative = true
				defer postfix.Push(c)
			case '.':
				if len(*infix) > 0 {
					nextc := infix.Shift()
					switch nextc {
					case '*', '+', '?':
						infix.Unshift(nextc, c)
						continue
					}
					infix.Unshift(nextc)
				}
				postfix.Push(c)
			case '*', '+', '?':
				if len(*postfix) == 0 {
					e = errors.Syntax("nothing to repeat")
					break
				}
				postfix.Push(c)
			case '(':
				var subinput container.Slice[base.Character]
				for len(*infix) != 0 {
					c = infix.Shift()
					if c == ')' {
						break
					}
					subinput.Push(c)
				}
				if c == ')' {
					s, e := Run(&subinput)
					if e != nil {
						return postfix, e
					}
					if len(*postfix) > 1 != alternative {
						infix.Unshift('.')
					}
					postfix.Push(*s...)
					alternative = false
					continue
				}
				e = errors.Syntax("unmatched opening parenthesis")
				break
			case ')':
				e = errors.Syntax("unmatched closing parenthesis")
				break
			}
		}
		return postfix, e
	}
}
