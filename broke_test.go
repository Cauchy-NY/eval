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
	{"1.5 * x", parser.Env{"x": 8}, float64(12)},
	{"1.5 * x", parser.Env{"x": 2}, float64(3)},
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
	{"want in [lang, \"php\"]", parser.Env{"want": "golang", "lang": "golang"}, true},
	{"want in [lang, \"php\"]", parser.Env{"want": "golang", "lang": "cpp"}, false},
	{"pron_predict > 0.86 && user_type not_in [\"big_v\", \"org\"]", parser.Env{"pron_predict": 0.97, "user_type": "normal"}, true},
	{"pron_predict > 0.86 && user_type not_in [\"big_v\", \"org\"]", parser.Env{"pron_predict": 0.97, "user_type": "big_v"}, false},
	{"pron_predict > 0.86 && user_type not_in [\"big_v\", \"org\"]", parser.Env{"pron_predict": 0.66, "user_type": "normal"}, false},
	// func test
	{"sqrt(num / pi)", parser.Env{"num": 87616.0, "pi": math.Pi}, float64(167.00011673013586)},
	{"sin(pi / 2)", parser.Env{"pi": math.Pi}, float64(1)},
	{"pow(x, 3) + pow(y, 3)", parser.Env{"x": 9.0, "y": 10.0}, float64(1729)},
	{"len(\"hello, world!\")", parser.Env{}, int64(13)},
	{"lower(\"GOLANG\")", parser.Env{}, "golang"},
	{"str_index(\"golang is a beautiful language\", x)", parser.Env{"x": "beautiful"}, int64(12)},
	{"contains(\"golang is a beautiful language\", x)", parser.Env{"x": "golang"}, true},
	{"contains(\"golang is a beautiful language\", x)", parser.Env{"x": "php"}, false},
	{"has_prefix(\"golang is a beautiful language\", x)", parser.Env{"x": "golang"}, true},
	{"has_prefix(\"golang is a beautiful language\", x)", parser.Env{"x": "php"}, false},
	{"has_suffix(\"golang is a beautiful language\", x)", parser.Env{"x": "language"}, true},
	{"has_suffix(\"golang is a beautiful language\", x)", parser.Env{"x": "beautiful"}, false},
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
