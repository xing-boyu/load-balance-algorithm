package rete_limit_algorithm

type weightEntry struct {
	item interface{}
	weight int
}

type RoundRobinAlg struct {
	items []*weightEntry
	n 		int
	maxW	int
	gcd		int
	i 		int
	cw		int
}

func (rra *RoundRobinAlg)Add(item interface{}, weight int) {
	weightEntryItem := &weightEntry{item:item, weight:weight}

	if weight > 0 {
		if rra.n == 0 {
			rra.maxW = weight
			rra.gcd = weight
			rra.i = -1
			rra.cw = 0
		} else {
			rra.gcd = gcd(weight, rra.gcd)
			if rra.maxW < weight {
				rra.maxW = weight
			}
		}
	}

	rra.items = append(rra.items, weightEntryItem)
	rra.n++
}

func (rra *RoundRobinAlg)Remove(item interface{}) {
	for i:=0; i<rra.n; i++ {
		if rra.items[i].item == item {
			rra.items = append(rra.items[:i], rra.items[i+1:]...)
			rra.n--
			break
		}
	}
}

func (rra *RoundRobinAlg)Next() (interface{}) {
	if rra.n == 0 {
		return nil
	}

	for {
		rra.i = (rra.i + 1) % rra.n
		if rra.i == 0 {
			rra.cw = rra.cw - rra.gcd
			if rra.cw <= 0 {
				rra.cw = rra.maxW
				if rra.cw == 0 {
					return nil
				}
			}
		}

		if rra.items[rra.i].weight >= rra.cw {
			return rra.items[rra.i].item
		}
	}
}

func gcd(x int, y int) (gcd int) {
	for {
		if x % y == 0 {
			gcd = y
			return
		} else {
			t := x % y
			x, y = y, t
		}
	}
}