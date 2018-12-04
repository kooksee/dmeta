package utils

import (
	"fmt"
	"github.com/kooksee/dmeta/internal/errs"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	MustNotError(err)

	// 处理header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.67 Safari/537.36")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return []byte(fmt.Sprintf("url: %s not found", url)), errs.NotFound
	}

	if resp.StatusCode != http.StatusOK {
		dt, err := ioutil.ReadAll(resp.Body)
		MustNotError(err)
		return nil, fmt.Errorf("url: %s get error, output: %s", url, dt)
	}

	dt, err := ioutil.ReadAll(resp.Body)
	MustNotError(err)

	return dt, nil
}
