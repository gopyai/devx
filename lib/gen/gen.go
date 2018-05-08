package gen

import (
	"fmt"
	"log"
)

type (
	joinQuery struct {
		q string
	}
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
		log.Println("Table:", tbl.name)

		// Start create table query
		q := fmt.Sprintf("CREATE TABLE %s(\nid INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,\n", tbl.name)

		var needIndex []string
		var needForeign []forKey

		// Define each column
		for _, field := range tbl.fields {
			log.Println("Field:", field.name)

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
			return err
		}
	}

	// TODO this query should be applicable for generated go file
	for _, s := range selects {
		j := s.Join
		q := &joinQuery{q: fmt.Sprintf("SELECT * FROM %s\n", j.name)}
		for _, c := range j.childs {
			q.composeJoin(c)
		}
		log.Printf("Query: %s\n%s\n", s.query, q.q)
	}

	return nil
}

func GenerateGo() error {
	return fmt.Errorf("TODO") // TODO
}

func (q *joinQuery) composeJoin(j *join) {
	j.self.name = fmt.Sprintf("%s_%s",
		j.parent.name, j.self.ref.(*Relation).name)

	switch j.type_ {
	case leftJoin:
		q.q += "LEFT JOIN"
	case rightJoin:
		q.q += "RIGHT JOIN"
	default:
		panic(fmt.Errorf("error join type"))
	}

	selfRef := j.self.ref.(*Relation)

	q.q += fmt.Sprintf(" %s %s ON %s.%s_id=%s.id\n",
		selfRef.relateWith.name,
		j.self.name,
		j.parent.name,
		selfRef.name,
		j.self.name,
	)

	for _, c := range j.self.childs {
		q.composeJoin(c)
	}
}
