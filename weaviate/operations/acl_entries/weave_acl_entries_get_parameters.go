package acl_entries


// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewWeaveACLEntriesGetParams creates a new WeaveACLEntriesGetParams object
// with the default values initialized.
func NewWeaveACLEntriesGetParams() WeaveACLEntriesGetParams {
	var (
		altDefault         = string("json")
		prettyPrintDefault = bool(true)
	)
	return WeaveACLEntriesGetParams{
		Alt: &altDefault,

		PrettyPrint: &prettyPrintDefault,
	}
}

// WeaveACLEntriesGetParams contains all the bound params for the weave acl entries get operation
// typically these are obtained from a http.Request
//
// swagger:parameters weave.aclEntries.get
type WeaveACLEntriesGetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Unique ACL entry ID.
	  Required: true
	  In: path
	*/
	ACLEntryID string
	/*Data format for the response.
	  In: query
	  Default: "json"
	*/
	Alt *string
	/*ID of the device to use.
	  Required: true
	  In: path
	*/
	DeviceID string
	/*Selector specifying which fields to include in a partial response.
	  In: query
	*/
	Fields *string
	/*Specifies the language code that should be used for text values in the API response.
	  In: query
	*/
	Hl *string
	/*API key. Your API key identifies your project and provides you with API access, quota, and reports. Required unless you provide an OAuth 2.0 token.
	  In: query
	*/
	Key *string
	/*OAuth 2.0 token for the current user.
	  In: query
	*/
	OauthToken *string
	/*Returns response with indentations and line breaks.
	  In: query
	  Default: true
	*/
	PrettyPrint *bool
	/*Available to use for quota purposes for server-side applications. Can be any arbitrary string assigned to a user, but should not exceed 40 characters. Overrides userIp if both are provided.
	  In: query
	*/
	QuotaUser *string
	/*IP address of the site where the request originates. Use this if you want to enforce per-user limits.
	  In: query
	*/
	UserIP *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *WeaveACLEntriesGetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rACLEntryID, rhkACLEntryID, _ := route.Params.GetOK("aclEntryId")
	if err := o.bindACLEntryID(rACLEntryID, rhkACLEntryID, route.Formats); err != nil {
		res = append(res, err)
	}

	qAlt, qhkAlt, _ := qs.GetOK("alt")
	if err := o.bindAlt(qAlt, qhkAlt, route.Formats); err != nil {
		res = append(res, err)
	}

	rDeviceID, rhkDeviceID, _ := route.Params.GetOK("deviceId")
	if err := o.bindDeviceID(rDeviceID, rhkDeviceID, route.Formats); err != nil {
		res = append(res, err)
	}

	qFields, qhkFields, _ := qs.GetOK("fields")
	if err := o.bindFields(qFields, qhkFields, route.Formats); err != nil {
		res = append(res, err)
	}

	qHl, qhkHl, _ := qs.GetOK("hl")
	if err := o.bindHl(qHl, qhkHl, route.Formats); err != nil {
		res = append(res, err)
	}

	qKey, qhkKey, _ := qs.GetOK("key")
	if err := o.bindKey(qKey, qhkKey, route.Formats); err != nil {
		res = append(res, err)
	}

	qOauthToken, qhkOauthToken, _ := qs.GetOK("oauth_token")
	if err := o.bindOauthToken(qOauthToken, qhkOauthToken, route.Formats); err != nil {
		res = append(res, err)
	}

	qPrettyPrint, qhkPrettyPrint, _ := qs.GetOK("prettyPrint")
	if err := o.bindPrettyPrint(qPrettyPrint, qhkPrettyPrint, route.Formats); err != nil {
		res = append(res, err)
	}

	qQuotaUser, qhkQuotaUser, _ := qs.GetOK("quotaUser")
	if err := o.bindQuotaUser(qQuotaUser, qhkQuotaUser, route.Formats); err != nil {
		res = append(res, err)
	}

	qUserIP, qhkUserIP, _ := qs.GetOK("userIp")
	if err := o.bindUserIP(qUserIP, qhkUserIP, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WeaveACLEntriesGetParams) bindACLEntryID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ACLEntryID = raw

	return nil
}

func (o *WeaveACLEntriesGetParams) bindAlt(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		var altDefault string = string("json")
		o.Alt = &altDefault
		return nil
	}

	o.Alt = &raw

	if err := o.validateAlt(formats); err != nil {
		return err
	}

	return nil
}

func (o *WeaveACLEntriesGetParams) validateAlt(formats strfmt.Registry) error {

	if err := validate.Enum("alt", "query", *o.Alt, []interface{}{"json"}); err != nil {
		return err
	}

	return nil
}

func (o *WeaveACLEntriesGetParams) bindDeviceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.DeviceID = raw

	return nil
}

func (o *WeaveACLEntriesGetParams) bindFields(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Fields = &raw

	return nil
}

func (o *WeaveACLEntriesGetParams) bindHl(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Hl = &raw

	return nil
}

func (o *WeaveACLEntriesGetParams) bindKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Key = &raw

	return nil
}

func (o *WeaveACLEntriesGetParams) bindOauthToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.OauthToken = &raw

	return nil
}

func (o *WeaveACLEntriesGetParams) bindPrettyPrint(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		var prettyPrintDefault bool = bool(true)
		o.PrettyPrint = &prettyPrintDefault
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("prettyPrint", "query", "bool", raw)
	}
	o.PrettyPrint = &value

	return nil
}

func (o *WeaveACLEntriesGetParams) bindQuotaUser(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.QuotaUser = &raw

	return nil
}

func (o *WeaveACLEntriesGetParams) bindUserIP(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.UserIP = &raw

	return nil
}
