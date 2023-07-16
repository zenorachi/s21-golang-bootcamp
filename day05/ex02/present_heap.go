package ph

type Present struct {
	Value int
	Size  int
}

type PresentHeap []Present

func (ph *PresentHeap) Len() int {
	return len(*ph)
}

func (ph *PresentHeap) Less(i, j int) bool {
	if (*ph)[i].Value == (*ph)[j].Value {
		return (*ph)[i].Size < (*ph)[j].Size
	}
	return (*ph)[i].Value > (*ph)[j].Value
}

func (ph *PresentHeap) Swap(i, j int) {
	(*ph)[i], (*ph)[j] = (*ph)[j], (*ph)[i]
}

func (ph *PresentHeap) Push(x interface{}) {
	present := x.(Present)
	*ph = append(*ph, present)
}

func (ph *PresentHeap) Pop() interface{} {
	old := *ph
	n := len(old)
	present := old[n-1]
	*ph = old[0 : n-1]
	return present
}
