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
		field.Enum("component").Values("ocsp", "nats", "cert-manager", "agent-worker", "notification-worker", "cert-manager-worker", "console"),
		field.String("version"),
		field.Enum("channel").Values("stable", "testing", "devel"),
	}
}

// Indexes of the Server.
func (Server) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("hostname", "component").Unique(),
	}
}
