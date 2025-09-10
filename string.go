package filter

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/gosimple/slug"
	"github.com/jinzhu/inflection"
)

// Default sets a default value for an empty string.
func Default(input, defaultValue string) string {
	if input == "" {
		return defaultValue
	}
	return input
}

// Trim strips leading and trailing whitespace from a string.
func Trim(input string) string {
	return strings.TrimSpace(input)
}

// Split splits a string into a slice of strings based on a delimiter.
func Split(input, delimiter string) []string {
	return strings.Split(input, delimiter)
}

// Replace replaces all occurrences of a substring with another string in the input string.
func Replace(input, old, replacement string) string {
	if old == "" {
		return input
	}
	return strings.ReplaceAll(input, old, replacement)
}

// Remove removes all occurrences of a substring from a string.
func Remove(input, toRemove string) string {
	return strings.ReplaceAll(input, toRemove, "")
}

// Append appends characters to the end of a string.
func Append(input, toAppend string) string {
	return input + toAppend
}

// Prepend prepends characters to the beginning of a string.
func Prepend(input, toPrepend string) string {
	return toPrepend + input
}

// Length returns the length of the input string.
func Length(input string) int {
	return utf8.RuneCountInString(input)
}

// Upper converts a string input to uppercase.
func Upper(input string) string {
	return strings.ToUpper(input)
}

// Lower converts a string input to lowercase.
func Lower(input string) string {
	return strings.ToLower(input)
}

// Titleize capitalizes the start of each part of the string.
func Titleize(input string) string {
	parts := toParts(input, defaultSpaceRunes, true)
	var result strings.Builder

	for i, part := range parts {
		if i > 0 {
			result.WriteByte(' ')
		}

		runes := []rune(part)
		for j, r := range runes {
			if j == 0 || (j > 0 && runes[j-1] == '-') {
				result.WriteRune(unicode.ToTitle(r))
			} else {
				result.WriteRune(r)
			}
		}
	}

	return result.String()
}

// Capitalize will capitalize the first letter of the string.
func Capitalize(input string) string {
	if input == "" {
		return ""
	}

	firstRune, size := utf8.DecodeRuneInString(input)
	capitalizedFirstRune := unicode.ToTitle(firstRune)

	return string(capitalizedFirstRune) + input[size:]
}

// Camelize converts a string to camelCase, lowercasing the first letter of the first segment and capitalizing the first letter of each subsequent segment.
func Camelize(input string) string {
	parts := toParts(input, defaultSpaceRunes, true)
	var builder strings.Builder

	for i, part := range parts {
		if acronym, ok := baseAcronyms[strings.ToUpper(part)]; ok {
			builder.WriteString(acronym)
			continue
		}

		var tempBuilder strings.Builder
		var capped bool
		for _, c := range part {
			if unicode.IsLetter(c) || unicode.IsDigit(c) {
				if i == 0 && !capped {
					tempBuilder.WriteRune(unicode.ToLower(c))
					capped = true
					continue
				}
				if !capped {
					capped = true
					tempBuilder.WriteRune(unicode.ToUpper(c))
					continue
				}
				tempBuilder.WriteRune(c)
			}
		}
		if tempBuilder.Len() > 0 {
			builder.WriteString(tempBuilder.String())
		}
	}
	return builder.String()
}

// Pascalize converts a string to PascalCase, capitalizing the first letter of each segment.
func Pascalize(input string) string {
	if input == "" {
		return ""
	}

	camelized := Camelize(input)

	r, size := utf8.DecodeRuneInString(camelized)
	return string(unicode.ToUpper(r)) + camelized[size:]
}

// Dasherize converts a string to a lowercased, dashed string, removing non-alphanumeric characters.
func Dasherize(input string) string {
	var result []string
	parts := toParts(input, defaultSpaceRunes, true)
	for _, part := range parts {
		var builder strings.Builder
		for _, c := range part {
			if unicode.IsLetter(c) || unicode.IsDigit(c) {
				builder.WriteRune(unicode.ToLower(c))
			}
		}
		if builder.Len() > 0 {
			result = append(result, builder.String())
		}
	}
	return strings.Join(result, "-")
}

// Slugify converts a string into a URL-friendly "slug", transliterating Unicode characters to ASCII and replacing or removing special characters.
func Slugify(input string) string {
	return slug.Make(input)
}

// Pluralize outputs the singular or plural version of a string based on the value of a number.
func Pluralize(count int, singular, plural string) string {
	// Handle the case when count is exactly one.
	if count == 1 {
		if singular != "" {
			return formatPluralizeWithCount(count, singular)
		}
		return formatPluralizeWithCount(count, inflection.Singular(plural))
	}

	// Handle the plural case.
	if plural != "" {
		return formatPluralizeWithCount(count, plural)
	}
	return formatPluralizeWithCount(count, inflection.Plural(singular))
}

// formatPluralizeWithCount formats the pluralization result string with the count if necessary.
func formatPluralizeWithCount(count int, result string) string {
	if strings.Contains(result, "%d") {
		return fmt.Sprintf(result, count)
	}
	return result
}

// Ordinalize converts a numeric input to its ordinal English version as a string.
func Ordinalize(number int) string {
	suffix := "th"
	switch number % 100 {
	case 11, 12, 13:
	default:
		switch number % 10 {
		case 1:
			suffix = "st"
		case 2:
			suffix = "nd"
		case 3:
			suffix = "rd"
		}
	}
	return strconv.Itoa(number) + suffix
}

// Truncate truncates a string to a specified length and appends "..." if it was longer.
func Truncate(input string, maxLength int) string {
	if maxLength <= 0 {
		return ""
	}

	if utf8.RuneCountInString(input) <= maxLength {
		return input
	}

	truncated := make([]rune, 0, maxLength)
	count := 0
	for _, r := range input {
		if count == maxLength {
			break
		}
		truncated = append(truncated, r)
		count++
	}

	return string(truncated) + "..."
}

// TruncateWords truncates a string to a specified number of words and appends "..." if it was longer.
func TruncateWords(input string, maxWords int) string {
	if maxWords <= 0 {
		return ""
	}

	var words []string
	var buf strings.Builder

	appendWord := func() {
		word := buf.String()
		onlyBoundaries := true
		for _, r := range word {
			if !unicode.IsSpace(r) && !unicode.IsPunct(r) {
				onlyBoundaries = false
				break
			}
		}
		if word != "" && !onlyBoundaries {
			words = append(words, word)
			buf.Reset()
		}
	}

	for _, r := range input {
		// Check if the character is a space or punctuation, treating them as word boundaries.
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			buf.WriteRune(r)
			appendWord()
		} else {
			buf.WriteRune(r)
		}
	}
	// Append any final word in the buffer.
	appendWord()

	if len(words) <= maxWords {
		return input
	}

	truncated := strings.Join(words[:maxWords], "")
	if len(truncated) > 0 {
		truncated = truncated[:len(truncated)-1]
	}
	return truncated + "..."
}

var defaultSpaceRunes = []rune{'_', ' ', ':', '-', '/'}

// isSpace checks if a rune is a word separator.
func isSpace(c rune, spaces []rune) bool {
	for _, r := range spaces {
		if r == c {
			return true
		}
	}
	return unicode.IsSpace(c)
}

// Appends parts of a string to a slice after trimming spaces and matching with base acronyms.
func appendPart(a []string, spaces []rune, ss ...string) []string {
	for _, s := range ss {
		s = strings.TrimSpace(s)
		for _, x := range spaces {
			s = strings.Trim(s, string(x))
		}
		if acronym, ok := baseAcronyms[strings.ToUpper(s)]; ok {
			s = acronym
		}
		if s != "" {
			a = append(a, s)
		}
	}
	return a
}

// toParts splits a string into identifiable parts, considering acronyms and separators.
func toParts(s string, spaces []rune, splitOnUpperCase bool) []string {
	parts := []string{}
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return parts
	}
	if _, ok := baseAcronyms[strings.ToUpper(s)]; ok {
		return []string{strings.ToUpper(s)}
	}
	var prev rune
	var x strings.Builder
	x.Grow(len(s))
	for _, c := range s {
		if !utf8.ValidRune(c) {
			continue
		}

		if isSpace(c, spaces) {
			parts = appendPart(parts, spaces, x.String())
			x.Reset()
			x.WriteRune(c)
			prev = c
			continue
		}

		// Modify condition to include splitOnUpperCase check
		if splitOnUpperCase && unicode.IsUpper(c) && !unicode.IsUpper(prev) && x.Len() > 0 {
			parts = appendPart(parts, spaces, x.String())
			x.Reset()
		}

		_, found := baseAcronyms[strings.ToUpper(x.String())]
		if unicode.IsUpper(c) && found {
			parts = appendPart(parts, spaces, x.String())
			x.Reset()
		}

		if unicode.IsLetter(c) || unicode.IsDigit(c) || unicode.IsPunct(c) || c == '`' {
			prev = c
			x.WriteRune(c)
			continue
		}

		parts = appendPart(parts, spaces, x.String())
		x.Reset()
		prev = c
	}
	parts = appendPart(parts, spaces, x.String())

	return parts
}
