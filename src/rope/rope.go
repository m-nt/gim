package rope

import (
	"fmt"
	"unicode/utf8"
)

const LEAF_LEN int = 6

type Rope struct {
	left   *Rope
	right  *Rope
	parent *Rope
	str    string
	count  int
}

func (rope *Rope) From_str(str *string) {
	str_len := utf8.RuneCountInString(*str)
	create_rope(rope, nil, str, 0, str_len-1)
}

func create_rope(rope *Rope, parent *Rope, str *string, L int, R int) {
	rope.parent = parent
	if (R - L) > LEAF_LEN {
		rope.count = (R - L) / 2
		mid := (L + R) / 2
		rope.left = &Rope{}
		rope.right = &Rope{}
		create_rope(rope.left, rope, str, L, mid)
		create_rope(rope.right, rope, str, mid+1, R)
	} else {
		rope.count = R - L
		rope.str = (*str)[L : R+1]
	}
}

func re_balance(rope *Rope) {
	str := rope.ToString()
	str_ln := utf8.RuneCountInString(*str)
	create_rope(rope, nil, str, 0, str_ln-1)
}

func (rope *Rope) Insert(index int, value string) {
	temp_rope := rope // this is the leaf that we merge the value with
	t_index := index
	for temp_rope.left != nil && temp_rope.right != nil {
		if t_index >= temp_rope.count {
			t_index = t_index - temp_rope.count - 1
			temp_rope = temp_rope.right
		} else {
			temp_rope = temp_rope.left
		}
	}
	new_data := temp_rope.str[:t_index] + value + temp_rope.str[t_index:]
	str_ln := utf8.RuneCountInString(new_data)
	temp_rope.str = ""
	create_rope(temp_rope, temp_rope, &new_data, 0, str_ln-1)
	re_balance(rope)
}

func (rope *Rope) Append(value string) {
	str := rope.ToString()
	str_ln := utf8.RuneCountInString(*str)
	rope.Insert(str_ln, value)
}

func (rope *Rope) Prepend(value string) {
	rope.Insert(0, value)
}

func (rope *Rope) Length() int {
	str := rope.ToString()
	return utf8.RuneCountInString(*str)
}

func (rope *Rope) Print() {
	rope.print()
	fmt.Printf("\n")
}

func (rope *Rope) print() {
	if rope == nil {
		return
	} else if rope.right == nil && rope.left == nil {
		fmt.Printf("%s", rope.str)
	}
	rope.left.print()
	rope.right.print()
}

func (rope *Rope) Depth() int {
	if rope == nil {
		return 0
	} else {
		L_depth := rope.left.Depth()
		R_depth := rope.right.Depth()
		if L_depth > R_depth {
			return L_depth + 1
		} else {
			return R_depth + 1
		}
	}
}

func (rope *Rope) ToString() *string {
	if rope == nil {
		return nil
	} else if rope.right == nil && rope.left == nil {
		return &rope.str
	}
	left_str := *rope.left.ToString()
	right_str := *rope.right.ToString()
	res := left_str + right_str
	return &res
}
