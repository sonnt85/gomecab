package gomecab

import (
	"fmt"
	"strings"
	"testing"
)

var str = "すもももももももものうち"

func parse(m *MeCab) {
	tg, err := m.NewTagger()
	if err != nil {
		panic(err)
	}
	defer tg.Destroy()
	lt, err := m.NewLattice(str)
	if err != nil {
		panic(err)
	}
	defer lt.Destroy()

	fmt.Println(tg.Parse(lt))
}

func parseToNode(m *MeCab) {
	tg, err := m.NewTagger()
	if err != nil {
		panic(err)
	}
	defer tg.Destroy()
	lt, err := m.NewLattice(str)
	if err != nil {
		panic(err)
	}
	defer lt.Destroy()

	node := tg.ParseToNode(lt)
	for {
		features := strings.Split(node.Feature(), ",")
		if features[0] == "名詞" {
			fmt.Println(fmt.Sprintf("%s %s", node.Surface(), node.Feature()))
		}
		if node.Next() != nil {
			break
		}
	}
}

func TestMecab2(t *testing.T) {
	m, err := New("-Owakati")
	if err != nil {
		panic(err)
	}
	defer m.Destroy()
	parse(m)
	parseToNode(m)
}

// func TestMecab1(t *testing.T) {
// 	tagger, err := NewTagger("-Okatakana") // Chuyển đổi sang Katakana
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer tagger.Destroy()

// 	result, err := tagger.Parse("漢字")
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(result) // Kết quả: カンジ
// }
