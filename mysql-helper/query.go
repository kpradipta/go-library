package mysqlhelper

import (
	"bytes"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func ApplySqlQuery(db *gorm.DB, filters map[string]interface{}) {
	for key, value := range filters {
		switch {
		case strings.Index(value.(string), "range(") == 0:
			query := fmt.Sprint(key, " BETWEEN ? and ?")
			parseStr := value.(string)[6:]
			splitStr := strings.Split(parseStr[:len(parseStr)-1], ",")
			db.Where(query, splitStr[0], splitStr[1])
		case strings.Index(value.(string), "in(") == 0:
			query := fmt.Sprint(key, " In (?)")
			parseStr := value.(string)[3:]
			splitStr := strings.Split(parseStr[:len(parseStr)-1], ",")
			db.Where(query, splitStr)
		default:
			db.Where(key, value)
		}
	}
}

func ApplySorting(db *gorm.DB, sorts map[string]interface{}) {
	b := new(bytes.Buffer)
	for key, value := range sorts {
		fmt.Fprintf(b, "%s %s ", key, value)
	}
	db.Order(b.String())
}
