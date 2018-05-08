package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gopyai/devx/lib/flow"
	"github.com/gopyai/devx/lib/gen"
)

func main() {
	// def1()
	// def2()
	def3()

	log.Println("Error:", gen.GenerateForMySQL("shadow", "shadow", "shadow"))
	log.Println("Error:", gen.GenerateGo())

	// defineFlow()
}

func def3() {
	obj := gen.WantTable("objective", "Objectives")
	obj.HasField("capability", "Capability to achieve by the end", "", gen.TEXT, true, gen.UNIQUE)
	obj.RelateWith(gen.User, "owner", "", false, gen.INDEX)

	obj.Select("hello")
}

func def1() {
	/*
	CREATE TABLE email (
		id    INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
		email VARCHAR(500) NOT NULL UNIQUE
	);

	INSERT INTO email (email) VALUES ('arief@btpn');
	INSERT INTO email (email) VALUES ('ana@home');
	*/

	email := gen.WantTable("email", "")
	email.HasField("email", "", "", gen.TEXT, true, gen.UNIQUE)

	/*
	CREATE TABLE usr (
		id       INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
		name     VARCHAR(500) NOT NULL,
		email_id INT UNSIGNED,

		INDEX (name),
		FOREIGN KEY (email_id) REFERENCES email (id)
			ON DELETE CASCADE
			ON UPDATE CASCADE
	);

	INSERT INTO usr (name, email_id) VALUES ('arief', 1);
	INSERT INTO usr (name, email_id) VALUES ('ana', 2);
	*/

	usr := gen.WantTable("usr", "")
	usr.HasField("name", "", "", gen.TEXT, true, gen.INDEX)
	usrEmail := usr.RelateWith(email, "email", "", false, gen.UNIQUE)

	/*
	CREATE TABLE objective (
		id         INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
		objective  VARCHAR(500) NOT NULL,
		owner_id   INT UNSIGNED,
		manager_id INT UNSIGNED,

		INDEX (objective),
		FOREIGN KEY (owner_id) REFERENCES usr (id)
			ON DELETE CASCADE
			ON UPDATE CASCADE,
		FOREIGN KEY (manager_id) REFERENCES usr (id)
			ON DELETE CASCADE
			ON UPDATE CASCADE
	);

	INSERT INTO objective (objective, owner_id, manager_id) VALUES ('Bekerja', 1, 2);
	INSERT INTO objective (objective, owner_id, manager_id) VALUES ('Mendidik', 2, 2);

	*/

	obj := gen.WantTable("objective", "")
	obj.HasField("objective", "", "", gen.TEXT, true, gen.INDEX)
	objOwner := obj.RelateWith(usr, "owner", "", false, gen.INDEX)
	objManager := obj.RelateWith(usr, "manager", "", false, gen.INDEX)
	obj.Unique(objOwner.Field, objManager.Field)

	/*
	SELECT *
		FROM objective

	LEFT JOIN usr obj_owner on objective.owner_id = obj_owner.id
	LEFT JOIN email obj_owner_email on obj_owner.email_id = obj_owner_email.id

	LEFT JOIN usr obj_mgr on objective.manager_id = obj_mgr.id
	LEFT JOIN email obj_mgr_email on obj_mgr.email_id = obj_mgr_email.id
	*/

	panicIf(obj.Select("ObjOwnerManager").
		LeftJoin(objOwner.Join().LeftJoin(usrEmail.Join())).
		LeftJoin(objManager.Join().LeftJoin(usrEmail.Join())).Err)
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
