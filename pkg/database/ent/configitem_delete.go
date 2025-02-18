// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/crowdsecurity/crowdsec/pkg/database/ent/configitem"
	"github.com/crowdsecurity/crowdsec/pkg/database/ent/predicate"
)

// ConfigItemDelete is the builder for deleting a ConfigItem entity.
type ConfigItemDelete struct {
	config
	hooks    []Hook
	mutation *ConfigItemMutation
}

// Where appends a list predicates to the ConfigItemDelete builder.
func (cid *ConfigItemDelete) Where(ps ...predicate.ConfigItem) *ConfigItemDelete {
	cid.mutation.Where(ps...)
	return cid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cid *ConfigItemDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cid.hooks) == 0 {
		affected, err = cid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ConfigItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cid.mutation = mutation
			affected, err = cid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cid.hooks) - 1; i >= 0; i-- {
			if cid.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cid *ConfigItemDelete) ExecX(ctx context.Context) int {
	n, err := cid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cid *ConfigItemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: configitem.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: configitem.FieldID,
			},
		},
	}
	if ps := cid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cid.driver, _spec)
}

// ConfigItemDeleteOne is the builder for deleting a single ConfigItem entity.
type ConfigItemDeleteOne struct {
	cid *ConfigItemDelete
}

// Exec executes the deletion query.
func (cido *ConfigItemDeleteOne) Exec(ctx context.Context) error {
	n, err := cido.cid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{configitem.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cido *ConfigItemDeleteOne) ExecX(ctx context.Context) {
	cido.cid.ExecX(ctx)
}
