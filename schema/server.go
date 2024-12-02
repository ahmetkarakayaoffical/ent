package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
	}
}

// Edges of the Server.
func (Server) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("release", Release.Type).Ref("servers").Unique(),
	}
}

// Indexes of the Server.
func (Server) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("hostname", "component").Unique(),
	}
}
