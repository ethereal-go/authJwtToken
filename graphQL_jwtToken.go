package authJwtToken

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/ethereal-go/base/root/database"
	"github.com/ethereal-go/ethereal"
	"github.com/ethereal-go/ethereal/root/config"
	"github.com/ethereal-go/ethereal/utils"
	"github.com/graphql-go/graphql"
	"net/http"
)

// set locale database
const (
	errorInputData = "Login or Password not valid"
)

// get type locale from configuration..
var locale = config.GetCnf("L18N.LOCALE").(string)

var jwtType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "JWTToken",
	Description: string(ethereal.ConstructorI18N().T(locale, "graphQL.JwtType.Description")),
	Fields: graphql.Fields{
		"token": &graphql.Field{
			Type:        graphql.String,
			Description: string(ethereal.ConstructorI18N().T(locale, "graphQL.JwtType.Token.Description")),
		},
	},
})

/**
/ Create Token
*/
var CreateJWTToken = graphql.Field{
	Type:        jwtType,
	Description: string(ethereal.ConstructorI18N().T(locale, "graphQL.JwtType.CreateJWTToken.Description")),
	Args: graphql.FieldConfigArgument{
		"login": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		var user database.User
		var generateToken string
		login, _ := params.Args["login"].(string)
		password, _ := params.Args["password"].(string)

		ethereal.App.Db.Where("email = ?", login).First(&user)

		if utils.CompareHashPassword([]byte(user.Password), []byte(password)) {
			claims := EtherealClaims{
				jwt.StandardClaims{
					ExpiresAt: 1,
					Issuer:    user.Email,
				},
			}
			// TODO add choose crypt via configuration!
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			generateToken, _ = token.SignedString(JWTKEY())

		} else {
			return nil, errors.New(errorInputData)
		}

		return struct {
			Token string `json:"token"`
		}{generateToken}, nil
	},
}

var AuthToken, _ = graphql.NewSchema(graphql.SchemaConfig{
	Mutation: authMutation,
})

var authMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "AuthMutation",
	Fields: graphql.Fields{
		"createJWTToken": &CreateJWTToken,
	},
})

// handler for get token if config in = "global" mode
func RegisterHandlerAuthCreateToken() {
	if config.GetCnf("AUTH.JWT_TOKEN").(string) == "global" {
		http.HandleFunc("/auth0/login", func(w http.ResponseWriter, r *http.Request) {
			var user database.User

			login := r.FormValue("login")
			password := r.FormValue("password")

			ethereal.App.Db.Where("email = ?", login).First(&user)

			if user.ID != 0 && utils.CompareHashPassword([]byte(user.Password), []byte(password)) {
				claims := EtherealClaims{
					jwt.StandardClaims{
						ExpiresAt: 1,
						Issuer:    user.Email,
					},
				}
				// TODO add choose crypt via configuration!
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

				generateToken, _ := token.SignedString(JWTKEY())
				w.Write([]byte(generateToken))
			} else {
				w.Write([]byte("Not found user or password not true."))
			}
		})
	}
}
