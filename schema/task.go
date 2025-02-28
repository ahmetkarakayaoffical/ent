package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Enum("type").Optional().Values("tasks.winget_install", "tasks.winget_update", "tasks.winget_delete", "tasks.execute_command", "tasks.reboot", "tasks.poweroff"),
		field.String("execute").Optional().Default(""),
		field.Time("when").Optional(),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type),
		edge.From("profile", Profile.Type).Unique().Ref("tasks"),
	}
}
