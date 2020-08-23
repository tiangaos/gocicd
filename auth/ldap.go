package auth

import (
	"fmt"
	"log"

	"github.com/go-ldap/ldap"
)

type LDAPConfig struct {
	Address      string
	BindUser     string
	BindPassword string
	SearchDN     string
}

type LDAPService struct {
	Conn   *ldap.Conn
	Config LDAPConfig
}

type LDAPUser struct {
	UID   string
	Email string
}

var ldapService *LDAPService

func init() {
	config := LDAPConfig{
		Address:      "ldap01s.daodao.com:389",
		BindUser:     "cn=manager,dc=daodao,dc=com",
		BindPassword: "daodao",
		SearchDN:     "ou=people,dc=daodao,dc=com",
	}
	var err error
	if ldapService, err = newLDAPService(config); err != nil {
		log.Fatal(err)
	}
}

func newLDAPService(config LDAPConfig) (*LDAPService, error) {
	conn, err := ldap.Dial("tcp", config.Address)
	if err != nil {
		return nil, err
	}
	err = conn.Bind(config.BindUser, config.BindPassword)
	if err != nil {
		return nil, err
	}
	return &LDAPService{Conn: conn, Config: config}, nil
}

// LDAP Login 登录
func (l *LDAPService) LoginByLDAP(uid, password string) (LDAPUser, error) {
	searchRequest := ldap.NewSearchRequest(
		l.Config.SearchDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=inetOrgPerson)(uid=%s))", uid),
		[]string{"dn", "mail"},
		nil,
	)
	sr, err := l.Conn.Search(searchRequest)
	if err != nil {
		return LDAPUser{}, err
	}
	if len(sr.Entries) != 1 {
		return LDAPUser{}, fmt.Errorf("User does not exist or too many entries returned")
	}
	email := sr.Entries[0].Attributes[0].Values[1]
	err = l.Conn.Bind(uid, password)
	if err != nil {
		return LDAPUser{}, err
	}
	return LDAPUser{UID: uid, Email: email}, nil
}
