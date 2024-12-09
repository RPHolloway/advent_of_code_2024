package main

import (
	"fmt"
	"os"
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

func test(disk_map []int) int {
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

	fmt.Println(checksum)
	return 0
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

	total := test(disk_map)
	fmt.Printf("Total: %d\r\n", total)
}
