package main

import "fmt"

type A struct {
	Id  int64 `sql:"id"`
	bId int64 `sql:"b_id"`
	B   *B
}

type B struct {
	Id  int64 `sql:"id"`
	cId int64 `sql:"c_id" json:"`
	C   *C
}

type C struct {
	Id   int64  `sql:"id"`
	Name string `sql:"name"`
}

var querySelect = `
SELECT
                 a.id,                
                 b.id		AS B__id,
                 c.name		AS B__C__name
        FROM a
        LEFT JOIN b ON a.b_id = b.id
        LEFT JOIN c ON b.c_id = c.id
;
`

var queryCreateAndInsert = `
BEGIN
;

CREATE TABLE IF NOT EXISTS c
(
    name    varchar(64),
    id      bigserial not null
				constraint c_pkey
				primary key
)
;

CREATE TABLE IF NOT EXISTS b
(
    id      bigserial not null
				constraint b_pkey
				primary key,
    c_id    BIGINT references c(id)
)
;

CREATE TABLE IF NOT EXISTS a
(
    id     bigserial not null
			constraint a_pkey
			primary key,
    b_id    BIGINT references b(id)
)
;

COMMIT
;

TRUNCATE TABLE C CASCADE;
`

func main() {
	count := 5
	db := setupSql()
	var empty interface{}
	_, err := db.Query(&empty, queryCreateAndInsert)
	if err != nil {
		panic(err)
	}
	for i := 0; i < count; i++ {
		name := fmt.Sprintf("john doe #%d", i)
		_, err := db.Query(empty, "insert into c (name) values (?);", name)
		if err != nil {
			panic(err)
		}
	}
	var c []C
	_, err = db.Query(&c, "select * from c;")
	if err != nil {
		panic(err)
	}

	for i := 0; i < count; i++ {
		_, err := db.Query(empty, "insert into b (c_id) values (?);", c[i].Id)
		if err != nil {
			panic(err)
		}
	}
	var b []B
	_, err = db.Query(&b, "select id from b;")
	if err != nil {
		panic(err)
	}

	for i := 0; i < count; i++ {
		_, err := db.Query(empty, "insert into a (b_id) values (?);", b[i].Id)
		if err != nil {
			panic(err)
		}
	}

	var a []A
	_, err = db.Query(&a, querySelect)
	if err != nil {
		panic(err)
	}
	for _, temp := range a {
		fmt.Printf("(a.id, b.id, c.name): (%v, %v, %v) \n", temp.Id, temp.B.Id, temp.B.C.Name)
	}
}
