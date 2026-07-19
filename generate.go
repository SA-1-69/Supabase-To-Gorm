package main

import (
    "gorm.io/driver/postgres"
    "gorm.io/gen"
    "gorm.io/gorm"
)

func main() {
    dsn := "postgresql://postgres:Poonchub%40123456@db.ppkbtiuegryphbcolkts.supabase.co:5432/postgres"

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