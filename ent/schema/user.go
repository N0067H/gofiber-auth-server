package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").
			Unique().
			NotEmpty().
			MaxLen(255),
		field.String("name").
			NotEmpty().
			MaxLen(100),
		field.String("password").
			NotEmpty().
			MaxLen(255),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
