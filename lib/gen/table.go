package gen

type (
	Table struct {
		name    string
		desc    string
		fields  []*Field
		uniques []*unique
	}

	Field struct {
		type_      FieldType
		tbl        *Table
		name       string
		desc       string
		isNotNull  bool
		isUnique   bool
		unit       string
		relateWith *Table
	}

	Relationship struct {
		*Field
	}

	Join struct {
		ref    interface{}
		name   string
		childs []*join
	}

	Select struct {
		*Join
		query string
	}

	FieldType byte

	unique struct {
		fields []*Field
	}

	join struct {
		type_  joinType
		parent *Join
		self   *Join
	}

	joinType byte
)

const (
	TEXT         FieldType = iota
	INT
	relationship
)

const (
	leftJoin  joinType = iota
	rightJoin
)

var (
	tables  []*Table
	selects []*Select
)

//
// Table
//

func WantTable(name, desc string) *Table {
	my := &Table{name: name, desc: desc}
	tables = append(tables, my)
	return my
}

func (my *Table) HasField(name, desc, unit string, type_ FieldType, isNotNull, isUnique bool) *Field {
	f := &Field{
		type_:     type_,
		tbl:       my,
		name:      name,
		desc:      desc,
		isNotNull: isNotNull,
		isUnique:  isUnique,
		unit:      unit,
	}
	my.fields = append(my.fields, f)
	return f
}

func (my *Table) RelateWith(tbl *Table, name, desc string, isNotNull, isUnique bool) *Relationship {
	f := &Field{
		type_:      relationship,
		tbl:        my,
		name:       name,
		desc:       desc,
		isNotNull:  isNotNull,
		isUnique:   isUnique,
		relateWith: tbl,
	}
	my.fields = append(my.fields, f)
	return &Relationship{f}
}

func (my *Table) Unique(f ...*Field) {
	if len(f) < 2 {
		panic("TODO at least 2 Fields must be defined") // TODO
	}
	my.uniques = append(my.uniques, &unique{f})
}

func (my *Table) Select(name string) *Join {
	s := &Select{query: name, Join: &Join{ref: my, name: my.name}}
	selects = append(selects, s)
	return s.Join
}

//
// Select
//

func (my *Join) LeftJoin(j *Join) *Join {
	// Check parent vs self table name
	var parentTableName string
	switch my.ref.(type) {
	case *Table:
		parentTableName = my.ref.(*Table).name
	case *Relationship:
		parentTableName = my.ref.(*Relationship).relateWith.name
	default:
		panic("TODO")
	}
	rel := j.ref.(*Relationship)
	if parentTableName != rel.tbl.name {
		panic("TODO Table query does not mismatch") // TODO
	}

	// Create the child
	c := &join{leftJoin, my, j}
	my.childs = append(my.childs, c)
	return my
}

func (my *Join) RightJoin(rel *Join) *Join {
	panic("TODO")
}

//
// Relationship
//

func (my *Relationship) Join() *Join {
	return &Join{ref: my}
}

// Define many to many relationship
func ManyToManyRelationship(name string, t ...*Table) {
	panic("TODO")
}
