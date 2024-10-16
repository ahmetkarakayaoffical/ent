package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Agent holds the schema definition for the Agent entity.
type Agent struct {
	ent.Schema
}

// Fields of the Agent.
func (Agent) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().Unique().StorageKey("oid"),
		field.String("os").NotEmpty(),
		field.String("hostname").NotEmpty(),
		field.String("version").NotEmpty(),
		field.String("ip").Default(""),
		field.String("mac").Default(""),
		field.Time("first_contact").Optional(),
		field.Time("last_contact").Optional(),
		field.Bool("enabled").Default(true),
		field.String("vnc").Optional().Default(""),
	}
}

// Edges of the Agent.
func (Agent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("computer", Computer.Type).Unique().Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("operatingsystem", OperatingSystem.Type).Unique().Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("systemupdate", SystemUpdate.Type).Unique().Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("antivirus", Antivirus.Type).Unique().Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("logicaldisks", LogicalDisk.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("apps", App.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("monitors", Monitor.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("shares", Share.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("printers", Printer.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("networkadapters", NetworkAdapter.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("deployments", Deployment.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("updates", Update.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
