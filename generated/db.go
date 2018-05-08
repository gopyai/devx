package generated

/*
2018/05/08 15:03:24 Table: user
2018/05/08 15:03:24 Field: email
2018/05/08 15:03:24 Field: token
2018/05/08 15:03:24 Field: token_renew
2018/05/08 15:03:24 Field: token_expire
2018/05/08 15:03:24 Will execute query:
CREATE TABLE user(
id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
email VARCHAR(500) NOT NULL UNIQUE,
token VARCHAR(500) NOT NULL UNIQUE,
token_renew DATE NOT NULL,
token_expire DATE NOT NULL,
INDEX(token_renew),
INDEX(token_expire)
)
 */

type (
	TableUser struct {
	}
)

var (
	User *TableUser
)

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

/*

2018/05/08 15:03:24 Table: user_role
2018/05/08 15:03:24 Field: user
2018/05/08 15:03:24 Field: role
2018/05/08 15:03:24 Will execute query:
CREATE TABLE user_role(
id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
user_id INT UNSIGNED NOT NULL,
role VARCHAR(500) NOT NULL,
UNIQUE(user_id,role),
INDEX(role),
FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE ON UPDATE CASCADE
)
2018/05/08 15:03:24 Table: email
2018/05/08 15:03:24 Field: email
2018/05/08 15:03:24 Will execute query:
CREATE TABLE email(
id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
email VARCHAR(500) NOT NULL UNIQUE
)
2018/05/08 15:03:24 Table: usr
2018/05/08 15:03:24 Field: name
2018/05/08 15:03:24 Field: email
2018/05/08 15:03:24 Will execute query:
CREATE TABLE usr(
id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
name VARCHAR(500) NOT NULL,
email_id INT UNSIGNED UNIQUE,
INDEX(name),
FOREIGN KEY (email_id) REFERENCES email(id) ON DELETE CASCADE ON UPDATE CASCADE
)
2018/05/08 15:03:24 Table: objective
2018/05/08 15:03:24 Field: objective
2018/05/08 15:03:24 Field: owner
2018/05/08 15:03:24 Field: manager
2018/05/08 15:03:24 Will execute query:
CREATE TABLE objective(
id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
objective VARCHAR(500) NOT NULL,
owner_id INT UNSIGNED,
manager_id INT UNSIGNED,
UNIQUE(owner_id,manager_id),
INDEX(objective),
FOREIGN KEY (owner_id) REFERENCES usr(id) ON DELETE CASCADE ON UPDATE CASCADE,
FOREIGN KEY (manager_id) REFERENCES usr(id) ON DELETE CASCADE ON UPDATE CASCADE
)
2018/05/08 15:03:24 Query: ObjOwnerManager
&{SELECT * FROM objective
LEFT JOIN usr objective_owner ON objective.owner_id=objective_owner.id
LEFT JOIN email objective_owner_email ON objective_owner.email_id=objective_owner_email.id
LEFT JOIN usr objective_manager ON objective.manager_id=objective_manager.id
LEFT JOIN email objective_manager_email ON objective_manager.email_id=objective_manager_email.id
}
2018/05/08 15:03:24 Error: <nil>
2018/05/08 15:03:24 Error: TODO

Process finished with exit code 0


 */
