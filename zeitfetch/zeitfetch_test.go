package zeitfetch

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchTimeFromWeb(t *testing.T) {
	tests := []struct {
		name           string
		html           string
		erwartet       string
		erwartetFehler bool
	}{
		{
			name:           "Korrektes Timeslot-Div",
			html:           `<html><body><div id="Timeslot">12:34:56</div></body></html>`,
			erwartet:       "12:34:56",
			erwartetFehler: false,
		},
		{
			name:           "Kein Timeslot-Div",
			html:           `<html><body><div id="Other">12:34:56</div></body></html>`,
			erwartet:       "",
			erwartetFehler: true,
		},
		{
			name:           "Leeres Timeslot-Div",
			html:           `<html><body><div id="Timeslot"></div></body></html>`,
			erwartet:       "",
			erwartetFehler: true,
		},
		{
			name:           "Timeslot-Div mit Leerzeichen",
			html:           `<html><body><div id="Timeslot"> 12:34:56 </div></body></html>`,
			erwartet:       " 12:34:56 ",
			erwartetFehler: false,
		},
	}

	for _, tt := range tests {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(tt.html))
		}))
		defer ts.Close()

		t.Run(tt.name, func(t *testing.T) {
			ergebnis, err := FetchTimeFromWeb(ts.URL)
			if tt.erwartetFehler {
				if err == nil {
					t.Errorf("Erwarteter Fehler, aber keiner erhalten. Ergebnis: %q", ergebnis)
				}
			} else {
				if err != nil {
					t.Errorf("Unerwarteter Fehler: %v", err)
				}
				if ergebnis != tt.erwartet {
					t.Errorf("Erwartet: %q, erhalten: %q", tt.erwartet, ergebnis)
				}
			}
		})
	}
}
