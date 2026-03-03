package btree

type BTree struct {
	//pointer to the tree's node
	root []uint64

	get func(uint16) BNode
	new func(BNode) uint64
	del func(uint64)
}
