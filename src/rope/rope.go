package rope

import (
	"fmt"
	"strconv"
	"strings"
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
	create_rope(rope, nil, str, 0, str_len)
}

func re_balance(rope *Rope) {
	str := rope.ToString()
	str_ln := utf8.RuneCountInString(*str)
	create_rope(rope, nil, str, 0, str_ln)
}

func (rope *Rope) Insert(index int, value string) {
	temp_rope := rope // this is the leaf that we merge the value with
	t_index := index
	for temp_rope.left != nil && temp_rope.right != nil {
		if t_index >= temp_rope.count {
			t_index = t_index - temp_rope.count
			temp_rope = temp_rope.right
		} else {
			temp_rope = temp_rope.left
		}
	}
	new_data := temp_rope.str[:t_index] + value + temp_rope.str[t_index:]
	str_ln := utf8.RuneCountInString(new_data)
	temp_rope.str = ""
	create_rope(temp_rope, temp_rope, &new_data, 0, str_ln)
	re_balance(rope)
}

func (rope *Rope) Concat(other *Rope) *Rope {
	new_root := &Rope{}
	new_root.count = rope.Length()
	new_root.left = rope
	new_root.right = other
	rope.parent = new_root
	other.parent = new_root
	re_balance(new_root)
	return new_root
}

func (rope *Rope) Append(value string) {
	str := rope.ToString()
	str_ln := utf8.RuneCountInString(*str)
	rope.Insert(str_ln, value)
}

func (rope *Rope) Prepend(value string) {
	rope.Insert(0, value)
}

func (rope *Rope) Delete(start int, end int) *Rope {
	start_right := rope.Split(start)
	end_rope := start_right.Split(end - start)
	*rope = *rope.Concat(end_rope)
	return start_right
}

func (rope *Rope) Length() int {
	str := rope.ToString()
	return utf8.RuneCountInString(*str)
}

func (rope *Rope) Print() {
	rope.print()
	fmt.Printf("\n")
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

func (rope *Rope) Path() {
	lines, _, _, _ := rope.path()
	for index, line := range lines {
		_ = index
		fmt.Printf("%s\n", line)
	}
}

// Credit : https://stackoverflow.com/questions/34012886/print-binary-tree-level-by-level-in-python
func (rope *Rope) path() ([]string, int, int, int) {
	if rope.right == nil && rope.left == nil {
		line := strconv.Itoa(rope.count)
		width := utf8.RuneCountInString(line)
		height := 1
		middle := width / 2
		return []string{line}, width, height, middle
	}
	if rope.right == nil {
		lines, n, p, x := rope.left.path()
		s := strconv.Itoa(rope.count)
		u := utf8.RuneCountInString(s)
		first_line := strings.Repeat(" ", x+1) + strings.Repeat("_", n-x-1) + s
		second_line := strings.Repeat(" ", x) + "/" + strings.Repeat(" ", n-x-1+u)
		shifted_lines := []string{}
		for _, line := range lines {
			shifted_lines = append(shifted_lines, line+strings.Repeat(" ", u))
		}
		return append([]string{first_line, second_line}, shifted_lines...), n + u, p + 2, n + u/2
	}
	if rope.left == nil {
		lines, n, p, x := rope.right.path()
		s := strconv.Itoa(rope.count)
		u := utf8.RuneCountInString(s)
		first_line := s + strings.Repeat("_", x) + strings.Repeat(" ", n-x)
		second_line := strings.Repeat(" ", u+x) + "\\" + strings.Repeat(" ", n-x-1)
		shifted_lines := []string{}
		for _, line := range lines {
			shifted_lines = append(shifted_lines, strings.Repeat(" ", u)+line)
		}
		return append([]string{first_line, second_line}, shifted_lines...), n + u, p + 2, u / 2
	}
	left, n, p, x := rope.left.path()
	right, m, q, y := rope.right.path()
	s := strconv.Itoa(rope.count)
	u := utf8.RuneCountInString(s)
	first_line := strings.Repeat(" ", x+1) + strings.Repeat("_", n-x-1) + s + strings.Repeat("_", y) + strings.Repeat(" ", m-y)
	second_line := strings.Repeat(" ", x) + "/" + strings.Repeat(" ", n-x-1+u+y) + "\\" + strings.Repeat(" ", m-y-1)
	if p < q {
		for i := range q - p {
			_ = i
			left = append(left, strings.Repeat(" ", n))
		}
	} else if q < p {
		for i := range p - q {
			_ = i
			right = append(right, strings.Repeat(" ", m))
		}
	}
	zipped_lines := []string{}
	for i := range len(left) {
		zipped_lines = append(zipped_lines, left[i]+strings.Repeat(" ", u)+right[i])
	}
	lines := append([]string{first_line, second_line}, zipped_lines...)
	return lines, n + m + u, max(p, q) + 2, n + u/2
}

func (rope *Rope) ToString() *string {
	if rope == nil {
		return nil
	} else if rope.right == nil && rope.left == nil {
		return &rope.str
	}
	if rope.left == nil {
		return rope.right.ToString()
	}
	if rope.right == nil {
		return rope.left.ToString()
	}
	left_str := *rope.left.ToString()
	right_str := *rope.right.ToString()
	res := left_str + right_str
	return &res
}

func (rope *Rope) print() {
	if rope == nil {
		return
	} else if rope.right == nil && rope.left == nil {
		fmt.Printf("%s", rope.str)
	}
	if rope.left != nil {
		rope.left.print()
	}
	if rope.right != nil {
		rope.right.print()
	}
}

func create_rope(rope *Rope, parent *Rope, str *string, L int, R int) {
	rope.parent = parent
	if (R - L) > LEAF_LEN {
		rope.count = (R - L) / 2
		mid := (L + R) / 2
		rope.left = &Rope{}
		rope.right = &Rope{}
		create_rope(rope.left, rope, str, L, mid)
		create_rope(rope.right, rope, str, mid, R)
	} else {
		rope.count = R - L
		rope.str = (*str)[L:R]
	}
}

func (rope *Rope) Split(index int) *Rope {
	temp_rope := rope // this is the leaf that we split from
	t_index := index
	direction := []int{} // 0 means left 1 means right
	for temp_rope.left != nil && temp_rope.right != nil {
		if t_index >= temp_rope.count {
			t_index = t_index - temp_rope.count
			temp_rope = temp_rope.right
			direction = append(direction, 0)
		} else {
			temp_rope = temp_rope.left
			direction = append(direction, 1)
		}
	}
	new_left := &Rope{}
	new_right := &Rope{}
	new_left.str = temp_rope.str[:t_index]
	new_left.count = t_index
	new_left.parent = temp_rope
	new_right.str = temp_rope.str[t_index:]
	new_right.count = utf8.RuneCountInString(new_right.str)
	new_right.parent = temp_rope
	temp_rope.str = ""
	temp_rope.count = new_left.count
	temp_rope.left = new_left
	temp_rope.right = new_right
	t_index = 0
	// new we can operate
	broken_ropes := []*Rope{}
	broken_ropes = append(broken_ropes, temp_rope.right)
	right_temp := temp_rope.right
	temp_rope.right = nil
	ind := len(direction)
	for temp_rope.parent != nil { // walk back up the nodes and remove right links
		temp_rope = temp_rope.parent
		if direction[ind-1] == 1 {
			temp_rope.count = temp_rope.count - right_temp.count
			broken_ropes = append(broken_ropes, temp_rope.right)
			right_temp = temp_rope.right
			temp_rope.right = nil
		}
		ind--
	}
	result_right_rope := broken_ropes[0]
	for i := 1; i < len(broken_ropes); i++ {
		result_right_rope = result_right_rope.Concat(broken_ropes[i])
	}
	re_balance(rope)
	return result_right_rope
}
