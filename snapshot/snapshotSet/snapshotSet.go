package snapshotSet

import "snapshot/snapshot"

type SnapshotSet interface {
	Add(num int)
	Remove(num int)
	Contains(num int) bool
	Iterator() snapshot.Snapshot
}

type SetStruct struct {
	SetMap map[int]bool
}

func SnapshotSetInitailise() SnapshotSet {
	return &SetStruct{SetMap: make(map[int]bool)}

}

func (s *SetStruct) Add(num int) {
	s.SetMap[num] = true
}

func (s *SetStruct) Remove(num int) {
	delete(s.SetMap, num)
}

func (s *SetStruct) Contains(num int) bool {
	_, ok := s.SetMap[num]
	return ok
}

func (s *SetStruct) Iterator() snapshot.Snapshot {
	tempSnapArr := make([]int, 0)
	for val, _ := range s.SetMap {
		tempSnapArr = append(tempSnapArr, val)
	}

	snapObj := snapshot.InitialiseSnapshot(tempSnapArr)
	return snapObj
}
