package router

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCustomersRouter(t *testing.T) {
	srv := httptest.NewServer(CustomersRouter())
	defer srv.Close()

	tt := []struct {
		name string
		id   string
	}{
		{
			"one",
			"7",
		},
		{
			"two",
			"12",
		},
		{
			"three",
			"26",
		},
		{
			"four",
			"",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			//path := fmt.Sprintf("%s/v1/customers/%s", srv.URL, tc.id)
			path := fmt.Sprintf("http://localhost:3080/v1/customers/%s", tc.id)
			t.Log(path)

			res, err := http.Get(path)
			if err != nil {
				t.Fatalf("could not send GET request: %v", err)
			}
			defer res.Body.Close()

			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status OK; got %v", res.Status)
			}

			if tc.id == "" {
				if res.StatusCode != http.StatusNotFound {
					t.Errorf("expected status not found Request; got %v", res.StatusCode)
				}
				return
			}

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}

			t.Log(string(b))
		})
	}
}
