package schema

import (
	"fmt"
	"net/url"

	"ent-mysql/ent/schema/schematype"
	"ent-mysql/enum"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Book holds the schema definition for the Book entity.
type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Int("age").
			Positive(),
		field.Float("rank").
			Optional(),
		field.Bool("active").
			Default(false),
		field.String("password").
			Sensitive(),
		field.String("name").
			Unique(),
		field.JSON("url", &url.URL{}).
			Optional(),
		field.JSON("strings", []string{}).
			Optional(),
		field.Int("state").
			GoType(enum.UserState(0)). // Use your custom UserRole Go type
			Validate(func(role int) error {
				if role < int(enum.On) || role > int(enum.Off) {
					return fmt.Errorf("invalid UserRole value: %d", role)
				}
				return nil
			}).
			Default(int(enum.On)), // Set a default value using the ordinal
		field.UUID("uuid", uuid.UUID{}).
			Default(uuid.New),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		schematype.SoftDeleteMixin{},
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("card", Card.Type),
	}
	/*return []ent.Edge{
		edge.To("groups", Group.Type),
		edge.To("friends", User.Type),
	}*/
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("age", "name").
			Unique(),
	}
}
