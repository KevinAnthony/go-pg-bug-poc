package main

type A struct {
	id  int64 `sql:"id"`
	bId int64 `sql:"b_id"`
	B   *B
}

type B struct {
	id  int64 `sql:"id"`
	cId int64 `sql:"c_id"`
	C   *C
}

type C struct {
	id   int64  `sql:"id"`
	Name string `sql:"name"`
}

var querySelect = `SELECT
                 a.id,
                 a.b_id
                 b.id          AS a__id
                 c.c_id       AS a__c_id
                 c.id          AS a__b__id
                 c.name    AS a__b__name
        FROM a
        LEFT JOIN b ON a.b_id = b.id
        LEFT JOIN c ON b.c_id = c.id
        WHERE ...
`

var queryCreateAndInsert = `
BEGIN
;

CREATE TABLE IF NOT EXISTS c
(
    name    varchar(64),
    id      BIGINT NOT NULL
)
;

CREATE TABLE IF NOT EXISTS b
(
    id      BIGINT NOT NULL,
    c_id    BIGINT references c(id)
)
;

CREATE TABLE IF NOT EXISTS a
(
    id      BIGINT NOT NULL,
    b_id    BIGINT references b(id)
)
;

COMMIT
;
`

func main() {
	db := setupSql()
	var empty interface{}
	_, err := db.Query(&empty, queryCreateAndInsert)
	if err != nil {
		panic(err)
	}
}
