package generated

/*
	// Define table

	obj := gen.WantTable("objective", "Objectives")
	obj.HasField("capability", "Capability to achieve by the end", "", gen.TEXT)
	relObjOwner := obj.HasOneToOneRelationshipWith(gen.User, "owner")

	kr := gen.WantTable("key_result", "Key Results")
	kr.HasField("key_result", "Deliverable which determine achievement of the objective", "", gen.TEXT)
	kr.HasField("prob_of_success", "Probability of success", "%", gen.INT)
	kr.HasField("achievement", "Achievement", "%", gen.INT)
	kr.HasField("retro", "Retrospective notes", "", gen.TEXT)
	kr.HasManyToOneRelationshipWith(obj, "obj")

	ki := gen.WantTable("key_init", "Key Initiatives")
	ki.HasField("key_init", "Initiative which is needed to deliver key result", "", gen.TEXT)
	ki.HasManyToOneRelationshipWith(kr, "kr")
	ki.HasOneToOneRelationshipWith(obj, "obj")

	itsp := gen.WantTable("itsp", "IT Strategic Plan")
	itsp.HasField("init", "Initiative to execute", "", gen.TEXT)
	itsp.HasOneToOneRelationshipWith(gen.User, "emit")

	gen.ManyToManyRelationship("ki_itsp", "ki", "itsp", ki, itsp)

	// Define query

	objOwner := obj.LeftJoin(relObjOwner)
*/

// TODO
func TODO() {
	s := ObjOwner.Select()
	s.Where(ObjOwner.AndWhere(
		ObjOwner.Where(ObjOwner.OwnerEmail, "=?", "karim@btpn.com"),
		ObjOwner.Where(ObjOwner.Achievement, "=100"),
	))

	cnt := s.Count()
	if cnt == 0 {

	} else if cnt > 0 {

	}
}
