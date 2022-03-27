package url

import (
	"context"
	"testing"
)

//Test then retrieve - test the shorten URL is stored and then retrived from database
func TestShortenAndRetrieve(t *testing.T) {
	testURL := "https://github.com/encoredev/encore"

	sp := ShortenParams{URL: testURL}
	resp, err := Shorten(context.Background(), &sp)
	if err != nil {
		t.Fatal(err)
	}
	wantURL := testURL
		if resp.URL != wantURL{
			t.Errorf("got %q, want %q", resp.URL, wantURL)
		}
	
	firstURL := resp
	gotURL, err := Get(context.Background(), firstURL.ID)
	if err != nil {
		t.Fatal(err)
	}
	if *gotURL != *firstURL {
		t.Errorf("got %v, want %v", *gotURL, *firstURL)
	}
			

}