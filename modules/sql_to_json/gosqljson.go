// [Ref]
// https://github.com/elgs/gosqljson
// https://stackoverflow.com/questions/19991541/dumping-mysql-tables-to-json-with-golang
// https://blog.csdn.net/newjueqi/article/details/52370786

package gosqljson

import (
	"database/sql"
	"encoding/json"
	//"fmt"
)

func GetJSON(db *sql.DB, sqlString string, sqlParams ...interface{}) (string, error) {
	rows, err := db.Query(sqlString, sqlParams...)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	// if rows.Next() {

	// 	for i := 0; i < count; i++ {
	// 		valuePtrs[i] = &values[i]
	// 	}
	// 	rows.Scan(valuePtrs...)
	// 	entry := make(map[string]interface{})
	// 	for i, col := range columns {
	// 		var v interface{}
	// 		val := values[i]
	// 		b, ok := val.([]byte)
	// 		if ok {
	// 			v = string(b)
	// 		} else {
	// 			v = val
	// 		}
	// 		entry[col] = v
	// 	}
	// 	tableData = append(tableData, entry)

	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}

	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
		// jsonData = "error"
	}
	// } else {
	// jsonData := "no_data"
	// }

	//fmt.Println(string(jsonData))
	return string(jsonData), nil
}
