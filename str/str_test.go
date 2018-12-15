package str

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestStringCount(t *testing.T)  {
	assert.Equal(t, 7, utf8.RuneCountInString("【12345你"))
	assert.Equal(t,0, strings.Count("23", "1"))
	assert.Equal(t, 3, strings.Count("23", ""))
	assert.Equal(t,1, strings.Count("123", "1"))
	assert.Equal(t, 2, strings.Count("1123", "1"))
	assert.Equal(t, 2, strings.Count("112312", "12"))
}

func TestStringContains(t *testing.T)  {
	assert.Equal(t, true, strings.Contains("112134", "11"))
	assert.Equal(t, false, strings.Contains("12134", "11"))
}

func TestStringContainsAny(t *testing.T)  {
	assert.Equal(t, true, strings.ContainsAny("112134", "15"))
	assert.Equal(t, true, strings.ContainsAny("112134", "14"))
	assert.Equal(t, false, strings.ContainsAny("112134", "【"))
	assert.Equal(t, true, strings.ContainsAny("1121【34", "【"))
}

func TestStringContainsRune(t *testing.T) {
	assert.Equal(t, true, strings.ContainsRune("123", '1'))
	assert.Equal(t, false, strings.ContainsRune("123", '5'))
	assert.Equal(t, false, strings.ContainsRune("123", '【'))
}

func TestStringLastIndex(t *testing.T)  {
	assert.Equal(t, 1, strings.LastIndex("0123", "12"))
	assert.Equal(t, -1, strings.LastIndex("0123", "234"))
	assert.Equal(t, 3, strings.LastIndex("012123", "12"))
	assert.Equal(t, 3, strings.LastIndex("012【123", "【"))
}

func TestStringIndexRune(t *testing.T)  {
	assert.Equal(t, 3, strings.IndexRune("01234", '3'))
	assert.Equal(t, -1, strings.IndexRune("01234", '【'))
}

func TestStringSplitN(t *testing.T) {
	assert.Equal(t, []string {"1", "2", "  3"}, strings.SplitN("1 2   3", " ", 3))
	assert.Equal(t, []string {"1", "2", "3"}, strings.SplitN("1 2 3", " ", 3))
	assert.Equal(t, []string {"1", "2 3"}, strings.SplitN("1 2 3", " ", 2))
	assert.Equal(t, []string {"1", "2", "3"}, strings.SplitN("1 2 3", " ", -1))
	assert.Equal(t, []string([]string(nil)), strings.SplitN("1 2 3", " ", 0))
	assert.Equal(t, []string {"123"}, strings.SplitN("123", " ", 2))
}

func TestStringSplitAfterN(t *testing.T) {
	// assert.Equal(t, []string {"1", "2", "3"}, strings.SplitAfterN("1 2 3", " ", 1))
	// TODO: understand this function
}

func TestStringSplit(t *testing.T) {
	assert.Equal(t, []string {"1", "2", "", "", "3"}, strings.Split("1 2   3", " "))
}

func TestStringFields(t *testing.T)  {
	assert.Equal(t, []string{"1", "2", "3", "999"}, strings.Fields("   1 2   3    999 "))
	assert.Equal(t, []string{}, strings.Fields("    "))
}

func TestStringFieldsFunc(t *testing.T) {
	assert.Equal(t, []string{"1", "2", "3"}, strings.FieldsFunc("1,2,3", func(r rune) bool {
		if r == ',' {
			return true
		}
		return false
	}))
}

func TestStringJoin(t *testing.T) {
	assert.Equal(t, "1,22,3", strings.Join([]string{"1", "22", "3"}, ","))
}

func TestStringHasPrefix(t *testing.T) {
	assert.Equal(t, true, strings.HasPrefix("1prefix1", "1pr"))
	assert.Equal(t, false, strings.HasPrefix("1p1refix1", "1pr"))
}

func TestStringHasSufix(t *testing.T) {
	assert.Equal(t, true, strings.HasSuffix("1prefix1", "ix1"))
	assert.Equal(t, false, strings.HasSuffix("1p1refix1", "1pr"))
}

func TestStringMap(t *testing.T) {
	//assert.Equal(t, "112233", strings.Map(func(r rune) rune {
	//	return r + ('a' - 'A')
	//}))
	// TODO
}

func TestStringRepeat(t *testing.T) {
	assert.Equal(t, "123123123", strings.Repeat("123", 3))
}

func TestStringToUpper(t *testing.T) {
	assert.Equal(t, "ABCD123ABC", strings.ToUpper("aBcd123aBc"))
}

func TestStringToLower(t *testing.T) {
	assert.Equal(t, "abcd123abc", strings.ToLower("AbCD123AbC"))
}

func TestStringToTitle(t *testing.T) {
	assert.Equal(t, "APPLE COMP", strings.ToTitle("apple comp"))
}

func TestStringTitle(t *testing.T) {
	assert.Equal(t, "Apple Comp", strings.Title("apple comp"))
}

//func Test()  {
//
//}
