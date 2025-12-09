package utils

import "os"

// ========================
// READ FILE
// ========================
func ReadFile(title string) string {
	data, err := os.ReadFile(title)
	if err != nil {
		panic("Error reading file")
	}
	return string(data)
}

// ========================
// PRINT GRID TO FILE
// ========================
func PrintGridToFile(grid [][]string, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	for y := range grid {
		for x := range grid[y] {
			_, err := f.WriteString(grid[y][x])
			if err != nil {
				return err
			}
		}
		_, err := f.WriteString("\n")
		if err != nil {
			return err
		}
	}
	return nil
}
