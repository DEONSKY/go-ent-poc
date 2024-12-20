// Code generated by ent, DO NOT EDIT.

package ent

import (
	"ent-mysql/ent/book"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Book is the model entity for the Book schema.
type Book struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeleteTime holds the value of the "delete_time" field.
	DeleteTime time.Time `json:"delete_time,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Author holds the value of the "author" field.
	Author       string `json:"author,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Book) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case book.FieldTitle, book.FieldAuthor:
			values[i] = new(sql.NullString)
		case book.FieldCreatedAt, book.FieldUpdatedAt, book.FieldDeleteTime:
			values[i] = new(sql.NullTime)
		case book.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Book fields.
func (b *Book) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case book.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				b.ID = *value
			}
		case book.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case book.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		case book.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				b.DeleteTime = value.Time
			}
		case book.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				b.Title = value.String
			}
		case book.FieldAuthor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field author", values[i])
			} else if value.Valid {
				b.Author = value.String
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Book.
// This includes values selected through modifiers, order, etc.
func (b *Book) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// Update returns a builder for updating this Book.
// Note that you need to call Book.Unwrap() before calling this method if this Book
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Book) Update() *BookUpdateOne {
	return NewBookClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Book entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Book) Unwrap() *Book {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Book is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Book) String() string {
	var builder strings.Builder
	builder.WriteString("Book(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete_time=")
	builder.WriteString(b.DeleteTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(b.Title)
	builder.WriteString(", ")
	builder.WriteString("author=")
	builder.WriteString(b.Author)
	builder.WriteByte(')')
	return builder.String()
}

// Books is a parsable slice of Book.
type Books []*Book
