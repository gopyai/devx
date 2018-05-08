package gen

import (
	"fmt"
)

var (
	tables []*Table
)

var (
	dataTypes = map[FieldType]string{
		TEXT: "VARCHAR(500)",
		INT:  "INT",
		DATE: "DATE",
	}
)

func GenerateForMySQL(dbName, dbUser, dbPass string) error {
	conn, err := sqlOpen(dbName, dbUser, dbPass)
	if err != nil {
		return err
	}

	type (
		forKey struct {
			colName  string
			relTable string
		}
	)
	for _, tbl := range tables {
		fmt.Println("Table:", tbl.name)

		// Start create table query
		q := fmt.Sprintf("CREATE TABLE %s(\nid INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,\n", tbl.name)

		var needIndex []string
		var needForeign []forKey

		// Define each column
		for _, field := range tbl.fields {
			fmt.Println("Field:", field.name)

			// Data type
			pre := " "
			var dataType string
			isRelationship := false
			switch t := field.type_; t {
			case relationship:
				pre = "_id "
				dataType = "INT UNSIGNED"
				isRelationship = true
				needForeign = append(needForeign, forKey{field.name, field.relateWith.name})
			default:
				dataType = dataTypes[t]
				if len(dataTypes) == 0 {
					panic(fmt.Errorf("error field type %d", t))
				}
			}
			q += field.name + pre + dataType

			// Not null
			if field.isNotNull {
				q += " NOT NULL"
			}

			// Index type
			if field.index == UNIQUE {
				q += " UNIQUE"
			} else if !isRelationship {
				needIndex = append(needIndex, field.name) // Pending index definition
			}
			q += ",\n"
		}

		// Define unique
		for _, u := range tbl.uniques {
			s := ""
			for i, f := range u {
				if i > 0 {
					s += ","
				}
				s += f.name
				switch f.type_ {
				case relationship:
					s += "_id"
				}
			}
			q += fmt.Sprintf("UNIQUE(%s),\n", s)
		}

		// Define index
		for _, colName := range needIndex {
			q += fmt.Sprintf("INDEX(%s),\n", colName)
		}

		// Define foreign key
		for _, foreign := range needForeign {
			q += fmt.Sprintf("FOREIGN KEY (%s_id) REFERENCES %s(id) ON DELETE CASCADE ON UPDATE CASCADE,\n",
				foreign.colName,
				foreign.relTable)
		}

		// Closing
		q = q[:len(q)-2] + "\n)"

		// Execute sql
		if err := sqlExec(conn, q); err != nil {
			fmt.Println(err)
			// return err
		}
	}

	return nil
}

func GenerateGo() error {
	return fmt.Errorf("TODO") // TODO
}
