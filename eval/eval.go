package eval

import (
	"context"
	"math"

	"github.com/PaesslerAG/gval"
	"github.com/knieriem/text/stringutil"
)

type Evaluator struct {
	gval.Evaluable
}

func NewEvaluator(expr string) (*Evaluator, error) {
	expr, err := stringutil.ConvertMethodChain(expr, ", ")
	if err != nil {
		return nil, err
	}
	e, err := gval.Arithmetic().NewEvaluable(expr)
	if err != nil {
		return nil, err
	}
	return &Evaluator{Evaluable: e}, nil
}

func (e *Evaluator) Eval(i int, x float64) (float64, error) {
	v, err := e.EvalFloat64(context.Background(), map[string]interface{}{
		"exp": func(x float64) float64 {
			return math.Exp(x)
		},
		"index": func(i int, a ...float64) float64 {
			return a[i]
		},

		"round": func(v, unit float64) float64 { return math.Round(v/unit) * unit },
		"clip": func(v, min, max float64) float64 {
			if v < min {
				return min
			} else if v > max {
				return max
			}
			return v
		},

		"x": x,
		"i": i,
	})
	return v, err
}
