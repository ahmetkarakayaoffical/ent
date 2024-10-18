package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// OrgMetadata holds the schema definition for the OrgMetadata entity.
type OrgMetadata struct {
	ent.Schema
}

// Fields of the OrgMetadata.
func (OrgMetadata) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
		field.String("description").Optional(),
	}
}
