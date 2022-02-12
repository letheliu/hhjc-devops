package dbFactory

import (
	"fmt"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/dbLink"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"gorm.io/gorm"
	"os"
	"path"
	"reflect"
	"strings"
	"time"
)

func ExportSqlFile(dblinkDto dbLink.DbLinkDto, dbSqlDto dbLink.DbSqlDto) interface{} {
	db, err := initDbLink(dblinkDto)
	if err != nil {
		return result.Error(err.Error())
	}

	workDir := path.Join(config.WorkSpace, dblinkDto.TenantId, dbSqlDto.FileName)
	ch := make(chan string)

	go ExportOne(db, workDir, ch, dblinkDto)

	return result.SuccessData("已提交导出，文件保存在" + workDir + ",请导完后下载")
}

func ExportOne(db *gorm.DB, workDir string, ch chan<- string, dblinkDto dbLink.DbLinkDto) {
	var fileName string

	//if flag.Tables {
	err := exportTables(fileName, db, dblinkDto)
	if err != nil {
		ch <- fmt.Sprintln("Error: ", workDir, "\t export tables throw, \t", err)
		return
	}
	//}

	//if flag.Views {
	//	err := exportViews(fileName, fields)
	//	if err != nil {
	//		ch <- fmt.Sprintln("Error: ", fields.DbName, "\t export views throw, \t", err)
	//		return
	//	}
	//}
	//
	//if flag.Funcs {
	//	err := exportFuncs(fileName, fields)
	//	if err != nil {
	//		ch <- fmt.Sprintln("Error: ", fields.DbName, "\t export funcs throw, \t", err)
	//		return
	//	}
	//}

	ch <- fmt.Sprintln("Export ", workDir, "\t success at \t", time.Now().Format("2006-01-02 15:04:05"))
}

func ExecuteWithDbConn(db *gorm.DB, sqlStr string, values []interface{}) ([]map[string]interface{}, error) {
	var (
		rows []map[string]interface{}
	)
	rs, err := db.Raw(sqlStr, values).Rows()
	if err != nil {
		return nil, err
	}
	cols, _ := rs.Columns()
	for rs.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}
		// Scan the result into the column pointers...
		if err := rs.Scan(columnPointers...); err != nil {
			continue
		}
		row := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			//m[colName] = string((*val).([]byte))
			//fmt.Println(reflect.TypeOf(*val).String())
			if *val != nil && reflect.TypeOf(*val).String() == "[]uint8" {
				row[colName] = string((*val).([]byte))
			} else if *val != nil && reflect.TypeOf(*val).String() == "time.Time" {
				row[colName] = date.GetTimeString((*val).(time.Time))
			} else {
				row[colName] = *val
			}
		}
		rows = append(rows, row)
	}

	return rows, nil
}

func exportTables(fileName string, db *gorm.DB, dblinkDto dbLink.DbLinkDto) error {
	sqlStr := "select CONSTRAINT_NAME,TABLE_NAME,COLUMN_NAME,REFERENCED_TABLE_SCHEMA," +
		"REFERENCED_TABLE_NAME,REFERENCED_COLUMN_NAME from information_schema.`KEY_COLUMN_USAGE` " +
		"where REFERENCED_TABLE_SCHEMA = ? "
	var values []interface{}
	values = append(values, dblinkDto.DbName)
	rows, err := ExecuteWithDbConn(db, sqlStr, values)
	if err != nil {
		return err
	}
	FKEYS := make(map[string]interface{})
	for i, _ := range rows {
		if _, ok := FKEYS[rows[i]["TABLE_NAME"].(string)+"."+rows[i]["CONSTRAINT_NAME"].(string)]; !ok {
			FKEYS[rows[i]["TABLE_NAME"].(string)+"."+rows[i]["CONSTRAINT_NAME"].(string)] = map[string]interface{}{
				"constraintName": rows[i]["CONSTRAINT_NAME"],
				"sourceCols":     make([]string, 0),
				"schema":         rows[i]["REFERENCED_TABLE_SCHEMA"],
				"tableName":      rows[i]["REFERENCED_TABLE_NAME"],
				"targetCols":     make([]string, 0),
			}
		}
		FKEYS[rows[i]["TABLE_NAME"].(string)+"."+rows[i]["CONSTRAINT_NAME"].(string)].(map[string]interface{})["sourceCols"] =
			append(FKEYS[rows[i]["TABLE_NAME"].(string)+"."+rows[i]["CONSTRAINT_NAME"].(string)].(map[string]interface{})["sourceCols"].([]string), rows[i]["COLUMN_NAME"].(string))
		FKEYS[rows[i]["TABLE_NAME"].(string)+"."+rows[i]["CONSTRAINT_NAME"].(string)].(map[string]interface{})["targetCols"] =
			append(FKEYS[rows[i]["TABLE_NAME"].(string)+"."+rows[i]["CONSTRAINT_NAME"].(string)].(map[string]interface{})["targetCols"].([]string), rows[i]["REFERENCED_COLUMN_NAME"].(string))
	}

	sqlStr = "select TABLE_NAME,ENGINE,ROW_FORMAT,AUTO_INCREMENT,TABLE_COLLATION,CREATE_OPTIONS,TABLE_COMMENT" +
		" from information_schema.`TABLES` where TABLE_SCHEMA = ? and TABLE_TYPE = ? order by TABLE_NAME"
	values = make([]interface{}, 0)
	values = append(values, dblinkDto.DbName, "BASE TABLE")
	rows, err = ExecuteWithDbConn(db, sqlStr, values)
	if err != nil {
		return err
	}
	for _, tbAl := range rows {
		sqlStr = "SELECT	`COLUMNS`.COLUMN_NAME,`COLUMNS`.COLUMN_TYPE,`COLUMNS`.IS_NULLABLE," +
			"`COLUMNS`.CHARACTER_SET_NAME,`COLUMNS`.COLUMN_DEFAULT,`COLUMNS`.EXTRA," +
			"`COLUMNS`.COLUMN_KEY,`COLUMNS`.COLUMN_COMMENT,`STATISTICS`.TABLE_NAME," +
			"`STATISTICS`.INDEX_NAME,`STATISTICS`.SEQ_IN_INDEX,`STATISTICS`.NON_UNIQUE," +
			"`COLUMNS`.COLLATION_NAME " +
			"FROM information_schema.`COLUMNS` " +
			"LEFT JOIN information_schema.`STATISTICS` ON " +
			"information_schema.`COLUMNS`.TABLE_NAME = `STATISTICS`.TABLE_NAME " +
			"AND information_schema.`COLUMNS`.COLUMN_NAME = information_schema.`STATISTICS`.COLUMN_NAME " +
			"AND information_schema.`STATISTICS`.table_schema = ? " +
			"where information_schema.`COLUMNS`.TABLE_NAME = ? and `COLUMNS`.table_schema = ?"
		values = make([]interface{}, 0)
		values = append(values, dblinkDto.DbName, tbAl["TABLE_NAME"], dblinkDto.DbName)
		colRs, err := ExecuteWithDbConn(db, sqlStr, values)
		if err != nil {
			return err
		}
		tableName := tbAl["TABLE_NAME"].(string)
		tableEngine := tbAl["ENGINE"].(string)
		//tableRowFormat := tbAl["ROW_FORMAT"]
		var tableAutoIncrement string
		if tbAl["AUTO_INCREMENT"] != nil {
			tableAutoIncrement = tbAl["AUTO_INCREMENT"].(string)
		}
		var tableCollation string
		if tbAl["TABLE_COLLATION"] != nil {
			tableCollation = tbAl["TABLE_COLLATION"].(string)
		}
		tableCharset := strings.Split(tableCollation, "_")[0]
		var tableCreateOptions string
		if tbAl["CREATE_OPTIONS"] != nil {
			tableCreateOptions = tbAl["CREATE_OPTIONS"].(string)
		}
		var tableComment string
		if tbAl["TABLE_COMMENT"] != nil {
			tableComment = tbAl["TABLE_COMMENT"].(string)
		}

		strExport := "DROP TABLE IF EXISTS `" + tbAl["TABLE_NAME"].(string) + "`;\n"
		strExport += "CREATE TABLE `" + tableName + "` (\n"

		priKey := make(map[string]interface{})
		colKey := make(map[string]interface{})
		mulKey := make(map[string]interface{})
		theTableColSet := make(map[string]int)
		var allFields []string
		var defaultValue string
		for _, colAl := range colRs {
			if _, ok := theTableColSet[colAl["COLUMN_NAME"].(string)]; !ok {
				theTableColSet[colAl["COLUMN_NAME"].(string)] = 1
				allFields = append(allFields, "`"+colAl["COLUMN_NAME"].(string)+"`")
				if colAl["COLUMN_DEFAULT"] != nil && len(colAl["COLUMN_DEFAULT"].(string)) > 0 {
					if colAl["COLUMN_DEFAULT"] == "CURRENT_TIMESTAMP" {
						defaultValue = colAl["COLUMN_DEFAULT"].(string)
					} else {
						defaultValue = "'" + colAl["COLUMN_DEFAULT"].(string) + "'"
					}
				}
				var charSet string
				if colAl["CHARACTER_SET_NAME"] != nil && colAl["CHARACTER_SET_NAME"] != tableCharset {
					charSet = " CHARACTER SET " + colAl["CHARACTER_SET_NAME"].(string)
				}
				var collation string
				if colAl["COLLATION_NAME"] != nil && colAl["COLLATION_NAME"] != tableCollation {
					collation = " COLLATE " + colAl["COLLATION_NAME"].(string)
				}
				var nullStr string
				if colAl["IS_NULLABLE"] != nil && colAl["IS_NULLABLE"] == "NO" {
					nullStr = " NOT NULL"
				}
				if colAl["COLUMN_DEFAULT"] != nil && len(colAl["COLUMN_DEFAULT"].(string)) > 0 {
					defaultValue = " DEFAULT " + defaultValue
				} else {
					if colAl["IS_NULLABLE"] != nil && colAl["IS_NULLABLE"] == "NO" {
						defaultValue = ""
					} else {
						defaultValue = " DEFAULT NULL"
					}
				}
				var space string
				if colAl["EXTRA"] != nil && len(colAl["EXTRA"].(string)) > 0 {
					space = " " + colAl["EXTRA"].(string)
				} else {
					space = ""
				}
				var cstr string
				if colAl["COLUMN_COMMENT"] != nil && len(colAl["COLUMN_COMMENT"].(string)) > 0 {
					cstr = " COMMENT '" + escape(colAl["COLUMN_COMMENT"].(string)) + "'"
				}
				strExport += "  `" + colAl["COLUMN_NAME"].(string) + "` " + colAl["COLUMN_TYPE"].(string) + charSet + collation +
					nullStr + defaultValue + space + cstr + ",\n"
			}
			if colAl["INDEX_NAME"] != nil && colAl["INDEX_NAME"].(string) == "PRIMARY" {
				if _, ok := priKey[colAl["INDEX_NAME"].(string)]; !ok {
					priKey[colAl["INDEX_NAME"].(string)] = make([]string, 0)
				}
				priKey[colAl["INDEX_NAME"].(string)] = append(priKey[colAl["INDEX_NAME"].(string)].([]string), colAl["COLUMN_NAME"].(string))
			} else if colAl["INDEX_NAME"] != nil && colAl["NON_UNIQUE"] == "0" {
				if _, ok := colKey[colAl["INDEX_NAME"].(string)]; !ok {
					colKey[colAl["INDEX_NAME"].(string)] = make([]string, 0)
				}
				colKey[colAl["INDEX_NAME"].(string)] = append(colKey[colAl["INDEX_NAME"].(string)].([]string), colAl["COLUMN_NAME"].(string))
			} else if colAl["INDEX_NAME"] != nil && colAl["NON_UNIQUE"] == "1" {
				if _, ok := mulKey[colAl["INDEX_NAME"].(string)]; !ok {
					mulKey[colAl["INDEX_NAME"].(string)] = make([]string, 0)
				}
				mulKey[colAl["INDEX_NAME"].(string)] = append(mulKey[colAl["INDEX_NAME"].(string)].([]string), colAl["COLUMN_NAME"].(string))
			}
		}
		for _, v := range priKey {
			strExport += "  PRIMARY KEY (`" + strings.Join(v.([]string), "`,`") + "`),\n"
		}
		for k, v := range colKey {
			strExport += "  UNIQUE KEY `" + k + "` (`" + strings.Join(v.([]string), "`,`") + "`),\n"
		}
		for k, v := range mulKey {
			strExport += "  KEY `" + k + "` (`" + strings.Join(v.([]string), "`,`") + "`),\n"
		}

		for k, v := range FKEYS {
			if strings.HasPrefix(k, tableName+".") {
				strExport += "  CONSTRAINT `" + v.(map[string]interface{})["constraintName"].(string) + "` FOREIGN KEY (`" +
					strings.Join(v.(map[string]interface{})["sourceCols"].([]string), "`,`") + "`) REFERENCES `" +
					v.(map[string]interface{})["tableName"].(string) + "` (`" +
					strings.Join(v.(map[string]interface{})["targetCols"].([]string), "`,`") + "`),\n"
			}
		}
		if strings.HasSuffix(strExport, ",\n") {
			strExport = strExport[:len(strExport)-2]
		}

		var incr string
		if len(tableAutoIncrement) > 0 {
			incr = " AUTO_INCREMENT=" + tableAutoIncrement
		}
		var colla string
		if len(tableCollation) > 0 {
			colla = " COLLATE=" + tableCollation
		}
		strExport += "\n) ENGINE=" + tableEngine + incr + " DEFAULT CHARSET=" +
			tableCharset + colla + " " + tableCreateOptions + " COMMENT='" + tableComment + "';\n\n"

		writeToFile(fileName, strExport, true) //表结构导出

		err = exportTableData(fileName, db, dblinkDto, tableName, allFields)
		if err != nil {
			return err
		}

	}
	return nil
}

func exportTableData(fileName string, db *gorm.DB, dblinkDto dbLink.DbLinkDto, tableName string, allFields []string) error {
	sqlStr := "select " + strings.Join(allFields, ",") + " from " + tableName
	recordsRs, err := ExecuteWithDbConn(db, sqlStr, make([]interface{}, 0))
	if err != nil {
		return err
	}
	for _, ele := range recordsRs {
		strExport := "INSERT INTO `" + tableName + "` (" //+strings.Join(allFields, ",")+") VALUES ("
		var ks []string
		var vs []string
		for k, v := range ele {
			ks = append(ks, "`"+k+"`")
			elStr := "''"
			if v == nil {
				elStr = "null"
			} else if len(v.(string)) > 0 {
				elStr = "'" + escape(v.(string)) + "'"
			}
			vs = append(vs, elStr)
		}
		strExport += strings.Join(ks, ",") + ") VALUES (" + strings.Join(vs, ",") + ");\n"
		writeToFile(fileName, strExport, true)
	}
	writeToFile(fileName, "\n", true)
	return nil
}

//func exportFuncs(fileName string, fields common.DbConnFields) error {
//	sqlStr := "select name,type,param_list,returns,body from mysql.proc where db = ? "
//	values := make([]interface{}, 0)
//	values = append(values, fields.DbName)
//
//	rs, err := db.ExecuteWithDbConn(sqlStr, values, fields)
//	if err != nil {
//		return err
//	}
//	fRs := rs["rows"].([]map[string]interface{})
//
//	for _, cstAl := range fRs {
//		var rets string
//		if cstAl["returns"] != nil && len(cstAl["returns"].(string)) > 0 {
//			rets = " RETURNS " + cstAl["returns"].(string)
//		}
//		sqlStr = "DROP PROCEDURE IF EXISTS `" + cstAl["name"].(string) + "`;\nDELIMITER ;;\n" +
//			"CREATE DEFINER=`root`@`%` " + cstAl["type"].(string) + " `" + cstAl["name"].(string) +
//			"`(" + cstAl["param_list"].(string) + ")" + rets + "\n" + cstAl["body"].(string) + "\n" +
//			";;\nDELIMITER ;\n\n"
//		writeToFile(fileName, sqlStr, true)
//	}
//	return nil
//}
//
//func exportViews(fileName string, fields common.DbConnFields) error {
//	sqlStr := "select TABLE_NAME, VIEW_DEFINITION from information_schema.VIEWS where TABLE_SCHEMA = ? "
//	values := make([]interface{}, 0)
//	values = append(values, fields.DbName)
//	rs, err := db.ExecuteWithDbConn(sqlStr, values, fields)
//	if err != nil {
//		return err
//	}
//	vRs := rs["rows"].([]map[string]interface{})
//	ps := make(map[string]string)
//	vName := make([]string, 0)
//	for _, v := range vRs {
//		ps["`"+v["TABLE_NAME"].(string)+"`"] = v["VIEW_DEFINITION"].(string)
//		vName = append(vName, "`"+v["TABLE_NAME"].(string)+"`")
//	}
//	rely1 := processRely(ps, &vName)
//	rely := processRely(ps, &rely1)
//	for _, al := range rely {
//		viewStr := strings.Replace(ps[al], "`"+fields.DbName+"`.", "", -1)
//		sqlStr = "DROP VIEW IF EXISTS " + al + ";\n" + "CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`%` " +
//			" SQL SECURITY DEFINER VIEW " + al + " AS " + viewStr + ";\n\n"
//		writeToFile(fileName, sqlStr, true)
//	}
//	return nil
//}

func processRely(params map[string]string, relyOld *[]string) []string {
	rely := make([]string, 0)
	for _, k := range *relyOld {
		for bl := range params {
			if strings.Index(params[k], bl) > -1 {
				if findInArray(&rely, bl) < 0 {
					if findInArray(&rely, k) < 0 {
						rely = append(rely, bl)
					} else {
						i := findInArray(&rely, k)
						lastStr := make([]string, len(rely)-i)
						copy(lastStr, rely[i:])
						rely = append(rely[:i], bl)
						rely = append(rely, lastStr...)
					}
				} else {
					if findInArray(&rely, k) > -1 {
						i := findInArray(&rely, k)
						j := findInArray(&rely, bl)
						if i < j {
							rely = append(rely[:j], rely[j+1:]...)
							lastStr := make([]string, len(rely)-i)
							copy(lastStr, rely[i:])
							rely = append(rely[:i], bl)
							rely = append(rely, lastStr...)
						}
					}
				}
			}
		}
		if findInArray(&rely, k) < 0 {
			rely = append(rely, k)
		}
	}
	return rely
}

func findInArray(arry *[]string, value string) int {
	if arry == nil {
		return -1
	} else {
		for index, v := range *arry {
			if v == value {
				return index
			}
		}
		return -1
	}
}

func writeToFile(name string, content string, append bool) {
	var fileObj *os.File
	var err error
	if append {
		fileObj, err = os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	} else {
		fileObj, err = os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	}
	if err != nil {
		fmt.Println("Failed to open the file", err.Error())
		os.Exit(2)
	}
	defer fileObj.Close()
	if _, err := fileObj.WriteString(content); err != nil {
		fmt.Println(err)
	}
}

func escape(source string) string {
	var j int
	if len(source) == 0 {
		return ""
	}
	tempStr := source[:]
	desc := make([]byte, len(tempStr)*2)
	for i := 0; i < len(tempStr); i++ {
		flag := false
		var escape byte
		switch tempStr[i] {
		case '\r':
			flag = true
			escape = '\r'
		case '\n':
			flag = true
			escape = '\n'
		case '\\':
			flag = true
			escape = '\\'
		case '\'':
			flag = true
			escape = '\''
		case '"':
			flag = true
			escape = '"'
		case '\032':
			flag = true
			escape = 'Z'
		default:
		}
		if flag {
			desc[j] = '\\'
			desc[j+1] = escape
			j = j + 2
		} else {
			desc[j] = tempStr[i]
			j = j + 1
		}
	}
	return string(desc[0:j])
}
