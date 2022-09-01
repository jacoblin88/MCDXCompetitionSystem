package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagAgent"
)

func main() {
	privateKey := "{\"D\": \"AkpaYnxWc3g7EMMublMot0xQEXQrEmuZMQC36hVOmos7nSSFELcAW3TjqrTDNVhzhau9S+2PweKFVL1UA69rL3cXZoWsJ67lSsm/oCAFeV0CinRg6yoTxmeWWwmieUDAzBjfUm43iG9HDZIFAQou4/WOJjesBK/7vwAvQvDgm1cbcjBWRP76xBBZ11+MDAV8SN5bLGtB6LXInyyS9Qzb/1Xb8iDb3j/1onL1apFRUTBH7HflGO1h+E6iOSrL05Dp+Lnjg6WkYTOVHHpP4u4Gyj0kHHpbk1QLvlNDnvNe1j9siwCQRNeEo9nj8dO0Syjela1So0Xm3rF2Wl63hn2EVgE=\", \"E\": \"AQABAA==\", \"N\": \"AskeFv+5gKrl0VgD/RnCjciQmg13aTCsOs7+pT3ezO765pHEx+gXGg7YqLhUt9ea+HrZSzTLUslQPWUBxnPPmLt37cu6Mfdo4kSj4AC8QX4ElHW9PXjZ/+4LzE/PLIxLPOMzsPvTe33dld3mRdTpZrG4AHvwnrL4DednzNzdH9phMjIfSuUWCcFiULluSSbX2KkDSEiRrQ/FLcOJGN/N2ByGwnD4ZZ+rzDIWw0V63H61bWUJQFmhYOMikr+wmhoY13uEizf5rlEdQOGD8szmINsLqhZZzQRoURQiCrCj/Zj6q7P28ROvvRZD/2KBspiV44ybcjTZ2WEPNRexZkPLkZs=\", \"Primes1\": \"AuSk5Yy/8YW0aVEB+SjdBnOAh6//iNEXlqpFadfwNZ1h9k4uUlSztVBT/dOaQ8LUo5LH67nVJk0XbQCUteSEKopD5sF+NAx+RSdZ/h2Eyn6bNhJyI9gLzqOz8rTfYOSj2+oWlTJaoR1sCFM2brWRjqX2nd/HI6grr178gBcnGssB\", \"Primes2\": \"AuEuFfaYhjUMb07Bcsj6IzwHcWQAO4Bw1GMrbq5eTZnKjtkn3ULOEAyEs3+cNoOHIJXdGJ09VNfBEwyVaMJnoCkP+ILj+3d/STz74XJzGnVK2c/pA9iuA+ROuFnVIaZpDQW15vyWk15/8K3bGKHKI9YoYzr1cGi3zHnoagujWqib\"}"
	publlicKey := "{\"N\": \"AqRzRA81LnCcUJdJCASXw4XhAmQczECCcY4zDiYKggmsTGAebcOXaq9HelkpYp1cN7fNhYHO+AoIjAoiBrkv9JwugE2ro9Htl6bmuWJN7LMVIEyrLxGM0XLoC1mEjNNqidg54tVwhkW3wWGFky+WOeCElKo+NBnHdAe1ekSSeQLLsKPbg1WRPZtky7aseexr+NZqKbAHPLQdsLhIddlOMRNnNllWM0q1BFUl1IeRQCAivPtrGrAy1XqVRKeT8xrLAQxfXGokmK/Goo8HZW9RZipqfJ9xbO3sgfEmabOxBnxF1PlA+YvURwQrvKA41af09/aJ9WlMXEtgGzfadEFn6N0=\", \"E\": \"AQABAA==\"}"
	f, err := FlagAgent.New("10.96.0.4", 8443, 3, privateKey, publlicKey)
	if err != nil {
		log.Panic(err)
		return
	}
	f.StartService()
}