package main

import "fmt"

/*给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/
func vaildParentheses(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	tempMap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []byte{}
	for i := 0; i < len(s); i++ {
		// 左括号存在
		if tempMap[s[i]] > 0 {
			// 栈中为空，或者栈的最后一个左括号没匹配上对应的右括号
			if len(stack) == 0 || tempMap[s[i]] != stack[len(stack)-1] {
				return false
			} else {
				// 出栈
				stack = stack[:len(stack)-1]
			}
		} else {
			//入栈
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func main() {
	s := "{[({{}}]}"
	ok := vaildParentheses(s)
	if ok {
		fmt.Println("vaild")
	} else {
		fmt.Println("unvaild")
	}
}
