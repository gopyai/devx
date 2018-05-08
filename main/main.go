package main

import (
	"fmt"
	"time"

	"github.com/gopyai/devx/lib/flow"
	"github.com/gopyai/devx/lib/gen"
)

func main() {
	// def1()
	// def2()
	// def3()

	// panicIf(gen.GenerateForMySQL("shadow", "shadow", "shadow"))
	// panicIf(gen.GenerateGo())

	// defineFlow()
}

func def1() {
	email := gen.WantTable("email", "")
	email.HasField("email", "", "", gen.TEXT, true, gen.UNIQUE)

	usr := gen.WantTable("usr", "")
	usr.HasField("name", "", "", gen.TEXT, true, gen.INDEX)
	usrEmail := usr.RelateWith(email, "email", "", false, gen.UNIQUE)

	obj := gen.WantTable("objective", "")
	obj.HasField("objective", "", "", gen.TEXT, true, gen.INDEX)
	objOwner := obj.RelateWith(usr, "owner", "", false, gen.INDEX)
	objManager := obj.RelateWith(usr, "manager", "", false, gen.INDEX)
	obj.Unique(objOwner.Field, objManager.Field)

	s := obj.Select()
	s.J().
		LeftJoin(objOwner.J().LeftJoin(usrEmail.J())).
		LeftJoin(objManager.J().LeftJoin(usrEmail.J()))
}

func def2() {
	obj := gen.WantTable("objective", "Objectives")
	obj.HasField("capability", "Capability to achieve by the end", "", gen.TEXT, true, gen.UNIQUE)
	obj.RelateWith(gen.User, "owner", "", false, gen.INDEX)

	kr := gen.WantTable("key_result", "Key Results")
	kr.HasField("key_result", "Deliverable which determine achievement of the objective", "", gen.TEXT, true, gen.UNIQUE)
	kr.HasField("prob_of_success", "Probability of success", "%", gen.INT, true, gen.INDEX)
	kr.HasField("achievement", "Achievement", "%", gen.INT, true, gen.INDEX)
	kr.HasField("retro", "Retrospective notes", "", gen.TEXT, false, gen.INDEX)
	kr.RelateWith(obj, "obj", "", false, gen.INDEX)

	ki := gen.WantTable("key_init", "Key Initiatives")
	ki.HasField("key_init", "Initiative which is needed to deliver key result", "", gen.TEXT, true, gen.UNIQUE)
	ki.RelateWith(kr, "kr", "", false, gen.INDEX)
	ki.RelateWith(obj, "obj", "", false, gen.INDEX)

	itsp := gen.WantTable("itsp", "IT Strategic Plan")
	itsp.HasField("init", "Initiative to execute", "", gen.TEXT, true, gen.UNIQUE)
	itsp.RelateWith(gen.User, "emit", "", false, gen.INDEX)

	_, _, err := gen.WantManyToManyRelationTable("ki_itsp", "KI and ITSP relationship", []*gen.ManyToManyRelation{
		{FieldName: "ki", FieldDesc: "Key Initiatives", FromTable: ki},
		{FieldName: "itsp", FieldDesc: "ITSP", FromTable: itsp},
	})
	panicIf(err)

	// Define query
}

func def3() {
	obj := gen.WantTable("objective", "Objectives")
	obj.HasField("capability", "Capability to achieve by the end", "", gen.TEXT, true, gen.UNIQUE)
	objOwner := obj.RelateWith(gen.User, "owner", "", false, gen.INDEX)

	s := obj.Select()
	s.J().LeftJoin(objOwner.J())
	fmt.Println(s.Query())
}

func defineFlow() {
	flow.EveryDayAt(8, 0, 0, func() {

		// Create objective task:

		// Loop every: 1 day at hh:mn:ss
		// ### Rule 1 ###
		// if:
		// - Within certain date and month
		// - no existing objective owned by Karim
		// then:
		// - create task to create objective with due date xxx.
		// ### Rule 2 ###
		// if:
		// - Within certain date and month
		// - There is some objective(s) owned by Karim
		// then:
		// - create task to create objective without due date.

		t := time.Now()
		_, mm, _ := t.Date()
		if mm >= 6 {
			return
		}

		fmt.Println("Create task")

	})

	flow.EveryDayAt(8, 0, 0, func() {
		// Check objective task:

		// ### Rule 1 ###
		// if:
		// - Within certain date and month
		// - There is some objective(s)
		// then:
		// - create task to check all existing objectives at once

	})
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
