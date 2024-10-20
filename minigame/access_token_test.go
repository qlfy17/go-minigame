package minigame

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAccessService_GetStableAccessToken(t *testing.T) {
	teardownServer := setupCleanServer()
	defer teardownServer(t)

	mux.HandleFunc("/cgi-bin/stable_token", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"access_token": "test access token","expires_in": 7200}`)
	})

	res, err := client.AccessToken.GetStableAccessToken(context.Background())
	if err != nil {
		t.Errorf("GetStableAccessToken(ctx): %v %v", err, res)
	}

	want := &AccessToken{
		AccessToken: "test access token",
		ExpiresIn:   7200,
	}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("GetStableAccessToken(ctx) mismatch (-got +want) %v %v", res, want)
	}
}
