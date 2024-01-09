package httpclient

import (
	"context"
	"io"
	"net/http"
)

// NewRequest returns a new Request given a context, method, URL, and optional body.
//
// If the provided body is also an io.Closer, the returned
// Request.Body is set to body and will be closed by the Client
// methods Do, Post, and PostForm, and Transport.RoundTrip.
//
// NewRequest returns a Request suitable for use with Client.Do or
// Transport.RoundTrip. To create a request for use with testing a
// Server Handler, either use the NewRequest function in the
// net/http/httptest package, use ReadRequest, or manually update the
// Request fields. See the Request type's documentation for the
// difference between inbound and outbound request fields.
//
// If body is of type *bytes.Buffer, *bytes.Reader, or
// *strings.Reader, the returned request's ContentLength is set to its
// exact value (instead of -1), GetBody is populated (so 307 and 308
// redirects can replay the body), and Body is set to NoBody if the
// ContentLength is 0.
func NewRequest(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return req, err
	}

	return req.WithContext(ctx), err
}
