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

func sortEntries(entries []Entry) []Entry {
	m1 := map[bool]int{true: 0, false: 1}
	m2 := map[bool]int{true: -1, false: 1}
	es := entries

	// removes extraneous loop which would cause O(n^2)
	for len(es) > 1 {
		first, rest := es[0], es[1:]
		for i, e := range rest {
			if (m1[e.Date == first.Date]*m2[e.Date < first.Date]*4 +
				m1[e.Description == first.Description]*m2[e.Description < first.Description]*2 +
				m1[e.Change == first.Change]*m2[e.Change < first.Change]*1) < 0 {
				es[0], es[i+1] = es[i+1], es[0]
				break
			}
		}
		es = es[1:]
	}

	return entries
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	var entriesCopy []Entry

	// simplifies copy of entries
	entriesCopy = append(entriesCopy, entries...)

	if len(entries) == 0 {
		if _, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil {
			return "", err
		}
	}

	entriesCopy = sortEntries(entriesCopy)

	// ledger header creation
	// renames `s` into header
	header, headerError := ledgerHeader(locale)
	if headerError != nil {
		return "", headerError
	}

	// creates channelError as it is always the same
	channelError := buildChannelPayload(ChannelPayload{e: errors.New("")})

	// Parallelism, always a great idea
	// uses ChannelPayload type to avoid passing wrongly built payloads
	channel := make(chan ChannelPayload)
	for key, entry := range entriesCopy {
		go func(key int, entry Entry) {
			if len(entry.Date) != 10 {
				channel <- channelError
			}
			d1, d2, d3, d4, d5 := entry.Date[0:4], entry.Date[4], entry.Date[5:7], entry.Date[7], entry.Date[8:10]
			if d2 != '-' || d4 != '-' {
				channel <- channelError
			}
			entryDescription := formetEntryDescription(entry.Description)
			var d string
			if locale == "nl-NL" {
				d = d5 + "-" + d3 + "-" + d1
			} else if locale == "en-US" {
				d = d3 + "/" + d5 + "/" + d1
			}
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
					channel <- struct {
						i int
						s string
						e error
					}{e: errors.New("")}
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
				if currency == "EUR" {
					a += "€"
				} else if currency == "USD" {
					a += "$"
				} else {
					channel <- channelError
				}
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
			channel <- ChannelPayload{i: key, s: d + strings.Repeat(" ", 10-len(d)) + " | " + entryDescription + " | " +
				strings.Repeat(" ", 13-al) + a + "\n"}
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

func formetEntryDescription(description string) string {
	if len(description) > 25 {
		return description[:22] + "..."
	} else {
		return description + strings.Repeat(" ", 25-len(description))
	}
}

type ChannelPayload struct {
	i int
	s string
	e error
}

func buildChannelPayload(value ChannelPayload) ChannelPayload {
	return ChannelPayload{
		i: value.i,
		s: value.s,
		e: value.e,
	}
}
