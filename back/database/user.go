package database


type SherpaUser struct {
    uid                 int
}


func AddUser(user SherpaUser) error {
    db := dbInstance.db
    q := `INSERT INTO SherpaUser
        (uid)
        VALUES ($1)
        ON CONFLICT (uid) DO NOTHING`

    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    _, err = tx.Exec(q, user.uid)
    if err != nil {
        return err
    }

    err = tx.Commit()
    return err
}


func init() {
    migrateSherpaUser()
}

func migrateSherpaUser() {
    New()
    db := dbInstance.db
    q := `CREATE TABLE IF NOT EXISTS SherpaUser (
        uid             INT PRIMARY KEY
    )`
    _, err := db.Exec(q)
    if err != nil {
        panic(err)
    }
}
