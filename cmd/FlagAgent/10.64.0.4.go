package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagAgent"
)

func main() {
	privateKey := "{\"D\": \"Age4NAroWGWlHuPF/hcUQP0DNOxY2UVV6zgG02PzmISqAzOKXsWkivvvldR4J16suxNwgQUVaRONqT+0XOQHe1nP97e8qh0y/wxt+++a+Y32fTKOVVSXFccjsuS0d1LIq45QSMqbtHfF5MrcuMAIPXbJao5pRDjEc8HtATRx1jBOtMc66b97Ox87jZQLG8zNLks/mNqZ1iYItDxPZrnncwhthcsloIsic80igTDHY1fOn2dsyyOFc8TsWd/i5GdmqTcl+O4KL/TMLcSjk0pbjPMFdNleMXvNOCj7nxCKkbKmNugVFw+LblxCgXAVlbz7iC++xYzZj49H4OngIjWg/NU=\", \"E\": \"AQABAA==\", \"N\": \"ArwAvQmlLZE4ds9IAHuDkui2E44CoJmn9eksLk7c8OYUWG4fO6YhWxY4S7ruxAzykAvBoxys9w3dQg5sURb2VSvKfbx+IXk1ZH/0nnNi89EzAYbB3jiMQmxmjipeLux/UZ8j+Ei6XuRZ27tEvlXnN7/2/r8yw/mYWyweXAkKnrGbqDO3ZuWAsNnW9on8uig5uc1T5C3IM4pEPFzU2c6yASMy8QsrdsqZiU49edTWaWAD3r6ql7JKNfG4cGy8rG54t37g12ORATDlYKbL9OYL0szPg7ePPL7Gn/II28EzfZ7vI4F2fAWmgN3ND3WP8v1DXcHh0FdseBniVinekihsSu0=\", \"Primes1\": \"AvZTdKETjNyoxakpE+Sya2WwxJq9FKfY2j4AAsKdOHflXEdt0RLiWDFK0UZNhZOAgvyZPBogArH2GMl7uRpcQOqcLS8+asHZCmU2Ux0eyeVHMMhQS413EcKr+D9N/Xrty1IpURKzIQnpaxK3thHrbmvxnVUWueyy1dJ348zQHQLD\", \"Primes2\": \"AsNi6ATa5ZjJBjCXi4Y5TcvXgRo8aIKq9YbHyJLuS3+MvsNtRX2dk7vw8B5JhSV6223Lm7g848onkCozcC8CzOiWSwEXoSiDUV4bxZ79sCgRnZpQbb9DQ3iCfpRA025cSmNwT1nZwIz1odl8lqNGagAOlThpNox2zJE6NDjz7UCP\"}"
	publlicKey := "{\"N\": \"AqRzRA81LnCcUJdJCASXw4XhAmQczECCcY4zDiYKggmsTGAebcOXaq9HelkpYp1cN7fNhYHO+AoIjAoiBrkv9JwugE2ro9Htl6bmuWJN7LMVIEyrLxGM0XLoC1mEjNNqidg54tVwhkW3wWGFky+WOeCElKo+NBnHdAe1ekSSeQLLsKPbg1WRPZtky7aseexr+NZqKbAHPLQdsLhIddlOMRNnNllWM0q1BFUl1IeRQCAivPtrGrAy1XqVRKeT8xrLAQxfXGokmK/Goo8HZW9RZipqfJ9xbO3sgfEmabOxBnxF1PlA+YvURwQrvKA41af09/aJ9WlMXEtgGzfadEFn6N0=\", \"E\": \"AQABAA==\"}"
	f, err := FlagAgent.New("10.64.0.4", 8443, 3, privateKey, publlicKey)
	if err != nil {
		log.Panic(err)
		return
	}
	f.StartService()
}