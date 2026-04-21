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

// Capitalize returns input with the first rune title-cased and the rest lower-cased.
// It matches Liquid's capitalize behavior.
func Capitalize(input string) string {
	if input == "" {
		return ""
	}
	runes := []rune(strings.ToLower(input))
	runes[0] = unicode.ToTitle(runes[0])
	return string(runes)
}

// Camelize converts input to camelCase.
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

// Pascalize converts input to PascalCase.
func Pascalize(input string) string {
	if input == "" {
		return ""
	}

	camelized := Camelize(input)

	r, size := utf8.DecodeRuneInString(camelized)
	return string(unicode.ToUpper(r)) + camelized[size:]
}

// Dasherize converts input to lowercase words joined by dashes.
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

// Slugify converts input to a URL-friendly slug.
func Slugify(input string) string {
	return slug.Make(input)
}

// Pluralize returns the singular or plural form selected by count.
// When the selected form contains %d, Pluralize formats count into it.
func Pluralize(count int, singular, plural string) string {
	result := plural
	if count == 1 {
		result = singular
		if result == "" {
			result = inflection.Singular(plural)
		}
	} else if result == "" {
		result = inflection.Plural(singular)
	}

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

// Truncate shortens input to maxLength runes.
// It uses "..." when ellipsis is omitted.
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

// TruncateWords shortens input to maxWords words.
// It uses "..." when ellipsis is omitted.
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

var (
	htmlScriptRe  = regexp.MustCompile(`(?is)<script.*?</script>`)
	htmlStyleRe   = regexp.MustCompile(`(?is)<style.*?</style>`)
	htmlCommentRe = regexp.MustCompile(`(?s)<!--.*?-->`)
	htmlTagRe     = regexp.MustCompile(`<[^>]*>`)
)

// Escape escapes HTML special characters in input.
func Escape(input string) string {
	return html.EscapeString(input)
}

// EscapeOnce escapes HTML special characters without double-escaping entities.
func EscapeOnce(input string) string {
	return html.EscapeString(html.UnescapeString(input))
}

// StripHTML removes tags, script blocks, style blocks, and comments from input.
func StripHTML(input string) string {
	s := htmlScriptRe.ReplaceAllString(input, "")
	s = htmlStyleRe.ReplaceAllString(s, "")
	s = htmlCommentRe.ReplaceAllString(s, "")
	s = htmlTagRe.ReplaceAllString(s, "")
	return s
}

var newlineReplacer = strings.NewReplacer("\r\n", "", "\r", "", "\n", "")

// StripNewlines removes newline characters from input.
func StripNewlines(input string) string {
	return newlineReplacer.Replace(input)
}

// TrimLeft removes leading whitespace from input.
// It matches Liquid's lstrip filter.
func TrimLeft(input string) string {
	return strings.TrimLeftFunc(input, unicode.IsSpace)
}

// TrimRight removes trailing whitespace from input.
// It matches Liquid's rstrip filter.
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

// Slice returns a substring or sub-slice starting at offset.
// Negative offsets count from the end. When length is omitted, Slice returns
// one element or rune.
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

// URLEncode percent-encodes a string for use in URLs.
func URLEncode(input string) string {
	return url.QueryEscape(input)
}

// URLDecode decodes a percent-encoded string.
func URLDecode(input string) (string, error) {
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

		if slices.Contains(spaces, c) || unicode.IsSpace(c) {
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
