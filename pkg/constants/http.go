package constants

import "net/http"

type HttpStatus int

func (s HttpStatus) Integer() int {
	return int(s)
}

const (
	HTTP_METHOD_GET     = "GET"
	HTTP_METHOD_POST    = "POST"
	HTTP_METHOD_PUT     = "PUT"
	HTTP_METHOD_DELETE  = "DELETE"
	HTTP_METHOD_PATCH   = "PATCH"
	HTTP_METHOD_OPTIONS = "OPTIONS"
	HTTP_METHOD_TRACE   = "TRACE"
	HTTP_METHOD_CONNECT = "CONNECT"
)

const (
	// Informational responses (100–199)
	HTTP_STATUS_CONTINUE            HttpStatus = http.StatusContinue
	HTTP_STATUS_SWITCHING_PROTOCOLS HttpStatus = http.StatusSwitchingProtocols
	HTTP_STATUS_PROCESSING          HttpStatus = http.StatusProcessing

	// Successful responses (200–299)
	HTTP_STATUS_OK                     HttpStatus = http.StatusOK
	HTTP_STATUS_CREATED                HttpStatus = http.StatusCreated
	HTTP_STATUS_ACCEPTED               HttpStatus = http.StatusAccepted
	HTTP_STATUS_NON_AUTHORITATIVE_INFO            = http.StatusNonAuthoritativeInfo
	HTTP_STATUS_NO_CONTENT             HttpStatus = http.StatusNoContent
	HTTP_STATUS_RESET_CONTENT          HttpStatus = http.StatusResetContent
	HTTP_STATUS_PARTIAL_CONTENT        HttpStatus = http.StatusPartialContent

	// Redirection messages (300–399)
	HTTP_STATUS_MULTIPLE_CHOICES   HttpStatus = http.StatusMultipleChoices
	HTTP_STATUS_MOVED_PERMANENTLY  HttpStatus = http.StatusMovedPermanently
	HTTP_STATUS_FOUND              HttpStatus = http.StatusFound
	HTTP_STATUS_SEE_OTHER          HttpStatus = http.StatusSeeOther
	HTTP_STATUS_NOT_MODIFIED       HttpStatus = http.StatusNotModified
	HTTP_STATUS_USE_PROXY          HttpStatus = http.StatusUseProxy
	HTTP_STATUS_TEMPORARY_REDIRECT HttpStatus = http.StatusTemporaryRedirect
	HTTP_STATUS_PERMANENT_REDIRECT HttpStatus = http.StatusPermanentRedirect

	// Client error responses (400–499)
	HTTP_STATUS_BAD_REQUEST                   HttpStatus = http.StatusBadRequest
	HTTP_STATUS_UNAUTHORIZED                  HttpStatus = http.StatusUnauthorized
	HTTP_STATUS_PAYMENT_REQUIRED              HttpStatus = http.StatusPaymentRequired
	HTTP_STATUS_FORBIDDEN                     HttpStatus = http.StatusForbidden
	HTTP_STATUS_NOT_FOUND                     HttpStatus = http.StatusNotFound
	HTTP_STATUS_METHOD_NOT_ALLOWED            HttpStatus = http.StatusMethodNotAllowed
	HTTP_STATUS_NOT_ACCEPTABLE                HttpStatus = http.StatusNotAcceptable
	HTTP_STATUS_PROXY_AUTHENTICATION_REQUIRED HttpStatus = http.StatusProxyAuthRequired
	HTTP_STATUS_REQUEST_TIMEOUT               HttpStatus = http.StatusRequestTimeout
	HTTP_STATUS_CONFLICT                      HttpStatus = http.StatusConflict
	HTTP_STATUS_GONE                          HttpStatus = http.StatusGone
	HTTP_STATUS_LENGTH_REQUIRED               HttpStatus = http.StatusLengthRequired
	HTTP_STATUS_PRECONDITION_FAILED           HttpStatus = http.StatusPreconditionFailed
	HTTP_STATUS_PAYLOAD_TOO_LARGE             HttpStatus = http.StatusRequestEntityTooLarge
	HTTP_STATUS_URI_TOO_LONG                  HttpStatus = http.StatusRequestURITooLong
	HTTP_STATUS_UNSUPPORTED_MEDIA_TYPE        HttpStatus = http.StatusUnsupportedMediaType
	HTTP_STATUS_RANGE_NOT_SATISFIABLE         HttpStatus = http.StatusRequestedRangeNotSatisfiable
	HTTP_STATUS_EXPECTATION_FAILED            HttpStatus = http.StatusExpectationFailed
	HTTP_STATUS_IM_A_TEAPOT                   HttpStatus = http.StatusTeapot // April Fools' joke status

	// Server error responses (500–599)
	HTTP_STATUS_INTERNAL_SERVER_ERROR      HttpStatus = http.StatusInternalServerError
	HTTP_STATUS_NOT_IMPLEMENTED            HttpStatus = http.StatusNotImplemented
	HTTP_STATUS_BAD_GATEWAY                HttpStatus = http.StatusBadGateway
	HTTP_STATUS_SERVICE_UNAVAILABLE        HttpStatus = http.StatusServiceUnavailable
	HTTP_STATUS_GATEWAY_TIMEOUT            HttpStatus = http.StatusGatewayTimeout
	HTTP_STATUS_HTTP_VERSION_NOT_SUPPORTED HttpStatus = http.StatusHTTPVersionNotSupported
)
