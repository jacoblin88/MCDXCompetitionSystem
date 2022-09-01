package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagAgent"
)

func main() {
	privateKey := "{\"D\": \"Aqzu3M9X2fHzDJXklkC3RXU29mHmoCcB/ocR6zpUsLTrsmkMaVG89KFwhpiOcRgzMGDb6NJtgFiy8A6TF+u1uDpJnlU6im8mbdNNqS8wTgZp+qKlIRZ5OO1+btUcVzV3la63RBHBFVS4qWTWGgMhEXEQciea8YRiob3f6f0I9k3+cQ6C39O7aN8tdS7AlGGHt04EBzt/C96mDpNTy91M2NpGh3fq6YGxZAk8RDKbWshDuDohyhZVYBaqeHUIHAbyZSxH6oRg6Xju5KmXjpUEXvEvMOW1ci8c3KHQCWJdDnf45dpYinRHp3Wh+UJjTxhcZRMagD8pXp3onPSoTJgraWk=\", \"E\": \"AQABAA==\", \"N\": \"ArwWFRt17iN6m90DvCp9vQmmakIEWcTCjWtHrVwehFRkr5V6NXJFsYYiIeHDiMg9c8rPdbsS/4mrKUnuuCYogS/+iomlRtm0JugGEaJ/XaKTgLfw0sX637GFAhneokYRsgzk1ZmxedUf2exPebg48utmEO9C7k+OtfsnwJmRtya+QPu0kBwtL0VcCxbpLsY/1GgxmRoSTVrmYWKQKj0Mqh+Cd4KgBKkb5VuiWGH3jx56XHcsogGrY0Kuv+xHktmR7Ry68v/L7OqXSMJiw652EbGuQ4tp3q2SAuZZCK+CWEmcs6AxVQR9Jvh4zLvA3LNUOjPjdfP8keSRaHEKaf5cIqM=\", \"Primes1\": \"AuzbP7QghZSHi2+a0eCUuE0ixRwxi9ng4hGYjI64SL38SsbALql7L7131RIEEcmTAImrAOU5SYt2OTU7XCeOQwTvm6qaOPFDrviATq4usGPOnfs0Fgzufa14+VZzzMj03+OMFCp2YNSsPdbxMVA273aBvvXd86JHW4SDm6EuTGQP\", \"Primes2\": \"AstJvTjX7rz+j2VlhjCekfjKmna7qQGuxCIjP1K6q861ORofuY2HmGlBDwfkwjOqsQWN/GZqQ47eiLTB3lvsR+ruPpYc3kd997lz0HkdYdCdp+J1gdBO3DQ6Ehnx9mEmaAxP9MkOya1XXAWsu+wG4hRlzbt8QB21dJ6buaDuLLQt\"}"
	publlicKey := "{\"N\": \"AqRzRA81LnCcUJdJCASXw4XhAmQczECCcY4zDiYKggmsTGAebcOXaq9HelkpYp1cN7fNhYHO+AoIjAoiBrkv9JwugE2ro9Htl6bmuWJN7LMVIEyrLxGM0XLoC1mEjNNqidg54tVwhkW3wWGFky+WOeCElKo+NBnHdAe1ekSSeQLLsKPbg1WRPZtky7aseexr+NZqKbAHPLQdsLhIddlOMRNnNllWM0q1BFUl1IeRQCAivPtrGrAy1XqVRKeT8xrLAQxfXGokmK/Goo8HZW9RZipqfJ9xbO3sgfEmabOxBnxF1PlA+YvURwQrvKA41af09/aJ9WlMXEtgGzfadEFn6N0=\", \"E\": \"AQABAA==\"}"
	f, err := FlagAgent.New("10.32.0.3", 8443, 3, privateKey, publlicKey)
	if err != nil {
		log.Panic(err)
		return
	}
	f.StartService()
}