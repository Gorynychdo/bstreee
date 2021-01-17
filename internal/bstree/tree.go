package bstree

import (
    "sync"

    log "github.com/sirupsen/logrus"
)

// Tree bynary search tree
type Tree struct {
    sync.RWMutex
    root *node
}

// NewTree BTree constructor from slice
func NewTree(slice []int) *Tree {
    var t Tree
    for _, value := range slice {
        t.root = t.root.insert(value)
    }
    t.log("tree initialized", 0)
    return &t
}

// Insert BTree value
func (t *Tree) Insert(value int) {
    t.Lock()
    t.root = t.root.insert(value)
    t.Unlock()
    t.log("insert value", value)
}

// Search BTree value
func (t *Tree) Search(value int) bool {
    t.RLock()
    _, n := t.root.search(value)
    t.RUnlock()
    t.log("search value", value)
    return n != nil
}

// Delete BTree value
func (t *Tree) Delete(value int) {
    t.Lock()
    t.root = t.root.delete(value)
    t.Unlock()
    t.log("delete value", value)
}

// dump BTree to slice
func (t *Tree) dump() []int {
    var s []int
    t.root.traverse(&s)
    return s
}

func (t *Tree) log(mes string, value int) {
    log.WithFields(log.Fields{
        "package": "bstree",
        "value":   value,
        "dump":    t.dump(),
    }).Debug(mes)
}

type node struct {
    value int
    left  *node
    right *node
}

// insert returns new changed node
func (n *node) insert(value int) *node {
    switch {
    case n == nil:
        return &node{value: value}
    case value < n.value:
        n.left = n.left.insert(value)
    case value > n.value:
        n.right = n.right.insert(value)
    }
    return n
}

// search returns node with it`s parent
func (n *node) search(value int) (parent, target *node) {
    target = n
    for {
        switch {
        case target == nil:
            return nil, nil
        case value == target.value:
            return parent, target
        case value < target.value:
            parent = target
            target = target.left
        case value > target.value:
            parent = target
            target = target.right
        }
    }
}

// delete returns new changed node
func (n *node) delete(value int) *node {
    parent, target := n.search(value)

    switch {
    case target == nil:
        return n
    // if no left child replace target with right child
    case target.left == nil:
        return n.reconstruct(parent, target, target.right)
    // if left child exists:
    // if no right child replace target with left child
    case target.right == nil:
        return n.reconstruct(parent, target, target.left)
    // if both children exists:
    // if right child of left child is empty:
    // replace target with left child
    case target.left.right == nil:
        target.left.right = target.right
        return n.reconstruct(parent, target, target.left)
    }

    // if right child of left child exists:
    // go to the edge of right nodes chain
    // replace target with right edge node
    // and right edge node replace with it`s left child
    prev := target.left
    next := target.left.right

    for {
        if next.right == nil {
            prev.right = next.left
            next.left = target.left
            next.right = target.right
            return n.reconstruct(parent, target, next)
        }
        prev = next
        next = next.right
    }
}

// reconstruct replace child for parent from current to replaced node
// return new changed node
func (n *node) reconstruct(parent, current, replaced *node) *node {
    switch {
    case parent == nil:
        return replaced
    case parent.left == current:
        parent.left = replaced
    case parent.right == current:
        parent.right = replaced
    }
    return n
}

// traverse node & dump values to slice
func (n *node) traverse(slice *[]int) {
    if n == nil {
        return
    }

    n.left.traverse(slice)
    *slice = append(*slice, n.value)
    n.right.traverse(slice)
}
