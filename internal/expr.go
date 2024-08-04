package internal

import (
	"strconv"

	"github.com/expr-lang/expr"
	"github.com/expr-lang/expr/vm"
)

var atoi = expr.Function(
	"atoi",
	func(params ...any) (any, error) {
		return strconv.Atoi(params[0].(string))
	},
	strconv.Atoi,
)

func Compile(exp string, env map[string]any, opts ...expr.Option) (*vm.Program, error) {
	opts = append(opts, expr.Env(env), atoi)
	return expr.Compile(exp, opts...)
}
