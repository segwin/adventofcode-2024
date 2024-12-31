package day9

import (
	"fmt"
	"slices"
)

type Solution struct {
	DiskMap []int
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 9:\n")

	uncompacted := ParseDiskMap(s.DiskMap)
	compactedByBlock := CompactBlocks(slices.Clone(uncompacted))

	fmt.Print("  PART 1:\n")
	fmt.Printf("    Compacted checksum: %d\n", Checksum(compactedByBlock))

	compactedByFile := CompactFiles(slices.Clone(uncompacted))

	fmt.Print("  PART 2:\n")
	fmt.Printf("    Compacted without fragmentation checksum: %d\n", Checksum(compactedByFile))

	return nil
}

// ParseDiskMap builds the memory layout described by the given disk map.
func ParseDiskMap(diskMap []int) (blocks []*int) {
	for i := range diskMap {
		if i%2 == 0 {
			// even entries describe blocks
			blocks = append(blocks, newFile(i/2, diskMap[i])...)
		} else {
			// odd entries describe free space
			blocks = append(blocks, newFreeSpace(diskMap[i])...)
		}
	}
	return blocks
}

// CompactBlocks returns a new memory layout that maximises free space to the right, but may fragment files.
func CompactBlocks(blocks []*int) []*int {
	lastIdx, ok := rightmostOccupiedBlock(blocks, 0, len(blocks))
	if !ok {
		return blocks // no occupied blocks, nothing to do
	}

	for i := 0; i < lastIdx; i++ {
		if blocks[i] != nil {
			continue // already occupied
		}

		blocks[i] = blocks[lastIdx]
		blocks[lastIdx] = nil

		// find next rightmost block
		lastIdx, ok = rightmostOccupiedBlock(blocks, i+1, lastIdx)
		if !ok {
			break // only free space remains
		}
	}

	return blocks
}

// CompactFiles returns a new memory layout that maximises free space to the right without fragmenting files.
func CompactFiles(blocks []*int) []*int {
	var (
		fileStart = len(blocks)
		fileEnd   int
		ok        bool
	)
	for {
		// find next rightmost file
		fileStart, fileEnd, ok = rightmostFile(blocks, 0, fileStart)
		if !ok {
			return blocks // no remaining files
		}

		// find free space that fits it
		freeStart, freeEnd, ok := leftmostFreeSpace(blocks, 0, fileStart, fileEnd-fileStart)
		if !ok {
			continue // no free space is suitable for this file
		}

		fileIdx := fileStart
		for i := freeStart; i < freeEnd; i++ {
			blocks[i] = blocks[fileIdx]
			blocks[fileIdx] = nil
			fileIdx++
		}
	}
}

func Checksum(blocks []*int) int {
	sum := 0
	for i, b := range blocks {
		if b == nil {
			continue
		}
		sum += i * (*b)
	}
	return sum
}

func rightmostOccupiedBlock(blocks []*int, start, end int) (idx int, ok bool) {
	for i := end - 1; i >= start; i-- {
		if block := blocks[i]; block != nil {
			return i, true // found last occupied block
		}
	}
	return -1, false // no occupied blocks
}

func rightmostFile(blocks []*int, start, end int) (fileStart, fileEnd int, ok bool) {
	lastIdx, ok := rightmostOccupiedBlock(blocks, start, end)
	if !ok {
		return start, end, false
	}

	fileStart, fileEnd = start, lastIdx+1
	for i := fileEnd - 1; i >= start; i-- {
		if blocks[i] == nil || *blocks[i] != *blocks[lastIdx] {
			fileStart = i + 1 // passed file's first block
			break
		}
	}

	return fileStart, fileEnd, true
}

func leftmostFreeSpace(blocks []*int, start, end int, minSize int) (freeStart, freeEnd int, ok bool) {
	for i := start; i < end; i++ {
		if block := blocks[i]; block != nil {
			freeStart = i + 1
			continue // not a free space
		}

		if (i+1)-freeStart >= minSize {
			return freeStart, i + 1, true // don't bother looking for additional free blocks, we already met minSize
		}
	}
	return -1, -1, false // no occupied blocks
}

func newFile(id int, size int) []*int {
	block := newFreeSpace(size)
	for i := range block {
		block[i] = &id
	}
	return block
}

func newFreeSpace(size int) []*int {
	return make([]*int, size)
}
