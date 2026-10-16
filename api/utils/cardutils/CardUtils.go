package cardutils

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/tracknumber"
	"github.com/globalpayments/go-sdk/api/paymentmethods/abstractions"
	"regexp"
	"strings"
)

var (
	AmexRegex          = regexp.MustCompile("^3[47]")
	MasterCardRegex    = regexp.MustCompile("^(?:5[1-6]|222[1-9]|22[3-9][0-9]|2[3-6][0-9]{2}|27[01][0-9]|2720)")
	VisaRegex          = regexp.MustCompile("^4")
	DinersClubRegex    = regexp.MustCompile("^3(?:0[0-5]|[68][0-9])")
	RouteClubRegex     = regexp.MustCompile("^(2014|2149)")
	DiscoverRegex      = regexp.MustCompile("^6(?:011|5[0-9]{2})")
	JcbRegex           = regexp.MustCompile("^(?:2131|1800|35\\d{3})")
	VoyagerRegex       = regexp.MustCompile("^70888[5-9]")
	WexRegex           = regexp.MustCompile("^(?:690046|707138)")
	FuelmanRegex       = regexp.MustCompile("^707649[0-9]")
	FleetwideRegex     = regexp.MustCompile("^707685[0-9]")
	StoredValueRegex   = regexp.MustCompile("^(?:600649|603261|603571|627600|639470)")
	ValueLinkRegex     = regexp.MustCompile("^(?:601056|603225)")
	HeartlandGiftRegex = regexp.MustCompile("^(?:502244|627720|708355)")
	UnionPayRegex      = regexp.MustCompile("^(?:62[0-8]|81[0-8])")

	trackOnePattern = regexp.MustCompile("%?[B0]?([\\d]+)\\^[^\\^]+\\^([\\d]{4})([^?]+)?/?")
	trackTwoPattern = regexp.MustCompile(";?([\\d]+)[=|[dD]](\\d{4})([^?]+)?/?")

	fleetBinMap     map[string]map[string]string
	regexMap        map[string]*regexp.Regexp
	readyLinkBinMap []string
)

func init() {
	regexMap = map[string]*regexp.Regexp{
		"Amex": AmexRegex,
		"MC":   MasterCardRegex,
		// ... other regex patterns
	}

	readyLinkBinMap = []string{
		"462766",
		"406498",
		// ... other bins
	}

	// fleet bin ranges initialization
	fleetBinMap = map[string]map[string]string{
		// Fleet bin map initializations
	}
}

func IsFleet(cardType, pan string) bool {
	if pan != "" {
		compareValue := pan[0:6]
		baseCardType := strings.TrimSuffix(cardType, "Fleet")

		if binRanges, ok := fleetBinMap[baseCardType]; ok {
			for lower, upper := range binRanges {
				if compareValue >= lower && compareValue <= upper {
					return true
				}
			}
		}
	}
	return false
}

func isReadyLink(pan string) bool {
	if pan != "" {
		compareValue := pan[0:6]
		for _, bin := range readyLinkBinMap {
			if bin == compareValue {
				return true
			}
		}
	}
	return false
}

func MapCardType(pan string) string {
	rvalue := "Unknown"
	if pan != "" {
		pan = strings.Replace(pan, " ", "", -1)
		pan = strings.Replace(pan, "-", "", -1)

		for key, pattern := range regexMap {
			if pattern.MatchString(pan) {
				rvalue = key
				break
			}
		}

		if rvalue != "Unknown" {
			if IsFleet(rvalue, pan) && rvalue != "FleetWide" {
				rvalue += "Fleet"
			} else if isReadyLink(pan) {
				rvalue += "ReadyLink"
			}
		}
	}
	return rvalue
}

func GetBaseCardType(cardType string) string {
	resultCardType := cardType
	for cardTypeKey := range regexMap {
		if strings.HasPrefix(strings.ToUpper(cardType), strings.ToUpper(cardTypeKey)) {
			return cardTypeKey
		}
	}
	return resultCardType
}

func ParseTrackData(paymentMethod abstractions.ITrackData) abstractions.ITrackData {
	trackData := paymentMethod.GetValue()
	matcher := trackTwoPattern.FindStringSubmatch(trackData)
	if len(matcher) > 0 {
		pan := matcher[1]
		expiry := matcher[2]
		discretionary := matcher[3]

		if discretionary != "" {
			if len(pan+expiry+discretionary) == 37 && strings.HasSuffix(strings.ToLower(discretionary), "f") {
				discretionary = discretionary[:len(discretionary)-1]
			}
		}

		paymentMethod.SetTrackNumber(tracknumber.TrackTwo)
		paymentMethod.SetPan(pan)
		paymentMethod.SetExpiry(expiry)
		paymentMethod.SetDiscretionaryData(discretionary)
		paymentMethod.SetTrackData(pan + "=" + expiry + discretionary)
	} else {
		matcher = trackOnePattern.FindStringSubmatch(trackData)
		if len(matcher) > 0 {
			paymentMethod.SetTrackNumber(tracknumber.TrackOne)
			paymentMethod.SetPan(matcher[1])
			paymentMethod.SetExpiry(matcher[2])
			paymentMethod.SetDiscretionaryData(matcher[3])
			paymentMethod.SetTrackData(strings.TrimLeft(matcher[0], "%"))
		}
	}

	return paymentMethod
}
