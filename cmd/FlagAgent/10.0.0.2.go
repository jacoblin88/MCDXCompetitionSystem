package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagAgent"
)

func main() {
	privateKey := "{\"D\": \"Aj+uHQ4fuj+H+KKXbLDSDE4ac4gsqzxj0Ra2VZZ10HIPSNWxwy/rs8v8FixQh/AQ2vrsxn87J3fjREt2+CU1MqT7ZWptnpWkOiNUpegnW4TtaSkJ9VIcxlPY4sNsbDs27DdTZefK+y5+culXknGGGEjbkE48hT808lhIVqbbPXte5dy7wMHUZn2qNn4ABrC0/eDsLrqa7no9hsMkeTPOyR7Kd6wveDYuJsetvcdMQ7tuQwjATDI/Eg4PEqLIWO4Xk2FPKJliAzxC7cNuSMAXbgDHk0nROa+GxNDOGdp3KfU+Ppw/m1DqCilyEOZyA6fav9tHsNtp0vi5NzyTKeRWM6k=\", \"E\": \"AQABAA==\", \"N\": \"Ask8D+bgo4XSpoz+IJDzfB6lmeYnhJL6uQRNB5Oc/QxaV/mDtTd/LPcGSPUVl7OTjzk1/1Vg3wSDvHzpPtH8d8M+bY/uI3ecKO1ItV3KyzVsauEErspXn1SO7NVwyCoCF2g2+F+zZ+Req5orOcVGPynB7LESLRnph1ELzMmfsnaWSRU02APvyoyLd3y8Pmm98JhV4s8VpJ9ph5JyYJ5KsgnUM2jbEFqQppIZII5xH5+Fs64NQbgRpEI3oSb1s0yqwHyboUFupuDjI8wadcCUKRcEmg+BkUcyV6ME3uuLBdr/6uVi2TfQqr7WQ52qx5VZ4OddcJsfSh+WOXBjW8WS1k8=\", \"Primes1\": \"AuPHkCeVNAL5/wVppuZQt1m8p4PNQam6LMIrxKpMpkq+BD8D0DCAxxmUoUo6achAtq0eePvcMoychboq2SjC1zv2InL5JFmbf1F51S3BpNtyElxZY0TaF8P0gebKVrgjxNIi+xj2qPOc8rkAChsn8tA1gqFogZ0skQ7NqfUNdDVr\", \"Primes2\": \"AuIqlD22q/m2aUnLNfkNlj2Dm1EDBGqeL5Z5Q/uD60VLP+VNah1wUQzWgR07J7s2JeJFB8NT7MQECaimYiO7pvzH6f4lmKpj1xz8MjIBuOcQYV9Of1PijOOzNzTnb0cIXWCk5pPY23EGgdwnNrZafJvfLLvteChhp/ulIFfzInet\"}"
	publlicKey := "{\"N\": \"AqRzRA81LnCcUJdJCASXw4XhAmQczECCcY4zDiYKggmsTGAebcOXaq9HelkpYp1cN7fNhYHO+AoIjAoiBrkv9JwugE2ro9Htl6bmuWJN7LMVIEyrLxGM0XLoC1mEjNNqidg54tVwhkW3wWGFky+WOeCElKo+NBnHdAe1ekSSeQLLsKPbg1WRPZtky7aseexr+NZqKbAHPLQdsLhIddlOMRNnNllWM0q1BFUl1IeRQCAivPtrGrAy1XqVRKeT8xrLAQxfXGokmK/Goo8HZW9RZipqfJ9xbO3sgfEmabOxBnxF1PlA+YvURwQrvKA41af09/aJ9WlMXEtgGzfadEFn6N0=\", \"E\": \"AQABAA==\"}"
	f, err := FlagAgent.New("10.0.0.2", 8443, 3, privateKey, publlicKey)
	if err != nil {
		log.Panic(err)
		return
	}
	f.StartService()
}