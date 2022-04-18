package pdf

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPdf_Basics(t *testing.T) {
	var book *Book
	var err error

	book, err = OpenBook("testdata/test.pdf")
	assert.NotNil(t, book.sectionBoundaries)
	assert.NotNil(t, book.sections)
	assert.NotNil(t, book.blankouts)

	assert.NotNil(t, book, "expected Book not to be nil")
	assert.NoError(t, err, "expected no error")

	err = book.LoadBlankoutsFromFile("testdata/blankouts.txt")
	assert.NoError(t, err)
	expectBlankouts := []string{
		"Something in section one.",
	}
	assert.Equal(t, book.GetBlankouts(), expectBlankouts, "expected blankouts to be equal")

	err = book.LoadSectionBoundaries("testdata/boundaries.json")
	assert.NoError(t, err)

	expectBoundaries := sectionList{
		{"ONE", "Fusce sagittis", "End of section one.", 1},
		{"TWO", "Aliquam erat", "End of section two.", 2},
	}
	assert.Equal(t, book.sectionBoundaries, expectBoundaries, "expected boundaries to be equal")

	err = book.Read()
	assert.NoError(t, err)
	assert.NotNil(t, book.buf)

	fixes := SectionFixList{
		{
			Match: regexp.MustCompile(`([^\s.].*?)?📒+`),
			Fix:   "\n\n$1⏱\n",
		},
		{
			Match: regexp.MustCompile(`[^.]?📒\s*`),
			Fix:   " ",
		},
		{
			Match: regexp.MustCompile(`\s+⏱\s*`),
			Fix:   " ",
		},
		{
			Match: regexp.MustCompile(`⏱\s*`),
			Fix:   "\n",
		},
	}

	err = book.ParseSections(fixes)
	assert.NoError(t, err)

	expectSections := map[string]string{
		"ONE": "Fusce sagittis, libero non molestie mollis, magna orci ultrices dolor, at vulputate neque nulla lacinia eros.",
		"TWO": `Aliquam erat volutpat. Nunc eleifend leo vitae magna. In id erat non orci commodo lobortis. Proin neque massa, cursus ut, gravida ut, lobortis eget, lacus. Sed diam. Praesent fermentum tempor tellus. Nullam tempus. Mauris ac felis vel velit tristique imperdiet. Donec at pede. Etiam vel neque nec dui dignissim bibendum. Vivamus id enim. Phasellus neque orci, porta a, aliquet quis, semper a, massa. Phasellus purus. Pellentesque tristique imperdiet tortor. Nam euismod tellus id ero. Nullam eu ante vel est convallis dignissim. Fusce suscipit, wisi nec facilisis facilisis, est dui fermentum leo, quis tempor ligula erat quis odio. Nunc porta vulputate tellus. Nunc rutrum turpis sed pede. Sed bibendum. Aliquam posuere. Nunc aliquet, augue nec adipiscing interdum, lacus tellus malesuada massa, quis varius mi purus non odio. Pellentesque condimentum, magna ut suscipit hendrerit, ipsum augue ornare nulla, non luctus diam neque sit amet urna. Curabitur vulputate vestibulum lorem. Fusce sagi, libero non molestie mollis, magna orci ultrices dolor, at vulputate neque nulla lacinia ebum. Sed id ligula quis est convallis tempor. Curabitur lacinia pulvinar nibh. Nam a sapien. Liquam erat volutpat. Nunc eleifend leo vitae magna. In id erat non orci commodo lobortis. Proin neque massa, cursus ut, gravida ut, lobortis eget, lacus. Sed diam. Praesent fermentum tempor tellus. Nullam tempus. Mauris ac felis vel velit tristique imperdiet. Donec at pede. Etiam vel neque nec dui dignissim bibendum. Vivamus id enim. Phasellus neque orci, porta a, aliquet quis, semper a, massa. Phasellus purus. Pellentesque tristique imperdiet tortor. Nam euismod tellus id erat.`,
	}

	for section, expect := range expectSections {
		got := book.GetSection(section)
		assert.Equal(t, expect, got, "expected section to be equal")
	}

	err = book.Close()
	assert.NoError(t, err)
	assert.Nil(t, book.buf)
	assert.Nil(t, book.sectionBoundaries)
	assert.Nil(t, book.sections)
	assert.Nil(t, book.blankouts)
}
