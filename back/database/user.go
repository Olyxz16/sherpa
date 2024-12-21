package database


type SherpaUser struct {
    uid                 int
}


func init() {
    migrateSherpaUser()
}

func migrateSherpaUser() {
    New()
    db := dbInstance.db
    q := `CREATE TABLE IF NOT EXISTS SherpaUser (
        uid             SERIAL PRIMARY KEY
    )`
    _, err := db.Exec(q)
    if err != nil {
        panic(err)
    }
}
