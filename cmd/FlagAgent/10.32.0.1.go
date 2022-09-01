package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagAgent"
)

func main() {
	privateKey := "{\"D\": \"AkNjdPzzwngJUYnDu2X+hJsQhZ7nZQjcwpXnzsM64HiOGu9/DxzTVlLnchxycjJKFywkhoLIharX2hkAjLhaBM8tPsHkHquKwALfjYTJi539F4qOsmOy5HrAgwEYbBSnJtteX+rDsb8aMw0B2KEmthMTVFHpn/dFL5XCWgayDaoR0laHogEKtAtAf4AvVPB+i+iydhuFvpY9FQ48rW6RkqU485GNgXZ/vfVhMTkwLGtzHtOVbSeba+kOnYFqJsuSHzo0de7llTRem68Pbhw9as9/My7bJSeBscOwm3SLQpzWVlax+LyzP6C7NZGIEec43kVyGVcy1sfgwQoP+tQfpC0=\", \"E\": \"AQABAA==\", \"N\": \"Apz+2QJPJiMuBCjahsp2WHNwL3Y7IhCPhkyEzCD1O00l+Ii9ox8ZiWhHLGY4LDvrOm9d0hwvEhzl6wXEML/Si6nuFw1gVOJkOfch4kgVevfFTIVk4p5pQWaGECiVACDKWXE6UTP8gpilylG/M+PIezKZ6ql7rkmhupd26VKRTGE6lWu71XVVoNQAeNm6uGwihT6R6CyVyCeYZZcifIg+2oPHSvdeOFEfl0BsVdSGjDrC2lAgUpnriiMlNsOTuIhfvzTD2dwKQrCh87HUjwA0EDOiAAsWD/M61xPEMZml7Dj8DoCo8CFgBzOfpD++mrnwT1TyNRx3hH5hs8iPjfrFXFk=\", \"Primes1\": \"As2ExSbY17Q+hu9TO0RFgwS5gyLBPsu1HpXfVHX9oMv4New8RcIgq/6o1JpcRZ9NCtmKRLmWvGk562HR5vjVNirxBVDbtv0iRLBwK0PM+T6iG7L9yqf/jLpG4kndA51A2trtFPtHhh6hJQOVZU/mr8DDNnBwmQkuyGf1i1NelnMD\", \"Primes2\": \"AsOO5C95Z2Qd0b+cegBjRaIR/Yl51CsouJdwXrkYiUdn1bAQhz+UFF3W8o6FHigqblDGsVagJkvseEzUs99YFBYMw2JlBjf87zZo7+RlsfMviF6iFxtNqkK3uuHEBFXUeanoOdduyVoigGT6ooAF1Wfc6FM+JTWRpyNysJuCSeZz\"}"
	publlicKey := "{\"N\": \"AqRzRA81LnCcUJdJCASXw4XhAmQczECCcY4zDiYKggmsTGAebcOXaq9HelkpYp1cN7fNhYHO+AoIjAoiBrkv9JwugE2ro9Htl6bmuWJN7LMVIEyrLxGM0XLoC1mEjNNqidg54tVwhkW3wWGFky+WOeCElKo+NBnHdAe1ekSSeQLLsKPbg1WRPZtky7aseexr+NZqKbAHPLQdsLhIddlOMRNnNllWM0q1BFUl1IeRQCAivPtrGrAy1XqVRKeT8xrLAQxfXGokmK/Goo8HZW9RZipqfJ9xbO3sgfEmabOxBnxF1PlA+YvURwQrvKA41af09/aJ9WlMXEtgGzfadEFn6N0=\", \"E\": \"AQABAA==\"}"
	f, err := FlagAgent.New("10.32.0.1", 8443, 3, privateKey, publlicKey)
	if err != nil {
		log.Panic(err)
		return
	}
	f.StartService()
}