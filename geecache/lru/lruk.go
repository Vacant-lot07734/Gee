package lru

type LruKCache struct {
	maxBytes  int64
	nbytes    int64
	k         int
	LruKCache map[string]lrukentry
}

type lrukentry struct {
	key    string
	value  Value
	record []int
}

func (node *lrukentry) times() int {
	return len(node.record)
}

func NewLruK(maxBytes int64, k int) *LruKCache {
	return &LruKCache{
		maxBytes: maxBytes,
		k:        k,
	}
}
