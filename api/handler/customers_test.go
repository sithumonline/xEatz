package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type x struct {
	CusID string
	FName string
	LName string
	Phone string
	Mail  string
}

func TestAddACustomer(t *testing.T) {
	tt := []struct {
		name string
		req  x
	}{
		{
			"one",
			x{
				"123",
				"Joo",
				"Mooo",
				"20254",
				"mm@ww.com",
			},
		},
		{
			"two",
			x{
				"254",
				"Foo",
				"Nuuo",
				"2363354",
				"fooo@ww.com",
			},
		},
		{
			"three",
			x{
				"3695",
				"Koolal",
				"Bonna",
				"882232",
				"koolal@ww.com",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out, err := json.Marshal(tc.req)
			if err != nil {
				t.Errorf("faile to marshel : %v", err)
			}
			req, err := http.NewRequest("GET", "localhost:3080/api/v1/customers", bytes.NewBuffer(out))
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			rec := httptest.NewRecorder()
			AddACustomer(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}

			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status OK; got %v", res.Status)
			}

			t.Log(string(b))
		})
	}
}
