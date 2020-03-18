package eval

import "github.com/Cauchy-NY/eval/parser"

func Parse(input string) (parser.Node, error) {
	expr, err := parser.Parse(input)
	if err != nil {
		return nil, err
	}
	return expr, nil
}
