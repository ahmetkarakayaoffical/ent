package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Sessions holds the schema definition for the Agent entity.
type Sessions struct {
	ent.Schema
}

// Fields of the Agent.
func (Sessions) Fields() []ent.Field {
	return []ent.Field{
		field.Text("token").NotEmpty().Unique(),
		field.Bytes("data").NotEmpty(),
		field.Time("expiry"),
	}
}

func (Sessions) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("token"),
		index.Fields("expiry").StorageKey("sessions_expiry_idx"),
	}
}
