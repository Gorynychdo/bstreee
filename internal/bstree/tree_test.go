package bstree

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestTree_Insert(t *testing.T) {
    tree := NewTree(nil)
    tree.Insert(4)
    tree.Insert(4)
    tree.Insert(2)
    tree.Insert(6)
    tree.Insert(1)
    tree.Insert(3)
    tree.Insert(5)
    tree.Insert(7)
    tree.Insert(2)
    assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, tree.dump())
}

func TestTree_Search(t *testing.T) {
    tree := NewTree([]int{5, 2, 7, 1, 4, 6, 9, 3})
    assert.True(t, tree.Search(5))
    assert.True(t, tree.Search(1))
    assert.True(t, tree.Search(9))
    assert.False(t, tree.Search(10))
}

func TestTree_Delete(t *testing.T) {
    tree := NewTree([]int{5, 2, 7, 1, 6, 9, 3, 8, 4})

    tree.Delete(10)
    assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, tree.dump())

    tree.Delete(6)
    assert.Equal(t, []int{1, 2, 3, 4, 5, 7, 8, 9}, tree.dump())

    tree.Delete(5)
    assert.Equal(t, []int{1, 2, 3, 4, 7, 8, 9}, tree.dump())

    tree.Delete(9)
    assert.Equal(t, []int{1, 2, 3, 4, 7, 8}, tree.dump())

    tree.Delete(3)
    tree.Delete(4)
    assert.Equal(t, []int{1, 2, 7, 8}, tree.dump())
}
