package parse

type SessionToken struct {
	token      string
	objectId   string
	restricted bool
	createdAt  string
	updatedAt  string
	expiresAt  string
}

var currentSession SessionToken

func Session(token string) SessionToken {
	sessionToken = SessionToken{token: token}

	// TODO:
	// fill out other attributes of SessionToken by
	// sending a GET request to "/1/sessions/me"

	return sessionToken
}

func (self *SessionToken) hasToken() bool {
	return len(self.token) > 0
}
