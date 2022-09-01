package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagAgent"
)

func main() {
	privateKey := "{\"D\": \"Ail90LGyi6I+Kl2aanWVCaZExOU55hwytQa+rlrLtvEG3UvOGzAQ0FNwl9ZxgNYpsw+JMrK8N3xhXRlpYhz0oCXAkDrmDPZrupI4jCqwZ4uQjAOkI5qE2trSBB1i1xuBkpg1Gcg5jrPo0SHJP92NeTgL5FRqnBa89Wb1pMdkibuo/dqpOUJLEnmTsY1oH7X5EXVh/Ar4NMooVpXcPhlSZ+STRJ896mp5nWdQGR7/fcGHa816M+7XalBEMninJCmCE45ChY37dswYyUs3GkYbcJWTaByUIOJCwrJ8ONg3gyeZWNtJcpDiPgXsruPI0X7sevMpn8T45MtZ9AE1IMlQ8wU=\", \"E\": \"AQABAA==\", \"N\": \"ArFeKgm0pYaRQER5w6cXJcbD/Ut47xYa+KkMMPvX4vUgUP9WuvSyAmAYxJ8wh5Vw6rr9nbacKkowA+uaVV3oAaSKBudxD6dxv3ozCUE3Sqg4GVFOWh1L2XSEkN0RDp3gRm3nBPK2wrEo6UipA2czU5ECSz9I/wmNWkCbJ74H3fBjm8Esse6a5vSjIRFMniuNiUhsLu6S3St6yO/ktukvlS4nE59A6qIXdCNPRvuavWoBIMvtXDybwfRIeDfl9N4VROPPnq4YX2G3+p82Xg2ibHyxCmrdTFNXml2VG5tVIk7NxXUrlB1RiMfbpvN0i8Mgu8Gi9SwjdjpypXIzE6CE+/U=\", \"Primes1\": \"AthC1dyE0HX8DhURCL3KoONcK/gdSqEevMQzcJ3tAkyqnPXVm6uyWr0aLHVOaNo3SsFn5EvCDusiCL6OCW2NOtuqNE7Oy9oP5iDzJqY5+/h7NAPu/6p8aC1AqwhORJuCwa4WkiP7ms+NAe4x5UZ7X+NdUbODadOZyKQBD8yHYCMP\", \"Primes2\": \"AtH1vyYKSc1ierYv9ekgeRsrh5M54CAXRtTCSLdkMr1vVP8BH+ZYlajxkQUWs1YDuIEjtsCLMUN/md3RTx1a3jvd+38nM3XhkJINtLAkX8IWGBcIMX/edJQMjS2GzQr+qvNGLnjK4uiVSvSLZUUeG8CaV9p4krt8s0zy0FI6fqC7\"}"
	publlicKey := "{\"N\": \"AqRzRA81LnCcUJdJCASXw4XhAmQczECCcY4zDiYKggmsTGAebcOXaq9HelkpYp1cN7fNhYHO+AoIjAoiBrkv9JwugE2ro9Htl6bmuWJN7LMVIEyrLxGM0XLoC1mEjNNqidg54tVwhkW3wWGFky+WOeCElKo+NBnHdAe1ekSSeQLLsKPbg1WRPZtky7aseexr+NZqKbAHPLQdsLhIddlOMRNnNllWM0q1BFUl1IeRQCAivPtrGrAy1XqVRKeT8xrLAQxfXGokmK/Goo8HZW9RZipqfJ9xbO3sgfEmabOxBnxF1PlA+YvURwQrvKA41af09/aJ9WlMXEtgGzfadEFn6N0=\", \"E\": \"AQABAA==\"}"
	f, err := FlagAgent.New("10.0.0.6", 8443, 3, privateKey, publlicKey)
	if err != nil {
		log.Panic(err)
		return
	}
	f.StartService()
}