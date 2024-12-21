package database

type GithubAuth struct {
    UserId              int
    Cookie              string
    Access_token        string
    Refresh_token       string
    Expires_in          float64
    Refresh_expires_in  float64
}



func AddGithubUser(auth GithubAuth) error {
    return nil 
}




func init() {
    migrateGithubAuth()
}

func migrateGithubAuth() {
    New()
    db := dbInstance.db
    
    allowed, err := isUserMigrated()
    if err != nil {
        panic(err)
    }

    if !allowed {
        migrateSherpaUser()
    }

    q := `CREATE TABLE IF NOT EXISTS GithubAuth (
        cookie          VARCHAR(255),
        access_token    VARCHAR(255),
        expires_in      FLOAT,
        refresh_token   VARCHAR(255),
        rt_expires_in   FLOAT,
        userId          INT,
        CONSTRAINT fk_userId FOREIGN KEY (userId)
        REFERENCES SherpaUser(uid)
    )`
    _, err = db.Exec(q)
    if err != nil {
        panic(err)
    }
}
func isUserMigrated() (bool, error) {
    db := dbInstance.db
    q := `SELECT EXISTS (
        SELECT FROM 
            pg_tables
        WHERE 
            schemaname = 'public' AND 
            tablename  = 'sherpauser'
    );`
    rows, err := db.Query(q)
    if err != nil {
        return false, err
    }
    defer rows.Close()

    var result bool
    rows.Next()
    rows.Scan(&result)
    
    return result, nil
}
