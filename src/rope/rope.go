package rope

import (
	"fmt"
	"unicode/utf8"
)

const LEAF_LEN int = 2

type Rope struct {
	left   *Rope
	right  *Rope
	parent *Rope
	str    string
	count  int
}

func (rope *Rope) From_str(str *string) {
	str_len := utf8.RuneCountInString(*str)
	left_index := 0
	rigth_index := str_len - 1
	go create_rope(rope, nil, str, left_index, rigth_index)
}
func (rope *Rope) Print() {
	if rope == nil {
		return
	} else if rope.right == nil && rope.left == nil {
		fmt.Printf("%s", rope.str)
	}
	rope.left.Print()
	rope.right.Print()
}
func create_rope(rope *Rope, parent *Rope, str *string, L int, R int) {
	rope.parent = parent
	if (R - L) > LEAF_LEN {
		rope.count = (R - L) / 2
		mid := (L + R) / 2
		rope.left = &Rope{}
		rope.right = &Rope{}
		go create_rope(rope.left, rope, str, L, mid)
		go create_rope(rope.right, rope, str, mid+1, R)
	} else {
		rope.count = R - L
		rope.str = (*str)[L : R+1]
	}
}
