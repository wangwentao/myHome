package main

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"myHome/gin/configs"
	"myHome/gin/utils"
)

var ctx = context.Background()

func testRedis() {
	fmt.Println("Redis Test")

	//val, err := configs.RedisStore.Ping(ctx).Result()
	val, err := configs.RedisStore.Set(ctx, "session_key", "test_user_session", 0).Result()
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}
	fmt.Println(val)
}

var PGCon *pgx.Conn

func testPgSQL() {
	fmt.Println("PostgreSQL Test")

	var err error
	PGCon, err = pgx.Connect(ctx, "postgres://postgres:123456@localhost:5432/postgres?sslmode=disable")
	utils.CheckErr(err)

	// configs.PGCon.Ping(ctx)
	// Add
	// pgAdd()

	// Delete
	// pgDel()

	// Update
	// pgUpdate()

	// Query
	//pgQuery()

	// ormQuery
	// scanQuery()

}

func pgAdd() {
	addSql := "insert into company values ($1,$2,$3,$4,$5)"
	cmdTag, err := PGCon.Exec(ctx, addSql, 3, "wangwu", 25, "shanghai", 4000)
	utils.CheckErr(err)
	fmt.Println("Insert Affected Rows: " + fmt.Sprintf("%d", cmdTag.RowsAffected()))
}

func pgDel() {
	delSql := "delete from company where name = $1"
	cmdTag, err := PGCon.Exec(ctx, delSql, "wangwu")
	utils.CheckErr(err)
	fmt.Println("Delete Affected Rows: " + fmt.Sprintf("%d", cmdTag.RowsAffected()))
}

func pgUpdate() {
	updateSql := "update company set salary = 6000 where name = $1"
	cmdTag, err := PGCon.Exec(ctx, updateSql, "wangwu")
	utils.CheckErr(err)
	// type change
	fmt.Println("Update Affected Rows: " + fmt.Sprintf("%d", cmdTag.RowsAffected()))
}

func pgQuery() {
	var id, age int
	var name, address string
	var salary float32
	querySql := "select * from company where name = $1"
	err := PGCon.QueryRow(ctx, querySql, "lisi").Scan(&id, &name, &age, &address, &salary)
	utils.CheckErr(err)

	fmt.Println(id, name, age, address, salary)

}

type Company struct {
	Id      int
	Name    string
	Age     int
	Address string
	Salary  float32
}

func scanQuery() {
	var comps []Company
	querySql := "select id,name,age,address,salary from company"

	err := pgxscan.Select(ctx, PGCon, &comps, querySql)
	utils.CheckErr(err)
	fmt.Println(comps[0])

	var comp Company
	getSql := "select id,name,age,address,salary from company where name = $1"
	err = pgxscan.Get(ctx, PGCon, &comp, getSql, "lisi")
	utils.CheckErr(err)
	fmt.Println(comp)

}

func testGORM() {

	//Insert
	//ormInsert()

	//Delete
	//ormDelete()

	//Update
	// ormUpdate()

	//Query
	ormQuery()

}

func ormInsert() {
	comp := Company{Id: 4, Name: "xiaoqi", Age: 29, Address: "beijing", Salary: 12000}
	res := configs.PGStore.Table("company").Create(&comp)
	utils.CheckErr(res.Error)
	fmt.Println("Update Affected Rows: " + fmt.Sprintf("%d", res.RowsAffected))
}

func ormDelete() {
	comp := Company{Id: 4, Name: "xiaoqi", Age: 29, Address: "beijing", Salary: 12000}
	res := configs.PGStore.Table("company").Delete(&comp)

	utils.CheckErr(res.Error)
	fmt.Println("Update Affected Rows: " + fmt.Sprintf("%d", res.RowsAffected))

}

func ormUpdate() {
	// 1
	/*comp := Company{Id: 4, Name: "xiaoqi", Age: 29, Address: "beijing", Salary: 8000}
	res := configs.PGStore.Table("company").Save(&comp)

	utils.CheckErr(res.Error)
	fmt.Println("Update Affected Rows: "+fmt.Sprintf("%d", res.RowsAffected))*/

	// 2
	/*var comp Company
	configs.PGStore.Table("company").First(&comp)
	comp.Salary = 500
	res := configs.PGStore.Table("company").Save(&comp)
	utils.CheckErr(res.Error)
	fmt.Println("Update Affected Rows: "+fmt.Sprintf("%d", res.RowsAffected))*/

	// 3 update column
	res := configs.PGStore.Table("company").Where("name=?", "xiaoqi").Update("salary", 12000)
	utils.CheckErr(res.Error)
	fmt.Println("Update Affected Rows: " + fmt.Sprintf("%d", res.RowsAffected))

}

func ormQuery() {

}

func main() {
	configs.InitSettings(ctx)

	defer configs.ReleaseSettings(ctx)

	//testRedis()
	//testPgSQL()

	// testGORM()

}
