package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"ent-mysql/ent/schema/schematype"

	"github.com/google/uuid"
)

type Card struct {
	ent.Schema
}

func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("card_no"),
	}
}

func (Card) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		schematype.SoftDeleteMixin{},
	}
}

/*
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type),
	}
	/*return []ent.Edge{
		edge.To("groups", Group.Type),
		edge.To("friends", User.Type),
	}*/
/*}*/
