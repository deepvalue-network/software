package grammar

import (
	"strings"
)

type rule struct {
	name     string
	sections []RuleSection
	index    int
}

func createRule(name string, sections []RuleSection, index int) Rule {
	out := rule{
		name:     name,
		sections: sections,
		index:    index,
	}

	return &out
}

// Name returns the name
func (obj *rule) Name() string {
	return obj.name
}

// Sections returns the sections
func (obj *rule) Sections() []RuleSection {
	return obj.sections
}

// Index returns the index
func (obj *rule) Index() int {
	return obj.index
}

// FindFirst finds the first occurence of the string, using the rule sections
func (obj *rule) FindFirst(str string) (string, bool, error) {
	out := []string{}
	for _, oneSection := range obj.sections {
		found, err := oneSection.FindFirst(str)
		if err != nil {
			continue
		}

		if found != "" {
			index := strings.Index(str, found)
			if index != 0 {
				return "", false, nil
			}

			out = append(out, found)
			str = str[len(found):]
			continue
		}

		return "", false, nil
	}

	return strings.Join(out, ""), true, nil
}

// FindConsecutives returns the concsecutive occurences of the string, using the rule sections:
func (obj *rule) FindConsecutives(str string) (string, bool, error) {
	out := []string{}
	for {
		code, found, err := obj.FindFirst(str)
		if err != nil {
			return "", false, err
		}

		if !found {
			break
		}

		str = strings.Replace(str, code, "", 1)
		out = append(out, code)
	}

	return strings.Join(out, ""), true, nil
}
