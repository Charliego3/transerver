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
	sm := map[string]string{dialect.Postgres: "varchar(5)"}
	return []ent.Field{
		field.String("code").Unique().Comment("区域缩写: CN").SchemaType(sm),
		field.String("area").Unique().Comment("区域编码: +86").SchemaType(sm),
		field.String("img").NotEmpty().Comment("国家图标"),
		field.JSON("name", struct {
			En string `json:"en"`
			Zh string `json:"zh"`
		}{}).SchemaType(map[string]string{
			dialect.Postgres: "jsonb",
		}).Comment(`名称: 多语言 -> {"en": "China", "zh": "中国"}`),
	}
}

// Edges of the Region.
func (Region) Edges() []ent.Edge {
	return nil
}
