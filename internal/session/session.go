package session

import (
	"net/http"
	"sync"

	"github.com/K-H-Tech/auto-tax-gov/internal/models"
)

// Session holds the state for a user authentication session.
// It is thread-safe and can be shared across goroutines.
type Session struct {
	cookies        []*http.Cookie
	captcha        *models.CaptchaData
	csrfToken      string
	redirectURL    string
	registrationID string
	authenticated  bool
	mu             sync.RWMutex
}

// New creates a new empty session.
func New() *Session {
	return &Session{
		cookies: make([]*http.Cookie, 0),
	}
}

// GetCookies returns a copy of the session cookies.
func (s *Session) GetCookies() []*http.Cookie {
	s.mu.RLock()
	defer s.mu.RUnlock()
	cookies := make([]*http.Cookie, len(s.cookies))
	copy(cookies, s.cookies)
	return cookies
}

// SetCookies replaces the session cookies.
func (s *Session) SetCookies(cookies []*http.Cookie) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cookies = cookies
}

// AddCookies appends cookies to the session.
func (s *Session) AddCookies(cookies []*http.Cookie) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cookies = append(s.cookies, cookies...)
}

// MergeCookies merges new cookies with existing ones, updating duplicates by name.
func (s *Session) MergeCookies(newCookies []*http.Cookie) {
	s.mu.Lock()
	defer s.mu.Unlock()

	cookieMap := make(map[string]*http.Cookie)
	for _, c := range s.cookies {
		cookieMap[c.Name] = c
	}
	for _, c := range newCookies {
		cookieMap[c.Name] = c
	}

	s.cookies = make([]*http.Cookie, 0, len(cookieMap))
	for _, c := range cookieMap {
		s.cookies = append(s.cookies, c)
	}
}

// GetCaptcha returns the current captcha data.
func (s *Session) GetCaptcha() *models.CaptchaData {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.captcha
}

// SetCaptcha sets the current captcha data and updates CSRF token.
func (s *Session) SetCaptcha(captcha *models.CaptchaData) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.captcha = captcha
	if captcha != nil {
		s.csrfToken = captcha.CSRFToken
	}
}

// GetCSRFToken returns the CSRF token.
func (s *Session) GetCSRFToken() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.csrfToken
}

// SetCSRFToken sets the CSRF token.
func (s *Session) SetCSRFToken(token string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.csrfToken = token
}

// GetRedirectURL returns the redirect URL after authentication.
func (s *Session) GetRedirectURL() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.redirectURL
}

// SetRedirectURL sets the redirect URL after authentication.
func (s *Session) SetRedirectURL(url string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.redirectURL = url
}

// GetRegistrationID returns the tax registration ID.
func (s *Session) GetRegistrationID() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.registrationID
}

// SetRegistrationID sets the tax registration ID.
func (s *Session) SetRegistrationID(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.registrationID = id
}

// IsAuthenticated returns true if the session is authenticated.
func (s *Session) IsAuthenticated() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.authenticated
}

// SetAuthenticated sets the authentication status.
func (s *Session) SetAuthenticated(status bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.authenticated = status
}

// IsActive returns true if the session has cookies.
func (s *Session) IsActive() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.cookies) > 0
}

// Reset clears all session state.
func (s *Session) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cookies = make([]*http.Cookie, 0)
	s.captcha = nil
	s.csrfToken = ""
	s.redirectURL = ""
	s.registrationID = ""
	s.authenticated = false
}

// Clone creates a copy of the session.
func (s *Session) Clone() *Session {
	s.mu.RLock()
	defer s.mu.RUnlock()

	clone := &Session{
		cookies:        make([]*http.Cookie, len(s.cookies)),
		csrfToken:      s.csrfToken,
		redirectURL:    s.redirectURL,
		registrationID: s.registrationID,
		authenticated:  s.authenticated,
	}
	copy(clone.cookies, s.cookies)
	if s.captcha != nil {
		clone.captcha = &models.CaptchaData{
			Key:        s.captcha.Key,
			ImageData:  s.captcha.ImageData,
			CSRFToken:  s.captcha.CSRFToken,
			AuthHeader: s.captcha.AuthHeader,
		}
	}
	return clone
}
