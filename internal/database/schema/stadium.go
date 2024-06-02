package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/privacy"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Stadium struct {
	ent.Schema
}

func (Stadium) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("team").NotEmpty(),
		field.String("fdcouk").NotEmpty(),
		field.String("city").NotEmpty(),
		field.String("stadium").NotEmpty().Unique(),
		field.Int("capacity"),
		field.Float("latitude"),
		field.Float("longitude"),
		field.String("country").NotEmpty(),
	}
}

func (Stadium) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.AlwaysAllowRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}

func (Stadium) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("weather", Weather.Type).Unique(),
	}
}
