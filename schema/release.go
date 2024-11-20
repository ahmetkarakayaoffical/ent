package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Release holds the schema definition for the Release entity.
type Release struct {
	ent.Schema
}

func (Release) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().Unique(),
		field.String("channel").Optional(),
		field.String("summary").Optional(),
		field.String("release_notes").Optional(),
		field.String("file_url").Optional(),
		field.String("checksum").Optional(),
		field.String("is_critical").Optional(),
	}
}

// Edges of the Release.
func (Release) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Agent.Type).Ref("release").Required(),
	}
}
