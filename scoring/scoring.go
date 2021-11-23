package scoring


import (
	"fmt"
	"sort"
)

/// Types for storing key/score pairs during decryption.

type ScoredItem struct {
	Item interface{}
	Score float64
}

type ScoredList struct {
	Items []ScoredItem
	sorted bool
}

/// Contains implementations for the STLlib sort interface for ScoredLists.

func (s *ScoredList) Len() int {
	return len(s.Items)
}

func (s *ScoredList) Swap(i, j int) {
	s.Items[i], s.Items[j] = s.Items[j], s.Items[i]
}

func (s *ScoredList) Less(i, j int) bool {
	return (s.Items[i].Score > s.Items[j].Score)
}

func (sList *ScoredList) Sort() {
	if sList.sorted {return}
	sort.Sort(sList)
	sList.sorted = true;
}

/// Useable sorting interface.

func (sList *ScoredList) Print() {
	sList.Sort()
	fmt.Printf("--------------------------------------------------------------------------------\n")
	for _, si := range sList.Items {
		fmt.Printf("%v : %6.4f\n", si.Item, si.Score)
	}
}

func (sList *ScoredList) AddItem(newItem interface{}, newScore float64) {
	sList.Items = append(sList.Items,ScoredItem{newItem, newScore})
	sList.sorted = false
}

func (sList *ScoredList) TopN(n int) ([]ScoredItem) {
	sList.Sort()
	return sList.Items[0:n]
}

func (sList *ScoredList) BottomN(n int) ([]ScoredItem) {
	sList.Sort()
	last := len(sList.Items) - 1
	bottom := make([]ScoredItem, n)
	for i := 0; i < n; i += 1 {
		bottom[i] = sList.Items[last-i]
	}
	return bottom
}
