package authJwtToken

/**
/ map locale middleware auth jwt token in documentation graphQL
*/

func GetLocale() map[string]map[string]string {
	return map[string]map[string]string{
		"en-US": map[string]string{
			"graphQL.JwtType.CreateJWTToken.Description": "Create new jwt-token",
			"graphQL.JwtType.Token.Description":          "Auth JWT Token",
		},
		"ru-RU": map[string]string{
			"graphQL.JwtType.CreateJWTToken.Description": "Создать новый jwt-token",
			"graphQL.JwtType.Token.Description":          "Токен",
		},
	}
}
