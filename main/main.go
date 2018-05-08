package main

import (
	"fmt"
	"time"

	"btpn/app/devx/lib/flow"
	"btpn/app/devx/lib/gen"
)

func main() {
	def1()
	//def2()

	panicIf(gen.GenerateForMySQL("shadow", "shadow", "shadow"))
	panicIf(gen.GenerateGo())

	//defineFlow()
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
	email.HasField("email", "", "", gen.TEXT, true, true)

	/*
	CREATE TABLE usr (
		id       INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
		name     VARCHAR(500) NOT NULL,
		email_id INT UNSIGNED,

		FOREIGN KEY (email_id) REFERENCES email (id)
			ON DELETE CASCADE
			ON UPDATE CASCADE
	);

	INSERT INTO usr (name, email_id) VALUES ('arief', 1);
	INSERT INTO usr (name, email_id) VALUES ('ana', 2);
	*/

	usr := gen.WantTable("usr", "")
	usr.HasField("name", "", "", gen.TEXT, true, false)
	usrEmail := usr.RelateWith(email, "email", "", false, false)

	/*
	CREATE TABLE objective (
		id         INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
		objective  VARCHAR(500) NOT NULL,
		owner_id   INT UNSIGNED,
		manager_id INT UNSIGNED,

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
	obj.HasField("objective", "", "", gen.TEXT, true, false)
	objOwner := obj.RelateWith(usr, "owner", "", false, false)
	objManager := obj.RelateWith(usr, "manager", "", false, false)

	/*
	SELECT *
		FROM objective

	LEFT JOIN usr obj_owner on objective.owner_id = obj_owner.id
	LEFT JOIN email obj_owner_email on obj_owner.email_id = obj_owner_email.id

	LEFT JOIN usr obj_mgr on objective.manager_id = obj_mgr.id
	LEFT JOIN email obj_mgr_email on obj_mgr.email_id = obj_mgr_email.id
	*/

	obj.Select("objOwnerManager").
		LeftJoin(objOwner.Join().LeftJoin(usrEmail.Join())).
		LeftJoin(objManager.Join().LeftJoin(usrEmail.Join()))
}

func def2() {
	// Define table

	obj := gen.WantTable("objective", "Objectives")
	obj.HasField("capability", "Capability to achieve by the end", "", gen.TEXT, true, true)
	obj.RelateWith(gen.User, "owner", "", false, false)

	kr := gen.WantTable("keyResult", "Key Results")
	kr.HasField("keyResult", "Deliverable which determine achievement of the objective", "", gen.TEXT, true, true)
	kr.HasField("probOfSuccess", "Probability of success", "%", gen.INT, true, false)
	kr.HasField("achievement", "Achievement", "%", gen.INT, true, false)
	kr.HasField("retro", "Retrospective notes", "", gen.TEXT, false, false)
	kr.RelateWith(obj, "obj", "", false, false)

	ki := gen.WantTable("keyInit", "Key Initiatives")
	ki.HasField("keyInit", "Initiative which is needed to deliver key result", "", gen.TEXT, true, true)
	ki.RelateWith(kr, "kr", "", false, false)
	ki.RelateWith(obj, "obj", "", false, true)

	itsp := gen.WantTable("itsp", "IT Strategic Plan")
	itsp.HasField("init", "Initiative to execute", "", gen.TEXT, true, true)
	itsp.RelateWith(gen.User, "emit", "", false, false)

	gen.ManyToManyRelationship("kiItsp", ki, itsp)

	// Define query

	// Generate

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
