package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		field.String("email").Optional(),
		field.String("phone").Optional(),
		field.String("csr").Optional(),
		field.String("certSerial").Optional(),
		field.Time("expiry").Optional(),
		field.Time("created").Optional().Default(time.Now),
		field.Time("modified").Optional().Default(time.Now).UpdateDefault(time.Now),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sessions", Sessions.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email").StorageKey("users_email_idx"),
	}
}
