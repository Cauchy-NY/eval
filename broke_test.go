package eval

import (
	"fmt"
	"github.com/Cauchy-NY/eval/parser"
	"math"
	"reflect"
	"testing"
)

var tests = []struct {
	expr string
	env  parser.Env
	want interface{}
}{
	// additional tests
	{"-1 + -x", parser.Env{"x": 1}, int64(-2)},
	{"-1 - x", parser.Env{"x": 1}, int64(-2)},
	// arithmetic tests
	{"1 + x", parser.Env{"x": 1}, int64(2)},
	{"1 - x", parser.Env{"x": 1}, int64(0)},
	{"a % 3", parser.Env{"a": 100}, int64(1)},
	{"a % 3", parser.Env{"a": -4}, int64(-1)},
	{"5 / 9 * (x - 32)", parser.Env{"x": 32}, int64(0)},
	{"5 / 9 * (x - 32)", parser.Env{"x": -40}, int64(0)}, // precision loss
	{"5 / 9 * (x - 32)", parser.Env{"x": 212}, int64(0)},
	{"5.0 / 9 * (x - 32)", parser.Env{"x": -40}, float64(-40)},
	{"5.0 / 9 * (x - 32)", parser.Env{"x": 212}, float64(100)},
	{"greet + name", parser.Env{"greet": "hello,", "name": " world"}, "hello, world"},
	// logical tests
	{"!true", parser.Env{}, false},
	{"false", parser.Env{}, false},
	{"!(a > 0)", parser.Env{"a": 2}, false},
	{"!(a > 0)", parser.Env{"a": -2}, true},
	{"a >= 10", parser.Env{"a": 21}, true},
	{"a >= 10", parser.Env{"a": 8}, false},
	{"a == 10", parser.Env{"a": 10}, true},
	{"a == 10", parser.Env{"a": 12}, false},
	{"a != 10", parser.Env{"a": 12}, true},
	{"a != 10", parser.Env{"a": 10}, false},
	{"name == \"Tom\"", parser.Env{"name": "Tom"}, true},
	{"name == \"Tom\"", parser.Env{"name": "Jim"}, false},
	{"name != \"Tom\"", parser.Env{"name": "Jim"}, true},
	{"name != \"Tom\"", parser.Env{"name": "Tom"}, false},
	{"2 < (a + b) && (a + b) <= 9", parser.Env{"a": 1, "b": 5}, true},
	{"2 < (a + b) && (a + b) <= 9", parser.Env{"a": 9, "b": 8}, false},
	{"2 < (a + b) && (a + b) <= 9", parser.Env{"a": 1, "b": 1}, false},
	{"2 < a || b < 9", parser.Env{"a": 3, "b": 10}, true},
	{"2 < a || b < 9", parser.Env{"a": 1, "b": 10}, false},
	{"2 < a || b < 9", parser.Env{"a": 3, "b": 8}, true},
	{"a < 10 == b > 6", parser.Env{"a": 1, "b": 8}, true},
	{"a < 10 == b > 6", parser.Env{"a": 12, "b": 5}, true},
	{"a < 10 == b > 6", parser.Env{"a": 1, "b": 5}, false},
	// func test
	{"sqrt(num / pi)", parser.Env{"num": 87616.0, "pi": math.Pi}, float64(167.00011673013586)},
	{"pow(x, 3.0) + pow(y, 3.0)", parser.Env{"x": 12.0, "y": 1.0}, float64(1729)},
	{"pow(x, 3.0) + pow(y, 3.0)", parser.Env{"x": 9.0, "y": 10.0}, float64(1729)},
}

func TestEval(t *testing.T) {
	var prevExpr string
	for _, test := range tests {
		// Print expr only when it changes.
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		got := expr.Eval(test.env)
		fmt.Printf("\t%v => %v\n", test.env, got)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s.Eval() in %v = %s, want %s\n",
				test.expr, test.env, got, test.want)
		}
	}
}
