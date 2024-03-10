package main

import (
	"fmt"
	"snapshot/snapshotSet"
)

/*

	Design a snapshot set functionality.

	The class should provide add, remove, contains and iterator functionality. The snapshot of the set is taken when the iterator function is called. The iterator always returns elements that were present in the snapshot even though the element might be added or removed from set after the snapshot.

	Once the snapshot is taken, the iterator of the class should only return values that were present in the set at time of snapshot.

	interface SnapshotSet {
	void add(int num);
	void remove(num);
	boolean contains(num);
	Iterator<Integer> iterator(); // the first call to this function should trigger a snapshot of the set
	}

	SnapshotSet set = new SnapshotSet();
	set.add(1);
	set.add(2);
	set.add(3);
	set.add(4);

	Iterator<Integer> itr1 = set.iterator(); // iterator snapshot
	set.remove(1);
	set.contains(1); // returns false; because 1 was removed.

	while(itr1.hasNext())
	System.out.print(itr1.next()); // Should print 1, 2, 3, 4 (in any order)

	Iterator<Integer> itr2 = set.iterator(); // iterator snapshot
	while(itr2.hasNext())
	System.out.print(itr2.next()); // Should print 2, 3, 4 (in any order)

*/

func main() {

	set := snapshotSet.SnapshotSetInitailise()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(4)

	it := set.Iterator()
	set.Remove(1)
	newIt := set.Iterator()

	fmt.Println("Before Removal")
	for it.HasNext() {
		fmt.Println(it.Next())
	}
	fmt.Println("After Removal")
	for newIt.HasNext() {
		fmt.Println(newIt.Next())
	}
}
