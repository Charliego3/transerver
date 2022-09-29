// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/transerver/accounts/internal/ent/rsa"
)

// Rsa is the model entity for the Rsa schema.
type Rsa struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Rsa) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case rsa.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Rsa", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Rsa fields.
func (r *Rsa) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case rsa.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this Rsa.
// Note that you need to call Rsa.Unwrap() before calling this method if this Rsa
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Rsa) Update() *RsaUpdateOne {
	return (&RsaClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Rsa entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Rsa) Unwrap() *Rsa {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Rsa is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Rsa) String() string {
	var builder strings.Builder
	builder.WriteString("Rsa(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Rsas is a parsable slice of Rsa.
type Rsas []*Rsa

func (r Rsas) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
