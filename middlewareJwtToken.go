package authJwtToken

import (
	"encoding/json"
	"github.com/ethereal-go/ethereal"
	"github.com/ethereal-go/ethereal/root/app"
	"github.com/ethereal-go/ethereal/root/config"
	"github.com/justinas/alice"
	"net/http"
)

type MiddlewareJWTToken struct {
	EtherealClaims
	StatusError    int
	ResponseError  string
	authenticated  bool
	responseWriter http.ResponseWriter
	included       bool // flag is enabled or disabled authJwtToken
}

func (m MiddlewareJWTToken) Add(where *[]alice.Constructor, application *app.Application) {
	confToken := config.GetCnf("AUTH.JWT_TOKEN").(string)

	if confToken == "local" {
		m.included = true
		*where = append(*where, func(handler http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				m.responseWriter = w
				if check, err := m.Verify(r); !check {
					m.ResponseError = handlerErrorToken(err).Error()
				} else {
					m.authenticated = true
				}
				ethereal.CtxStruct(application, m)
				handler.ServeHTTP(w, r)
			})
		})
	} else if confToken == "global" {
		// check jwt token all queries..
		*where = append(*where, func(handler http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				if check, err := m.Verify(r); !check {
					json.NewEncoder(w).Encode(handlerErrorToken(err).Error())
					w.WriteHeader(m.StatusError)
				}
			})
		})
	}
}

func GetMiddlewareJwtToken() MiddlewareJWTToken {
	return MiddlewareJWTToken{
		ResponseError: http.StatusText(http.StatusNetworkAuthenticationRequired),
		StatusError:   http.StatusNetworkAuthenticationRequired,
	}
}
