package gomecab

import (
	"strings"
	"testing"
)

var str = "すもももももももものうち"

func parse(t *testing.T, m *MeCab) {
	t.Helper()
	tg, err := m.NewTagger()
	if err != nil {
		t.Fatal(err)
	}
	defer tg.Destroy()
	lt, err := m.NewLattice(str)
	if err != nil {
		t.Fatal(err)
	}
	defer lt.Destroy()

	out := tg.Parse(lt)
	if out == "" {
		t.Error("Parse returned empty string")
	}
	t.Logf("parse: %s", out)
}

func parseToNode(t *testing.T, m *MeCab) {
	t.Helper()
	tg, err := m.NewTagger()
	if err != nil {
		t.Fatal(err)
	}
	defer tg.Destroy()
	lt, err := m.NewLattice(str)
	if err != nil {
		t.Fatal(err)
	}
	defer lt.Destroy()

	node := tg.ParseToNode(lt)
	for {
		features := strings.Split(node.Feature(), ",")
		if features[0] == "名詞" {
			t.Logf("noun: %s %s", node.Surface(), node.Feature())
		}
		if node.Next() != nil {
			break
		}
	}
}

func TestMecab2(t *testing.T) {
	m, err := New("-Owakati")
	if err != nil {
		t.Fatal(err)
	}
	defer m.Destroy()
	parse(t, m)
	parseToNode(t, m)
}
