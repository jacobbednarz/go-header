package models

import (
	"regexp"

	"github.com/go-header/messages"
)

type Rule struct {
	Template      string `yaml:"template"`
	PathMatcher   string `yaml:"path_matcher"`
	AuthorMatcher string `yaml:"author_matcher"`
	authorMatcher *regexp.Regexp
	pathMatcher   *regexp.Regexp
}

func (r *Rule) Compile() messages.ErrorList {
	result := messages.NewErrorList()
	var err error
	if r.PathMatcher != "" {
		if r.pathMatcher, err = regexp.Compile(r.PathMatcher); err != nil {
			result.Append(err)
		}
	}
	if r.AuthorMatcher != "" {
		if r.authorMatcher, err = regexp.Compile(r.AuthorMatcher); err != nil {
			result.Append(err)
		}
	}
	return result
}

func (r Rule) Match(s *Source) bool {
	if r.pathMatcher != nil {
		if !r.pathMatcher.MatchString(s.Path) {
			return false
		}
	}
	if r.authorMatcher != nil {
		if !r.authorMatcher.MatchString(s.Author) {
			return false
		}
	}
	return true
}