package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

var timeSchema = map[string]string{
	dialect.Postgres: "timestamptz",
}

func NewNullTime() *sql.NullTime {
	return &sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
}

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("create_at").
			StructTag(`json:"createAt,omitempty"`).
			Immutable().
			Default(time.Now).
			SchemaType(timeSchema),

		field.Time("update_at").
			StructTag(`json:"updateAt,omitempty"`).
			Optional().
			GoType(&sql.NullTime{}).
			UpdateDefault(NewNullTime).
			SchemaType(timeSchema),
	}
}
