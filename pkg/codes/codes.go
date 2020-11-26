package codes

type Code struct {
	Message     string `yaml:"message"`
	Description string `yaml:"description"`
}

var Codes = map[string]Code{
	"100": {
		Message:     "Continue",
		Description: "The server has received the request headers, and the client should proceed to send the request body.",
	},
	"101": {
		Message:     "Switching Protocols",
		Description: "The requester has asked the server to switch protocols.",
	},
	"102": {
		Message:     "Processing",
		Description: "This code indicates that the server has received and is processing the request, but no response is available yet. This prevents the client from timing out and assuming the request was lost.",
	},
	"103": {
		Message:     "Early Hints",
		Description: "Used to return some response headers before final HTTP Message.",
	},
	"200": {
		Message:     "OK",
		Description: "The request is OK (this is the standard response for successful HTTP requests).",
	},
	"201": {
		Message:     "Created",
		Description: "The request has been fulfilled, and a new resource is created.",
	},
	"202": {
		Message:     "Accepted",
		Description: "The request has been accepted for processing, but the processing has not been completed.",
	},
	"203": {
		Message:     "Non-Authoritative Information",
		Description: "The request has been successfully processed, but is returning information that may be from another source.",
	},
	"204": {
		Message:     "No Content",
		Description: "The request has been successfully processed, but is not returning any content.",
	},
	"205": {
		Message:     "Reset Content",
		Description: "The request has been successfully processed, but is not returning any content, and requires that the requester reset the document view.",
	},
	"206": {
		Message:     "Partial Content",
		Description: "The server is delivering only part of the resource due to a range header sent by the client.",
	},
	"207": {
		Message:     "Multi-Status",
		Description: "The Message body that follows is by default an XML Message and can contain a number of separate response codes, depending on how many sub-requests were made.",
	},
	"208": {
		Message:     "Already Reported",
		Description: "The members of a DAV binding have already been enumerated in a preceding part of the (multistatus) response, and are not being included again.",
	},
	"218": {
		Message:     "This is fine (Apache Web Server)",
		Description: "Used as a catch-all error condition for allowing response bodies to flow through Apache when ProxyErrorOverride is enabled.",
	},
	"226": {
		Message:     "IM Used",
		Description: "The server has fulfilled a request for the resource, and the response is a representation of the result of one or more instance-manipulations applied to the current instance.",
	},
	"300": {
		Message:     "Multiple Choices",
		Description: "A link list. The user can select a link and go to that location. Maximum five addresses.",
	},
	"301": {
		Message:     "Moved Permanently",
		Description: "The requested page has moved to a new URL.",
	},
	"302": {
		Message:     "Found",
		Description: "The requested page has moved temporarily to a new URL.",
	},
	"303": {
		Message:     "See Other",
		Description: "The requested page can be found under a different URL.",
	},
	"304": {
		Message:     "Not Modified",
		Description: "Indicates the requested page has not been modified since last requested.",
	},
	"306": {
		Message:     "Switch Proxy",
		Description: "No longer used. Originally meant 'Subsequent requests should use the specified proxy.'",
	},
	"307": {
		Message:     "Temporary Redirect",
		Description: "The requested page has moved temporarily to a new URL.",
	},
	"308": {
		Message:     "Resume Incomplete",
		Description: "Used in the resumable requests proposal to resume aborted PUT or POST requests.",
	},
	"400": {
		Message:     "Bad Request",
		Description: "The request cannot be fulfilled due to bad syntax.",
	},
	"401": {
		Message:     "Unauthorized",
		Description: "The request was a legal request, but the server is refusing to respond to it. For use when authentication is possible but has failed or not yet been provided.",
	},
	"402": {
		Message:     "Payment Required",
		Description: "Not yet implemented by RFC standards, but reserved for future use.",
	},
	"403": {
		Message:     "Forbidden",
		Description: "The request was a legal request, but the server is refusing to respond to it.",
	},
	"404": {
		Message:     "Not Found",
		Description: "The requested page could not be found but may be available again in the future.",
	},
	"405": {
		Message:     "Method Not Allowed",
		Description: "A request was made of a page using a request method not supported by that page.",
	},
	"406": {
		Message:     "Not Acceptable",
		Description: "The server can only generate a response that is not accepted by the client.",
	},
	"407": {
		Message:     "Proxy Authentication Required",
		Description: "The client must first authenticate itself with the proxy.",
	},
	"408": {
		Message:     "Request Timeout",
		Description: "The server timed out waiting for the request.",
	},
	"409": {
		Message:     "Conflict",
		Description: "The request could not be completed because of a conflict in the request.",
	},
	"410": {
		Message:     "Gone",
		Description: "The requested page is no longer available.",
	},
	"411": {
		Message:     "Length Required",
		Description: "The 'Content-Length' is not defined. The server will not accept the request without it.",
	},
	"412": {
		Message:     "Precondition Failed",
		Description: "The precondition given in the request evaluated to false by the server.",
	},
	"413": {
		Message:     "Request Entity Too Large",
		Description: "The server will not accept the request, because the request entity is too large.",
	},
	"414": {
		Message:     "Request-URI Too Long",
		Description: "The server will not accept the request, because the URL is too long. Occurs when you convert a POST request to a GET request with a long query information.",
	},
	"415": {
		Message:     "Unsupported Media Type",
		Description: "The server will not accept the request, because the media type is not supported.",
	},
	"416": {
		Message:     "Requested Range Not Satisfiable",
		Description: "The client has asked for a portion of the file, but the server cannot supply that portion.",
	},
	"417": {
		Message:     "Expectation Failed",
		Description: "The server cannot meet the requirements of the Expect request-header field.",
	},
	"418": {
		Message: "I'm a teapot",
		Description: `Any attempt to brew coffee with a teapot should result in the error code '418 I'm a teapot'. The resulting entity body MAY be short and stout.
         ________                                       __     
        |        \                                     |  \    
         \▓▓▓▓▓▓▓▓ ______   ______   ______   ______  _| ▓▓_   
           | ▓▓   /      \ |      \ /      \ /      \|   ▓▓ \  
           | ▓▓  |  ▓▓▓▓▓▓\ \▓▓▓▓▓▓\  ▓▓▓▓▓▓\  ▓▓▓▓▓▓\\▓▓▓▓▓▓  
           | ▓▓  | ▓▓    ▓▓/      ▓▓ ▓▓  | ▓▓ ▓▓  | ▓▓ | ▓▓ __ 
           | ▓▓  | ▓▓▓▓▓▓▓▓  ▓▓▓▓▓▓▓ ▓▓__/ ▓▓ ▓▓__/ ▓▓ | ▓▓|  \
           | ▓▓   \▓▓     \\▓▓    ▓▓ ▓▓    ▓▓\▓▓    ▓▓  \▓▓  ▓▓
            \▓▓    \▓▓▓▓▓▓▓ \▓▓▓▓▓▓▓ ▓▓▓▓▓▓▓  \▓▓▓▓▓▓    \▓▓▓▓ 
                                   | ▓▓                        
                                   | ▓▓                        
                                    \▓▓                        

`,
	},
	"419": {
		Message:     "Page Expired (Laravel Framework)",
		Description: "Used by the Laravel Framework when a CSRF Token is missing or expired.",
	},
	"420": {
		Message:     "Method Failure (Spring Framework)",
		Description: "A deprecated response used by the Spring Framework when a method has failed.",
	},
	"421": {
		Message:     "Misdirected Request",
		Description: "The request was directed at a server that is not able to produce a response (for example because a connection reuse).",
	},
	"422": {
		Message:     "Unprocessable Entity",
		Description: "The request was well-formed but was unable to be followed due to semantic errors.",
	},
	"423": {
		Message:     "Locked",
		Description: "The resource that is being accessed is locked.",
	},
	"424": {
		Message:     "Failed Dependency",
		Description: "The request failed due to failure of a previous request (e.g., a PROPPATCH).",
	},
	"426": {
		Message:     "Upgrade Required",
		Description: "The client should switch to a different protocol such as TLS/1.0, given in the Upgrade header field.",
	},
	"428": {
		Message:     "Precondition Required",
		Description: "The origin server requires the request to be conditional.",
	},
	"429": {
		Message:     "Too Many Requests",
		Description: "The user has sent too many requests in a given amount of time. Intended for use with rate limiting schemes.",
	},
	"431": {
		Message:     "Request Header Fields Too Large",
		Description: "The server is unwilling to process the request because either an individual header field, or all the header fields collectively, are too large.",
	},
	"440": {
		Message:     "Login Time-out",
		Description: "The client's session has expired and must log in again. (IIS)",
	},
	"444": {
		Message:     "Connection Closed Without Response",
		Description: "A non-standard status code used to instruct nginx to close the connection without sending a response to the client, most commonly used to deny malicious or malformed requests.",
	},
	"449": {
		Message:     "Retry With",
		Description: "The server cannot honour the request because the user has not provided the required information. (IIS)",
	},
	"450": {
		Message:     "Blocked by Windows Parental Controls",
		Description: "The Microsoft extension code indicated when Windows Parental Controls are turned on and are blocking access to the requested webpage.",
	},
	"451": {
		Message:     "Unavailable For Legal Reasons",
		Description: "A server operator has received a legal demand to deny access to a resource or to a set of resources that includes the requested resource.",
	},
	"494": {
		Message:     "Request Header Too Large",
		Description: "Used by nginx to indicate the client sent too large of a request or header line that is too long.",
	},
	"495": {
		Message:     "SSL Certificate Error",
		Description: "An expansion of the 400 Bad Request response code, used when the client has provided an invalid client certificate.",
	},
	"496": {
		Message:     "SSL Certificate Required",
		Description: "An expansion of the 400 Bad Request response code, used when a client certificate is required but not provided.",
	},
	"497": {
		Message:     "HTTP Request Sent to HTTPS Port",
		Description: "An expansion of the 400 Bad Request response code, used when the client has made a HTTP request to a port listening for HTTPS requests.",
	},
	"498": {
		Message:     "Invalid Token (Esri)",
		Description: "Returned by ArcGIS for Server. Code 498 indicates an expired or otherwise invalid token.",
	},
	"499": {
		Message:     "Client Closed Request",
		Description: "A non-standard status code introduced by nginx for the case when a client closes the connection while nginx is processing the request.",
	},
	"500": {
		Message:     "Internal Server Error",
		Description: "An error has occurred in a server side script, a no more specific Message is suitable.",
	},
	"501": {
		Message:     "Not Implemented",
		Description: "The server either does not recognize the request method, or it lacks the ability to fulfill the request.",
	},
	"502": {
		Message:     "Bad Gateway",
		Description: "The server was acting as a gateway or proxy and received an invalid response from the upstream server.",
	},
	"503": {
		Message:     "Service Unavailable",
		Description: "The server is currently unavailable (overloaded or down).",
	},
	"504": {
		Message:     "Gateway Timeout",
		Description: "The server was acting as a gateway or proxy and did not receive a timely response from the upstream server.",
	},
	"505": {
		Message:     "HTTP Version Not Supported",
		Description: "The server does not support the HTTP protocol version used in the request.",
	},
	"506": {
		Message:     "Variant Also Negotiates",
		Description: "Transparent content negotiation for the request results in a circular reference.",
	},
	"507": {
		Message:     "Insufficient Storage",
		Description: "The server is unable to store the representation needed to complete the request.",
	},
	"508": {
		Message:     "Loop Detected",
		Description: "The server detected an infinite loop while processing the request (sent instead of 208 Already Reported).",
	},
	"509": {
		Message:     "Bandwidth Limit Exceeded",
		Description: "The server has exceeded the bandwidth specified by the server administrator; this is often used by shared hosting providers to limit the bandwidth of customers.",
	},
	"510": {
		Message:     "Not Extended",
		Description: "Further extensions to the request are required for the server to fulfil it.",
	},
	"511": {
		Message:     "Network Authentication Required",
		Description: "The client needs to authenticate to gain network access.",
	},
	"520": {
		Message:     "Unknown Error",
		Description: "The 520 error is used as a 'catch-all response for when the origin server returns something unexpected', listing connection resets, large headers, and empty or invalid responses as common triggers.",
	},
	"521": {
		Message:     "Web Server Is Down",
		Description: "The origin server has refused the connection from Cloudflare.",
	},
	"522": {
		Message:     "Connection Timed Out",
		Description: "Cloudflare could not negotiate a TCP handshake with the origin server.",
	},
	"523": {
		Message:     "Origin Is Unreachable",
		Description: "Cloudflare could not reach the origin server; for example, if the DNS records for the origin server are incorrect.",
	},
	"524": {
		Message:     "A Timeout Occurred",
		Description: "Cloudflare was able to complete a TCP connection to the origin server, but did not receive a timely HTTP response.",
	},
	"525": {
		Message:     "SSL Handshake Failed",
		Description: "Cloudflare could not negotiate a SSL/TLS handshake with the origin server.",
	},
	"526": {
		Message:     "Invalid SSL Certificate",
		Description: "Used by Cloudflare and Cloud Foundry's gorouter to indicate failure to validate the SSL/TLS certificate that the origin server presented.",
	},
	"527": {
		Message:     "Railgun Listener to Origin Error",
		Description: "Error 527 indicates that the request timed out or failed after the WAN connection had been established.",
	},
	"530": {
		Message:     "Origin DNS Error",
		Description: "Error 530 indicates that the requested host name could not be resolved on the Cloudflare network to an origin server.",
	},
	"598": {
		Message:     "Network Read Timeout Error",
		Description: "Used by some HTTP proxies to signal a network read timeout behind the proxy to a client in front of the proxy.",
	},
	"1xx": {
		Message:     "Information",
		Description: "1xx codes are often interim responses for sharing connection status information. Not intended for final request or response action.",
	},
	"2xx": {
		Message:     "Successful",
		Description: "2xx codes indicate successful responses usually this means the action requested by the client was received, understood and accepted successfully.",
	},
	"3xx": {
		Message:     "Redirection",
		Description: "3xx codes are a class of responses that suggest the User-Agent must follow another course of action to obtain the complete requested resource.",
	},
	"4xx": {
		Message:     "Client Error",
		Description: "4xx codes generally are error responses specifying an issue at the client’s end. Potentially a network issue.",
	},
	"5xx": {
		Message:     "Server Error",
		Description: "5xx error codes indicate that an error or unresolvable request occurred on the server side, whether that is a proxy or the origin host.",
	},
}
