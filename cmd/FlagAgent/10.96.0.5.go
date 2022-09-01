package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagAgent"
)

func main() {
	privateKey := "{\"D\": \"AkJLbnGDPLBXk/CvHB1M5dWUsTuatfYnX7tXkqMPHn/kvwNRZkqpp+iKqkjOR+9M7s02TdEUSX4qFrA7Kbtx+FYegENk6/TRuHmuEeILZFAobzpbqV1kxyl/HwpD7+HoEMGG3UNKpdW1L75Zw6Us14D5QQc9QHEiKbzBdw047yqU4JWvZM2+uATDR+kiVToVDHxPrEBRv4/jVti9vIuNwXh9GCteM04/1ixFvZsotVngZDe+AUzo62jum13roCaescA+uiPGEIvEjEC9Q+SQeNqqobFRm73nJrD0ewkQE7dGoolcrkn2Rmy4QaptY4Tp/DyhLIJ3szQPfvaQhQQGkgE=\", \"E\": \"AQABAA==\", \"N\": \"ApiN9QrNNx2FgO/rZ6mAUNloL+URtDbqVgdmhpFGMsJ5X9tUXXKfzrB9lJj1OFsU+w63LqoDsMbvgRl5CtiN/Au1/W2jOuj3xyZQpKzY8CPIMTYsniPLZxqE93SFaZ3XnEZlhrT/OBfUU+EFMgo3MnYVTarXzKREH+saWnN8zbHN3oFPChRNnlv6L2kHP47eFQ7Umq3pY0LBFUhTr4TBmm9ebQKPdSaFNl7DgwH9knLkhyJNHzTdpmaC1P9FjZ4wRp2ribccej08OZCXgeYiuk6n60dh6EG/eNYRilTBPjPzuHeKL9qWkoDnown+4FBceuLsSCeuX0Ua8yH65P9gMuc=\", \"Primes1\": \"Ash5PXvX+blGjlDwOd7164SbAsRq0qKK19OvIwrnVqI3u/6Z+9SSkDM4q6XJleX19qRYfEq5CHnzW1gTcyerkY3dXlm7RumCxpCzacm4fqVc7FqbrXVrcOOg8PgmFMexq8yMgUrYP0zAoBBZbuegbCjwgMsxrNaPjgu5pSptKJD7\", \"Primes2\": \"AsLO+N45tRDGEylvN9GsIIcXhYkH2llw1n/Dk2ZmrG8+jP8D/TQGjbWqKCu881GSBOoM0OR+VGE3F1dqUiDAAXNwehwdMt2n3GtC6vH6QgUGH7JcbiWcLXaAPqpdvzBCMBpfP58NhJKGMCUTFnY4u/WODn+EWaUUco7Cg8p+jboF\"}"
	publlicKey := "{\"N\": \"AqRzRA81LnCcUJdJCASXw4XhAmQczECCcY4zDiYKggmsTGAebcOXaq9HelkpYp1cN7fNhYHO+AoIjAoiBrkv9JwugE2ro9Htl6bmuWJN7LMVIEyrLxGM0XLoC1mEjNNqidg54tVwhkW3wWGFky+WOeCElKo+NBnHdAe1ekSSeQLLsKPbg1WRPZtky7aseexr+NZqKbAHPLQdsLhIddlOMRNnNllWM0q1BFUl1IeRQCAivPtrGrAy1XqVRKeT8xrLAQxfXGokmK/Goo8HZW9RZipqfJ9xbO3sgfEmabOxBnxF1PlA+YvURwQrvKA41af09/aJ9WlMXEtgGzfadEFn6N0=\", \"E\": \"AQABAA==\"}"
	f, err := FlagAgent.New("10.96.0.5", 8443, 3, privateKey, publlicKey)
	if err != nil {
		log.Panic(err)
		return
	}
	f.StartService()
}