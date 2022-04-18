package game

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// SectionFixList ...
type SectionFixList []SectionFix

var (
	sectionHeaderRE = regexp.MustCompile(`=+ START (\w+) =+\n\n(.*?)\n\n=+ END \w+ =+`)
)

const (
	startFmt = "\n\n===== START %v =====\n\n%v"
	endFmt   = "%v\n\n===== END %v =====\n\n"

	unknownSection = "unknown section"
)

// LoadSectionFixes ...
func (b *Book) LoadSectionFixes(path string) error {
	f, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("unable to read section fixes: %w", err)
	}

	err = json.Unmarshal(f, &b.sectionFixes)
	if err != nil {
		return fmt.Errorf("unable to unmarshal section fixes: %w", err)
	}
	return nil
}

// ParseSections ...
func (b *Book) ParseSections() error {
	if b.buf == nil {
		return fmt.Errorf("buffer is nil, have you called Read()?")
	}
	txt := b.buf.String()

	for _, f := range b.sectionFixes {
		txt = f.Match.ReplaceAllString(txt, f.Fix)
	}

	txt = strings.ReplaceAll(txt, "\n", " - ")

	for _, r := range b.blankouts {
		txt = strings.ReplaceAll(txt, r, " ")
	}

	sort.Sort(b.sectionBoundaries)

	for _, sec := range b.sectionBoundaries {
		txt = strings.ReplaceAll(txt, sec.Start, fmt.Sprintf(startFmt, sec.Name, sec.Start))
		txt = strings.ReplaceAll(txt, sec.End, fmt.Sprintf(endFmt, "", sec.Name))
	}

	findings := sectionHeaderRE.FindAllStringSubmatch(txt, -1)

	for _, sectionParts := range findings {
		fixed := strings.ReplaceAll(sectionParts[2], "-   -  - ", " ")
		fixed = strings.ReplaceAll(fixed, " - ", "\n")
		b.sections[sectionParts[1]] = strings.TrimSpace(fixed)
	}

	return nil
}

// GetSection ...
func (b *Book) GetSection(sec string) string {
	s, ok := b.sections[sec]
	if !ok {
		return unknownSection
	}
	return s
}
