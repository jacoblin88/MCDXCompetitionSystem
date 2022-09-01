package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagAgent"
)

func main() {
	privateKey := "{\"D\": \"Ahi2OksI00et8uZM1mS1I6nm1uo3wHUPRIBdcYB+kQywK+82VTTwDm3gZzL0cdfHAnfsnMl59Qu5AixKDFL9i5IQLF0rW4q5pleroiD3Qm8lDaRQ7X6i4HqYpiuU/klQAhMxSRTjG/y00jQVhVZ8/BMv9mkD3Aqrk3YmeiHNgrIWKTqs0r/EVnkDFZVxrMzASXEgNz3/O+IpCEd83P1Cpi097mNCv+K4XQn7tCWZ8c37aLwcYe8oFAqf7yiw9rMQk7DqquL6EFZj20GgSCNZsUZG/KpXJmZ1OK6zCDWB0tR0YzBrU1Y9iwCb4DUxIGC2O7gHvYzZzNbd+9UsJXUa8gE=\", \"E\": \"AQABAA==\", \"N\": \"AsRnKfNOXkT1dmoqvWy00dY6g8oQWc7wln8ZFJy4FbVnphLSYFeWTU0EsLCy+F6nF/76XzAU0QyrDoHmj9WWhnnwLhzs9dHXbQKbaCUMlfYsezsA/NHgzZDkhII9hx5u/8n2EvFqI+H5Itxp2MJKtm2HeC0GLgaR2uxgnzWSvF7odavEqwkBF9lg7mlbXcaSjOMGPDJyVpi4Ct5Zoy5k2yH5R27a1dRIghRQQQbWVkSjNo0bt9HMFuL0Bz1N+lBEe5CHWXxlmv44/wrE0cv5drg9UDg2fCUPzBk8Clh2tI21vE5u7BDbDWpRGi7Lub5WH48M1YJXg7WsNyxjxv3qgVk=\", \"Primes1\": \"At9I8ygbQ+XAlCWwyJ88w9ooqJaR3gU26L7lg2Z7ijH3pI6JbW1c0FDMfhuFxVH9cn9ywG2avd0hMnRtoPhTCyrhOHD4aCUVO4HbWh66EUqjnwhozCGvdTgRHrwPubbmA3pf83M8r/wUDIItIi2i1HbejiCHuKkv1v7WwEobQsa5\", \"Primes2\": \"AuEt6l2YBSbBJEBad3qZDElBDRqZfHqRKBKlad++uYWXPBVvuux87P4bZykAuN3IvGWno+9EN+nTBKNeup/ijB7nXXGPklA0+7MIr8BuTZvoJzktfRTuElDaHTcrbi/wnXfdBNL/evCcR2yyNPRM+Utb4dEWdROAV+Y292o4JD+h\"}"
	publlicKey := "{\"N\": \"AqRzRA81LnCcUJdJCASXw4XhAmQczECCcY4zDiYKggmsTGAebcOXaq9HelkpYp1cN7fNhYHO+AoIjAoiBrkv9JwugE2ro9Htl6bmuWJN7LMVIEyrLxGM0XLoC1mEjNNqidg54tVwhkW3wWGFky+WOeCElKo+NBnHdAe1ekSSeQLLsKPbg1WRPZtky7aseexr+NZqKbAHPLQdsLhIddlOMRNnNllWM0q1BFUl1IeRQCAivPtrGrAy1XqVRKeT8xrLAQxfXGokmK/Goo8HZW9RZipqfJ9xbO3sgfEmabOxBnxF1PlA+YvURwQrvKA41af09/aJ9WlMXEtgGzfadEFn6N0=\", \"E\": \"AQABAA==\"}"
	f, err := FlagAgent.New("10.96.0.1", 8443, 3, privateKey, publlicKey)
	if err != nil {
		log.Panic(err)
		return
	}
	f.StartService()
}