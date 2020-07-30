package s11

import (
	"fmt"
	"math/big"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	if hasPathSum(getTree(), 2) {

	}
}

//     1
//   /
//  2
func getTree() *TreeNode {
	node1 := &TreeNode{
		Val: 1,
	}
	node2 := &TreeNode{
		Val: 1,
	}
	node1.Left = node2
	return nil
}

func TestCombination(t *testing.T) {
	res := combination(2, 2)
	fmt.Printf("res: %d\n", res)
}

func TestGenerate(t *testing.T) {
	res := generate(10)
	for _, v := range res {
		fmt.Printf("%v\n", v)
	}
}

func TestBigInt(t *testing.T) {
	na := big.NewInt(3)
	nb := big.NewInt(4)
	nc := na.Mul(na, nb)
	fmt.Printf("na: %p, nb: %p, nc: %p\n", na, nb, nc)
	fmt.Printf("na: %d, nb: %d, nc: %d\n", na, nb, nc)

	ncC := *nc
	fmt.Printf("ncC: %d\n", &ncC)
	fmt.Printf("ncC: %d, na: %d\n", ncC.Mul(&ncC, na), na)
}

func TestFactorial(t *testing.T) {
	res := factorial(25)
	fmt.Printf("%d\n", res)
	for i := 0; i < len(factorialCache); i++ {
		fmt.Printf("p: %p, %2d: %d\n", factorialCache[int64(i)], i, factorialCache[int64(i)])
	}
}
