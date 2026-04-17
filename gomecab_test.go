package gomecab

import (
	"fmt"
	"strings"
	"testing"
)

var str = "すもももももももものうち"

func parse(t *testing.T, m *MeCab) {
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

	fmt.Println(tg.Parse(lt))
}

func parseToNode(t *testing.T, m *MeCab) {
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
			fmt.Printf("%s %s\n", node.Surface(), node.Feature())
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

// func TestMecab1(t *testing.T) {
// 	tagger, err := NewTagger("-Okatakana") // Convert to Katakana
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer tagger.Destroy()

// 	result, err := tagger.Parse("漢字")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	fmt.Println(result) // Result: カンジ
// }
