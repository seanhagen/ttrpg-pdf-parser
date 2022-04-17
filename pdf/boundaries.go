package pdf

import (
	"encoding/json"
	"fmt"
	"os"
)

type sectionBoundary struct {
	Name  string
	Start string
	End   string
	Ord   int
}

type sectionList []sectionBoundary

// Len ...
func (sl sectionList) Len() int {
	return len(sl)
}

// Less  ...
func (sl sectionList) Less(i, j int) bool {
	return sl[i].Ord < sl[j].Ord
}

// Swap ...
func (sl sectionList) Swap(i, j int) {
	sl[i], sl[j] = sl[j], sl[i]
}

// LoadSectionBoundaries ...
func (b *Book) LoadSectionBoundaries(path string) error {
	f, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("unable to read section boundaries from file: %w", err)
	}

	err = json.Unmarshal(f, &b.sectionBoundaries)
	if err != nil {
		return fmt.Errorf("unable to unmarshal section boundaries: %w", err)
	}
	return nil
}
