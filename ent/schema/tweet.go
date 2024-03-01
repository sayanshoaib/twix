package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"net/url"
	"time"
)

// Tweet holds the schema definition for the Tweet entity.
type Tweet struct {
	ent.Schema
}

// Fields of the Tweet.
func (Tweet) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("tweet_id", &uuid.UUID{}),
		field.String("content").NotEmpty(),
		field.JSON("media", &url.URL{}),
		field.String("place").Optional().Nillable(),
		field.Int("reply_count").Positive().Default(0),
		field.Int("retweet_count").Positive().Default(0),
		field.Int("rate_count").Positive().Default(0),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()).UpdateDefault(time.Now()),
	}
}

// Edges of the Tweet.
func (Tweet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("tweets").
			Unique(),
		edge.To("comments", Comment.Type),
		edge.To("rates", Rate.Type),
	}
}
