package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/duo-labs/webauthn.io/session"
	"github.com/duo-labs/webauthn/protocol"
	"github.com/duo-labs/webauthn/webauthn"
	"github.com/patrickmn/go-cache"
	rtv1 "github.com/sonr-io/sonr/internal/blockchain/x/registry/types"
	"github.com/sonr-io/sonr/pkg/config"
)

const (
	REGISTRATION_SESSION_KEY   = "registration_session"
	AUTHENTICATION_SESSION_KEY = "authentication_session"
)

// WebAuthn manages the WebAuthn interface
type WebAuthn struct {
	instance *webauthn.WebAuthn

	ctx      context.Context
	cache    *cache.Cache
	config   *config.Config
	sessions *session.Store
}

// NewWebauthn creates a new WebAuthn instance with the given configuration
func NewWebauthn(ctx context.Context, config *config.Config) (*WebAuthn, error) {
	// Create a WebAuthn instance
	web, err := webauthn.New(config.WebauthnConfig())
	if err != nil {
		return nil, err
	}

	// Create a new Session Store
	sessionStore, err := session.NewStore()
	if err != nil {
		return nil, err
	}

	// return a new WebAuthn instance
	return &WebAuthn{
		instance: web,
		ctx:      ctx,
		cache:    cache.New(5*time.Minute, 10*time.Minute),
		config:   config,
		sessions: sessionStore,
	}, nil
}

// SaveRegistrationSession saves the registration session for the given user
func (wan *WebAuthn) SaveRegistrationSession(r *http.Request, w http.ResponseWriter, username string, creator string) (*protocol.CredentialCreation, error) {
	// Create Blank WhoIs
	whoIs := blankWhoIs(username, creator)
	wan.cache.Set(registerCacheKey(username), whoIs, cache.DefaultExpiration)

	// generate PublicKeyCredentialCreationOptions, session data
	options, sessionData, err := wan.instance.BeginRegistration(
		whoIs,
		registerOptions(whoIs),
		webauthn.WithAuthenticatorSelection(authSelect()),
		webauthn.WithConveyancePreference(conveyancePreference()),
	)
	if err != nil {
		return nil, err
	}

	// TODO: replace with SaveWebauthnSession below when it works
	wan.cache.Set(registerSessionCacheKey(username), *sessionData, cache.DefaultExpiration)

	// store session data as marshaled JSON
	// err = wan.sessions.SaveWebauthnSession(REGISTRATION_SESSION_KEY, *sessionData, r, w)
	// if err != nil {
	// 	return nil, err
	// }

	return options, nil
}

// FinishRegistrationSession returns the registration session for the given user
func (w *WebAuthn) FinishRegistrationSession(r *http.Request, username string) (*webauthn.Credential, error) {
	// get user
	wi, found := w.cache.Get(registerCacheKey(username))
	if !found {
		return nil, errors.New("user not found")
	}
	whois := wi.(*rtv1.WhoIs)

	// TODO: replace with GetWebauthnSession below when it works
	cacheResult, found := w.cache.Get(registerSessionCacheKey(username))
	if !found {
		return nil, errors.New("could not find session data")
	}

	sessionData, ok := cacheResult.(webauthn.SessionData)
	if !ok {
		return nil, errors.New("found unexpected data type in cache")
	}

	// sessionData, err := w.sessions.GetWebauthnSession(REGISTRATION_SESSION_KEY, r)
	// if err != nil {
	// 	return nil, err
	// }

	credential, err := w.instance.FinishRegistration(whois, sessionData, r)
	if err != nil {
		return nil, err
	}
	w.cache.Delete(registerSessionCacheKey(username))
	return credential, nil
}

// SaveAuthenticationSession saves the login session for the given user
func (wan *WebAuthn) SaveAuthenticationSession(r *http.Request, w http.ResponseWriter, whoIs *rtv1.WhoIs) (*protocol.CredentialAssertion, error) {
	// generate PublicKeyCredentialRequestOptions, session data
	wan.cache.Set(authenticationCacheKey(whoIs.Name+".snr"), whoIs, cache.DefaultExpiration)
	options, sessionData, err := wan.instance.BeginLogin(whoIs)
	if err != nil {
		return nil, err
	}

	// TODO: replace with SaveWebauthnSession below when it works
	wan.cache.Set(AUTHENTICATION_SESSION_KEY, *sessionData, cache.DefaultExpiration)

	// store session data as marshaled JSON
	// err = wan.sessions.SaveWebauthnSession("authentication", sessionData, r, w)
	// if err != nil {
	// 	return nil, err
	// }
	return options, nil
}

// FinishAuthenticationSession returns the registration session for the given user
func (w *WebAuthn) FinishAuthenticationSession(r *http.Request, username string) (*webauthn.Credential, error) {
	// get user
	x, found := w.cache.Get(authenticationCacheKey(username))
	if !found {
		return nil, errors.New("user not found")
	}
	whois := x.(*rtv1.WhoIs)

	// TODO: replace with GetWebauthnSession below when it works
	cacheResult, found := w.cache.Get(AUTHENTICATION_SESSION_KEY)
	if !found {
		return nil, errors.New("could not find session data")
	}

	sessionData, ok := cacheResult.(webauthn.SessionData)
	if !ok {
		return nil, errors.New("found unexpected data type in cache")
	}

	// sessionData, err := w.sessions.GetWebauthnSession(AUTHENTICATION_SESSION_KEY, r)
	// if err != nil {
	// 	return nil, err
	// }

	credential, err := w.instance.FinishLogin(whois, sessionData, r)
	if err != nil {
		return nil, err
	}
	w.cache.Delete(AUTHENTICATION_SESSION_KEY)
	return credential, nil
}

// ----------------
// HELPER FUNCTIONS
// ----------------

// authenticationCacheKey is a helper function to create a cache key for the given user
func authenticationCacheKey(username string) string {
	// the username needs the suffix since it's included in the webauthn ID during fetch
	uname := username
	if !strings.HasSuffix(username, ".snr") {
		uname = username + ".snr"
	}
	return fmt.Sprintf("%s_authentication", uname)
}

// authSelect is a helper function to create a AuthenticatorSelection
func authSelect() protocol.AuthenticatorSelection {
	return protocol.AuthenticatorSelection{
		AuthenticatorAttachment: protocol.AuthenticatorAttachment("platform"),
		RequireResidentKey:      protocol.ResidentKeyUnrequired(),
		UserVerification:        protocol.VerificationRequired,
	}
}

// blankWhoIs is a helper function to create a blank WhoIs
func blankWhoIs(username, creator string) *rtv1.WhoIs {
	return &rtv1.WhoIs{
		Name:        username,
		Did:         "",
		Document:    nil,
		Creator:     creator,
		Credentials: make([]*rtv1.Credential, 0),
	}
}

// conveyancePreference is a helper function to create a ConveyancePreference
func conveyancePreference() protocol.ConveyancePreference {
	return protocol.ConveyancePreference(protocol.PreferNoAttestation)
}

// registerCacheKey is a helper function to create a cache key for a registration session
func registerCacheKey(username string) string {
	return fmt.Sprintf("%s_registration", username)
}

func registerSessionCacheKey(username string) string {
	return fmt.Sprintf("%s_sessiondata", username)
}

// registerOptions is a helper function to create a PublicKeyCredentialCreationOptions
func registerOptions(whois *rtv1.WhoIs) func(credCreationOpts *protocol.PublicKeyCredentialCreationOptions) {
	return func(credCreationOpts *protocol.PublicKeyCredentialCreationOptions) {
		credCreationOpts.CredentialExcludeList = whois.CredentialExcludeList()
	}
}
