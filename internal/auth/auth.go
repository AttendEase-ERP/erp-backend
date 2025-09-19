package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/clerk/clerk-sdk-go/v2/user"
)

func InitClerk(secret string) {
	clerk.SetKey(secret)
}

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler := clerkhttp.WithHeaderAuthorization()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, ok := clerk.SessionClaimsFromContext(r.Context())
			if !ok {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		}))

		handler.ServeHTTP(w, r)
	})
}

// Extract session claims
func GetSessionClaimsFromContext(ctx context.Context) (*clerk.SessionClaims, error) {
	claims, ok := clerk.SessionClaimsFromContext(ctx)
	if !ok {
		return nil, errors.New("no session claims in context")
	}

	return claims, nil
}

// Extract user email from session claim
func GetUserEmail(ctx context.Context) (string, error) {
	claims, err := GetSessionClaimsFromContext(ctx)
	if err != nil {
		return "", err
	}

	usr, err := user.Get(ctx, claims.Subject)
	if err != nil {
		return "", errors.New("no user found")
	}

	return usr.EmailAddresses[0].EmailAddress, nil
}
