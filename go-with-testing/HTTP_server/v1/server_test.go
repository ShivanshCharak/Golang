package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()
	PlayerServer(response, request)
	t.Run("returns pepper; score", func(t *testing.T) {
		got := response.Body.String()
		want := "20"
		if got != want {
			t.Errorf("got %q wanted %q", got, want)
		}
	})
}
