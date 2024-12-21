package database


type GithubAuth struct {
    UserId              int
    Cookie              string
    Access_token        string
    Refresh_token       string
    Expires_in          int
    Refresh_expires_in  int
}


func AddGithubUser(auth GithubAuth) error {
    db := dbInstance.db
    
    user := SherpaUser{uid: auth.UserId}
    err := AddUser(user)
    if err != nil {
        return err
    }

    q := `INSERT INTO GithubAuth
        (cookie, access_token, expires_in, refresh_token, rt_expires_in, userId)
        VALUES ($1, $2, $3, $4, $5, $6)
    `
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    _, err = tx.Exec(q,
        auth.Cookie, 
        auth.Access_token, 
        auth.Expires_in, 
        auth.Refresh_token,
        auth.Expires_in,
        auth.UserId)
    if err != nil {
        return err
    }

    err = tx.Commit()
    return err
}

func TokenFromCookie(cookie string) (string, error) {
    db := dbInstance.db
    q := `SELECT access_token FROM GithubAuth
        WHERE cookie=$1`

    rows, err := db.Query(q, cookie)
    if err != nil {
        return "", err
    }
    defer rows.Close()

    var access_token string
    rows.Next()
    err = rows.Scan(&access_token)
    if err != nil {
        return "", err
    }

    return access_token, nil
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
