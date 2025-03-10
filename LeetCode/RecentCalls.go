package main

type RecentCounter struct {
	recentReq []int
}

func ConstructorA() RecentCounter {
	newRC := RecentCounter{}
	newRC.recentReq = []int{}
	return newRC
}

func (this *RecentCounter) Ping(t int) int {
	this.recentReq = append(this.recentReq, t)
	l := 0
	for ; l < len(this.recentReq) && this.recentReq[l] < t-3000; l++ {
	}
	this.recentReq = this.recentReq[l:]
	return len(this.recentReq)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
