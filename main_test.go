package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestValidHook(t *testing.T) {
	url := "/hooks"
	handler := http.HandlerFunc(hookHandler)

	w := httptest.NewRecorder()

	req, err := http.NewRequest("POST", url, postBody())
	if err != nil {
		t.Errorf("Bad request: %v", err)
	}

	handler.ServeHTTP(w, req)

	if w.Code != 201 {
		t.Errorf("Status code was not 201. Was: %v", w.Code)
	}
}

func postBody() *strings.Reader {
	return strings.NewReader(`
    {
      "id": "c16423cc-681b-4d4d-bf70-e2bb26889848",
      "created_at": "2015-01-01T00:00:00Z",
      "delivery_attempt": 1,
      "notification_id": "2bb9f97b-b26c-4fae-b8d1-4439ab6d428a",
      "payload": {
        "type": "new-block",
        "block_chain": "bitcoin",
        "block": {
          "hash": "0000000000000000329fcc0e6cd53b7ae32b6c00e3f8bf356ea898f5b23a73e7",
          "previous_block_hash": "0000000000000000185e56fbad963648dc1916e20e9ccd07e4308b3013812659",
          "height": 314645,
          "confirmations": 1,
          "merkle_root": "5994b79a9bd567a9c4199bea91529fd0d61b75c0b7fff2e7e7d60112c4987001",
          "time": "2014-08-09T03:33:10Z",
          "nonce": 2324252789,
          "difficulty": 19729645940.577133,
          "bits": "1837ba62",
          "transaction_hashes": [
            "62278007bbcced2a2e9d3cc118f447ebd13d9e8119d50bfb56ad98050aafadfd",
            "602a2eb5918b3d75819aec57c26f8bf211ab743e97c4afdd30b6918339d7cc84"
          ],
          "chain_received_at": "0001-01-01T00:00:00Z"
        }
      }
    }
`)
}
