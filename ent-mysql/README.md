### Example ent ORM for fiber with Postgres

A sample program how to connect ent ORM

## Generate New Migration (If Its Required)

```sh
go generate ./ent 
go run -mod=mod ent/migrate/main.go mig1
```

## Execute Versioned Migrations by Atlas

```sh
atlas migrate apply \                   
  --dir "file://ent/migrate/migrations" \
  --url "postgresql://myuser:mypassword@127.0.0.1/mydatabase?search_path=public&sslmode=disable"
```

### How to start (If no ent dir)
Execute command first
```bash
go run -mod=mod entgo.io/ent/cmd/ent new Book
```
go to `./ent/schema/book.go` and add fields(you want) to Book Schema
```go
// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.String("author").NotEmpty(),
	}
}
```
Execute command
```bash
go generate ./ent
```

#### Endpoints

| Method | URL         | Description     |
|--------|-------------|-----------------|
| GET    | /book       | All Books Info  |
| GET    | /book:id    | One Book Info   |
| POST   | /create     | One Book Add    |
| PUT    | /update/:id | One Book Update |
| DELETE | /delete/:id | One Book Delete |

### 