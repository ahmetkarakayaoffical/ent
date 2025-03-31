package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// SoftwarePackage holds the schema definition for the Package entity.
type SoftwarePackage struct {
	ent.Schema
}

// Fields of the Package.
func (SoftwarePackage) Fields() []ent.Field {
	return []ent.Field{
		field.String("package_id").NotEmpty(),
		field.String("name").NotEmpty(),
		field.Enum("source").Values("winget", "flatpak"),
	}
}
