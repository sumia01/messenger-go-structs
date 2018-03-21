package messenger

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func doRequest(method string, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	return http.DefaultClient.Do(req)
}

func doThreadRequest(method string, url string, body io.Reader) error {
	resp, err := doRequest(method, url, body)
	if err != nil {
		return errors.Wrap(err, "doThreadRequest - doRequest()")
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return errors.Wrap(err, "ioutil.ReadAll fail")
	}

	if resp.StatusCode != http.StatusOK {
		return errors.Wrapf(err, "doThreadRequest response != 200: %v", string(respBody))
	}
	return nil
}
