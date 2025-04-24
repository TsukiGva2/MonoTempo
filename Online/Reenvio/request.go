package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	/* i love this library */
	backoff "github.com/cenkalti/backoff"
)

const (
	REQUEST_TIMEOUT = 20 * time.Second
)

type Form map[string]string
type RawForm []byte

func SimpleRawRequest(url string, data RawForm, contentType string) (err error) {

	var res *http.Response

	bf := backoff.NewExponentialBackOff()
	bf.MaxElapsedTime = REQUEST_TIMEOUT

	err = backoff.Retry(
		func() (err error) {
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))

			if err != nil {
				err = fmt.Errorf("error creating request: %s", err)

				return
			}

			req.Header.Set("Content-Type", contentType)

			res, err = http.DefaultClient.Do(req)

			return
		},

		bf,
	)

	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("error connecting to '%s': got HTTP %d", url, res.StatusCode)

		return
	}

	_, err = io.ReadAll(res.Body)

	if err != nil {
		err = fmt.Errorf("error reading response body: %s", err)

		return
	}

	return
}
