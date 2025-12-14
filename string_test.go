package filter

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefault(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		defaultValue string
		expected     string
	}{
		{"Empty Input", "", "default value", "default value"},
		{"Non-Empty Input", "actual value", "default value", "actual value"},
		{"Empty Default Value", "", "", ""},
		{"Non-Empty Input With Empty Default", "actual value", "", "actual value"},
		{"Both Empty", "", "", ""},
		{"Numerical Input and Default", "123", "456", "123"},
		{"Numerical Empty Input", "", "456", "456"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Default(tt.input, tt.defaultValue)
			if actual != tt.expected {
				t.Errorf("Test '%s' failed. Expected '%s', got '%s'", tt.name, tt.expected, actual)
			}
		})
	}
}

func TestTrim(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Leading Spaces", "  leading space", "leading space"},
		{"Trailing Spaces", "trailing space  ", "trailing space"},
		{"Leading and Trailing Spaces", "  both sides  ", "both sides"},
		{"Tab Characters", "\t\tstart and end\t\t", "start and end"},
		{"New Line Characters", "\nstart and end\n", "start and end"},
		{"Mixed Whitespace Characters", " \t\n mixed whitespace \n\t ", "mixed whitespace"},
		{"No Whitespace", "nowhitespace", "nowhitespace"},
		{"Only Whitespace", "   \t\n  ", ""},
		{"Empty String", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Trim(tt.input)
			if actual != tt.expected {
				t.Errorf("Test '%s' failed. Expected '%s', got '%s'", tt.name, tt.expected, actual)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		delimiter string
		expected  []string
	}{
		{"Single Space Delimiter", "a b c", " ", []string{"a", "b", "c"}},
		{"Comma Delimiter", "a,b,c", ",", []string{"a", "b", "c"}},
		{"No Delimiter Present", "abc", ",", []string{"abc"}},
		{"Multiple Delimiters", "a,,b,c", ",", []string{"a", "", "b", "c"}},
		{"Delimiter At Start", ",a,b,c", ",", []string{"", "a", "b", "c"}},
		{"Delimiter At End", "a,b,c,", ",", []string{"a", "b", "c", ""}},
		{"Empty String Input", "", ",", []string{""}},
		{"Empty String Delimiter", "abc", "", []string{"a", "b", "c"}},
		{"Multi-character Delimiter", "a<>b<>c", "<>", []string{"a", "b", "c"}},
		{"New Line Delimiter", "a\nb\nc", "\n", []string{"a", "b", "c"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Split(tt.input, tt.delimiter)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Test '%s' failed. Expected %+v, got %+v", tt.name, tt.expected, actual)
			}
		})
	}
}

func TestReplace(t *testing.T) {
	tests := []struct {
		input    string
		old      string
		new      string
		expected string
	}{
		{"hello world", "world", "gopher", "hello gopher"},
		{"hello hello hello", "hello", "hi", "hi hi hi"},
		{"", "hello", "hi", ""},
		{"hello world", "", "hi", "hello world"},
		{"hello world", "world", "", "hello "},
		{"hello world world", "world", "gopher", "hello gopher gopher"},
		{"123-456", "-", ":", "123:456"},
		{"foobarfoobar", "bar", "baz", "foobazfoobaz"},
	}

	for _, tt := range tests {
		actual := Replace(tt.input, tt.old, tt.new)
		if actual != tt.expected {
			t.Errorf("Replace(%q, %q, %q) = %q, expected %q", tt.input, tt.old, tt.new, actual, tt.expected)
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		input    string
		toRemove string
		expected string
	}{
		{"hello world", "world", "hello "},
		{"hello world", "l", "heo word"},
		{"an apple a day", "a", "n pple  dy"},
		{"spaces    everywhere", " ", "spaceseverywhere"},
		{"123-456-789", "-", "123456789"},
		{"no occurrence", "z", "no occurrence"},
		{"empty string", "", "empty string"},
		{"remove empty", "hello", "remove empty"},
	}

	for _, tt := range tests {
		actual := Remove(tt.input, tt.toRemove)
		if actual != tt.expected {
			t.Errorf("Remove(%q, %q) = %q, expected %q", tt.input, tt.toRemove, actual, tt.expected)
		}
	}
}

func TestAppend(t *testing.T) {
	tests := []struct {
		input    string
		toAppend string
		expected string
	}{
		{"hello", " world", "hello world"},
		{"", "world", "world"},
		{"hello", "", "hello"},
		{"123", "456", "123456"},
		{"", "", ""},
		{"multi", "-line\nnew line", "multi-line\nnew line"},
		{"special", " chars\n\t", "special chars\n\t"},
	}

	for _, tt := range tests {
		actual := Append(tt.input, tt.toAppend)
		if actual != tt.expected {
			t.Errorf("Append(%q, %q) = %q, expected %q", tt.input, tt.toAppend, actual, tt.expected)
		}
	}
}

func TestPrepend(t *testing.T) {
	tests := []struct {
		input     string
		toPrepend string
		expected  string
	}{
		{"world", "hello ", "hello world"},
		{"world", "", "world"},
		{"", "hello", "hello"},
		{"456", "123", "123456"},
		{"", "", ""},
		{"line\nnew line", "multi-", "multi-line\nnew line"},
		{"chars\n\t", "special ", "special chars\n\t"},
	}

	for _, tt := range tests {
		actual := Prepend(tt.input, tt.toPrepend)
		if actual != tt.expected {
			t.Errorf("Prepend(%q, %q) = %q, expected %q", tt.input, tt.toPrepend, actual, tt.expected)
		}
	}
}

func TestLength(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello world", 11},
		{"", 0},
		{"こんにちは", 5},  // Japanese for "hello"
		{"😊😂🥺😍😒😘", 6}, // 6 emoji characters
		{"1234567890", 10},
		{"special chars\n\t", 15},
		{"multi-line\nnew line", 19},
		{"with spaces ", 12},
		{" leading space", 14},
		{"trailing space ", 15},
	}

	for _, tt := range tests {
		actual := Length(tt.input)
		if actual != tt.expected {
			t.Errorf("Length(%q) = %d, expected %d", tt.input, actual, tt.expected)
		}
	}
}

func TestUpper(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "HELLO WORLD"},
		{"", ""},
		{"こんにちは", "こんにちは"},       // Japanese characters remain unchanged
		{"ıstanbul", "ISTANBUL"}, // Turkish dotless i (ı) to I
		{"éclair", "ÉCLAIR"},     // French e with acute accent
		{"niño", "NIÑO"},         // Spanish n with tilde
		{"опера", "ОПЕРА"},       // Cyrillic script (Russian)
		{"άδικος", "ΆΔΙΚΟΣ"},     // Greek with accent
		{"µ", "Μ"},               // Micro sign (µ) to Greek Capital Letter Mu (Μ)
	}

	for _, tt := range tests {
		actual := Upper(tt.input)
		if actual != tt.expected {
			t.Errorf("Upper(%q) = %q, expected %q", tt.input, actual, tt.expected)
		}
	}
}

func TestLower(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"HELLO WORLD", "hello world"},
		{"", ""},
		{"こんにちは", "こんにちは"},                               // Japanese characters remain unchanged
		{"SS", "ss"},                                     // German sharp s (ß) lowercase is "ss", but uppercase "SS" doesn't convert back to "ß"
		{"ISTANBUL", "istanbul"},                         // Turkish I to lowercase
		{"ÉCLAIR", "éclair"},                             // French E with acute accent
		{"NIÑO", "niño"},                                 // Spanish N with tilde
		{"ОПЕРА", "опера"},                               // Cyrillic script (Russian)
		{"ΆΔΙΚΟΣ", "άδικοσ"},                             // Greek with accent
		{"Μ", "μ"},                                       // Greek Capital Letter Mu (Μ) to micro sign (μ)
		{"ABC123", "abc123"},                             // Mix of letters and numbers
		{"[Special*Characters]", "[special*characters]"}, // Special characters
	}

	for _, tt := range tests {
		actual := Lower(tt.input)
		if actual != tt.expected {
			t.Errorf("Lower(%q) = %q, expected %q", tt.input, actual, tt.expected)
		}
	}
}

func TestTitleize(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"hello world", "Hello World"},
		{"enjoy your coffee", "Enjoy Your Coffee"},
		{"123 abc", "123 Abc"},
		{"special $characters$", "Special Characters"},
		{"use_the_force", "Use The Force"},
		{"camelCase", "Camel Case"},
		{"mixedCASE Words", "Mixed CASE Words"},
		{"non-ascii èéêë ēėę", "Non Ascii Èéêë Ēėę"},
		{"with-dashes-and spaces", "With Dashes And Spaces"},
		{"multiple   spaces", "Multiple Spaces"},
		{"tabs\tand\nnewlines", "Tabs And Newlines"},
		{"*greeting*", "*greeting*"},
		{"hello *marvelous* world!", "Hello *marvelous* World!"},
		{"user/profile", "User Profile"},
		{"user_id", "User ID"},
	}

	for _, tt := range tests {
		actual := Titleize(tt.input)
		if actual != tt.expected {
			t.Errorf("Titleize(%q) = %q, expected %q", tt.input, actual, tt.expected)
		}
	}
}

func TestCapitalize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty string", "", ""},
		{"Single lowercase character", "x", "X"},
		{"Single uppercase character", "X", "X"},
		{"Lowercase", "capitalize", "Capitalize"},
		{"Uppercase all", "CAPITALIZE", "CAPITALIZE"},
		{"Mixed case", "cAPITALIZE", "CAPITALIZE"},
		{"Numbers leading", "123start", "123start"},
		{"Special characters leading", "*special", "*special"},
		{"Whitespace leading", " capitalize", " capitalize"},
		{"Non-English characters", "ñandú", "Ñandú"},
		{"Hyphenated compound", "multi-word-example", "Multi-word-example"},
		{"Snake_case input", "snake_case_input", "Snake_case_input"},
		{"Unicode characters", "привет мир", "Привет мир"},
		{"With apostrophe", "o'clock", "O'clock"},
		{"All lowercase", "alldown", "Alldown"},
		{"Emoji present", "😊emoji", "😊emoji"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Capitalize(tt.input)
			require.Equal(t, tt.expected, actual, fmt.Sprintf("Output for '%s' should match expected value '%s' but got '%s'", tt.input, tt.expected, actual))
		})
	}
}

func TestCamelize(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"john smith", "johnSmith"},
		{"profile_id", "profileID"},
		{"profile-id", "profileID"},
		{"profile id", "profileID"},
		{"Profile_Id", "profileID"},
		{"Happy to meet you!", "happyToMeetYou"},
		{"**welcome**", "welcome"},
		{"I've seen a movie! Have you?", "iveSeenAMovieHaveYou"},
		{"This is *sample* text", "thisIsSampleText"},
		{"alpha_beta", "alphaBeta"},
		{"user/profile", "userProfile"},
		{"example_link", "exampleLink"},
		{"profiles", "profiles"},
		{"Actions", "actions"},
		{"members", "members"},
		{"version 2 update", "version2Update"},
		{"user-profile", "userProfile"},
		{"apples", "apples"},
		{"UserAccounts", "userAccounts"},
		{"résumé operation", "résuméOperation"},
		{"happy 😊 day", "happyDay"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			actual := Camelize(tt.input)
			require.Equal(t, tt.expected, actual, fmt.Sprintf("Output for '%s' should match expected value '%s' but got '%s'", tt.input, tt.expected, actual))
		})
	}
}
func TestDasherize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Handle Empty String", "", ""},
		{"Convert Pathlike String", "admin/AreaID", "admin-area-id"},
		{"Name with Middle Initial", "Grace H. Hopper", "grace-h-hopper"},
		{"String With Invalid Characters", "Text with unusual *&characters*", "text-with-unusual-characters"},
		{"Ends With Special Characters", "Ending special characters!**", "ending-special-characters"},
		{"Starts With Special Characters", "**Starting special characters", "starting-special-characters"},
		{"Multiple Consecutive Spaces", "Multiple    spaces here", "multiple-spaces-here"},
		{"Contains Plus Character", "Phrase with + character", "phrase-with-character"},
		{"Includes Malformed UTF8", "String with bad utf8 \250", "string-with-bad-utf8"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Dasherize(tt.input)
			require.Equal(t, tt.expected, actual, fmt.Sprintf("Output for '%s' should match expected value '%s' but got '%s'", tt.input, tt.expected, actual))
		})
	}
}

func TestSlugify(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hellö Wörld хелло ворлд", "hello-world-khello-vorld"},
		{"影師", "ying-shi"},
		{"This & that", "this-and-that"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := Slugify(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestPluralize(t *testing.T) {
	tests := []struct {
		count    int
		singular string
		plural   string
		expected string
	}{
		{1, "cat", "", "cat"},
		{2, "cat", "", "cats"},
		{2, "CAT", "", "CATS"},
		{1, "mouse", "", "mouse"},
		{2, "mouse", "mice", "mice"},
		{0, "sheep", "", "sheep"},
		{2, "sheep", "", "sheep"},
		{1, "person", "", "person"},
		{2, "person", "", "people"},
		{1, "city", "", "city"},
		{2, "city", "", "cities"},
		{2, "foot", "feet", "feet"},
		{-1, "apple", "", "apples"},
		{2, "$dollar", "", "$dollars"},
		{2, "mother-in-law", "", "mother-in-laws"},
		{1, "%d cat", "%d cats", "1 cat"},
		{2, "%d cat", "%d cats", "2 cats"},
		{1, "%d mouse", "%d mice", "1 mouse"},
		{2, "%d mouse", "%d mice", "2 mice"},
	}

	for _, tt := range tests {
		actual := Pluralize(tt.count, tt.singular, tt.plural)
		if actual != tt.expected {
			t.Errorf("Pluralize(%d, %q, %q) = %q, expected %q", tt.count, tt.singular, tt.plural, actual, tt.expected)
		}
	}
}

func TestOrdinalize(t *testing.T) {
	tests := []struct {
		number   int
		expected string
	}{
		{1, "1st"},
		{2, "2nd"},
		{3, "3rd"},
		{4, "4th"},
		{11, "11th"},
		{12, "12th"},
		{13, "13th"},
		{21, "21st"},
		{22, "22nd"},
		{23, "23rd"},
		{101, "101st"},
		{111, "111th"},
		{112, "112th"},
		{113, "113th"},
		{121, "121st"},
		{-1, "-1th"},
	}

	for _, tt := range tests {
		actual := Ordinalize(tt.number)

		if actual != tt.expected {
			t.Errorf("Ordinalize(%d) = %q, expected %q", tt.number, actual, tt.expected)
		}
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		maxLength  int
		wantOutput string
	}{
		{"Shorter String", "Hello", 10, "Hello"},
		{"Equal Length String", "Hello", 5, "Hello"},
		{"Longer String", "Hello, world!", 5, "Hello..."},
		{"Exact Boundary With Space", "Hello world", 5, "Hello..."},
		{"Multibyte Characters", "Hello, 世界", 8, "Hello, 世..."},
		{"Emoji Characters", "😊😊😊😊", 2, "😊😊..."},
		{"Zero MaxLength", "Hello", 0, ""},
		{"Negative MaxLength", "Hello", -1, ""},
		{"MaxLength One", "Hello", 1, "H..."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := Truncate(tt.input, tt.maxLength); gotOutput != tt.wantOutput {
				t.Errorf("Truncate(%q, %d) = %q, want %q", tt.input, tt.maxLength, gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestTruncateWords(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		maxWords   int
		wantOutput string
	}{
		{"Shorter String", "Hello world", 3, "Hello world"},
		{"Equal Length String", "Hello world", 2, "Hello world"},
		{"Longer String", "Hello beautiful world", 2, "Hello beautiful..."},
		{"Empty String", "", 2, ""},
		{"Only Spaces", "    ", 2, "    "},
		{"Zero MaxWords", "Hello world", 0, ""},
		{"Negative MaxWords", "Hello world", -1, ""},
		{"One Word Input", "Hello", 1, "Hello"},
		{"MaxWords One", "Hello world", 1, "Hello..."},
		{"Punctuation Handling", "Hello, world! How are you?", 3, "Hello, world! How..."},
		{"Long String", "This is a longer string with many words", 5, "This is a longer string..."},
		{"Emoji Characters", "🌟✨🌟 Sparkling stars", 2, "🌟✨🌟 Sparkling..."},
		{"Multibyte Characters", "你好，世界 Hello world", 3, "你好，世界 Hello..."},
		{"Mixed Emoji Words", "😊 World 🌍 is beautiful 🏞️", 3, "😊 World 🌍..."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := TruncateWords(tt.input, tt.maxWords); gotOutput != tt.wantOutput {
				t.Errorf("TruncateWords(%q, %d) = %q, want %q", tt.input, tt.maxWords, gotOutput, tt.wantOutput)
			}
		})
	}
}

// Benchmark tests for string operations

func BenchmarkCamelize(b *testing.B) {
	input := "hello_world_this_is_a_long_test_string"
	for b.Loop() {
		_ = Camelize(input)
	}
}

func BenchmarkTruncateWords(b *testing.B) {
	input := "This is a sample text with many words that we want to truncate at some point to test performance "
	for i := range 10 {
		input += fmt.Sprintf("word%d ", i)
	}
	b.ResetTimer()
	for b.Loop() {
		_ = TruncateWords(input, 50)
	}
}

func BenchmarkTruncate(b *testing.B) {
	input := "This is a very long string that needs to be truncated for testing purposes and performance benchmarking"
	for b.Loop() {
		_ = Truncate(input, 50)
	}
}

func BenchmarkSlugify(b *testing.B) {
	input := "Hello World This Is A Test String"
	for b.Loop() {
		_ = Slugify(input)
	}
}

func BenchmarkTitleize(b *testing.B) {
	input := "hello_world this-is a_test string"
	for b.Loop() {
		_ = Titleize(input)
	}
}

func BenchmarkDasherize(b *testing.B) {
	input := "HelloWorld ThisIsA TestString"
	for b.Loop() {
		_ = Dasherize(input)
	}
}

func BenchmarkPluralize(b *testing.B) {
	for b.Loop() {
		_ = Pluralize(5, "cat", "")
	}
}
