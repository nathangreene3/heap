package heap

import "testing"

func TestSortedInts(t *testing.T) {
	var (
		less Lesser        = func(x, y interface{}) bool { return x.(int) < y.(int) }
		exp  []interface{} = []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		rec  []interface{} = New(less, exp...).Sorted()
	)

	if len(exp) != len(rec) {
		t.Errorf("\nexpected length %d\nreceived %d\n", len(exp), len(rec))
	} else {
		for i := 0; i < len(exp); i++ {
			if e, r := exp[i].(int), rec[i].(int); e != r {
				t.Errorf("\nexpected %d\nreceived %d\n", e, r)
			}
		}
	}
}
