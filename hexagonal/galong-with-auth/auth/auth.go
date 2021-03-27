package authentication

import (
	"backend-service/models"
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"regexp"
	"strconv"
)

type Identity struct {
	Id          int
	Roles []*models.Role
	Permissions []*models.Permission
}

func (ac *AuthContext) newIdentity(t *jwt.Token) (*Identity, error) {
	claims := t.Claims.(jwt.MapClaims)
	id, err := strconv.Atoi(claims["jti"].(string))
	if err != nil {
		return nil, err
	}
	user, err := models.Accounts(qm.Where("Id=?", id), qm.Load("RoleUsers.Role"), qm.Load("RoleUsers.Role.Permissions")).One(context.TODO(), ac.PsqlDb)
	if err != nil {
		return nil, err
	}
	var permissions []*models.Permission
	var roles []*models.Role
	for _, ru := range user.R.UserRoleUsers {
		if ru.R.Role.IsActive {
			roles = append(roles, ru.R.Role)
			permissions = append(permissions, ru.R.Role.R.Permissions...)
		}
	}
	return &Identity{Id: id, Permissions: permissions, Roles: roles}, nil
}

func (i *Identity) HasAccessTo(path string) bool {
	pathBytes := []byte(path)
	for _, perm := range i.Permissions {
		reg, err := regexp.Compile(perm.Route.String)
		if err != nil {
			return false
		}
		if reg.Match(pathBytes) {
			return true
		}
	}
	return false
}


func (ac *AuthContext) NewJWTToken(claims jwt.Claims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := ac.VConfig.GetString("auth.jwt.secret")
	return t.SignedString([]byte(secret))
}

func (ac *AuthContext) validateToken(token string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(ac.VConfig.GetString("auth.jwt.secret")), nil
	})
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, errors.New("invalid jwt token")
	}
	return t, nil
}
