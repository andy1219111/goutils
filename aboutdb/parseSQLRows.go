package aboutdb

import (
	"database/sql"
)

//ParseSQLRows 将sql.Rows中的数据解析成map数组
func ParseSQLRows(rows *sql.Rows) ([]map[string]sql.RawBytes, error) {
	var result []map[string]sql.RawBytes
	//查询的列名
	fields, _ := rows.Columns()
	for rows.Next() {
		values := make([]sql.RawBytes, len(fields))
		scanArgs := make([]interface{}, len(fields))

		for i := range scanArgs {
			scanArgs[i] = &values[i]
		}
		err := rows.Scan(scanArgs...)
		if err != nil {
			return result, err
		}

		rowData := make(map[string]sql.RawBytes)
		for i := 0; i < len(fields); i++ {
			rowData[fields[i]] = values[i]
		}
		result = append(result, rowData)
	}
	return result, nil
}
