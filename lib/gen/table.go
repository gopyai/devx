package gen

import (
	"fmt"
	"log"
)

type (
	Table struct {
		name    string
		desc    string
		fields  []*Field
		uniques [][]*Field
	}

	Field struct {
		name       string
		desc       string
		unit       string
		type_      FieldType
		isNotNull  bool
		index      IndexType
		tbl        *Table
		relateWith *Table
	}

	Relation struct {
		*Field
	}

	Join struct {
		Err    error
		ref    interface{}
		name   string
		childs []*join
	}

	Select struct {
		*Join
		query string
	}

	ManyToManyRelation struct {
		FieldName string
		FieldDesc string
		FromTable *Table
	}

	FieldType byte
	IndexType byte

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
	DATE
	relationship
)

const (
	leftJoin  joinType = iota
	rightJoin
)

const (
	INDEX  IndexType = iota
	UNIQUE
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

func (my *Table) HasField(name, desc, unit string, type_ FieldType, isNotNull bool, index IndexType) *Field {
	f := &Field{
		type_:     type_,
		tbl:       my,
		name:      name,
		desc:      desc,
		isNotNull: isNotNull,
		index:     index,
		unit:      unit,
	}
	my.fields = append(my.fields, f)
	return f
}

func (my *Table) RelateWith(tbl *Table, name, desc string, isNotNull bool, index IndexType) *Relation {
	f := &Field{
		type_:      relationship,
		tbl:        my,
		name:       name,
		desc:       desc,
		isNotNull:  isNotNull,
		index:      index,
		relateWith: tbl,
	}
	my.fields = append(my.fields, f)
	return &Relation{f}
}

func (my *Table) Unique(f ...*Field) error {
	if len(f) < 2 {
		return fmt.Errorf("must be at least 2 fields")
	}
	my.uniques = append(my.uniques, f)
	return nil
}

func (my *Table) Select(name string) *Join {
	s := &Select{query: name, Join: &Join{ref: my, name: my.name}}
	selects = append(selects, s)
	return s.Join
}

//
// Select
//

func (my *Join) doJoin(j *Join, t joinType) {
	// If child join is error, then this join will be error too
	if j.Err != nil {
		my.Err = j.Err
	}
	// Return early for error
	if my.Err != nil {
		return
	}

	// Check parent vs self table name
	var parentTableName string
	switch my.ref.(type) {
	case *Table:
		parentTableName = my.ref.(*Table).name
	case *Relation:
		parentTableName = my.ref.(*Relation).relateWith.name
	default:
		panic(fmt.Errorf("error ref type"))
	}
	rel := j.ref.(*Relation)
	if parentTableName != rel.tbl.name {
		my.Err = fmt.Errorf("error join using unmatched relationship")
		log.Println(my.Err)
	}

	// Create the child
	c := &join{t, my, j}
	my.childs = append(my.childs, c)
}

func (my *Join) LeftJoin(j *Join) *Join {
	my.doJoin(j, leftJoin)
	return my
}

func (my *Join) RightJoin(j *Join) *Join {
	my.doJoin(j, rightJoin)
	return my
}

//
// Relationship
//

func (my *Relation) Join() *Join {
	return &Join{ref: my}
}

// Define many to many relationship
func WantManyToManyRelationTable(tblName, tblDesc string, relInfo []*ManyToManyRelation) (tbl *Table, rels []*Relation, err error) {
	var flds []*Field
	tbl = WantTable(tblName, tblDesc)
	for _, i := range relInfo {
		r := tbl.RelateWith(i.FromTable, i.FieldName, i.FieldDesc, true, INDEX)
		rels = append(rels, r)
		flds = append(flds, r.Field)
	}
	err = tbl.Unique(flds...)
	if err != nil {
		tables = tables[:len(tables)-1] // Remove last table entry if error, to cancer table creation
	}
	return
}
