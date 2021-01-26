package hashtree

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

// we must also split data, create a tree, create a compact tree, and pass the shuffled data to it, to get it back in order
// when passing an invalid amount of blocks to the CreateHashTree, returns an error (1, for example.)
func createTreeAndTest(t *testing.T, text string, delimiter string, height int) {

	shuf := func(v [][]byte) {
		f := reflect.Swapper(v)
		n := len(v)
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for i := 0; i < n; i++ {
			f(r.Intn(n), r.Intn(n))
		}
	}

	splittedData := bytes.Split([]byte(text), []byte(delimiter))
	splittedDataLength := len(splittedData)
	splittedDataLengthPowerOfTwo := int(math.Pow(2, math.Ceil(math.Log(float64(splittedDataLength))/math.Log(2))))
	tree, treeErr := NewBuilder().Create().WithBlocks(splittedData).Now()

	if tree == nil {
		t.Errorf("the returned instance was expected to be an instance, nil returned")
		return
	}

	if treeErr != nil {
		t.Errorf("the returned error was expected to be nil, valid error returned: %s", treeErr.Error())
		return
	}

	secondTree, secondTreeErr := NewBuilder().Create().WithBlocks(splittedData).Now()
	if secondTreeErr != nil {
		t.Errorf("the returned error was expected to be nil, valid error returned: %s", secondTreeErr.Error())
		return
	}

	if tree.Head().String() != secondTree.Head().String() {
		t.Errorf("the tree hashes changed even if they were build with the same data: First: %s, Second: %s", tree.Head().String(), secondTree.Head().String())
		return
	}

	treeHeight := tree.Height()
	if treeHeight != height {
		t.Errorf("the binary tree's height should be %d because it contains %d data blocks, %d given", height, len(splittedData), treeHeight)
		return
	}

	treeLength := tree.Length()
	if treeLength != splittedDataLengthPowerOfTwo {
		t.Errorf("the HashTree should have a length of %d, %d given", splittedDataLengthPowerOfTwo, treeLength)
		return
	}

	compact := tree.Compact()
	compactLength := compact.Length()
	if splittedDataLengthPowerOfTwo != compactLength {
		t.Errorf("the CompactHashTree should have a length of %d, %d given", splittedDataLengthPowerOfTwo, compactLength)
		return
	}

	if !tree.Head().Compare(compact.Head()) {
		t.Errorf("the HashTree root hash: %x is not the same as the CompactHashTree root hash: %x", tree.Head().Bytes(), compact.Head().Bytes())
		return
	}

	shuffledData := make([][]byte, len(splittedData))
	copy(shuffledData, splittedData)
	shuf(shuffledData)

	reOrderedSplittedData, reOrderedSplittedDataErr := tree.Order(shuffledData)
	if reOrderedSplittedDataErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", reOrderedSplittedDataErr.Error())
		return
	}

	if !reflect.DeepEqual(splittedData, reOrderedSplittedData) {
		t.Errorf("the re-ordered data is invalid")
		return
	}

	treeJS, err := json.Marshal(tree)
	if err != nil {
		t.Errorf("the error was expected to be nil")
		return
	}

	treeFromJS, err := NewBuilder().Create().WithJS(treeJS).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil")
		return
	}

	if tree.Head().String() != treeFromJS.Head().String() {
		t.Errorf("there is an error while building an HashTree from JSON")
		return
	}
}

func TestHashTree_Success(t *testing.T) {
	createTreeAndTest(t, "this|is", "|", 2)                                                                                                                       //2 blocks
	createTreeAndTest(t, "this|is|some|data|separated|by|delimiters|asfsf", "|", 4)                                                                               //8 blocks
	createTreeAndTest(t, "this|is|some|data|separated|by|delimiters|asfsf|another", "|", 5)                                                                       //9 blocks, rounded up to 16
	createTreeAndTest(t, "this|is|some|data|separated|by|delimiters|asfsf|another|lol", "|", 5)                                                                   //10 blocks, rounded up to 16
	createTreeAndTest(t, "this|is|some|data|separated|by|delimiters|asfsf|asfasdf|asdfasdf|asdfasdf|asdfasdf|asdfasdf|asdfasdf|asdfasfd|sdfasd", "|", 5)          //16 blocks
	createTreeAndTest(t, "this|is|some|data|separated|by|delimiters|asfsf|asfasdf|asdfasdf|asdfasdf|asdfasdf|asdfasdf|asdfasdf|asdfasfd|sdfasd|dafgsagf", "|", 6) //17 blocks, rounded up to 32
}

func TestHashTree_withOneBlock_returnsError(t *testing.T) {

	//variables:
	text := "this"
	delimiter := "|"

	splittedData := bytes.Split([]byte(text), []byte(delimiter))
	tree, treeErr := NewBuilder().Create().WithBlocks(splittedData).Now()
	orderedData, orderedDataErr := tree.Order(splittedData)

	if orderedDataErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", orderedDataErr.Error())
		return
	}

	if !reflect.DeepEqual(orderedData, splittedData) {
		t.Errorf("the ordered data was invalid")
		return
	}

	if treeErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned")
		return
	}

	if tree.Height() != 2 {
		t.Errorf("the height of the tree was edxpected to be 2, %d returned", tree.Height())
		return
	}

	if tree.Length() != 2 {
		t.Errorf("the length of the tree was edxpected to be 2, %d returned", tree.Length())
		return
	}
}

func TestHashTree_convertToJSON_backAndForth_Success(t *testing.T) {

	//variables:
	jsEmpty := new(hashtree)
	r := rand.New(rand.NewSource(99))
	blks := [][]byte{
		[]byte("this"),
		[]byte("is"),
		[]byte("some"),
		[]byte("blocks"),
		[]byte(fmt.Sprintf("some rand number to make it unique: %d", r.Int())),
	}

	//execute:
	h, htErr := NewBuilder().Create().WithBlocks(blks).Now()
	if htErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", htErr.Error())
		return
	}

	if h == nil {
		t.Errorf("the returned instance was expected to be valid, nil returned")
	}

	js, jsErr := json.Marshal(h)
	if jsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", jsErr.Error())
	}

	newErr := json.Unmarshal(js, jsEmpty)
	if newErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", newErr.Error())
	}

	backJS, backJSErr := json.Marshal(jsEmpty)
	if backJSErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", backJSErr.Error())
	}

	if !reflect.DeepEqual(js, backJS) {
		t.Errorf("the json conversion (back and forth) did not succeed.  \n Expected: %v, \n Returned: %v\n", js, backJS)
	}

	comp := h.Compact()
	js, err := json.Marshal(comp)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
	}

	reCompact := new(compact)
	err = json.Unmarshal(js, reCompact)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
	}

	reJS, _ := json.Marshal(reCompact)
	if !reflect.DeepEqual(js, reJS) {
		t.Errorf("the json conversion (back and forth) did not succeed.  \n Expected: %v, \n Returned: %v\n", js, reJS)
	}
}
