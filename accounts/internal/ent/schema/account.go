package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/transerver/commons/mixin"
	"github.com/transerver/commons/types/enums"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").Unique().Immutable().StructTag(`json:"userId"`),
		field.String("username").NotEmpty(),
		field.String("region").NotEmpty(),
		field.String("area").NotEmpty(),
		field.String("phone").Unique(),
		field.String("email").Unique(),
		field.String("avatar").Optional(),
		field.Bytes("password").NotEmpty().Sensitive(),
		field.Uint8("pwd_level").Default(0).StructTag(`json:"passwordLevel"`),
		field.String("platform").NotEmpty(),
		field.Uint8("state").
			GoType(enums.UserState(0)).
			Default(uint8(enums.UserUnverified)),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return nil
}

func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}
