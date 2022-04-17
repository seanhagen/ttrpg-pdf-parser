package pdf

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var (
	sectionHeaderRE = regexp.MustCompile(`=+ START (\w+) =+\n\n(.*?)\n\n=+ END \w+ =+`)
	sectionFixRE    = regexp.MustCompile(`([^\s.].*?)?📒+`)
	newlineFixRE    = regexp.MustCompile(`[^.]?📒\s*`)
	thirdFixRE      = regexp.MustCompile(`\s+⏱\s*`)
	fourthFixRE     = regexp.MustCompile(`⏱\s*`)
)

const (
	startFmt = "\n\n===== START %v =====\n\n%v"
	endFmt   = "%v\n\n=========== END %v ===========\n\n"

	unknownSection = "unknown section"

	spacedFix = `
⏱
`
)

// ParseSections ...
func (b *Book) ParseSections() error {
	if b.buf == nil {
		return fmt.Errorf("buffer is nil, have you called Read()?")
	}

	txt := b.buf.String()

	txt = sectionFixRE.ReplaceAllString(txt, "\n\n$1⏱\n")
	txt = newlineFixRE.ReplaceAllString(txt, " ")
	txt = thirdFixRE.ReplaceAllString(txt, " ")
	txt = fourthFixRE.ReplaceAllString(txt, "\n")
	// txt = strings.ReplaceAll(txt, spacedFix, " ")

	txt = strings.ReplaceAll(txt, "\n", " - ")

	// fmt.Printf(
	// 	"############################## full pdf:\n\n  %v\n\n##############################end of full pdf\n", txt)

	for _, r := range b.blankouts {
		txt = strings.ReplaceAll(txt, r, " ")
	}

	sort.Sort(b.sectionBoundaries)

	for _, sec := range b.sectionBoundaries {
		// fmt.Printf("section %v\nstarts from '%v'\ngoes to '%v'\n\n", sec.Name, sec.Start, sec.End)
		txt = strings.ReplaceAll(txt, sec.Start, fmt.Sprintf(startFmt, sec.Name, sec.Start))
		txt = strings.ReplaceAll(txt, sec.End, fmt.Sprintf(endFmt, sec.End, sec.Name))
	}

	findings := sectionHeaderRE.FindAllStringSubmatch(txt, -1)

	for _, sectionParts := range findings {
		fixed := strings.ReplaceAll(sectionParts[2], "-   -  - ", " ")
		fixed = strings.ReplaceAll(fixed, " - ", "\n")
		b.sections[sectionParts[1]] = fixed
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
