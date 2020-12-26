package go_type_test

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	t.Run("Basic operation", func(t *testing.T) {
		m := map[string]int{
			"Adam": 20,
			"Eve":  18,
		}

		if m["Adam"] != 20 {
			t.Fatal("expected 20 but got", m["Adam"])
		}

		_, ok := m["Rob"]
		if ok == true {
			t.Fatal(`m["Rob"] doen't exist `, ok)
		}

		delete(m, "Eve")

		_, ok = m["Eve"]
		if ok == true {
			t.Fatal(`m["Eve"] should have been deleted`)
		}
	})

	t.Run("Distinct", func(t *testing.T) {
		users := []string{
			"Adam",
			"Eve",
			"Rob",
			"Kine",
			"Rob",
			"Eve",
		}

		uniq := make([]string, 0, len(users))
		m := make(map[string]struct{})
		for _, v := range users {
			if _, ok := m[v]; ok {
				continue
			}
			uniq = append(uniq, v)
			// 空の struct にすることで、メモリアロケーション無しになる
			m[v] = struct{}{}
		}

		if !reflect.DeepEqual(uniq, []string{"Adam", "Eve", "Rob", "Kine"}) {
			t.Fatal("expected ", []string{"Adam", "Eve", "Rob", "Kine"}, " but got ", uniq)
		}
	})
}
