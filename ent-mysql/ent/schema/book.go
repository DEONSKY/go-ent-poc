package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"ent-mysql/ent/schema/schematype"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("title").NotEmpty(),
		field.String("author").NotEmpty(),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return nil
}

func (Book) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		schematype.SoftDeleteMixin{},
	}
}
