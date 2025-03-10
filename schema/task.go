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
		field.Enum("type").Values("winget_install", "winget_update", "winget_delete", "add_registry_key", "add_registry_key_value", "remove_registry_key", "remove_registry_key_value", "environment", "package", "remote_file", "local_user", "local_group", "execute_command", "reboot", "poweroff"),
		field.String("execute").Optional().Default(""),
		field.String("package_id").Optional().Default(""),
		field.String("package_name").Optional().Default(""),
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
