package tscsv

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

//CsvTable csv表格
type CsvTable struct {
	FileName string
	Rows     []CsvRow
}

//CsvRow csv中的一行
type CsvRow struct {
	Row map[string]string
}

//IntValue 得到int型数据
func (c *CsvRow) IntValue(field string) (int, error) {
	var value int
	var err error
	if value, err = strconv.Atoi(c.Row[field]); err != nil {
		return value, err
	}
	return value, nil
}

//StringValue 得到string型数据
func (c *CsvRow) StringValue(field string) string {
	value, ok := c.Row[field]
	if ok {
		return value
	}
	return ""
}

//LoadCSVFile 加载csv文件
func LoadCSVFile(fileName string, withFieldName bool) (*CsvTable, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if reader == nil {
		return nil, fmt.Errorf("get csv error,file：%s", fileName)
	}
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("read csv error,file：%s", fileName)
	}

	colNum := len(records[0])
	rowNum := len(records)
	var allRecords []CsvRow
	//数据从第几行开始
	start := 0
	if withFieldName {
		start = 1
	}
	//从第二行开始算  默认第一行为标题
	for i := start; i < rowNum; i++ {
		row := make(map[string]string)
		for k := 0; k < colNum; k++ {
			//使用第一行
			if withFieldName {
				row[records[0][k]] = records[i][k]
			} else {
				key := strconv.Itoa(k)
				row[key] = records[i][k]
			}
		}
		record := CsvRow{row}
		allRecords = append(allRecords, record)
	}
	var result = &CsvTable{
		fileName,
		allRecords,
	}
	return result, nil
}
