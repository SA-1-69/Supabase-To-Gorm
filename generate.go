package main

import (
    "gorm.io/driver/postgres"
    "gorm.io/gen"
    "gorm.io/gorm"
)

func main() {
    dsn := "<Your Data Source Name>"

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }

    g := gen.NewGenerator(gen.Config{
        OutPath: "./models",
        Mode: gen.WithDefaultQuery | gen.WithQueryInterface,
    })

    g.UseDB(db)

    g.ApplyBasic(
        g.GenerateAllTable()...,
    )

    g.Execute()
}
