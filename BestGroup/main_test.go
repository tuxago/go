// Get /players/{name}/wins
//Post /players/{name}/wins

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequesting(t *testing.T) {
	t.Run("call request with only /players/", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/", nil)
		response := httptest.NewRecorder()
		PlayerServer(response, request)
		got := response.Body.String()

		want := "No player name called"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("call request with only /players", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players", nil)
		response := httptest.NewRecorder()
		PlayerServer(response, request)
		got := response.Body.String()

		want := "No player name called"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestRecordingWinsAndRetrievingThem(t *testing.T) {

	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response := httptest.NewRecorder()
		PlayerServer(response, request)
		got := response.Body.String()

		want := "20"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("returns Salt's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Salt", nil)
		response := httptest.NewRecorder()
		PlayerServer(response, request)
		got := response.Body.String()

		want := "0"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("update Salt's wins", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/Salt", nil)
		response := httptest.NewRecorder()
		PlayerServer(response, request)
		request, _ = http.NewRequest(http.MethodGet, "/players/Salt", nil)
		response = httptest.NewRecorder()
		PlayerServer(response, request)
		want := "1"
		got := response.Body.String()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("update Pepper's wins twice", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		response := httptest.NewRecorder()
		PlayerServer(response, request)
		PlayerServer(response, request)
		request, _ = http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response = httptest.NewRecorder()
		PlayerServer(response, request)
		want := "22"
		got := response.Body.String()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("call request with non-existing player", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Jacques", nil)
		response := httptest.NewRecorder()
		PlayerServer(response, request)
		got := response.Body.String()

		want := "Player Jacques doesn't exist"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})



}
