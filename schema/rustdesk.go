package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RustDesk holds the schema definition for the RustDesk entity.
type RustDesk struct {
	ent.Schema
}

// Fields of the RustDesk.
func (RustDesk) Fields() []ent.Field {
	return []ent.Field{
		field.String("custom_rendezvous_server").Optional().Default(""),
		field.String("relay_server").Optional().Default(""),
		field.String("api_server").Optional().Default(""),
		field.String("key").Optional().Default(""),
	}
}

// Edges of the RustDesk.
func (RustDesk) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).Unique().Ref("rustdesk"),
	}
}
