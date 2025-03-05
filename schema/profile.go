package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Profile holds the schema definition for the Profile entity.
type Profile struct {
	ent.Schema
}

// Fields of the Profile.
func (Profile) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Bool("apply_to_all").Default(false),
		field.Enum("type").Values("winget").Optional().Default("winget"),
	}
}

// Edges of the Profile.
func (Profile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type),
		edge.To("tasks", Task.Type),
	}
}
