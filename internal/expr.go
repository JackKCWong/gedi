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

var toint = expr.Function(
	"toint",
	func(params ...any) (any, error) {
		return strconv.Atoi(params[0].(string))
	},
	strconv.Atoi,
)

var tofloat = expr.Function(
	"tofloat",
	func(params ...any) (any, error) {
		return strconv.ParseFloat(params[0].(string), 64)
	},
	new(func(string) float64),
)

var tostr = expr.Function(
	"tostr",
	func(params ...any) (any, error) {
		return fmt.Printf("%s", params[0])
	},
	new(func(any) string),
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

var now = time.Now()
var within = expr.Function(
	"within",
	func(params ...any) (any, error) {
		var err error
		var dt time.Time
		switch arg := params[0].(type) {
		case time.Time:
			dt = arg
		case string:
			dt, err = dateparse.ParseAny(arg)
			if err != nil {
				return false, fmt.Errorf("failed to parse time: %w", err)
			}
		default:
			return false, fmt.Errorf("expecting a time.Time or string but was: %T", params[0])
		}

		dur, ok := params[1].(time.Duration)
		if !ok {
			return false, fmt.Errorf("expecting a time.Duration but was: %T", dur)
		}

		return dt.Compare(now.Add(dur)) >= 0, nil
	},
	new(func(time.Time, time.Duration) (bool, error)),
	new(func(string, time.Duration) (bool, error)),
)

func parseTimes(params ...any) (time.Time, time.Time, error) {
	var err error
	var dt1, dt2 time.Time
	switch dt := params[0].(type) {
	case time.Time:
		dt1 = dt
	case string:
		dt1, err = dateparse.ParseAny(dt)
		if err != nil {
			return time.Time{}, time.Time{}, fmt.Errorf("failed to parse time: %w", err)
		}
	default:
		return time.Time{}, time.Time{}, fmt.Errorf("expecting a time.Time or string but was: %T", params[0])
	}

	switch dt := params[1].(type) {
	case time.Time:
		dt2 = dt
	case string:
		dt2, err = dateparse.ParseAny(dt)
		if err != nil {
			return time.Time{}, time.Time{}, fmt.Errorf("failed to parse time: %w", err)
		}
	default:
		return time.Time{}, time.Time{}, fmt.Errorf("expecting a time.Time or string but was: %T", params[0])
	}

	return dt1, dt2, nil
}

var before = expr.Function(
	"before",
	func(params ...any) (any, error) {
		dt1, dt2, err := parseTimes(params...)
		if err != nil {
			return false, err
		}

		return dt1.Compare(dt2) <= 0, nil
	},
	new(func(time.Time, time.Time) (bool, error)),
	new(func(string, time.Time) (bool, error)),
	new(func(time.Time, string) (bool, error)),
	new(func(string, string) (bool, error)),
)

var after = expr.Function(
	"after",
	func(params ...any) (any, error) {
		dt1, dt2, err := parseTimes(params...)
		if err != nil {
			return false, err
		}

		return dt1.Compare(dt2) >= 0, nil
	},
	new(func(time.Time, time.Time) (bool, error)),
	new(func(string, time.Time) (bool, error)),
	new(func(time.Time, string) (bool, error)),
	new(func(string, string) (bool, error)),
)

func Compile(exp string, env map[string]any, opts ...expr.Option) (*vm.Program, error) {
	env["now"] = now
	env["msec"] = time.Millisecond
	env["sec"] = time.Second
	env["min"] = time.Minute
	env["hour"] = time.Hour
	env["day"] = 24 * time.Hour
	env["week"] = 7 * 24 * time.Hour
	env["month"] = 30 * 24 * time.Hour
	env["year"] = 365 * 24 * time.Hour
	env["now"] = now

	opts = append(opts, expr.Env(env),
		atoi,
		toint,
		tofloat,
		tostr,
		localtime,
		utctime,
		tztime,
		within,
		before,
		after,
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
