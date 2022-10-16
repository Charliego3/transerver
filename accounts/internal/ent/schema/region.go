package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Region holds the schema definition for the Region entity.
type Region struct {
	ent.Schema
}

// Fields of the Region.
func (Region) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").StructTag(`json:"code,omitempty"`),
		field.String("area").StructTag(`json:"area,omitempty"`),
		field.String("img").StructTag(`json:"img,omitempty"`),
		field.JSON("name", struct {
			En string `json:"en"`
			Zh string `json:"zh"`
		}{}).SchemaType(map[string]string{
			dialect.Postgres: "json",
		}),
	}
}

// Edges of the Region.
func (Region) Edges() []ent.Edge {
	return nil
}
