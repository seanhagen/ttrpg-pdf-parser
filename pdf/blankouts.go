package pdf

import (
	"bufio"
	"fmt"
	"os"
)

// LoadBlankoutsFromFile ...
func (b *Book) LoadBlankoutsFromFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("unable to read blankouts from file: %w", err)
	}

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		b.blankouts = append(b.blankouts, sc.Text())
	}

	return sc.Err()
}
