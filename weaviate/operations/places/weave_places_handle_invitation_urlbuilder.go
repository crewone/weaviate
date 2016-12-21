package places


// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	"strings"

	"github.com/go-openapi/swag"
)

// WeavePlacesHandleInvitationURL generates an URL for the weave places handle invitation operation
type WeavePlacesHandleInvitationURL struct {
	PlaceID string

	Action      string
	Alt         *string
	Fields      *string
	Hl          *string
	Key         *string
	Member      string
	OauthToken  *string
	PrettyPrint *bool
	QuotaUser   *string
	UserIP      *string

	// avoid unkeyed usage
	_ struct{}
}

// Build a url path and query string
func (o *WeavePlacesHandleInvitationURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/places/{placeId}/handleInvitation"

	placeID := o.PlaceID
	if placeID != "" {
		_path = strings.Replace(_path, "{placeId}", placeID, -1)
	} else {
		return nil, errors.New("PlaceID is required on WeavePlacesHandleInvitationURL")
	}
	result.Path = _path

	qs := make(url.Values)

	action := o.Action
	if action != "" {
		qs.Set("action", action)
	}

	var alt string
	if o.Alt != nil {
		alt = *o.Alt
	}
	if alt != "" {
		qs.Set("alt", alt)
	}

	var fields string
	if o.Fields != nil {
		fields = *o.Fields
	}
	if fields != "" {
		qs.Set("fields", fields)
	}

	var hl string
	if o.Hl != nil {
		hl = *o.Hl
	}
	if hl != "" {
		qs.Set("hl", hl)
	}

	var key string
	if o.Key != nil {
		key = *o.Key
	}
	if key != "" {
		qs.Set("key", key)
	}

	member := o.Member
	if member != "" {
		qs.Set("member", member)
	}

	var oauthToken string
	if o.OauthToken != nil {
		oauthToken = *o.OauthToken
	}
	if oauthToken != "" {
		qs.Set("oauth_token", oauthToken)
	}

	var prettyPrint string
	if o.PrettyPrint != nil {
		prettyPrint = swag.FormatBool(*o.PrettyPrint)
	}
	if prettyPrint != "" {
		qs.Set("prettyPrint", prettyPrint)
	}

	var quotaUser string
	if o.QuotaUser != nil {
		quotaUser = *o.QuotaUser
	}
	if quotaUser != "" {
		qs.Set("quotaUser", quotaUser)
	}

	var userIP string
	if o.UserIP != nil {
		userIP = *o.UserIP
	}
	if userIP != "" {
		qs.Set("userIp", userIP)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *WeavePlacesHandleInvitationURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *WeavePlacesHandleInvitationURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *WeavePlacesHandleInvitationURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on WeavePlacesHandleInvitationURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on WeavePlacesHandleInvitationURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *WeavePlacesHandleInvitationURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
