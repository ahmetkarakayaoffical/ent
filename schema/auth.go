package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Authentication holds the schema definition for the Authentication entity.
type Authentication struct {
	ent.Schema
}

// Fields of the Authentication.
func (Authentication) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("use_certificates").Optional().Default(true),
		field.Bool("allow_register").Optional().Default(true),
		field.Bool("use_OIDC").Optional().Default(false),
		field.Enum("OIDC_provider").Values("authentik", "keycloak", "zitadel").Optional(),
		field.String("OIDC_server").Optional().Default(""),
		field.String("OIDC_client_id").Optional().Default(""),
		field.String("OIDC_role").Optional().Default(""),
		field.String("OIDC_cookie_encription_key").Optional().Default(""),
		field.String("OIDC_keycloak_public_key").Optional().Default(""),
		field.Bool("OIDC_auto_create_account").Optional().Default(true),
		field.Bool("OIDC_auto_approve").Optional().Default(true),
	}
}
