package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Site holds the schema definition for the Site entity.
type Site struct {
	ent.Schema
}

// Fields of the Site.
func (Site) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Optional(),
		field.Bool("is_default").Optional(),
	}
}

// Edges of the Site.
func (Site) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).Ref("sites").Unique(),
		edge.To("agents", Agent.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("profiles", Profile.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
