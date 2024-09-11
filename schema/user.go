package schema

import (
	"time"

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
		field.String("id").NotEmpty().Unique().StorageKey("uid"),
		field.String("name"),
		field.String("email"),
		field.String("phone"),
		field.Time("created").Default(time.Now()),
		field.Time("modififed"),
	}
}
