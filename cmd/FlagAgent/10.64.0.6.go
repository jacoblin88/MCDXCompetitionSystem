package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagAgent"
)

func main() {
	privateKey := "{\"D\": \"AkkcFRReGNLu5OWLaxlZDdPRBitQRlF78ZdbPgQgTVkNcdlZuhMHOObxpXlL5Cs/wU4h8/Kr+GD3ZA80ouC8Yao+KJ5pICFYiocnjXzXdZ+j7CNg78CPkHzOHgA3n/xRLFwpLeWzVqFn44k6IGM+1LygLQ/R49oTHmacanRoC8j/rUIRL0COnhHWDjIH8U7Y+7kBjKQG3OeLfO97Dgb0rkDQULTokoQm23nbFvfrZJt6R7rLd0neD6+NOjw5GLBsLPALGWpHv9zhRPeBxaK9sQ1NO+RIDkAuHjTCUAAqW0EnXNN6xOHhEex0F4VWumRqyfPFtK/O512mU5uhT9nSIGE=\", \"E\": \"AQABAA==\", \"N\": \"ArQnkApZX+shElTI1TFwJgKHs0TJxC9NYVOOs6T3CxAow1kDOofFSKK77rPRT6cFZUdMclGL6EDVz8gtwq3x7m4Jf8FV4kzbRJG2MqJLkJEDdoqAjMcgPH2A0PkV1SWM9mMoVNkCY+uArs9ZdKyitZvdcQDgHf4kWbn0ExrcKA1ckrd4Ssz48e0aYCOJQx1uKp6cqiCRkJFv6xeNAd1Lnlwlqstj+IXst/xR7Qf9e7UVHj4paJYFgIKN4ft0YZlu8MefrDBrlYCW0OeLwScsDtjLGOAXVT9lQmL+BPpyg7TUnLy5+krWzMB9+YAVY25cRsqIjPKd4zGvbS7mqtKXSB8=\", \"Primes1\": \"AsHTq8pNEaFV+QoDL7Jh5pPybS9PPXedviGCw9j6cPEV6E6gNiOGh+Ky8D82YcX8+YklnEjXfoY7tGlEfKamBwoFPH4znp29DFDCaU64pWCsD2+iSTPnOgL3oYY6Yo2OwHGiovmLEDiL1Jfg4Px8K5KFl+sxCWZQ6hX9Ox0+xiNL\", \"Primes2\": \"Au3xLI1D+faHnNnDk/DYYlUVoVpYZf3wV02Teg/hBL94sO2xzu4YD3pMqwR1OBmas/FMRGmAVV9AhsCBzYX1F3UjQheknqryO6tMYVkrYpFE96LCLPmRzYygc1umqrTuV1XZrgr8Azi10Qnmhv3X/1w6n+1cDdqp46EtFFxNm9X9\"}"
	publlicKey := "{\"N\": \"AqRzRA81LnCcUJdJCASXw4XhAmQczECCcY4zDiYKggmsTGAebcOXaq9HelkpYp1cN7fNhYHO+AoIjAoiBrkv9JwugE2ro9Htl6bmuWJN7LMVIEyrLxGM0XLoC1mEjNNqidg54tVwhkW3wWGFky+WOeCElKo+NBnHdAe1ekSSeQLLsKPbg1WRPZtky7aseexr+NZqKbAHPLQdsLhIddlOMRNnNllWM0q1BFUl1IeRQCAivPtrGrAy1XqVRKeT8xrLAQxfXGokmK/Goo8HZW9RZipqfJ9xbO3sgfEmabOxBnxF1PlA+YvURwQrvKA41af09/aJ9WlMXEtgGzfadEFn6N0=\", \"E\": \"AQABAA==\"}"
	f, err := FlagAgent.New("10.64.0.6", 8443, 3, privateKey, publlicKey)
	if err != nil {
		log.Panic(err)
		return
	}
	f.StartService()
}