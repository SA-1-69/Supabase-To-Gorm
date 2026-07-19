# Supabase to GORM

Generate **GORM models** from an existing **Supabase (PostgreSQL)** database using **GORM Gen**.

This project connects to your Supabase PostgreSQL database and automatically generates GORM models from existing tables, eliminating the need to manually write model structs.

---

## Features

- Generate GORM models from an existing Supabase database
- Automatic model generation from PostgreSQL schema
- Supports generating all tables or selected tables
- Easy integration with existing Go projects
- No manual model creation required

---

## Requirements

- Go 1.24 or Go 1.25 (Recommended)
- A Supabase project
- Database password
- Direct database connection

> **Note**
>
> `gorm/gen` may not yet be fully compatible with Go 1.26 due to dependency compatibility issues. If you encounter build errors, use Go 1.24 or Go 1.25.

---

## Installation

Clone the repository.

```bash
git clone https://github.com/your-username/supabase-to-gorm.git
cd supabase-to-gorm
```

Install the required dependencies.

```bash
go mod tidy
```

Or install them manually.

```bash
go get gorm.io/gorm
go get gorm.io/gen
go get gorm.io/driver/postgres
```

---

## Configure Database Connection

### 1. Get the Connection String

1. Open your **Supabase Dashboard**
2. Select your project
3. Click **Connect**
4. Under **Get connected**, choose **Direct connection**
5. Copy the **Connection string**

Example:

```text
postgresql://postgres:Poonchub%40123456@db.ppkbtiuegryphbcolkts.supabase.co:5432/postgres
```

> **Note**
>
> The password inside the connection string is URL-encoded.
>
> Common examples:
>
> - `@` → `%40`
> - `#` → `%23`
> - `$` → `%24`
> - `!` → `%21`

---

### 2. Update the DSN

Open `generate.go` and replace the `dsn` variable.

Using PostgreSQL URI (Recommended)

```go
dsn := "postgresql://postgres:YOUR_PASSWORD@db.xxxxxxxxx.supabase.co:5432/postgres?sslmode=require"
```

Or using GORM parameter format

```go
dsn := "host=db.xxxxxxxxx.supabase.co user=postgres password=YOUR_PASSWORD dbname=postgres port=5432 sslmode=require"
```

---

## Usage

Install dependencies.

```bash
go mod tidy
```

Generate GORM models.

```bash
go run generate.go
```

If the generation is successful, a `models` directory will be created containing all generated models.

---

## Generated Output

Example output:

```text
.
├── generate.go
├── go.mod
├── go.sum
├── models
│   ├── users.gen.go
│   ├── products.gen.go
│   ├── orders.gen.go
│   ├── categories.gen.go
│   └── ...
└── README.md
```

---

## Generate Specific Tables

Generate all tables.

```go
g.ApplyBasic(
    g.GenerateAllTable()...,
)
```

Generate only selected tables.

```go
g.ApplyBasic(
    g.GenerateModel("users"),
    g.GenerateModel("products"),
    g.GenerateModel("orders"),
)
```

---

## Troubleshooting

### Go 1.26 Build Error

If you see an error similar to:

```text
invalid array length -delta * delta
```

or

```text
undefined: gorm.Stmt
```

It is likely caused by dependency incompatibility between `gorm/gen` and Go 1.26.

Recommended solutions:

- Use Go 1.24 or Go 1.25
- Update project dependencies

```bash
go get -u gorm.io/gorm
go get -u gorm.io/gen
go get -u gorm.io/driver/postgres
go mod tidy
```

---

### Connection Failed

Verify the following:

- Database host
- Username
- Password
- SSL mode (`require`)
- Direct Connection is being used
- Firewall or network restrictions

---

## Tech Stack

- Go
- GORM
- GORM Gen
- PostgreSQL
- Supabase

---

## References

- GORM — https://gorm.io/
- GORM Gen — https://gorm.io/gen/
- Supabase — https://supabase.com/
- PostgreSQL — https://www.postgresql.org/

---

## License

This project is licensed under the MIT License.