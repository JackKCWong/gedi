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

func Compile(exp string, env map[string]any) (*vm.Program, error) {
	return expr.Compile(exp, expr.Env(env), atoi)
}
