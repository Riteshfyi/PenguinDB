package memory

import (
	"os"

	"PenguinDB/storage/btree"
)

// represents a db file
type KV struct {
	Path string
	fp   *os.File
	tree btree.BTree
	mmap struct {
		file   int      //file size
		total  int      //mmap size
		chunks [][]byte //multiple mmaps, can be non-continous
	}
	page struct {
		flushed uint64   //database size in number of pages
		temp    [][]byte //newloy allocated pages
	}
}
  