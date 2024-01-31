package requester

import (
	"context"
	"fmt"
	"github.com/valyala/fasthttp"
)

func SendHttpRequest(ctx context.Context, url, method string, body []byte, headers map[string]string) (respBody []byte, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)

	req.Header.SetMethod(method)
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	req.SetBody(body)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err = fasthttp.Do(req, resp)
	if err != nil {
		return nil, fmt.Errorf("failed to send requester: %w", err)
	}

	if resp.StatusCode() > 301 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return resp.Body(), nil
}
