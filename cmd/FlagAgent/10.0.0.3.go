package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagAgent"
)

func main() {
	privateKey := "{\"D\": \"Apkzya45BjaqHwxJXZ70CkBLovtIaIRFoLWkDgTis56Y+ibXsOrZQSRfxP6vTcZwGnwGcwFqPazn32CdGQcNh6UgpY1WGAuw3krwJcgFYqfiKwE7WGVyUleMZ07mW9Q2wr2B6EGRuXe+pUrHkv2kNRqnK3CQn4YkTuyJRVCYNYhvBaYm7ek9cX4PMt7IOazSp4gLWp0Csr6bo0JDGOjxzU0sqn7a2Ko1Ono0UVGNL6Gjjad33tTbMpPAArpstohbPl8huQeBUwwAWRt+mKnQaS1a0DQypLGlUWbbhdyj6EbLvLVhZrzyi0rivRxSe7vLXWmpogETHiv5YwtqFiZw+FE=\", \"E\": \"AQABAA==\", \"N\": \"ArDgfm2p+0yjC/m0xpiOhApBPF48BlpXZU9RUvZf8bf/tvnlwsJR6EbTizmIe5hRk+sPg5B8nHyoj3NZjxVyW9KrKOiUMu7V0c895onOrwOLPqb5I43zKgImaNUIjvRz+EfH/vgqlh7bUTlKy1pzKok3kOB2jfj9sPb1bH3+RNBToBHaxJu+mwIkApo3SS2nRR1tAW1W3rLqENTH3+3SfWO0XRXkdEY5dyMZZeaKrLzEFtVRiNG6iAtzrGLrqIXEPQwRfvhQQ2SrWIGxk+te7+U94UA/ZjEFoc9yGtyZLHK+hC7A7T1vx5jPjiKPMGxkn8ZYaz8Wq2GWQlebmrud1Ok=\", \"Primes1\": \"AtwwPJmQRelu3M5DyMf9V8yDknY0G8johA1arCu789APZ9ZtC1o9g/sGp64j78uUrSfpV6t/7XghrnYIJPCruvlimgQCQFNoZxpOvE0WF460MWMt2sHvvSk3snYOAZMzjMJGPSGREDvjO/sAFwpTbI4ckEuuHVtC1uxiV5spVQN9\", \"Primes2\": \"As2k8KqAKC9/C0F2FxWOnkrKpV2WR9FEWxm00xFj26U2M1XskOSBt4ifuLWE9mWAqoit9YijzY+QWyJ5LMoFfmvyTxXKlNuAFjzftnDN3m1YvDqJh57K1SQlRGMaSTtZZXgj/WcDNou+UPy8gQ0LUmUwggdWG0xBOPO4ti84Rbrd\"}"
	publlicKey := "{\"N\": \"AqRzRA81LnCcUJdJCASXw4XhAmQczECCcY4zDiYKggmsTGAebcOXaq9HelkpYp1cN7fNhYHO+AoIjAoiBrkv9JwugE2ro9Htl6bmuWJN7LMVIEyrLxGM0XLoC1mEjNNqidg54tVwhkW3wWGFky+WOeCElKo+NBnHdAe1ekSSeQLLsKPbg1WRPZtky7aseexr+NZqKbAHPLQdsLhIddlOMRNnNllWM0q1BFUl1IeRQCAivPtrGrAy1XqVRKeT8xrLAQxfXGokmK/Goo8HZW9RZipqfJ9xbO3sgfEmabOxBnxF1PlA+YvURwQrvKA41af09/aJ9WlMXEtgGzfadEFn6N0=\", \"E\": \"AQABAA==\"}"
	f, err := FlagAgent.New("10.0.0.3", 8443, 3, privateKey, publlicKey)
	if err != nil {
		log.Panic(err)
		return
	}
	f.StartService()
}