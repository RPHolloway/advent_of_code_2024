package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"
)

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func isFileEntry(i int) bool {
	return i%2 == 0
}

func test_method1(disk_map []int) int {
	checksum := 0
	file_idx := 0
	last_map_idx := len(disk_map) - 1

	for map_idx, entry := range disk_map {
		if map_idx > last_map_idx {
			break
		}

		if isFileEntry(map_idx) {
			file_id := map_idx / 2
			field_size := entry

			for x := 0; x < field_size; x++ {
				checksum += file_id * file_idx
				file_idx++
			}
		} else {
			// Free space
			// Get last file
			file_id := last_map_idx / 2
			field_size := entry

			for x := 0; x < field_size; x++ {
				checksum += file_id * file_idx
				file_idx++

				// Reduce the number of files at the end
				disk_map[last_map_idx]--
				// If we run out of files and still have free space move to the next entry
				if disk_map[last_map_idx] == 0 {
					last_map_idx -= 2
					file_id = last_map_idx / 2
				}
			}
		}
	}

	return checksum
}

func test_method2(disk_map []int) int {
	checksum := 0
	last_map_idx := len(disk_map) - 1

	var files, free []int
	for map_idx, entry := range disk_map {
		if isFileEntry(map_idx) {
			files = append(files, entry)
		} else {
			free = append(free, entry)
		}
	}

	slices.Reverse(files)
	file_count := len(files) - 1

	for i, file_size := range files {
		for j, free_space := range free {
			file_id := last_map_idx/2 - i

			if j >= (file_count - i) {
				file_idx := 0
				for x := 0; x < j*2; x++ {
					file_idx += disk_map[x]
				}

				for x := 0; x < file_size; x++ {
					// calculate checksum
					checksum += file_id * file_idx
					file_idx++
				}
				break
			}

			if file_size <= free_space {
				file_idx := disk_map[0]
				for x := 0; x < j*2; x++ {
					file_idx += disk_map[x+1]
				}

				for x := 0; x < file_size; x++ {
					// calculate checksum
					checksum += file_id * file_idx
					file_idx++
				}

				free[j] -= file_size
				disk_map[j*2] += file_size
				disk_map[j*2+1] -= file_size
				break
			}
		}
	}

	return checksum
}

func main() {
	var disk_map []int

	// Read input
	//data, _ := os.ReadFile("example.txt")
	data, _ := os.ReadFile("input.txt")
	input := string(data)

	// Parse input
	for _, c := range input {
		v, _ := strconv.Atoi(string(c))
		disk_map = append(disk_map, v)
	}

	defer timeTrack(time.Now())

	//method1 := test_method1(disk_map)
	method2 := test_method2(disk_map)

	//fmt.Printf("Method 1: %d\r\n", method1)
	fmt.Printf("Method 2: %d\r\n", method2)
}
