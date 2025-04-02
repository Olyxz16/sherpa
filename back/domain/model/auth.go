package model

type AuthSource string;
const (
	Github AuthSource = "github.com"
)

type Auth struct {
	uid					int
	user				*User
	source				AuthSource
    access_token        string
    refresh_token       string
    expires_in          int
    refresh_expires_in  int
}

func NewAuth(uid int, user *User, source AuthSource, access_token, refresh_token string, expires_in, refresh_expires_in int) *Auth {
	return &Auth{
		uid: uid,
		user: user,
		source: source,
		access_token: access_token,
		refresh_token: refresh_token,
		expires_in: expires_in,
		refresh_expires_in: refresh_expires_in,
	}
}

func (a *Auth) GetAuthID() int {
	return a.uid
}

func (a *Auth) GetUser() *User {
	return a.user
}

func (a *Auth) GetPlatormSource() AuthSource {
	return a.source
}

func (a *Auth) GetAccessToken() string {
	return a.access_token
}

func (a *Auth) GetRefreshToken() string {
	return a.refresh_token
}

func (a *Auth) GetExpiresIn() int {
	return a.expires_in
}

func (a *Auth) GetRefreshTokenExpireIn() int {
	return a.refresh_expires_in
}

func(a *Auth) Refresh(access_token, refresh_token string, expires_in, refresh_expires_in int) {
	a.access_token = access_token;
	a.refresh_token = refresh_token;
	a.expires_in = expires_in;
	a.refresh_expires_in = refresh_expires_in;
}
