package sql2struct

import (
	"fmt"
	"os"
	"text/template"

	"github.com/go-programming-tour-book/tour/internal/word"
)

/**
把得到的数据库信息按照特定的规则转为Go结构体: 模板渲染的方案
 */

/**
const structTpl 为定义的预定义模板

生成的原型如下：
type 大写驼峰的表名称 struct {
	// 注释
	字段名	字段类型
	// 注释
	字段名	字段类型
	...
}

func (model 大写驼峰的表名称) TableName() string {
	return "表名称"
}
 */
const structTpl = `type {{.TableName | ToCamelCase}} struct {
{{range .Columns}}	{{ $length := len .Comment}} {{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
	{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
{{end}}}
func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}`

// 对模板渲染对象进行声明
type StructTemplate struct {
	structTpl string
}
// 存储转换后的Go结构体的所有字段信息
type StructColumn struct {
	Name 	string
	Type 	string
	Tag  	string
	Comment string
}
// 存储最终用于渲染的模板对象信息
type StructTemplateDB struct {
	TableName string
	Columns []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTpl}
}
// 处理模板对象
// 查询 COLUMNS 表所组装得到的 tbColumns 进行分解和转换
func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tag := fmt.Sprintf("`"+"json:"+"\"%s\""+"`", column.ColumnName)
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     tag,
			Comment: column.ColumnComment,
		})
	}
	return tplColumns
}
/**
对模板渲染的自定义函数和模板对象进行处理
原理：
1. 声明一个名为 sql2struct 的新模板对象
2. 定义了自定义函数 ToCamelCase 并与 word.UnderscoreToUpperCamelCase 方法进行绑定
3. 组装符合预定义模板的模板对象，再调用 Execute 方法进行渲染
 */
func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToUpperCamelCase,
	}).Parse(t.structTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}
	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}

	return nil
}