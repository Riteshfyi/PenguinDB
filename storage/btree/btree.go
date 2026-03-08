package btree

type BTree struct {
	//pointer to the tree's node
	root []uint64

	get func(uint64) BNode //get a page from the disk
	new func(BNode) uint64 //allocate a page to the object
	del func(uint64)       //free the corresponding page
}
