package internal

import (
	"cmp"
	"fmt"
	"time"

	"github.com/expr-lang/expr"
)

var greaterThan = expr.Function(
	"gt",
	func(params ...any) (any, error) {
		if len(params) != 2 {
			return false, fmt.Errorf("expecting 2 arguments but was: %d", len(params))
		}

		{
			v1, ok1 := params[0].(int)
			v2, ok2 := params[1].(int)
			if ok1 && ok2 {
				return v1 > v2, nil
			}
		}

		{
			v1, ok1 := params[0].(int8)
			v2, ok2 := params[1].(int8)
			if ok1 && ok2 {
				return v1 > v2, nil
			}
		}

		{
			v1, ok1 := params[0].(int16)
			v2, ok2 := params[1].(int16)
			if ok1 && ok2 {
				return v1 > v2, nil
			}
		}

		{
			v1, ok1 := params[0].(int32)
			v2, ok2 := params[1].(int32)
			if ok1 && ok2 {
				return v1 > v2, nil
			}
		}

		{
			v1, ok1 := params[0].(int64)
			v2, ok2 := params[1].(int64)
			if ok1 && ok2 {
				return v1 > v2, nil
			}
		}

		{
			v1, ok1 := params[0].(float64)
			v2, ok2 := params[1].(float64)
			if ok1 && ok2 {
				return v1 > v2, nil
			}
		}

		{
			v1, ok1 := params[0].(float32)
			v2, ok2 := params[1].(float32)
			if ok1 && ok2 {
				return v1 > v2, nil
			}
		}


		{
			v1, ok1 := params[0].(string)
			v2, ok2 := params[1].(string)
			if ok1 && ok2 {
				return cmp.Compare(v1, v2) > 0, nil
			}
		}

		{
			v1, ok1 := params[0].(time.Time)
			v2, ok2 := params[1].(time.Time)
			if ok1 && ok2 {
				return v1.Compare(v2) > 0, nil
			}
		}

		{
			v1, ok1 := params[0].(time.Duration)
			v2, ok2 := params[1].(time.Duration)
			if ok1 && ok2 {
				return v1 > v2, nil
			}
		}

		return false, fmt.Errorf("uncomparable types: %T and %T", params[0], params[1])
	},
	new (func(any, any) (bool, error)),
)


var greaterOrEqual = expr.Function(
	"ge",
	func(params ...any) (any, error) {
		if len(params) != 2 {
			return false, fmt.Errorf("expecting 2 arguments but was: %d", len(params))
		}

		{
			v1, ok1 := params[0].(int)
			v2, ok2 := params[1].(int)
			if ok1 && ok2 {
				return v1 >= v2, nil
			}
		}

		{
			v1, ok1 := params[0].(int8)
			v2, ok2 := params[1].(int8)
			if ok1 && ok2 {
				return v1 >= v2, nil
			}
		}

		{
			v1, ok1 := params[0].(int16)
			v2, ok2 := params[1].(int16)
			if ok1 && ok2 {
				return v1 >= v2, nil
			}
		}

		{
			v1, ok1 := params[0].(int32)
			v2, ok2 := params[1].(int32)
			if ok1 && ok2 {
				return v1 >= v2, nil
			}
		}

		{
			v1, ok1 := params[0].(int64)
			v2, ok2 := params[1].(int64)
			if ok1 && ok2 {
				return v1 >= v2, nil
			}
		}

		{
			v1, ok1 := params[0].(float64)
			v2, ok2 := params[1].(float64)
			if ok1 && ok2 {
				return v1 >= v2, nil
			}
		}

		{
			v1, ok1 := params[0].(float32)
			v2, ok2 := params[1].(float32)
			if ok1 && ok2 {
				return v1 >= v2, nil
			}
		}


		{
			v1, ok1 := params[0].(string)
			v2, ok2 := params[1].(string)
			if ok1 && ok2 {
				return cmp.Compare(v1, v2) >= 0, nil
			}
		}

		{
			v1, ok1 := params[0].(time.Time)
			v2, ok2 := params[1].(time.Time)
			if ok1 && ok2 {
				return v1.Compare(v2) >= 0, nil
			}
		}

		{
			v1, ok1 := params[0].(time.Duration)
			v2, ok2 := params[1].(time.Duration)
			if ok1 && ok2 {
				return v1 >= v2, nil
			}
		}

		return false, fmt.Errorf("uncomparable types: %T and %T", params[0], params[1])
	},
	new (func(any, any) (bool, error)),
)

var lessThan = expr.Function(
	"lt",
	func(params ...any) (any, error) {
		if len(params) != 2 {
			return false, fmt.Errorf("expecting 2 arguments but was: %d", len(params))
		}

		{
			v1, ok1 := params[0].(int)
			v2, ok2 := params[1].(int)
			if ok1 && ok2 {
				return v1 < v2, nil
			}
		}

		{
			v1, ok1 := params[0].(int8)
			v2, ok2 := params[1].(int8)
			if ok1 && ok2 {
				return v1 < v2, nil
			}
		}

		{
			v1, ok1 := params[0].(int16)
			v2, ok2 := params[1].(int16)
			if ok1 && ok2 {
				return v1 < v2, nil
			}
		}

		{
			v1, ok1 := params[0].(int32)
			v2, ok2 := params[1].(int32)
			if ok1 && ok2 {
				return v1 < v2, nil
			}
		}

		{
			v1, ok1 := params[0].(int64)
			v2, ok2 := params[1].(int64)
			if ok1 && ok2 {
				return v1 < v2, nil
			}
		}

		{
			v1, ok1 := params[0].(float64)
			v2, ok2 := params[1].(float64)
			if ok1 && ok2 {
				return v1 < v2, nil
			}
		}

		{
			v1, ok1 := params[0].(float32)
			v2, ok2 := params[1].(float32)
			if ok1 && ok2 {
				return v1 < v2, nil
			}
		}


		{
			v1, ok1 := params[0].(string)
			v2, ok2 := params[1].(string)
			if ok1 && ok2 {
				return cmp.Compare(v1, v2) < 0, nil
			}
		}

		{
			v1, ok1 := params[0].(time.Time)
			v2, ok2 := params[1].(time.Time)
			if ok1 && ok2 {
				return v1.Compare(v2) < 0, nil
			}
		}

		{
			v1, ok1 := params[0].(time.Duration)
			v2, ok2 := params[1].(time.Duration)
			if ok1 && ok2 {
				return v1 < v2, nil
			}
		}

		return false, fmt.Errorf("uncomparable types: %T and %T", params[0], params[1])
	},
	new (func(any, any) (bool, error)),
)

var lessOrEqual = expr.Function(
	"le",
	func(params ...any) (any, error) {
		if len(params) != 2 {
			return false, fmt.Errorf("expecting 2 arguments but was: %d", len(params))
		}

		{
			v1, ok1 := params[0].(int)
			v2, ok2 := params[1].(int)
			if ok1 && ok2 {
				return v1 <= v2, nil
			}
		}

		{
			v1, ok1 := params[0].(int8)
			v2, ok2 := params[1].(int8)
			if ok1 && ok2 {
				return v1 <= v2, nil
			}
		}

		{
			v1, ok1 := params[0].(int16)
			v2, ok2 := params[1].(int16)
			if ok1 && ok2 {
				return v1 <= v2, nil
			}
		}

		{
			v1, ok1 := params[0].(int32)
			v2, ok2 := params[1].(int32)
			if ok1 && ok2 {
				return v1 <= v2, nil
			}
		}

		{
			v1, ok1 := params[0].(int64)
			v2, ok2 := params[1].(int64)
			if ok1 && ok2 {
				return v1 <= v2, nil
			}
		}

		{
			v1, ok1 := params[0].(float64)
			v2, ok2 := params[1].(float64)
			if ok1 && ok2 {
				return v1 <= v2, nil
			}
		}

		{
			v1, ok1 := params[0].(float32)
			v2, ok2 := params[1].(float32)
			if ok1 && ok2 {
				return v1 <= v2, nil
			}
		}


		{
			v1, ok1 := params[0].(string)
			v2, ok2 := params[1].(string)
			if ok1 && ok2 {
				return cmp.Compare(v1, v2) <= 0, nil
			}
		}

		{
			v1, ok1 := params[0].(time.Time)
			v2, ok2 := params[1].(time.Time)
			if ok1 && ok2 {
				return v1.Compare(v2) <= 0, nil
			}
		}

		{
			v1, ok1 := params[0].(time.Duration)
			v2, ok2 := params[1].(time.Duration)
			if ok1 && ok2 {
				return v1 <= v2, nil
			}
		}

		return false, fmt.Errorf("uncomparable types: %T and %T", params[0], params[1])
	},
	new (func(any, any) (bool, error)),
)
