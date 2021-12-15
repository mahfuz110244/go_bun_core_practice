package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	toml "github.com/pelletier/go-toml"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func dbconn(dbconfig *toml.Tree) *bun.DB {

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v search_path=%v sslmode=disable", dbconfig.Get("dbhost").(string), dbconfig.Get("dbuser").(string), dbconfig.Get("dbpass").(string), dbconfig.Get("dbname").(string), dbconfig.Get("schema").(string))

	sqldb, err := sql.Open("postgres", dsn)
	//	db := bun.NewDB(sqldb, sqlitedialect.New())
	//sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	ctx := context.Background()

	var status []Status
	count, err := db.NewSelect().Model(&status).Limit(20).ScanAndCount(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(status, count)

	if err != nil {
		log.Fatal(err.Error())
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("DB Connected...")
	}

	/*var user User1

	db.NewCreateTable().
		Model(user)
	*/

	db.SetMaxOpenConns(SetMaxOpenConns)
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true)))

	return db
}

//gorm
/*func dbconn(dbconfig *toml.Tree) *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v search_path=%v sslmode=disable", dbconfig.Get("dbhost").(string), dbconfig.Get("dbuser").(string), dbconfig.Get("dbpass").(string), dbconfig.Get("dbname").(string), dbconfig.Get("schema").(string))
	//db, err := sql.Open("postgres", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	} else {
		fmt.Println("DB Connected...")
	}

	db.AutoMigrate(&User1{})
	return db
}
*/

/*end postgresql connection*/
