package internal

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/araddon/dateparse"
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

var localtime = expr.Function(
	"localtime",
	func(params ...any) (any, error) {
		return dateparse.ParseLocal(params[0].(string))
	},
	dateparse.ParseLocal,
)

var utctime = expr.Function(
	"utctime",
	func(params ...any) (any, error) {
		return dateparse.ParseIn(params[0].(string), time.UTC)
	},
	dateparse.ParseIn,
)

var tztime = expr.Function(
	"tztime",
	func(params ...any) (any, error) {
		tz, err := TimeZoneStrToLocation(params[1].(string))
		if err != nil {
			tz, err = time.LoadLocation(params[1].(string))
			if err != nil {
				return time.Time{}, fmt.Errorf("invalid timezone string: %w", err)
			}
		}

		return dateparse.ParseIn(params[0].(string), tz)
	},
	dateparse.ParseIn,
)

var within = expr.Function(
	"within",
	func(params ...any) (any, error) {
		dt, ok := params[0].(time.Time)
		if !ok {
			return false, fmt.Errorf("expecting first parameter to be a time.Time but was: %T", params[0])
		}

		dur, ok := params[1].(string)
		if !ok {
			return false, fmt.Errorf("invalid duration string: %+v", dur)
		}

		return Within(dt, dur)
	},
	Within,
)

var now = time.Now()

func Within(dt time.Time, durstr string) (bool, error) {
	dur, err := time.ParseDuration(durstr)

	if err != nil {
		return false, fmt.Errorf("invalid duration: %w", err)
	}

	return dt.After(now.Add(dur)), nil
}

func Compile(exp string, env map[string]any, opts ...expr.Option) (*vm.Program, error) {
	env["now"] = now
	opts = append(opts, expr.Env(env),
		atoi,
		localtime,
		utctime,
		tztime,
		within,
	)

	return expr.Compile(exp, opts...)
}

func TimeZoneStrToLocation(tzStr string) (*time.Location, error) {
	if !strings.HasPrefix(tzStr, "UTC+") && !strings.HasPrefix(tzStr, "UTC-") {
		return nil, fmt.Errorf("invalid timezone string: %s", tzStr)
	}

	offset := strings.TrimPrefix(tzStr, "UTC")
	hours, err := strconv.Atoi(offset)
	if err != nil {
		return nil, fmt.Errorf("failed to parse offset string: %w", err)
	}

	return time.FixedZone(tzStr, hours*3600), nil
}
