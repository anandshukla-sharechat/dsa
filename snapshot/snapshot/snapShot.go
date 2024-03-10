package snapshot

type SnapshotStruct struct {
	SnapSlice []int
	Index     int
}

type Snapshot interface {
	HasNext() bool
	Next() int
}

func InitialiseSnapshot(nums []int) Snapshot {
	return &SnapshotStruct{SnapSlice: nums, Index: 0}
}

func (s *SnapshotStruct) HasNext() bool {
	if s.Index >= len(s.SnapSlice) {
		return false
	}
	return true
}

func (s *SnapshotStruct) Next() int {
	if s.Index >= len(s.SnapSlice) {
		return -1
	}
	nextVal := s.SnapSlice[s.Index]
	s.Index = s.Index + 1
	return nextVal
}
