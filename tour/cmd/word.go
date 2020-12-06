package cmd

import (
	"github.com/go-programming-tour-book/tour/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

/**
 * 用户放置单词格式转换的子命令word

运行效果如下：
go run main.go help word
该命令支持各种单词格式转换，模式如下：
1: 全部单词转为大写
2: 全部单词转为小写
3：下划线单词转为大写驼峰单词
4: 下划线单词转为小写驼峰单词
5: 驼峰单词转为下划线单词

Usage:
   word [flags]

Flags:
  -h, --help         help for word
  -m, --mode int8    请输入单词转换的模式
  -s, --str string   请输入单词内容
 */

var str string
var mode int8

const (
	ModeUpper                      = iota + 1 // 全部单词转为大写
	ModeLower                                 // 全部单词转为小写
	ModeUnderscoreToUpperCamelcase            // 下划线单词转为大写驼峰单词
	ModeUnderscoreToLowerCamelcase            // 下划线单词转为小写驼峰单词
	ModeCamelcaseToUnderscore                 // 驼峰单词转为下划线单词
)

var desc = strings.Join([]string{
	"该命令支持各种单词格式转换，模式如下：",
	"1: 全部单词转为大写",
	"2: 全部单词转为小写",
	"3：下划线单词转为大写驼峰单词",
	"4: 下划线单词转为小写驼峰单词",
	"5: 驼峰单词转为下划线单词",
}, "\n")

var wordCmd = &cobra.Command{
	Use: "word",			// go run main.go help word (此处的word为cli中的 help word)
	Short: "单词格式转换",
	Long: desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		// go run main.go word -s=hucy -m=1
		//2020/12/06 12:29:24 输出结果：HUCY
		case ModeUpper:
			content = word.ToUpper(str)
		// go run main.go word -s=HUCY -m=2
		//2020/12/06 12:29:38 输出结果：hucy
		case ModeLower:
			content = word.ToLower(str)
		// go run main.go word -s=hucy -m=3
		//2020/12/06 12:30:09 输出结果：Hucy
		case ModeUnderscoreToUpperCamelcase:
			content = word.UnderscoreToUpperCamelCase(str)
		// go run main.go word -s=HUCY -m=4
		//2020/12/06 12:30:22 输出结果：hUCY
		case ModeUnderscoreToLowerCamelcase:
			content = word.UnderscoreToLowerCamelCase(str)
		// go run main.go word -s=HuCY -m=5
		//2020/12/06 12:30:43 输出结果：hu_c_y
		case ModeCamelcaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		// go run main.go word -s=HuCY -m=6
		//2020/12/06 12:34:18 暂不支持该转换模式，请执行 help word查看帮助文档
		//exit status 1
		default:
			log.Fatalf("暂不支持该转换模式，请执行 help word查看帮助文档")
		}
		log.Printf("输出结果：%s", content)
	},
}
/**
在 VarP 系列的方法中，第一个参数为需要绑定的变量，第二个参数为接收该参数的完整的命令标志，
第三个参数为对应的短标识，第四个参数为默认值，第五个参数为使用说明。
 */
func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}