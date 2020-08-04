// 380
package s38

type RandomizedSet struct {
	m map[int]struct{}
}

/** Initialize your data structure here. */
func Constructor1() RandomizedSet {
	rs := RandomizedSet{
		m: make(map[int]struct{}, 1000),
	}

	return rs
}

func (this *RandomizedSet) isExist(val int) bool {
	_, ok := this.m[val]
	return ok
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if this.isExist(val) {
		return false
	}
	this.m[val] = struct{}{}
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if !this.isExist(val) {
		return false
	}
	delete(this.m, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	var result int
	for k := range this.m {
		result = k
	}
	return result
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
