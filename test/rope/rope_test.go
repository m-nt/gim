package rope_test

import (
	"testing"
	"unicode/utf8"

	"github.com/m-nt/gim/src/rope"
)

func Test_create_rope(t *testing.T) {
	rope := rope.Rope{}
	str := "THIS IS A TEST FILE FOR CREATING ROPE"
	rope.From_str(&str)
	if *rope.ToString() != str {
		t.Fail()
	}
}

func Test_rope_length(t *testing.T) {
	rope := rope.Rope{}
	str := "THIS IS A TEST FILE FOR CREATING ROPE"
	rope.From_str(&str)
	if rope.Length() != utf8.RuneCountInString(str) {
		t.Fail()
	}
}
func Test_rope_append(t *testing.T) {
	rope := rope.Rope{}
	str := "THIS IS A TEST FILE FOR CREATING ROPE"
	rope.From_str(&str)
	rope.Append("APPENDED")
	str_len := utf8.RuneCountInString(str)
	if (*rope.ToString())[str_len:] != "APPENDED" {
		t.Fail()
	}
}
func Test_rope_prepend(t *testing.T) {
	rope := rope.Rope{}
	str := "THIS IS A TEST FILE FOR CREATING ROPE"
	rope.From_str(&str)
	pre_str := "PREPENDED"
	rope.Prepend(pre_str)
	if (*rope.ToString())[:utf8.RuneCountInString(pre_str)] != "PREPENDED" {
		t.Fail()
	}
}

func Test_rope_depth(t *testing.T) {
	leaf_len := rope.LEAF_LEN
	rope := rope.Rope{}
	str := "THIS IS A TEST FILE FOR CREATING ROPE"
	rope.From_str(&str)
	str_len := utf8.RuneCountInString(str)
	depth := 1
	for str_len > leaf_len {
		str_len = str_len / 2
		depth++
	}
	if depth != rope.Depth() {
		t.Fail()
	}
}
