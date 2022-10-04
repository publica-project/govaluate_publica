package govaluate_publica

import (
	"fmt"
	"testing"
)

func TestWorking(t *testing.T) {
	query := "[456] == true || [12345] == true"

	exp, err := NewEvaluableExpression(query)
	if err != nil {
		panic(err)
	}

	dids := []string{"12345"}

	res, err := exp.Evaluate(dids)
	if err != nil {
		panic(err)
	}
	if res.(bool) != true {
		panic("not true")
	}
}

// V0    	 5785722	       187.2 ns/op	     352 B/op	       3 allocs/op
// V1    	63652362	        17.84 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSimpleQuery(b *testing.B) {
	query := "12345 == true"

	exp, err := NewEvaluableExpression(query)
	if err != nil {
		panic(err)
	}

	dids := []string{"12345"}

	for i := 0; i < 40; i++ {
		dids = append(dids, fmt.Sprintf("%d", i))
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := exp.Evaluate(dids)
		if err != nil {
			panic(err)
		}
	}
}
