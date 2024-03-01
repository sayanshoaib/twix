// Code generated by ent, DO NOT EDIT.

package rate

import (
	"time"
	"twix/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Rate {
	return predicate.Rate(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Rate {
	return predicate.Rate(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Rate {
	return predicate.Rate(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Rate {
	return predicate.Rate(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Rate {
	return predicate.Rate(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Rate {
	return predicate.Rate(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Rate {
	return predicate.Rate(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Rate {
	return predicate.Rate(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Rate {
	return predicate.Rate(sql.FieldLTE(FieldID, id))
}

// RateID applies equality check predicate on the "rate_id" field. It's identical to RateIDEQ.
func RateID(v *uuid.UUID) predicate.Rate {
	return predicate.Rate(sql.FieldEQ(FieldRateID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Rate {
	return predicate.Rate(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Rate {
	return predicate.Rate(sql.FieldEQ(FieldUpdatedAt, v))
}

// RateIDEQ applies the EQ predicate on the "rate_id" field.
func RateIDEQ(v *uuid.UUID) predicate.Rate {
	return predicate.Rate(sql.FieldEQ(FieldRateID, v))
}

// RateIDNEQ applies the NEQ predicate on the "rate_id" field.
func RateIDNEQ(v *uuid.UUID) predicate.Rate {
	return predicate.Rate(sql.FieldNEQ(FieldRateID, v))
}

// RateIDIn applies the In predicate on the "rate_id" field.
func RateIDIn(vs ...*uuid.UUID) predicate.Rate {
	return predicate.Rate(sql.FieldIn(FieldRateID, vs...))
}

// RateIDNotIn applies the NotIn predicate on the "rate_id" field.
func RateIDNotIn(vs ...*uuid.UUID) predicate.Rate {
	return predicate.Rate(sql.FieldNotIn(FieldRateID, vs...))
}

// RateIDGT applies the GT predicate on the "rate_id" field.
func RateIDGT(v *uuid.UUID) predicate.Rate {
	return predicate.Rate(sql.FieldGT(FieldRateID, v))
}

// RateIDGTE applies the GTE predicate on the "rate_id" field.
func RateIDGTE(v *uuid.UUID) predicate.Rate {
	return predicate.Rate(sql.FieldGTE(FieldRateID, v))
}

// RateIDLT applies the LT predicate on the "rate_id" field.
func RateIDLT(v *uuid.UUID) predicate.Rate {
	return predicate.Rate(sql.FieldLT(FieldRateID, v))
}

// RateIDLTE applies the LTE predicate on the "rate_id" field.
func RateIDLTE(v *uuid.UUID) predicate.Rate {
	return predicate.Rate(sql.FieldLTE(FieldRateID, v))
}

// ReactionEQ applies the EQ predicate on the "reaction" field.
func ReactionEQ(v Reaction) predicate.Rate {
	return predicate.Rate(sql.FieldEQ(FieldReaction, v))
}

// ReactionNEQ applies the NEQ predicate on the "reaction" field.
func ReactionNEQ(v Reaction) predicate.Rate {
	return predicate.Rate(sql.FieldNEQ(FieldReaction, v))
}

// ReactionIn applies the In predicate on the "reaction" field.
func ReactionIn(vs ...Reaction) predicate.Rate {
	return predicate.Rate(sql.FieldIn(FieldReaction, vs...))
}

// ReactionNotIn applies the NotIn predicate on the "reaction" field.
func ReactionNotIn(vs ...Reaction) predicate.Rate {
	return predicate.Rate(sql.FieldNotIn(FieldReaction, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Rate {
	return predicate.Rate(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Rate {
	return predicate.Rate(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Rate {
	return predicate.Rate(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Rate {
	return predicate.Rate(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Rate {
	return predicate.Rate(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Rate {
	return predicate.Rate(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Rate {
	return predicate.Rate(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Rate {
	return predicate.Rate(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Rate {
	return predicate.Rate(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Rate {
	return predicate.Rate(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Rate {
	return predicate.Rate(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Rate {
	return predicate.Rate(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Rate {
	return predicate.Rate(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Rate {
	return predicate.Rate(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Rate {
	return predicate.Rate(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Rate {
	return predicate.Rate(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasTweet applies the HasEdge predicate on the "tweet" edge.
func HasTweet() predicate.Rate {
	return predicate.Rate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TweetTable, TweetColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTweetWith applies the HasEdge predicate on the "tweet" edge with a given conditions (other predicates).
func HasTweetWith(preds ...predicate.Tweet) predicate.Rate {
	return predicate.Rate(func(s *sql.Selector) {
		step := newTweetStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Rate {
	return predicate.Rate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Rate {
	return predicate.Rate(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Rate) predicate.Rate {
	return predicate.Rate(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Rate) predicate.Rate {
	return predicate.Rate(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Rate) predicate.Rate {
	return predicate.Rate(sql.NotPredicates(p))
}
