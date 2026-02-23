package filter

import (
	"encoding/base64"
	"fmt"
	"html"
	"net/url"
	"reflect"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/gosimple/slug"
	"github.com/jinzhu/inflection"
)

var defaultSpaceRunes = []rune{'_', ' ', ':', '-', '/'}

// Default returns defaultValue if input is nil, false, or empty string.
func Default(input, defaultValue any) any {
	if input == nil {
		return defaultValue
	}
	if b, ok := input.(bool); ok && !b {
		return defaultValue
	}
	if s, ok := input.(string); ok && s == "" {
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
	result := strings.Builder{}
	result.Grow(len(input))

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

// Capitalize capitalizes the first letter and lowercases the rest.
// This matches Liquid's capitalize behavior.
func Capitalize(input string) string {
	if input == "" {
		return ""
	}
	runes := []rune(strings.ToLower(input))
	runes[0] = unicode.ToTitle(runes[0])
	return string(runes)
}

// Camelize converts a string to camelCase. It lowercases the first letter
// of the first segment and capitalizes the first letter of each subsequent segment.
func Camelize(input string) string {
	parts := toParts(input, defaultSpaceRunes, true)
	builder := strings.Builder{}
	builder.Grow(len(input))

	for i, part := range parts {
		if acronym, ok := baseAcronyms[strings.ToUpper(part)]; ok {
			builder.WriteString(acronym)
			continue
		}

		tempBuilder := strings.Builder{}
		capped := false
		for _, c := range part {
			if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
				continue
			}

			if !capped {
				if i == 0 {
					tempBuilder.WriteRune(unicode.ToLower(c))
				} else {
					tempBuilder.WriteRune(unicode.ToUpper(c))
				}
				capped = true
			} else {
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

// Dasherize converts a string to a lowercased, dashed string,
// removing non-alphanumeric characters.
func Dasherize(input string) string {
	parts := toParts(input, defaultSpaceRunes, true)
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		builder := strings.Builder{}
		builder.Grow(len(part))
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

// Slugify converts a string into a URL-friendly "slug", transliterating
// Unicode characters to ASCII and replacing or removing special characters.
func Slugify(input string) string {
	return slug.Make(input)
}

// Pluralize returns the singular or plural form of a string based on count.
// If the form string contains "%d", the count is substituted into the result.
// Otherwise, only the appropriate form string is returned without the count.
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

// Truncate truncates a string to maxLength characters (including ellipsis).
// An optional ellipsis string can be provided (default "...").
func Truncate(input string, maxLength int, ellipsis ...string) string {
	omission := "..."
	if len(ellipsis) > 0 {
		omission = ellipsis[0]
	}
	if maxLength <= 0 {
		return ""
	}
	runes := []rune(input)
	if len(runes) <= maxLength {
		return input
	}
	omissionRunes := []rune(omission)
	if maxLength <= len(omissionRunes) {
		return string(omissionRunes[:maxLength])
	}
	return string(runes[:maxLength-len(omissionRunes)]) + omission
}

// TruncateWords truncates a string to a specified number of words.
// An optional ellipsis string can be provided (default "...").
func TruncateWords(input string, maxWords int, ellipsis ...string) string {
	omission := "..."
	if len(ellipsis) > 0 {
		omission = ellipsis[0]
	}
	if maxWords <= 0 {
		return ""
	}

	var words []string
	buf := strings.Builder{}

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
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			buf.WriteRune(r)
			appendWord()
		} else {
			buf.WriteRune(r)
		}
	}
	appendWord()

	if len(words) <= maxWords {
		return input
	}

	truncated := strings.Join(words[:maxWords], "")
	if len(truncated) > 0 {
		truncated = truncated[:len(truncated)-1]
	}
	return truncated + omission
}

// Package-level compiled regexps for HTML processing.
var (
	htmlScriptRe  = regexp.MustCompile(`(?is)<script.*?</script>`)
	htmlStyleRe   = regexp.MustCompile(`(?is)<style.*?</style>`)
	htmlCommentRe = regexp.MustCompile(`(?s)<!--.*?-->`)
	htmlTagRe     = regexp.MustCompile(`<[^>]*>`)
)

// Escape converts <, >, &, ", ' to HTML entities.
func Escape(input string) string {
	return html.EscapeString(input)
}

// EscapeOnce converts <, >, &, ", ' to HTML entities without double-escaping.
// Already escaped entities like &amp; &lt; &#39; are preserved.
func EscapeOnce(input string) string {
	return html.EscapeString(html.UnescapeString(input))
}

// StripHTML removes all HTML tags, script blocks, style blocks, and comments from the input.
func StripHTML(input string) string {
	s := htmlScriptRe.ReplaceAllString(input, "")
	s = htmlStyleRe.ReplaceAllString(s, "")
	s = htmlCommentRe.ReplaceAllString(s, "")
	s = htmlTagRe.ReplaceAllString(s, "")
	return s
}

// newlineReplacer is a package-level replacer for stripping newlines,
// avoiding repeated allocation on each call.
var newlineReplacer = strings.NewReplacer("\r\n", "", "\r", "", "\n", "")

// StripNewlines removes all newline characters (\n, \r\n, \r) from the input.
func StripNewlines(input string) string {
	return newlineReplacer.Replace(input)
}

// TrimLeft removes leading whitespace from a string.
// Liquid equivalent: lstrip.
func TrimLeft(input string) string {
	return strings.TrimLeftFunc(input, unicode.IsSpace)
}

// TrimRight removes trailing whitespace from a string.
// Liquid equivalent: rstrip.
func TrimRight(input string) string {
	return strings.TrimRightFunc(input, unicode.IsSpace)
}

// ReplaceFirst replaces the first occurrence of old with replacement.
func ReplaceFirst(input, old, replacement string) string {
	if old == "" {
		return input
	}
	return strings.Replace(input, old, replacement, 1)
}

// ReplaceLast replaces the last occurrence of old with replacement.
func ReplaceLast(input, old, replacement string) string {
	if old == "" {
		return input
	}
	idx := strings.LastIndex(input, old)
	if idx < 0 {
		return input
	}
	return input[:idx] + replacement + input[idx+len(old):]
}

// RemoveFirst removes the first occurrence of a substring.
func RemoveFirst(input, toRemove string) string {
	return ReplaceFirst(input, toRemove, "")
}

// RemoveLast removes the last occurrence of a substring.
func RemoveLast(input, toRemove string) string {
	return ReplaceLast(input, toRemove, "")
}

// Slice extracts a substring or sub-slice.
// For strings: returns substring starting at offset with optional length.
// For slices: returns sub-slice starting at offset with optional length.
// Negative offset counts from end (-1 = last element).
// If length is omitted, returns single character/element.
func Slice(input any, offset int, length ...int) (any, error) {
	switch v := input.(type) {
	case string:
		return sliceString(v, offset, length...), nil
	default:
		rv := reflect.ValueOf(input)
		if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
			return nil, fmt.Errorf("%w: expected string or slice, got %T", ErrUnsupportedType, input)
		}
		return sliceReflect(rv, offset, length...), nil
	}
}

// sliceString extracts a substring from s starting at offset with optional size.
// Negative offset counts from end. Returns single rune if size is omitted.
func sliceString(s string, offset int, length ...int) string {
	runes := []rune(s)
	n := len(runes)
	if n == 0 {
		return ""
	}
	if offset < 0 {
		offset = n + offset
	}
	if offset < 0 || offset >= n {
		return ""
	}
	size := 1
	if len(length) > 0 {
		size = length[0]
	}
	if size <= 0 {
		return ""
	}
	end := min(offset+size, n)
	return string(runes[offset:end])
}

// sliceReflect extracts a sub-slice from rv starting at offset with optional size.
// Negative offset counts from end. Returns single element if size is omitted.
func sliceReflect(rv reflect.Value, offset int, length ...int) []any {
	n := rv.Len()
	if n == 0 {
		return []any{}
	}
	if offset < 0 {
		offset = n + offset
	}
	if offset < 0 || offset >= n {
		return []any{}
	}
	size := 1
	if len(length) > 0 {
		size = length[0]
	}
	if size <= 0 {
		return []any{}
	}
	end := min(offset+size, n)
	result := make([]any, end-offset)
	for i := offset; i < end; i++ {
		result[i-offset] = rv.Index(i).Interface()
	}
	return result
}

// UrlEncode percent-encodes a string for use in URLs.
func UrlEncode(input string) string {
	return url.QueryEscape(input)
}

// UrlDecode decodes a percent-encoded string.
func UrlDecode(input string) (string, error) {
	result, err := url.QueryUnescape(input)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrInvalidArguments, err)
	}
	return result, nil
}

// Base64Encode encodes a string to standard Base64.
func Base64Encode(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

// Base64Decode decodes a standard Base64 string.
func Base64Decode(input string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrInvalidArguments, err)
	}
	return string(b), nil
}

// isSpace checks if a rune is a word separator.
func isSpace(c rune, spaces []rune) bool {
	return slices.Contains(spaces, c) || unicode.IsSpace(c)
}

// appendPart appends parts of a string to a slice after trimming spaces.
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

// toParts splits a string into parts, considering acronyms and separators.
func toParts(s string, spaces []rune, splitOnUpperCase bool) []string {
	parts := []string{}
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return parts
	}
	if _, ok := baseAcronyms[strings.ToUpper(s)]; ok {
		return []string{strings.ToUpper(s)}
	}
	prev := rune(0)
	x := strings.Builder{}
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

// formatPluralizeWithCount formats the pluralization result with count if needed.
func formatPluralizeWithCount(count int, result string) string {
	if strings.Contains(result, "%d") {
		return fmt.Sprintf(result, count)
	}
	return result
}
