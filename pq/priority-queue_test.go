package pq

import (
	"sort"
	"strconv"
	"testing"

	"github.com/Narayana08918/alien/models"
)

func TestPriorityQueue(t *testing.T) {
	pq := New()
	l := 10
	e := make([]int, 0, l)

	for i := 0; i < l; i++ {
		p := (i + 5) * 10
		e = append(e, p)

		city := &models.City{
			Out: genrateOutVals(p),
		}

		pq.Push(city)
	}

	sort.Ints(e)

	for i := l - 1; i >= 0; i-- {
		r := pq.Pop()

		if len(r.Out) != e[i] {
			t.Errorf("incorrect result: expected: %v, got: %v", e[i], len(r.Out))
		}
	}

	if pq.Size() != 0 {
		t.Errorf("incorrect result: expected: %v, got: %v", 0, pq.Size())
	}
}

func genrateOutVals(n int) map[string]string {
	out := make(map[string]string)
	for i := 0; i < n; i++ {
		out["dir-"+strconv.Itoa(i)] = "test-" + strconv.Itoa(i)
	}

	return out
}

func TestPriorityQueueRandom(t *testing.T) {
	pq := New()

	cities := []*models.City{
		{
			Name: "5",
			Out:  genrateOutVals(5),
		},
		{
			Name: "10",
			Out:  genrateOutVals(10),
		},
		{
			Name: "2",
			Out:  genrateOutVals(2),
		},
		{
			Name: "1",
			Out:  genrateOutVals(1),
		},
		{
			Name: "22",
			Out:  genrateOutVals(22),
		},
		{
			Name: "55",
			Out:  genrateOutVals(55),
		},
	}

	resName := []string{"55", "22", "10", "5", "2", "1"}

	for _, city := range cities {
		pq.Push(city)

	}

	for i := 0; i < len(cities); i++ {
		r := pq.Pop()

		if r.Name != resName[i] {
			t.Errorf("incorrect result: expected: %v, got: %v", resName[i], r.Name)
		}
	}
}
