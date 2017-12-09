package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const MB = 1024 * 1024

var RangedRequestNotSupportedError = errors.New("Ranged request not supported (Server did not respond with 206 Status)")
var ResponseFromServerWasGZiped = errors.New("HTTP response was gzip encoded. Ranges may not match those requested.")

var ClientNoCompression = &http.Client{
	Transport: &http.Transport{},
}

type URLNotFoundError string

func (url URLNotFoundError) Error() string {
	return "404 Error on URL: " + string(url)
}

// This class provides the implementation of BlockSourceRequester for BlockSourceBase
// this simplifies creating new BlockSources that satisfy the requirements down to
// writing a request function
type HttpRequester struct {
	Client *http.Client
	Url    string
}

func (r *HttpRequester) DoRequest(startOffset int64, endOffset int64) (data []byte, err error) {
	rangedRequest, err := http.NewRequest("GET", r.Url, nil)

	if err != nil {
		return nil, fmt.Errorf("Error creating request for \"%v\": %v", r.Url, err)
	}

	rangeSpecifier := fmt.Sprintf("bytes=%v-%v", startOffset, endOffset-1)
	rangedRequest.ProtoAtLeast(1, 1)
	rangedRequest.Header.Add("Range", rangeSpecifier)
	rangedRequest.Header.Add("Accept-Encoding", "identity")
	rangedResponse, err := r.Client.Do(rangedRequest)

	if err != nil {
		return nil, fmt.Errorf("Error executing request for \"%v\": %v", r.Url, err)
	}

	defer rangedResponse.Body.Close()

	if rangedResponse.StatusCode == 404 {
		return nil, URLNotFoundError(r.Url)
	} else if rangedResponse.StatusCode != 206 {
		return nil, RangedRequestNotSupportedError
	} else if strings.Contains(
		rangedResponse.Header.Get("Content-Encoding"),
		"gzip",
	) {
		return nil, ResponseFromServerWasGZiped
	} else {
		buf := bytes.NewBuffer(make([]byte, 0, endOffset-startOffset))
		_, err = buf.ReadFrom(rangedResponse.Body)

		if err != nil {
			err = fmt.Errorf(
				"Failed to read response body for %v (%v-%v): %v",
				r.Url,
				startOffset, endOffset-1,
				err,
			)
		}

		data = buf.Bytes()

		if int64(len(data)) != endOffset-startOffset {
			err = fmt.Errorf(
				"Unexpected response length %v (%v): %v",
				r.Url,
				endOffset-startOffset+1,
				len(data),
			)
		}

		return
	}
}

func (r *HttpRequester) IsFatal(err error) bool {
	return true
}
