// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"twix/ent/rate"
	"twix/ent/tweet"
	"twix/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// RateCreate is the builder for creating a Rate entity.
type RateCreate struct {
	config
	mutation *RateMutation
	hooks    []Hook
}

// SetRateID sets the "rate_id" field.
func (rc *RateCreate) SetRateID(u *uuid.UUID) *RateCreate {
	rc.mutation.SetRateID(u)
	return rc
}

// SetReaction sets the "reaction" field.
func (rc *RateCreate) SetReaction(r rate.Reaction) *RateCreate {
	rc.mutation.SetReaction(r)
	return rc
}

// SetCreatedAt sets the "created_at" field.
func (rc *RateCreate) SetCreatedAt(t time.Time) *RateCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *RateCreate) SetNillableCreatedAt(t *time.Time) *RateCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *RateCreate) SetUpdatedAt(t time.Time) *RateCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *RateCreate) SetNillableUpdatedAt(t *time.Time) *RateCreate {
	if t != nil {
		rc.SetUpdatedAt(*t)
	}
	return rc
}

// SetTweetID sets the "tweet" edge to the Tweet entity by ID.
func (rc *RateCreate) SetTweetID(id int) *RateCreate {
	rc.mutation.SetTweetID(id)
	return rc
}

// SetNillableTweetID sets the "tweet" edge to the Tweet entity by ID if the given value is not nil.
func (rc *RateCreate) SetNillableTweetID(id *int) *RateCreate {
	if id != nil {
		rc = rc.SetTweetID(*id)
	}
	return rc
}

// SetTweet sets the "tweet" edge to the Tweet entity.
func (rc *RateCreate) SetTweet(t *Tweet) *RateCreate {
	return rc.SetTweetID(t.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (rc *RateCreate) SetUserID(id int) *RateCreate {
	rc.mutation.SetUserID(id)
	return rc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (rc *RateCreate) SetNillableUserID(id *int) *RateCreate {
	if id != nil {
		rc = rc.SetUserID(*id)
	}
	return rc
}

// SetUser sets the "user" edge to the User entity.
func (rc *RateCreate) SetUser(u *User) *RateCreate {
	return rc.SetUserID(u.ID)
}

// Mutation returns the RateMutation object of the builder.
func (rc *RateCreate) Mutation() *RateMutation {
	return rc.mutation
}

// Save creates the Rate in the database.
func (rc *RateCreate) Save(ctx context.Context) (*Rate, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RateCreate) SaveX(ctx context.Context) *Rate {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RateCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RateCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RateCreate) defaults() {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := rate.DefaultCreatedAt
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		v := rate.DefaultUpdatedAt
		rc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RateCreate) check() error {
	if _, ok := rc.mutation.RateID(); !ok {
		return &ValidationError{Name: "rate_id", err: errors.New(`ent: missing required field "Rate.rate_id"`)}
	}
	if _, ok := rc.mutation.Reaction(); !ok {
		return &ValidationError{Name: "reaction", err: errors.New(`ent: missing required field "Rate.reaction"`)}
	}
	if v, ok := rc.mutation.Reaction(); ok {
		if err := rate.ReactionValidator(v); err != nil {
			return &ValidationError{Name: "reaction", err: fmt.Errorf(`ent: validator failed for field "Rate.reaction": %w`, err)}
		}
	}
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Rate.created_at"`)}
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Rate.updated_at"`)}
	}
	return nil
}

func (rc *RateCreate) sqlSave(ctx context.Context) (*Rate, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RateCreate) createSpec() (*Rate, *sqlgraph.CreateSpec) {
	var (
		_node = &Rate{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(rate.Table, sqlgraph.NewFieldSpec(rate.FieldID, field.TypeInt))
	)
	if value, ok := rc.mutation.RateID(); ok {
		_spec.SetField(rate.FieldRateID, field.TypeUUID, value)
		_node.RateID = value
	}
	if value, ok := rc.mutation.Reaction(); ok {
		_spec.SetField(rate.FieldReaction, field.TypeEnum, value)
		_node.Reaction = value
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(rate.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.SetField(rate.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := rc.mutation.TweetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rate.TweetTable,
			Columns: []string{rate.TweetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.tweet_rates = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rate.UserTable,
			Columns: []string{rate.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_rates = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RateCreateBulk is the builder for creating many Rate entities in bulk.
type RateCreateBulk struct {
	config
	err      error
	builders []*RateCreate
}

// Save creates the Rate entities in the database.
func (rcb *RateCreateBulk) Save(ctx context.Context) ([]*Rate, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Rate, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RateMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation types %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RateCreateBulk) SaveX(ctx context.Context) []*Rate {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RateCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RateCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
