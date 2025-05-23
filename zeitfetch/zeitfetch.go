package zeitfetch

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// FetchTimeFromWeb extrahiert die Zeit aus einem <div id="Timeslot">...</div> mit HTML-Parser
func FetchTimeFromWeb(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		return "", fmt.Errorf("Fehler beim Parsen des HTML: %v", err)
	}

	var timeslot string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == "Timeslot" {
					if n.FirstChild != nil {
						timeslot = n.FirstChild.Data
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	if timeslot == "" {
		return "", fmt.Errorf("Kein <div id=Timeslot> gefunden oder leer")
	}
	return timeslot, nil
}
