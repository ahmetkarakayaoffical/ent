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
		field.Enum("type").Values("winget_install", "winget_update", "winget_delete", "add_registry_key", "update_registry_key_default_value", "add_registry_key_value", "remove_registry_key", "remove_registry_key_value", "environment", "package", "remote_file", "local_user", "local_group", "execute_command", "reboot", "poweroff"),
		field.String("package_id").Optional().Default(""),
		field.String("package_name").Optional().Default(""),
		field.String("registry_key").Optional().Default(""),
		field.String("registry_key_value_name").Optional().Default(""),
		field.Enum("registry_key_value_type").Values("String", "Binary", "DWord", "QWord", "MultiString", "ExpandString").Optional(),
		field.String("registry_key_value_data").Optional().Default(""),
		field.Bool("registry_hex").Optional().Default(false),
		field.Bool("registry_force").Optional().Default(false),
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
