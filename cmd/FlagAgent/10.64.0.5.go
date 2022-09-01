package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagAgent"
)

func main() {
	privateKey := "{\"D\": \"AkvnhDHByV/uj3fNmKtDjzHH5aQCz2blA/lStXSxelY3v0qvOS4BNTawJD+V7upmZT7boqJQkNC3rkkEP8stuZdkhFdYSQ90O9S4ZAOvHps2vmPDzGjuIFurOCAFSgouealMN5HEESP1O+eGlXfGGmTHakUpDknVjHvBSi/FSAqNOCtP99rpD3FQaSo5cdhJpFl9xXXqMfzgiESqS4miVnaV/xJPfFajXBr5Sj1rQ/B+Az7P7+CsTQl3R8FjxD5RadZc/SxL6Myoblg+DZ4mcB7D3V6BBzM/lJglVpkm7YB1E3EEcUlBbtO307nrO8rt7VU3lQw68fOMgMHnUnLCh+k=\", \"E\": \"AQABAA==\", \"N\": \"ArdFjvc4FMP8j67vxdWd6ppitHZGj6iEXspBk0lbmWt5GwtVBiH1CH8wAYGoLabYwIc3/X2hPh5koADHRukHmYI43p7aCrsKGCwt62vNu5FmVnkaRLmZQ9ugxv+DcqKfj6ZEbKt/3y2zBMBQQLj3lEgJRYtJs8OPrGlFyIijlCO1GTPxl6X1ptnomVrH0E0ekhHC3SRPoLUYJOvAoimasc3T5qNGV7zlv18zhQDCzrj2Jf2hUWy6dLLfV61rt0f5sosDEHhRHEkZowLDvWYBKOWrN2r97MFR9U1LQSfciWYlKgVDnOVxL/Lj5k9yD2NJLM2KDKkgpr/w+ZASCGPUI9M=\", \"Primes1\": \"AuQ3YTHryo+/nPO+Shj6KX/oTeszHV7EPZAqrTmilpS823Umwu5ryrrsILfBDrsTxVoLlVmv5tgg+5ks88l4iIDb+fzLHu/ZO9QthF1U4i7cYCbMpXyqg3dpbtD2t8IPoftIiSykGHTMOzIU3d/VtTZU8cZL+3Ik/jpyJXD11dGF\", \"Primes2\": \"As2VbdHeyq/v5vOOkHv0gWC3RKWMvCNVuBkYdoqML4t49UvhRkePab6pl+grTsTxAtlGHiBYEWGk1aa/NiXB4gM3eT2w+i186uFIvghXoW5aet9AHH5sIr6+bZBIaAKGzkvVJckJfui7b2mn91LCaBghftXlBoN5+xJbEWsXXXN3\"}"
	publlicKey := "{\"N\": \"AqRzRA81LnCcUJdJCASXw4XhAmQczECCcY4zDiYKggmsTGAebcOXaq9HelkpYp1cN7fNhYHO+AoIjAoiBrkv9JwugE2ro9Htl6bmuWJN7LMVIEyrLxGM0XLoC1mEjNNqidg54tVwhkW3wWGFky+WOeCElKo+NBnHdAe1ekSSeQLLsKPbg1WRPZtky7aseexr+NZqKbAHPLQdsLhIddlOMRNnNllWM0q1BFUl1IeRQCAivPtrGrAy1XqVRKeT8xrLAQxfXGokmK/Goo8HZW9RZipqfJ9xbO3sgfEmabOxBnxF1PlA+YvURwQrvKA41af09/aJ9WlMXEtgGzfadEFn6N0=\", \"E\": \"AQABAA==\"}"
	f, err := FlagAgent.New("10.64.0.5", 8443, 3, privateKey, publlicKey)
	if err != nil {
		log.Panic(err)
		return
	}
	f.StartService()
}