package gen

// Predefined tables

var (
	User     *Table
	UserRole *Table
)

func init() {
	User = WantTable("user", "Predefined user table")
	User.HasField("email", "Email address", "", TEXT, true, UNIQUE)
	User.HasField("token", "Token for login", "", TEXT, true, UNIQUE)
	User.HasField("token_renew", "Date to renew the token", "", DATE, true, INDEX)
	User.HasField("token_expire", "Expiration date of token", "", DATE, true, INDEX)

	UserRole = WantTable("user_role", "Predefined role table")
	UserRole.Unique(
		UserRole.RelateWith(User, "user", "User", true, INDEX).Field,
		UserRole.HasField("role", "Role name", "", TEXT, true, INDEX),
	)
}
