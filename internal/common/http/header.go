package common

type Header string

const (
	// https://datatracker.ietf.org/doc/html/rfc7540#section-8.1.2
	// Just as in HTTP/1.x, header field names are strings of ASCII
	// characters that are compared in a case-insensitive fashion. However,
	// header field names MUST be converted to lowercase prior to their
	// encoding in HTTP/2.  A request or response containing uppercase
	// header field names MUST be treated as malformed (Section 8.1.2.6).

	// HTTP Header Standard
	HeaderRequestID              Header = `x-request-id`
	HeaderAPIKey                 Header = `x-api-key`
	HeaderContentType            Header = `content-type`
	HeaderAccept                 Header = `accept`
	HeaderAcceptLanguage         Header = `accept-language`
	HeaderAcceptEncoding         Header = `accept-encoding`
	HeaderAcceptCharset          Header = `accept-charset`
	HeaderAcceptRanges           Header = `accept-ranges`
	HeaderAcceptDatetime         Header = `accept-datetime`
	HeaderAuthorization          Header = `authorization`
	HeaderUserAgent              Header = `user-agent`
	HeaderXForwardedFor          Header = `x-forwarded-for`
	HeaderXForwardedHost         Header = `x-forwarded-host`
	HeaderXForwardedProto        Header = `x-forwarded-proto`
	HeaderXForwardedServer       Header = `x-forwarded-server`
	HeaderXForwardedServerHeader Header = `x-forwarded-server-header`
	HeaderXForwardedServerPort   Header = `x-forwarded-server-port`

	// Custom HTTP Header
	HeaderAppLang       Header = `x-app-lang`
	HeaderAppDebug      Header = `x-app-debug`
	HeaderCreationSteps Header = `x-create-step`

	// Lang Header
	HeaderLangEN Header = `en`
	HeaderLangID Header = `id` // #nosec G101 -- This is a header name, not a credential
)

func (h Header) String() string {
	return string(h)
}
