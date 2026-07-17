package ledger

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

type ChannelMessage struct {
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

	// creates channelErrorMessage as it is always the same
	channelErrorMessage := ChannelMessage{e: errors.New("")}

	// Parallelism, always a great idea
	// uses ChannelPayload type to avoid passing wrongly built payloads
	channelMessages := make(chan ChannelMessage)

	for key, entry := range entriesCopy {
		go func(key int, entry Entry) {
			if len(entry.Date) != 10 || isValidDateSeparator(entry.Date) != nil {
				channelMessages <- channelErrorMessage
			}

			entryDescription := formatEntryDescription(entry.Description)
			entryDate := formatDate(entry.Date, locale)

			if !isValidCurrency(currency) {
				channelMessages <- channelErrorMessage
			}

			entryChangeFormatted := formatEntryChange(entry.Change, currency, locale)

			var entryChangeFormattedLength int
			for range entryChangeFormatted {
				entryChangeFormattedLength++
			}
			channelMessages <- ChannelMessage{
				i: key,
				s: entryDate + strings.Repeat(" ", 10-len(entryDate)) + " | " + entryDescription + " | " + strings.Repeat(" ", 13-entryChangeFormattedLength) + entryChangeFormatted + "\n",
			}
		}(key, entry)
	}
	ss := make([]string, len(entriesCopy))
	for range entriesCopy {
		message := <-channelMessages
		if message.e != nil {
			return "", message.e
		}
		ss[message.i] = message.s
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

func isValidDateSeparator(date string) error {
	if date[4] != '-' || date[7] != '-' {
		return errors.New("invalid date separator")
	}

	return nil
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

func isValidCurrency(currency string) bool {
	switch currency {
	case "EUR", "USD":
		return true
	default:
		return false
	}
}

func formatAmmountFor(locale string, currency string, number string, isNegative bool) string {
	negative := ""
	if isNegative {
		negative = "-"
	}

	switch {
	case currency == "EUR" && locale == "nl-NL":
		return "€ " + negative + number + " "
	case currency == "USD" && locale == "nl-NL":
		return "$ " + negative + number + " "
	case currency == "EUR" && locale == "en-US":
		if isNegative {
			return "(€" + number + ")"
		} else {
			return "€" + number
		}
	case currency == "USD" && locale == "en-US":
		if isNegative {
			return "($" + number + ")"
		} else {
			return "$" + number + " "
		}
	default:
		return ""
	}
}

func formatEntryChange(change int, currency string, locale string) string {
	isNegativeChange := false
	valueInCents := 0.0
	if change < 0 {
		isNegativeChange = true
		valueInCents = float64(change*-1) / 100.
	} else {
		valueInCents = float64(change) / 100.
	}

	var p *message.Printer
	var languageToUse language.Tag
	switch locale {
	case "nl-NL":
		languageToUse = language.Dutch
	case "en-US":
		languageToUse = language.AmericanEnglish
	}

	p = message.NewPrinter(languageToUse)

	formattedChange := ""
	if valueInCents < 0 {
		formattedChange = p.Sprintf("%.2f", number.Decimal(valueInCents, number.FormatWidth(3), number.Pad('0')))
	} else {
		formattedChange = p.Sprintf("%.2f", number.Decimal(valueInCents))
	}

	formattedChange = formatAmmountFor(locale, currency, formattedChange, isNegativeChange)

	return formattedChange
}

func formatEntryDescription(description string) string {
	if len(description) > 25 {
		return description[:22] + "..."
	} else {
		return description + strings.Repeat(" ", 25-len(description))
	}
}
