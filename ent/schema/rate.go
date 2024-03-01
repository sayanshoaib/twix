package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Rate holds the schema definition for the Rate entity.
type Rate struct {
	ent.Schema
}

// Fields of the Rate.
func (Rate) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("rate_id", &uuid.UUID{}),
		field.Enum("reaction").
			Values("like", "funny", "love", "sad"),
		field.Time("created_at").
			Default(time.Now()),
		field.Time("updated_at").
			Default(time.Now()).
			UpdateDefault(time.Now()),
	}
}

// Edges of the Rate.
func (Rate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tweet", Tweet.Type).
			Ref("rates").
			Unique(),
		edge.From("user", User.Type).
			Ref("rates").
			Unique(),
	}
}
