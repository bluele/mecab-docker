package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bluele/mecab-golang"
)

func parseToNode(m *mecab.MeCab) {
	tg, err := m.NewTagger()
	if err != nil {
		panic(err)
	}
	defer tg.Destroy()
	lt, err := m.NewLattice("すもももももももものうち")
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

func main() {
	log.Println("load mecab")
	m, err := mecab.New("-Owakati")
	if err != nil {
		panic(err)
	}
	defer m.Destroy()
	parseToNode(m)
}
