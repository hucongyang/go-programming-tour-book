package sql2struct

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/**
要获取表中列的信息，即需要访问 information_schema 数据库中的 COLUMNS 表，
在程序中进行连接、查询、数据组装等处理
 */

// 数据库连接的核心对象
type DBModel struct {
	DBEngine *sql.DB
	DBInfo *DBInfo
}
// 存储连接mysql基本信息
type DBInfo struct {
	DBType	string
	Host 	string
	Username 	string
	Password 	string
	Charset 	string
}
// 存储columns表中需要的一些字段
type TableColumn struct {
	ColumnName		string
	DataType 		string
	IsNullable		string
	ColumnKey		string
	ColumnType 		string
	ColumnComment	string
}
// 数据库DataType字段的类型与 Go结构体中的类型不是完全一致，
// 用以下做类型转换
var DBTypeToStructType = map[string]string{
	"int":        "int32",
	"tinyint":    "int8",
	"smallint":   "int",
	"mediumint":  "int64",
	"bigint":     "int64",
	"bit":        "int",
	"bool":       "bool",
	"enum":       "string",
	"set":        "string",
	"varchar":    "string",
	"char":       "string",
	"tinytext":   "string",
	"mediumtext": "string",
	"text":       "string",
	"longtext":   "string",
	"blob":       "string",
	"tinyblob":   "string",
	"mediumblob": "string",
	"longblob":   "string",
	"date":       "time.Time",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"time":       "time.Time",
	"float":      "float64",
	"double":     "float64",
}

// 初始化方法
func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}
// 连接数据库具体方法
func (m *DBModel) Connect() error {
	var err error
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/information_schema?charset=%s&parseTime=True&loc=Local",
		m.DBInfo.Username,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset,
	)
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}
	return nil
}
// 完成数据库的连接方法后，针对 COLUMNS 表进行查询和数据组装
func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	query := "SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, " +
		"IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT " +
		"FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? "
	rows, err := m.DBEngine.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("没有数据")
	}
	defer rows.Close()
	
	var columns []*TableColumn
	for rows.Next() {
		var column TableColumn
		err := rows.Scan(&column.ColumnName, &column.DataType,
			&column.ColumnKey, &column.IsNullable,
			&column.ColumnType, &column.ColumnComment)
		if err != nil {
			return nil, err
		}
		columns = append(columns, &column)
	}
	return columns, nil
}