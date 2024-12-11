package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Server holds the schema definition for the Server entity.
type Server struct {
	ent.Schema
}

// Fields of the Server.
func (Server) Fields() []ent.Field {
	return []ent.Field{
		field.String("hostname"),
		field.String("arch"),
		field.String("os"),
		field.String("version"),
		field.Enum("channel").Values("stable", "testing", "devel"),
		field.Enum("update_status").Values("Success", "Error", "Pending", "In Progress").Optional(),
		field.String("update_message").Optional(),
		field.Time("update_when").Optional(),
	}
}

// Indexes of the Server.
func (Server) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("hostname", "arch", "os", "version", "channel").Unique(),
	}
}
