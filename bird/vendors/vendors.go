package vendors

import "container/heap"

type Vendor struct {
	VendorName string
	Price      float64
	Region     int
}

type MinHeap []Vendor

func (v MinHeap) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v MinHeap) Len() int {
	return len(v)
}

func (v MinHeap) Less(i, j int) bool {
	return v[i].Price < v[j].Price
}

func (v *MinHeap) Push(x any) {
	*v = append(*v, x.(Vendor))
}

func (v *MinHeap) Pop() any {
	oldLen := len(*v)
	old := *v
	x := old[oldLen-1]
	*v = old[:oldLen-1]
	return x

}

var VendorInfoMap map[int][]Vendor

// 00 denotes all
func InitializeVendorInfo() {
	vendorInit := []Vendor{{"TweetMobile", 0.04, 0}, {"Swanizon", 0.02, 32}, {"Swanizon", 0.03, 33}, {"BeakOn", 0.01, 31}, {"BeakOn", 0.01, 32}}
	// for +31
	markmap := make(map[int]bool)
	VendorInfoMap = make(map[int][]Vendor)
	for _, val := range vendorInit {
		markmap[val.Region] = true
	}
	for region, _ := range markmap {
		newlist := make(MinHeap, 0)
		heap.Init(&newlist)
		for _, vendors := range vendorInit {
			if region == vendors.Region {
				heap.Push(&newlist, vendors)
			}
			if vendors.Region == 0 {
				heap.Push(&newlist, vendors)
			}
		}
		VendorInfoMap[region] = newlist
	}
}
