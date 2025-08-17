package user

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	accessTokenCookieName  = "access_token"
	refreshTokenCookieName = "refresh_token"
)

func (h *Handler) setTokenCookies(c *gin.Context, accessToken, refreshToken string) {
	accessCookie := &http.Cookie{
		Name:     accessTokenCookieName,
		Value:    accessToken,
		Expires:  time.Now().Add(15 * time.Minute),
		Path:     "/api",
		Domain:   "", // automatically set based on domain
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode, // option: LaxMode or strict
	}

	refreshCookie := &http.Cookie{
		Name:     refreshTokenCookieName,
		Value:    refreshToken,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		Path:     "/api/v1/auth",
		Domain:   "",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(c.Writer, accessCookie)
	http.SetCookie(c.Writer, refreshCookie)
}

func (h *Handler) clearTokenCookies(c *gin.Context) {
	accessCookie := &http.Cookie{
		Name:     accessTokenCookieName,
		Value:    "",
		Expires:  time.Unix(0, 0),
		Path:     "/api",
		Domain:   "",
		Secure:   true,
		HttpOnly: true,
	}

	refreshCookie := &http.Cookie{
		Name:     refreshTokenCookieName,
		Value:    "",
		Expires:  time.Unix(0, 0),
		Path:     "/api/v1/auth",
		Domain:   "",
		Secure:   true,
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, accessCookie)
	http.SetCookie(c.Writer, refreshCookie)
}
