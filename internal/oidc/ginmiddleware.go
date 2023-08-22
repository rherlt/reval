package oidc

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zitadel/oidc/v2/pkg/client/rp"
	"github.com/zitadel/oidc/v2/pkg/oidc"
)

var relyingParty rp.RelyingParty
var userInfoCache map[string]*oidc.UserInfo

const OidcUserClaimsKey = "OidcUserClaims"
const OidcUserIsAuthenticatedKey = "OidcUserIsAuthenticated"

func OidcAuthMiddleware(authority, audience string) gin.HandlerFunc {
	userInfoCache = map[string]*oidc.UserInfo{}

	_, err := rp.Discover(authority, http.DefaultClient)
	if err != nil {
		fmt.Printf("error discovering server %s:  %s", authority, err.Error())
	}

	relyingParty, err = rp.NewRelyingPartyOIDC(authority, audience, "", "", []string{})
	if err != nil {
		fmt.Printf("error new relying partiy oidc %s:  %s", authority, err.Error())
	}

	return func(c *gin.Context) {

		c.Set(OidcUserIsAuthenticatedKey, false)
		c.Set(OidcUserClaimsKey, new(map[string]any))

		ok, token := checkTokenType(c)
		if !ok {
			c.Set(OidcUserIsAuthenticatedKey, false)
			c.Set(OidcUserClaimsKey, new(map[string]any))
			return
		}

		var claims *oidc.AccessTokenClaims
		payload, err := oidc.ParseToken(token, &claims)
		if err != nil {
			c.Set(OidcUserIsAuthenticatedKey, false)
			c.Set(OidcUserClaimsKey, new(map[string]any))
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		err = verifyAccessToken(c, token, payload, claims, relyingParty)
		if err != nil {

			//remove token from cache, if present
			delete(userInfoCache, token)

			c.Set(OidcUserIsAuthenticatedKey, false)
			c.Set(OidcUserClaimsKey, new(map[string]any))
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		userInfo := userInfoCache[token]
		if userInfo == nil {

			//get userInfo from authority userInfo endpoint
			userInfo, err = rp.Userinfo(token, "bearer", claims.Subject, relyingParty)
			if err != nil {
				c.Set(OidcUserIsAuthenticatedKey, false)
				c.Set(OidcUserClaimsKey, new(map[string]any))
				c.AbortWithError(http.StatusUnauthorized, err)
				return
			}
			//if token was validated successfully, add it to cache
			userInfoCache[token] = userInfo
		}

		c.Set(OidcUserIsAuthenticatedKey, true)
		c.Set(OidcUserClaimsKey, userInfo.Claims)
	}
}

func checkTokenType(c *gin.Context) (bool, string) {
	auth := c.Request.Header.Get("authorization")
	if auth == "" {
		//http.Error(w, "auth header missing", http.StatusUnauthorized)
		c.AbortWithStatus(http.StatusUnauthorized)

		return false, ""
	}
	if !strings.HasPrefix(auth, oidc.PrefixBearer) {
		//http.Error(w, "invalid header", http.StatusUnauthorized)
		c.AbortWithStatus(http.StatusUnauthorized)

		return false, ""
	}
	return true, strings.TrimPrefix(auth, oidc.PrefixBearer)
}

func verifyAccessToken(c *gin.Context, token string, payload []byte, claims *oidc.AccessTokenClaims, rp rp.RelyingParty) error {

	if err := oidc.CheckSignature(c, token, payload, claims, rp.IDTokenVerifier().SupportedSignAlgs(), rp.IDTokenVerifier().KeySet()); err != nil {
		return err
	}

	if err := oidc.CheckSubject(claims); err != nil {
		return err
	}

	if err := oidc.CheckIssuer(claims, rp.Issuer()); err != nil {
		return err
	}

	if err := oidc.CheckAudience(claims, rp.OAuthConfig().ClientID); err != nil {
		return err
	}

	if err := oidc.CheckExpiration(claims, rp.IDTokenVerifier().Offset()); err != nil {
		return err
	}

	if err := oidc.CheckIssuedAt(claims, rp.IDTokenVerifier().MaxAgeIAT(), rp.IDTokenVerifier().Offset()); err != nil {
		return err
	}

	if err := oidc.CheckAuthTime(claims, rp.IDTokenVerifier().MaxAge()); err != nil {
		return err
	}

	return nil
}
