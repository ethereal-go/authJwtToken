# authJwtToken
 Extension authentication jwt token
 
 ## Example use 
 
 ```
 
import (
	"github.com/agoalofalife/ethereal"
	"github.com/ethereal-go/authJwtToken"
)

func main() {

	ethereal.ConstructorMiddleware().AddMiddleware(authJwtToken.GetMiddlewareJwtToken())
	authJwtToken.RegisterHandlerAuthCreateToken()
	ethereal.Mutations().Add("createToken", &authJwtToken.CreateJWTToken)
	ethereal.I18nGraphQL().Merge(authJwtToken.GetLocale()).Fill()
	ethereal.Start()
}

 ```
