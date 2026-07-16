package ledger

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

type ChannelPayload struct {
	i int
	s string
	e error
}

func sortEntries(entries []Entry) []Entry {
	m1 := map[bool]int{true: 0, false: 1}
	m2 := map[bool]int{true: -1, false: 1}
	entriesCopy := entries

	// removes extraneous loop which would cause O(n^2)
	for len(entriesCopy) > 1 {
		first, rest := entriesCopy[0], entriesCopy[1:]
		for i, e := range rest {
			if (m1[e.Date == first.Date]*m2[e.Date < first.Date]*4 +
				m1[e.Description == first.Description]*m2[e.Description < first.Description]*2 +
				m1[e.Change == first.Change]*m2[e.Change < first.Change]*1) < 0 {
				entriesCopy[0], entriesCopy[i+1] = entriesCopy[i+1], entriesCopy[0]
				break
			}
		}
		entriesCopy = entriesCopy[1:]
	}

	return entries
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	var entriesCopy []Entry

	if len(entries) == 0 {
		if _, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil {
			return "", err
		}
	}

	// simplifies copy of entries
	entriesCopy = sortEntries(append(entriesCopy, entries...))

	// ledger header creation
	// renames `s` into header
	header, headerError := ledgerHeader(locale)
	if headerError != nil {
		return "", headerError
	}

	// creates channelError as it is always the same
	channelError := ChannelPayload{e: errors.New("")}

	// Parallelism, always a great idea
	// uses ChannelPayload type to avoid passing wrongly built payloads
	channel := make(chan ChannelPayload)
	for key, entry := range entriesCopy {
		go func(key int, entry Entry) {
			if len(entry.Date) != 10 {
				channel <- channelError
			}
			d2, d4 := entry.Date[4], entry.Date[7]
			if d2 != '-' || d4 != '-' {
				channel <- channelError
			}
			entryDescription := formatEntryDescription(entry.Description)
			formattedDate := formatDate(entry.Date, locale)
			negative := false
			cents := entry.Change
			if cents < 0 {
				cents = cents * -1
				negative = true
			}
			var a string
			if locale == "nl-NL" {
				if currency == "EUR" {
					a += "€"
				} else if currency == "USD" {
					a += "$"
				} else {
					channel <- ChannelPayload{e: errors.New("")}
				}
				a += " "
				centsStr := strconv.Itoa(cents)
				switch len(centsStr) {
				case 1:
					centsStr = "00" + centsStr
				case 2:
					centsStr = "0" + centsStr
				}
				rest := centsStr[:len(centsStr)-2]
				var parts []string
				for len(rest) > 3 {
					parts = append(parts, rest[len(rest)-3:])
					rest = rest[:len(rest)-3]
				}
				if len(rest) > 0 {
					parts = append(parts, rest)
				}
				if negative {
					a += "-"
				}
				for i := len(parts) - 1; i >= 0; i-- {
					a += parts[i] + "."
				}
				a = a[:len(a)-1]
				a += ","
				a += centsStr[len(centsStr)-2:]
				a += " "
			} else if locale == "en-US" {
				if negative {
					a += "("
				}
				switch currency {
				case "EUR":
					a += "€"
				case "USD":
					a += "$"
				default:
					channel <- channelError
				}
				centsStr := fmt.Sprintf("%03s", strconv.Itoa(cents))
				rest := centsStr[:len(centsStr)-2]
				var parts []string
				for len(rest) > 3 {
					parts = append(parts, rest[len(rest)-3:])
					rest = rest[:len(rest)-3]
				}
				if len(rest) > 0 {
					parts = append(parts, rest)
				}
				for i := len(parts) - 1; i >= 0; i-- {
					a += parts[i] + ","
				}
				a = a[:len(a)-1]
				a += "."
				a += centsStr[len(centsStr)-2:]
				if negative {
					a += ")"
				} else {
					a += " "
				}
			} else {
				channel <- channelError
			}
			var al int
			for range a {
				al++
			}
			channel <- ChannelPayload{
				i: key,
				s: formattedDate + strings.Repeat(" ", 10-len(formattedDate)) + " | " + entryDescription + " | " + strings.Repeat(" ", 13-al) + a + "\n",
			}
		}(key, entry)
	}
	ss := make([]string, len(entriesCopy))
	for range entriesCopy {
		v := <-channel
		if v.e != nil {
			return "", v.e
		}
		ss[v.i] = v.s
	}
	for i := range len(entriesCopy) {
		header += ss[i]
	}
	return header, nil
}

func ledgerHeader(locale string) (string, error) {
	switch locale {
	case "nl-NL":
		header := fmt.Sprintf("Datum%6s| Omschrijving%14s| Verandering%2s\n", "", "", "")
		return header, nil

	case "en-US":
		header := fmt.Sprintf("Date%7s| Description%15s| Change%7s\n", "", "", "")

		return header, nil

	default:
		return "", errors.New("")
	}
}

func formatDate(date string, locale string) string {
	d1, d2, d3 := date[0:4], date[5:7], date[8:10]
	switch locale {
	case "nl-NL":
		return d3 + "-" + d2 + "-" + d1
	case "en-US":
		return d2 + "/" + d3 + "/" + d1
	default:
		return date
	}
}

func formatEntryDescription(description string) string {
	if len(description) > 25 {
		return description[:22] + "..."
	} else {
		return description + strings.Repeat(" ", 25-len(description))
	}
}

func formatNegativeNumberFor(locale string, number string) string {
	switch locale {
	case "nl-NL":
		return "-" + number
	case "en-US":
		return "(" + number + ")"
	default:
		return "-" + number
	}
}
