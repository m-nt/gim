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
	go create_rope(rope, nil, str, 0, str_len-1)
}
func (rope *Rope) Print() {
	// fmt.Printf("%+v\n", rope)
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

func (rope *Rope) Insert(index int, value *string) {
	temp_rope := rope // this is the leaf that we merge the value with
	t_index := index
	for temp_rope.left != nil && rope.right != nil {
		if t_index >= temp_rope.count {
			temp_rope = temp_rope.right
			t_index = t_index - temp_rope.count
		} else {
			temp_rope = temp_rope.left
		}
	}
	new_data := temp_rope.str[:t_index] + *value + temp_rope.str[t_index:]
	str_ln := utf8.RuneCountInString(new_data)
	temp_rope.str = ""
	fmt.Printf("new str: %s\n", new_data)
	go create_rope(temp_rope, temp_rope, &new_data, 0, str_ln-1)
}
