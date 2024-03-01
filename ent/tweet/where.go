// Code generated by ent, DO NOT EDIT.

package tweet

import (
	"time"
	"twix/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Tweet {
	return predicate.Tweet(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Tweet {
	return predicate.Tweet(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Tweet {
	return predicate.Tweet(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Tweet {
	return predicate.Tweet(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Tweet {
	return predicate.Tweet(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Tweet {
	return predicate.Tweet(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Tweet {
	return predicate.Tweet(sql.FieldLTE(FieldID, id))
}

// TweetID applies equality check predicate on the "tweet_id" field. It's identical to TweetIDEQ.
func TweetID(v *uuid.UUID) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldTweetID, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldContent, v))
}

// Place applies equality check predicate on the "place" field. It's identical to PlaceEQ.
func Place(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldPlace, v))
}

// ReplyCount applies equality check predicate on the "reply_count" field. It's identical to ReplyCountEQ.
func ReplyCount(v int) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldReplyCount, v))
}

// RetweetCount applies equality check predicate on the "retweet_count" field. It's identical to RetweetCountEQ.
func RetweetCount(v int) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldRetweetCount, v))
}

// RateCount applies equality check predicate on the "rate_count" field. It's identical to RateCountEQ.
func RateCount(v int) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldRateCount, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldUpdatedAt, v))
}

// TweetIDEQ applies the EQ predicate on the "tweet_id" field.
func TweetIDEQ(v *uuid.UUID) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldTweetID, v))
}

// TweetIDNEQ applies the NEQ predicate on the "tweet_id" field.
func TweetIDNEQ(v *uuid.UUID) predicate.Tweet {
	return predicate.Tweet(sql.FieldNEQ(FieldTweetID, v))
}

// TweetIDIn applies the In predicate on the "tweet_id" field.
func TweetIDIn(vs ...*uuid.UUID) predicate.Tweet {
	return predicate.Tweet(sql.FieldIn(FieldTweetID, vs...))
}

// TweetIDNotIn applies the NotIn predicate on the "tweet_id" field.
func TweetIDNotIn(vs ...*uuid.UUID) predicate.Tweet {
	return predicate.Tweet(sql.FieldNotIn(FieldTweetID, vs...))
}

// TweetIDGT applies the GT predicate on the "tweet_id" field.
func TweetIDGT(v *uuid.UUID) predicate.Tweet {
	return predicate.Tweet(sql.FieldGT(FieldTweetID, v))
}

// TweetIDGTE applies the GTE predicate on the "tweet_id" field.
func TweetIDGTE(v *uuid.UUID) predicate.Tweet {
	return predicate.Tweet(sql.FieldGTE(FieldTweetID, v))
}

// TweetIDLT applies the LT predicate on the "tweet_id" field.
func TweetIDLT(v *uuid.UUID) predicate.Tweet {
	return predicate.Tweet(sql.FieldLT(FieldTweetID, v))
}

// TweetIDLTE applies the LTE predicate on the "tweet_id" field.
func TweetIDLTE(v *uuid.UUID) predicate.Tweet {
	return predicate.Tweet(sql.FieldLTE(FieldTweetID, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Tweet {
	return predicate.Tweet(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Tweet {
	return predicate.Tweet(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldContainsFold(FieldContent, v))
}

// PlaceEQ applies the EQ predicate on the "place" field.
func PlaceEQ(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldPlace, v))
}

// PlaceNEQ applies the NEQ predicate on the "place" field.
func PlaceNEQ(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldNEQ(FieldPlace, v))
}

// PlaceIn applies the In predicate on the "place" field.
func PlaceIn(vs ...string) predicate.Tweet {
	return predicate.Tweet(sql.FieldIn(FieldPlace, vs...))
}

// PlaceNotIn applies the NotIn predicate on the "place" field.
func PlaceNotIn(vs ...string) predicate.Tweet {
	return predicate.Tweet(sql.FieldNotIn(FieldPlace, vs...))
}

// PlaceGT applies the GT predicate on the "place" field.
func PlaceGT(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldGT(FieldPlace, v))
}

// PlaceGTE applies the GTE predicate on the "place" field.
func PlaceGTE(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldGTE(FieldPlace, v))
}

// PlaceLT applies the LT predicate on the "place" field.
func PlaceLT(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldLT(FieldPlace, v))
}

// PlaceLTE applies the LTE predicate on the "place" field.
func PlaceLTE(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldLTE(FieldPlace, v))
}

// PlaceContains applies the Contains predicate on the "place" field.
func PlaceContains(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldContains(FieldPlace, v))
}

// PlaceHasPrefix applies the HasPrefix predicate on the "place" field.
func PlaceHasPrefix(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldHasPrefix(FieldPlace, v))
}

// PlaceHasSuffix applies the HasSuffix predicate on the "place" field.
func PlaceHasSuffix(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldHasSuffix(FieldPlace, v))
}

// PlaceIsNil applies the IsNil predicate on the "place" field.
func PlaceIsNil() predicate.Tweet {
	return predicate.Tweet(sql.FieldIsNull(FieldPlace))
}

// PlaceNotNil applies the NotNil predicate on the "place" field.
func PlaceNotNil() predicate.Tweet {
	return predicate.Tweet(sql.FieldNotNull(FieldPlace))
}

// PlaceEqualFold applies the EqualFold predicate on the "place" field.
func PlaceEqualFold(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldEqualFold(FieldPlace, v))
}

// PlaceContainsFold applies the ContainsFold predicate on the "place" field.
func PlaceContainsFold(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldContainsFold(FieldPlace, v))
}

// ReplyCountEQ applies the EQ predicate on the "reply_count" field.
func ReplyCountEQ(v int) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldReplyCount, v))
}

// ReplyCountNEQ applies the NEQ predicate on the "reply_count" field.
func ReplyCountNEQ(v int) predicate.Tweet {
	return predicate.Tweet(sql.FieldNEQ(FieldReplyCount, v))
}

// ReplyCountIn applies the In predicate on the "reply_count" field.
func ReplyCountIn(vs ...int) predicate.Tweet {
	return predicate.Tweet(sql.FieldIn(FieldReplyCount, vs...))
}

// ReplyCountNotIn applies the NotIn predicate on the "reply_count" field.
func ReplyCountNotIn(vs ...int) predicate.Tweet {
	return predicate.Tweet(sql.FieldNotIn(FieldReplyCount, vs...))
}

// ReplyCountGT applies the GT predicate on the "reply_count" field.
func ReplyCountGT(v int) predicate.Tweet {
	return predicate.Tweet(sql.FieldGT(FieldReplyCount, v))
}

// ReplyCountGTE applies the GTE predicate on the "reply_count" field.
func ReplyCountGTE(v int) predicate.Tweet {
	return predicate.Tweet(sql.FieldGTE(FieldReplyCount, v))
}

// ReplyCountLT applies the LT predicate on the "reply_count" field.
func ReplyCountLT(v int) predicate.Tweet {
	return predicate.Tweet(sql.FieldLT(FieldReplyCount, v))
}

// ReplyCountLTE applies the LTE predicate on the "reply_count" field.
func ReplyCountLTE(v int) predicate.Tweet {
	return predicate.Tweet(sql.FieldLTE(FieldReplyCount, v))
}

// RetweetCountEQ applies the EQ predicate on the "retweet_count" field.
func RetweetCountEQ(v int) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldRetweetCount, v))
}

// RetweetCountNEQ applies the NEQ predicate on the "retweet_count" field.
func RetweetCountNEQ(v int) predicate.Tweet {
	return predicate.Tweet(sql.FieldNEQ(FieldRetweetCount, v))
}

// RetweetCountIn applies the In predicate on the "retweet_count" field.
func RetweetCountIn(vs ...int) predicate.Tweet {
	return predicate.Tweet(sql.FieldIn(FieldRetweetCount, vs...))
}

// RetweetCountNotIn applies the NotIn predicate on the "retweet_count" field.
func RetweetCountNotIn(vs ...int) predicate.Tweet {
	return predicate.Tweet(sql.FieldNotIn(FieldRetweetCount, vs...))
}

// RetweetCountGT applies the GT predicate on the "retweet_count" field.
func RetweetCountGT(v int) predicate.Tweet {
	return predicate.Tweet(sql.FieldGT(FieldRetweetCount, v))
}

// RetweetCountGTE applies the GTE predicate on the "retweet_count" field.
func RetweetCountGTE(v int) predicate.Tweet {
	return predicate.Tweet(sql.FieldGTE(FieldRetweetCount, v))
}

// RetweetCountLT applies the LT predicate on the "retweet_count" field.
func RetweetCountLT(v int) predicate.Tweet {
	return predicate.Tweet(sql.FieldLT(FieldRetweetCount, v))
}

// RetweetCountLTE applies the LTE predicate on the "retweet_count" field.
func RetweetCountLTE(v int) predicate.Tweet {
	return predicate.Tweet(sql.FieldLTE(FieldRetweetCount, v))
}

// RateCountEQ applies the EQ predicate on the "rate_count" field.
func RateCountEQ(v int) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldRateCount, v))
}

// RateCountNEQ applies the NEQ predicate on the "rate_count" field.
func RateCountNEQ(v int) predicate.Tweet {
	return predicate.Tweet(sql.FieldNEQ(FieldRateCount, v))
}

// RateCountIn applies the In predicate on the "rate_count" field.
func RateCountIn(vs ...int) predicate.Tweet {
	return predicate.Tweet(sql.FieldIn(FieldRateCount, vs...))
}

// RateCountNotIn applies the NotIn predicate on the "rate_count" field.
func RateCountNotIn(vs ...int) predicate.Tweet {
	return predicate.Tweet(sql.FieldNotIn(FieldRateCount, vs...))
}

// RateCountGT applies the GT predicate on the "rate_count" field.
func RateCountGT(v int) predicate.Tweet {
	return predicate.Tweet(sql.FieldGT(FieldRateCount, v))
}

// RateCountGTE applies the GTE predicate on the "rate_count" field.
func RateCountGTE(v int) predicate.Tweet {
	return predicate.Tweet(sql.FieldGTE(FieldRateCount, v))
}

// RateCountLT applies the LT predicate on the "rate_count" field.
func RateCountLT(v int) predicate.Tweet {
	return predicate.Tweet(sql.FieldLT(FieldRateCount, v))
}

// RateCountLTE applies the LTE predicate on the "rate_count" field.
func RateCountLTE(v int) predicate.Tweet {
	return predicate.Tweet(sql.FieldLTE(FieldRateCount, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasComments applies the HasEdge predicate on the "comments" edge.
func HasComments() predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CommentsTable, CommentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentsWith applies the HasEdge predicate on the "comments" edge with a given conditions (other predicates).
func HasCommentsWith(preds ...predicate.Comment) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		step := newCommentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRates applies the HasEdge predicate on the "rates" edge.
func HasRates() predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RatesTable, RatesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRatesWith applies the HasEdge predicate on the "rates" edge with a given conditions (other predicates).
func HasRatesWith(preds ...predicate.Rate) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		step := newRatesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Tweet) predicate.Tweet {
	return predicate.Tweet(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Tweet) predicate.Tweet {
	return predicate.Tweet(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Tweet) predicate.Tweet {
	return predicate.Tweet(sql.NotPredicates(p))
}
