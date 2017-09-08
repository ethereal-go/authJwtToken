package authJwtToken

/**
/ map locale middleware auth jwt token in documentation graphQL
*/

func GetLocale() map[string]map[string]string {
	return map[string]map[string]string{
		"en-US": map[string]string{
			"graphQL.JwtType.CreateJWTToken.Description": "Create new jwt-token",
			"graphQL.JwtType.Token.Description":          "Auth JWT Token",
			"graphQL.JwtType.Description" : "The type field holds the token and is designed for local authentication",
		},
		"ru-RU": map[string]string{
			"graphQL.JwtType.CreateJWTToken.Description": "Создать новый jwt-token",
			"graphQL.JwtType.Token.Description":          "Токен",
			"graphQL.JwtType.Description" : "Тип хранит в себе поле токен и создан для локальной аутентификации",
		},
	}
}
