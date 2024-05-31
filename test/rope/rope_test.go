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
func Test_rope_insert(t *testing.T) {
	rope := rope.Rope{}
	str := "THIS IS A TEST FILE FOR CREATING ROPE"
	rope.From_str(&str)
	rope.Insert(3, "INSERTED")
	rope.Print()
	if (*rope.ToString())[3:11] != "INSERTED" {
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

func Test_rope_concat(t *testing.T) {
	str1 := "FIRST PART"
	str2 := " AND SECOND PART"

	rope1 := rope.Rope{}
	rope1.From_str(&str1)
	rope2 := rope.Rope{}
	rope2.From_str(&str2)

	new_rope := rope1.Concat(&rope2)

	if new_rope.Length() != rope1.Length()+rope2.Length() {
		t.Fail()
	}
}
func Test_rope_split(t *testing.T) {
	str := "SPLIT MESPLIT ME"
	rope := &rope.Rope{}
	rope.From_str(&str)
	right := rope.Split(8)
	if *rope.ToString() != *right.ToString() {
		t.Fail()
	}
}

func Test_rope_delete(t *testing.T) {
	str := "THIS IS A TEST <TO_BE_DELETED> FILE FOR CREATING ROPE"
	rope := &rope.Rope{}
	rope.From_str(&str)
	deleted := rope.Delete(15, 30)
	if *deleted.ToString() != "<TO_BE_DELETED>" {
		t.Fail()
	}
}
func Test_rope_path(t *testing.T) {
	str := "THIS IS A TEST FILE FOR CREATING ROPE"
	rope := &rope.Rope{}
	rope.From_str(&str)
	rope.Path()
}
