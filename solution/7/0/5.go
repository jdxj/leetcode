// 705
package s70

type MyHashSet struct {
	set map[int]struct{}
}

/** Initialize your data structure here. */
func ConstructorI() MyHashSet {
	return MyHashSet{
		set: make(map[int]struct{}, 1000),
	}
}

func (this *MyHashSet) Add(key int) {
	this.set[key] = struct{}{}
}

func (this *MyHashSet) Remove(key int) {
	delete(this.set, key)
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	_, exist := this.set[key]
	return exist
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
