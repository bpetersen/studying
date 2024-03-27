package skiplist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func intCompare(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func stringCompare(a string, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// S                -> H                -> nil
// S      -> D      -> H      -> L      -> nil
// S -> B -> D -> F -> H -> J -> L -> N -> nil
// Height = 2
func getTestSkipList() (*SkipList[string], map[string]*slNode[string]) {
	var nNode = slNode[string]{
		data: "N",
		next: []*slNode[string]{nil},
	}
	var lNode = slNode[string]{
		data: "L",
		next: []*slNode[string]{&nNode, nil},
	}
	var jNode = slNode[string]{
		data: "J",
		next: []*slNode[string]{&lNode},
	}
	var hNode = slNode[string]{
		data: "H",
		next: []*slNode[string]{&jNode, &lNode, nil},
	}
	var fNode = slNode[string]{
		data: "F",
		next: []*slNode[string]{&hNode},
	}
	var dNode = slNode[string]{
		data: "D",
		next: []*slNode[string]{&fNode, &hNode},
	}
	var bNode = slNode[string]{
		data: "B",
		next: []*slNode[string]{&dNode},
	}
	var sentinelNode = slNode[string]{
		data: "sentinel",
		next: []*slNode[string]{&bNode, &dNode, &hNode},
	}
	var list = SkipList[string]{
		comparer: stringCompare,
		sentinel: &sentinelNode,
	}

	var mapOfNodes = make(map[string]*slNode[string])
	mapOfNodes["N"] = &nNode
	mapOfNodes["L"] = &lNode
	mapOfNodes["J"] = &jNode
	mapOfNodes["H"] = &hNode
	mapOfNodes["F"] = &fNode
	mapOfNodes["D"] = &dNode
	mapOfNodes["B"] = &bNode
	mapOfNodes["sentinel"] = &sentinelNode
	return &list, mapOfNodes
}

func TestComparerLess(t *testing.T) {
	actual := intCompare(1, 2)
	assert.Less(t, actual, 0)
}

func TestComparerGreater(t *testing.T) {
	actual := intCompare(2, 1)
	assert.Greater(t, actual, 0)
}

func TestComparerEqual(t *testing.T) {
	actual := intCompare(2, 2)
	assert.Equal(t, actual, 0)
}

func TestGeneric(t *testing.T) {
	actual := genericCompare[string](stringCompare, "A", "B")
	assert.Less(t, actual, 0)

	actual = genericCompare[int](intCompare, 1, 2)
	assert.Less(t, actual, 0)
}

func genericCompare[T any](comparer func(a T, b T) int, a T, b T) int {
	return comparer(a, b)
}

func TestGetPredNode(t *testing.T) {
	var list, nodes = getTestSkipList()

	var actualNode = list.getPredNode("F")
	fmt.Printf("actual node: %v", actualNode)

	assert.Equal(t, nodes["D"].data, actualNode.data)
	fmt.Println(list)

}

func TestFindNodeEqual(t *testing.T) {
	var list, nodes = getTestSkipList()

	var actual, err = list.Find("F")
	assert.NoError(t, err)

	assert.Equal(t, nodes["F"].data, actual)

}

func TestFindNodeGreaterThan(t *testing.T) {
	var list, nodes = getTestSkipList()

	var actual, err = list.Find("G")
	assert.NoError(t, err)

	assert.Equal(t, nodes["H"].data, actual)
}

// S                -> H                -> nil
// S      -> D      -> H      -> L      -> nil
// S -> B -> D -> F -> H -> J -> L -> N -> nil
// Height = 2
func TestAddNode(t *testing.T) {
	list, nodes := getTestSkipList()

	fmt.Println("BEFORE ADD")
	fmt.Println(list)

	newNode, wasAdded := list.addNode("G", 1)

	fmt.Println("AFTER ADD")
	fmt.Println(list)

	assert.Equal(t, true, wasAdded)
	assert.Equal(t, "G", newNode.data)

	assert.Equal(t, "H", newNode.next[0].data)
	assert.Equal(t, "H", newNode.next[1].data)

	assert.Equal(t, "G", nodes["F"].next[0].data)
	assert.Equal(t, "G", nodes["D"].next[1].data)

	actual, err := list.Find("G")
	assert.NoError(t, err)
	assert.Equal(t, "G", actual)
}

func TestAddNodeHeight0(t *testing.T) {
	list, nodes := getTestSkipList()

	newNode, wasAdded := list.addNode("G", 0)

	assert.Equal(t, true, wasAdded)
	assert.Equal(t, "G", newNode.data)

	assert.Equal(t, "H", newNode.next[0].data)

	assert.Equal(t, "G", nodes["F"].next[0].data)

	actual, err := list.Find("G")
	assert.NoError(t, err)
	assert.Equal(t, "G", actual)
}

func TestAddNodeHeightResizeSentinel(t *testing.T) {
	list, _ := getTestSkipList()

	fmt.Println("BEFORE ADD")
	fmt.Println(list)
	newNode, wasAdded := list.addNode("A", 4)
	fmt.Println("AFTER ADD")
	fmt.Println(list)

	assert.Equal(t, true, wasAdded)
	assert.Equal(t, "A", newNode.data)

	assert.Equal(t, "B", newNode.next[0].data)
	assert.Equal(t, "D", newNode.next[1].data)
	assert.Equal(t, "H", newNode.next[2].data)
	assert.Nil(t, newNode.next[3])
	assert.Nil(t, newNode.next[4])

	assert.Equal(t, "A", list.sentinel.next[0].data)
	assert.Equal(t, "A", list.sentinel.next[1].data)
	assert.Equal(t, "A", list.sentinel.next[2].data)
	assert.Equal(t, "A", list.sentinel.next[3].data)
	assert.Equal(t, "A", list.sentinel.next[4].data)

	actual, err := list.Find("A")
	assert.NoError(t, err)
	assert.Equal(t, "A", actual)
}
