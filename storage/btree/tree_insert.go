package btree

import "bytes"

func treeInsert(tree *BTree, node BNode, key []byte, val []byte) BNode {
	new := BNode{data: make([]byte, 2*BTREE_PAGE_SIZE)}

	idx := nodeLookupLE(node, key)

	switch node.btype() {
	case BNODE_LEAF:
		if bytes.Equal(key, node.getKey(idx)) {
			leafUpdate(new, node, idx, key, val)
		} else {
			leafInsert(new, node, idx+1, key, val)
		}

	case BNODE_NODE:
		nodeInsert(tree, new, node, idx, key, val)

	default:
		panic("bad node!")
	}
	return new
}
