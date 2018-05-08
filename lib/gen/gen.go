package gen

import (
	"fmt"
)

func GenerateForMySQL(dbName, dbUser, dbPass string) error {
	for _, t := range tables {
		fmt.Println("create table:", t.name)
		fmt.Println("id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY")
		for _, f := range t.fields {
			fmt.Print(f.name)
			switch f.type_ {
			case TEXT:
				fmt.Print(" varchar(500)")
			case INT:
				fmt.Print(" int")
			case relationship:
				fmt.Print("_id int unsigned, relationship with ", f.relateWith.name, "(id)")
			default:
				panic("TODO")
			}

			if f.isNotNull {
				fmt.Print(" not null")
			}

			if f.isUnique {
				fmt.Print(" unique")
			}

			fmt.Println()
		}
		fmt.Println("###")
	}

	for _, s := range selects {
		fmt.Print(s.query, " => ")
		j := s.Join
		fmt.Println("SELECT FROM", j.name)
		for _, c := range j.childs {
			composeJoin(c)
		}
		fmt.Println("###")
	}

	return nil // TODO
}

func composeJoin(j *join) {
	j.self.name = fmt.Sprintf("%s_%s",
		j.parent.name, j.self.ref.(*Relationship).name)

	switch j.type_ {
	case leftJoin:
		fmt.Print("LEFT JOIN")
	case rightJoin:
		fmt.Print("RIGHT JOIN")
	default:
		panic("TODO")
	}

	parentRef := j.parent.ref
	switch parentRef.(type) {
	case (*Table):
	case (*Relationship):
	default:
		panic("TODO")
	}

	selfRef := j.self.ref.(*Relationship)

	fmt.Printf(" %s %s ON %s.%s_id=%s.id\n",
		selfRef.relateWith.name,
		j.self.name,
		j.parent.name,
		selfRef.name,
		j.self.name,
	)

	for _, c := range j.self.childs {
		composeJoin(c)
	}

}

func GenerateGo() error {
	return nil // TODO
}


