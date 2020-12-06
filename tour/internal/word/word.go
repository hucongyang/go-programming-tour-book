package word

import (
	"strings"
	"unicode"
)
/**
单词全部转为大写
 */
func ToUpper(s string) string {
	return strings.ToUpper(s)
}
/**
单词全部转为小写
 */
func ToLower(s string) string {
	return strings.ToLower(s)
}
/**
下划线单词转大写驼峰单词
思路：1. 下划线替换为空格字符
	 2. 所有字符修改为其对应的首字母大写形式
	 3. 把空格字符串替换为空
 */
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}
/**
下划线单词转小写驼峰单词
 */
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}
/**
驼峰单词转下划线单词
思路：1. 直接使用 govalidator 库提供的转换方法把大写或小写驼峰单词转为下划线单词
	 2. 将字符转为小写的同时添加下划线，如果当前字符不是小写字母、下划线或数字，那么在处理时将对 segment 置空
 */
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}