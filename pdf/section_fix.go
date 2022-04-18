package pdf

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// SectionFix ...
type SectionFix struct {
	Match *regexp.Regexp
	Fix   string
}

// UnmarshalJSON ...
func (sf *SectionFix) UnmarshalJSON(b []byte) error {
	var parts map[string]string
	err := json.Unmarshal(b, &parts)
	if err != nil {
		return fmt.Errorf("unable to unmarshal: %w", err)
	}

	re, err := regexp.Compile(parts["Match"])
	if err != nil {
		return fmt.Errorf("unable to compile regexp: %w", err)
	}

	*sf = SectionFix{
		Match: re,
		Fix:   parts["Fix"],
	}

	return nil
}
