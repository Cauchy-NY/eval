package v1

import (
	"fmt"
	"math"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		// additional tests
		{"-1 + -x", Env{"x": 1}, "-2"},
		{"-1 - x", Env{"x": 1}, "-2"},
		// arithmetic tests
		{"1 + x", Env{"x": 1}, "2"},
		{"1 - x", Env{"x": 1}, "0"},
		{"a % 3", Env{"a": 100}, "1"},
		{"a % 3", Env{"a": -4}, "-1"},
		{"a % 3", Env{"a": 7.3}, "1"}, // force convert to int64 and calculate
		{"sqrt(num / pi)", Env{"num": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
		// logical tests
		{"a >= 10", Env{"a": 21}, "1"},
		{"a >= 10", Env{"a": 8}, "0"},
		{"a == 10", Env{"a": 10}, "1"},
		{"a == 10", Env{"a": 12}, "0"},
		{"2 < (a + b) && (a + b) <= 9", Env{"a": 1, "b": 5}, "1"},
		{"2 < (a + b) && (a + b) <= 9", Env{"a": 9, "b": 8}, "0"},
		{"2 < (a + b) && (a + b) <= 9", Env{"a": 1, "b": 1}, "0"},
		{"2 < a || b < 9", Env{"a": 3, "b": 10}, "1"},
		{"2 < a || b < 9", Env{"a": 1, "b": 10}, "0"},
		{"2 < a || b < 9", Env{"a": 3, "b": 8}, "1"},
		{"((a + 2) == 3) > 0", Env{"a": 1}, "1"},
		{"((a + 2) == 3) > 0", Env{"a": 2}, "0"},
		{"((a + b) == 6) <= 0", Env{"a": 6, "b": 5}, "1"},
		{"((a + b) == 6) <= 0", Env{"a": 1, "b": 5}, "0"},
		{"a < 10 == b > 6", Env{"a": 1, "b": 8}, "1"},
		{"a < 10 == b > 6", Env{"a": 12, "b": 5}, "1"},
		{"a < 10 == b > 6", Env{"a": 1, "b": 5}, "0"},
		{"!(a > 0)", Env{"a": 2}, "0"},
		{"!(a > 0)", Env{"a": -2}, "1"},
		{"!true", Env{}, "0"},
		{"false", Env{}, "0"},
	}
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
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}

/*
// output
-1 + -x
	map[x:1] => -2

-1 - x
	map[x:1] => -2

1 + x
	map[x:1] => 2

1 - x
	map[x:1] => 0

a % 3
	map[a:100] => 1
	map[a:-4] => -1
	map[a:7.3] => 1

sqrt(num / pi)
	map[num:87616 pi:3.141592653589793] => 167

pow(x, 3) + pow(y, 3)
	map[x:12 y:1] => 1729
	map[x:9 y:10] => 1729

5 / 9 * (F - 32)
	map[F:-40] => -40
	map[F:32] => 0
	map[F:212] => 100

a >= 10
	map[a:21] => 1
	map[a:8] => 0

a == 10
	map[a:10] => 1
	map[a:12] => 0

2 < (a + b) && (a + b) <= 9
	map[a:1 b:5] => 1
	map[a:9 b:8] => 0
	map[a:1 b:1] => 0

2 < a || b < 9
	map[a:3 b:10] => 1
	map[a:1 b:10] => 0
	map[a:3 b:8] => 1

((a + 2) == 3) > 0
	map[a:1] => 1
	map[a:2] => 0

((a + b) == 6) <= 0
	map[a:6 b:5] => 1
	map[a:1 b:5] => 0

a < 10 == b > 6
	map[a:1 b:8] => 1
	map[a:12 b:5] => 1
	map[a:1 b:5] => 0

!(a > 0)
	map[a:2] => 0
	map[a:-2] => 1

!true
	map[] => 0

false
	map[] => 0
*/

func TestErrors(t *testing.T) {
	for _, test := range []struct{ expr, wantErr string }{
		{"x << 2", "unexpected binary op \"<<\""},
		{"math.Pi", "unexpected '.'"},
		{`"hello"`, "unexpected '\"'"},
		{"log(10)", `unknown function "log"`},
		{"sqrt(1, 2)", "call to sqrt has 2 args, want 1"},
	} {
		expr, err := Parse(test.expr)
		if err == nil {
			vars := make(map[Var]bool)
			err = expr.Check(vars)
			if err == nil {
				t.Errorf("unexpected success: %s", test.expr)
				continue
			}
		}
		fmt.Printf("%-20s%v\n", test.expr, err) // (for book)
		if err.Error() != test.wantErr {
			t.Errorf("got error %s, want %s", err, test.wantErr)
		}
	}
}

/*
//!+errors
x % 2               unexpected '%'
math.Pi             unexpected '.'
"hello"             unexpected '"'
log(10)             unknown function "log"
sqrt(1, 2)          call to sqrt has 2 args, want 1
//!-errors
*/
