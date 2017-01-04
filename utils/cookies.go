package utils

import(
	"time"
	"net/http"
)

const(
	sessionLenght = 24 * 3 * time.Hour //Tenemos las propiedades de horas, pero no propiedades más grandes 
	sessionCookieName = "GopherSession"
	sessionIdLenght = 20
)

type Session struct{
	Id string
	UserId int
	Expiry time.Time
}

func GenerateID(key string, lenght int) string{
	return "HolaMundo!"
}

func NewSession(w http.ResponseWriter) *Session{
	expiry := time.Now().Add(sessionIdLenght)
	
	session := &Session{
		Id: GenerateID("sess", sessionIdLenght),
		Expiry: expiry,
	}
	/*
	cookie := http.Cookie{
		Name : sessionCookieName,
		Value : Session.Id,
		Expiries: expiry,
	}

	http.SetCookie(w, &cookie)
	*/

	return session
}























