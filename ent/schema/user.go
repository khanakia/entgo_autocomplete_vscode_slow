package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// field.String("id").
		// 	DefaultFunc(func() string {
		// 		// An example of a dumb ID generator - use a production-ready alternative instead.
		// 		return publicid.Must()
		// 	}).Unique(),

		// field.Time("created_at").
		// 	Optional().
		// 	Default(time.Now),

		// field.Time("updated_at").
		// 	Optional().
		// 	Default(time.Now).
		// 	UpdateDefault(time.Now),

		// field.Int("id_num").SchemaType(map[string]string{
		// 	dialect.Postgres: postgres.TypeBigSerial,
		// }).Unique(),

		field.String("email").Unique(),
		field.String("api_key").Optional(),
		field.Bool("status").Optional().Default(false),
	}
}

// Edges of the User.
// func (User) Edges() []ent.Edge {
// 	return []ent.Edge{
// 		edge.To("items", Item.Type),
// 	}
// }
