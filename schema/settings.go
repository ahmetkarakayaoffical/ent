package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Settings holds the schema definition for the Settings entity.
type Settings struct {
	ent.Schema
}

// Fields of the Settings.
func (Settings) Fields() []ent.Field {
	return []ent.Field{
		field.String("language").Optional(),
		field.String("organization").Optional(),
		field.String("postal_address").Optional(),
		field.String("postal_code").Optional(),
		field.String("locality").Optional(),
		field.String("province").Optional(),
		field.String("state").Optional(),
		field.String("country").Optional(),
		field.String("smtp_server").Optional(),
		field.Int("smtp_port").Optional().Default(587),
		field.String("smtp_user").Optional(),
		field.String("smtp_password").Optional(),
		field.String("smtp_auth").Optional().Default("PLAIN"),
		field.Bool("smtp_tls").Optional().Default(false),
		field.Bool("smtp_starttls").Optional().Default(true),
		field.String("message_from").Optional(),
		field.Time("created").Optional().Default(time.Now),
		field.Time("modified").Optional().Default(time.Now).UpdateDefault(time.Now),
	}
}
