package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagAgent"
)

func main() {
	privateKey := "{\"D\": \"AlUQERyCgsHVgLSKAITZah1e5zgbehpzCowMXrecELSi69K9+v2oniF78ocxQ7/TsVFbvCBZWfgtxGQOfyn6t0/eDqyVVZHHMOB/JksqJh75EkYF2ofMOfBkwRq4wP8scQ063COA8TvDyEKCQSB2O3anG4eYv4J0xxglvRO5YYJ+STuqiMfUfjo6FpWEvssTRLT5F6QazeaF5LDIaZUa0BFHnhi8/ULNTM4NbOLqQv5MSi+3CkOhOaOX34JqiKZO+F8dBTmm9PBQstBHIm9JTABGaOH6+pW3u9SFePtUid90/WXe3yGHbx68jlrPqCOP58nWm0N/dqA3HenTBciYfQk=\", \"E\": \"AQABAA==\", \"N\": \"AriD5RgWBieJvGsXB9P0tHfjGDMQs9ErhKgh6odfwgzTQc7r9WwjyfV+cf2kiBaitipKQSjPlBk/cYfS8zN5OD3MSMo832EVZnr62LZqC1AOYY1tRCnqK75oFIjXLq2X9+FCMhAGAVhBt9hMj32krO2TYB3/VN55Xa3ENOs8cxk6xWcJZxbXy5i6EEF+ZMuE3NM8xqoGFFHGM1dB/asgkYkmI/vhjAW4TEGuM6UkiIRYo9+P652Nuw/wwj+NvneDH55FXkYWnTuKiezq/zKglXN5DiaQqhZHUuS7BboCrOxqeiNHcdd982GnmH9xNKypqAfM2nwTY9aGapazbnWVwSM=\", \"Primes1\": \"AsR9M2pffPTOQ0kiQr0KlpE8ftAJ6WHXKFIUjq65X1E8G7MRbnJbVWx/2PL5yTI6iSt9z25IKR0Y2I7s0Jv6fcoLgptSMAPbGskhspg6STrKx8xqdLbfB7GLhqzxg8LDij5T08x0XI5BXNwwWE2jt5wlikGWQpHAG3K0A+kEyD2t\", \"Primes2\": \"AvBmTOEDq58s4rxy+qV+KWXCqy5+8b8ilzWwo5Q9972eT0YhulWQb8YHlQ1I/7nTkJy00AEPcV3tL1+thSBhP74Y/3o4xxUXWZpm7qmiKbCUn9wp3zFPT0aMOjS7Q8Usn0m5eFTqApbr6dxrUBa0ApDuiNBHCgn4iypPUhua/zQP\"}"
	publlicKey := "{\"N\": \"AqRzRA81LnCcUJdJCASXw4XhAmQczECCcY4zDiYKggmsTGAebcOXaq9HelkpYp1cN7fNhYHO+AoIjAoiBrkv9JwugE2ro9Htl6bmuWJN7LMVIEyrLxGM0XLoC1mEjNNqidg54tVwhkW3wWGFky+WOeCElKo+NBnHdAe1ekSSeQLLsKPbg1WRPZtky7aseexr+NZqKbAHPLQdsLhIddlOMRNnNllWM0q1BFUl1IeRQCAivPtrGrAy1XqVRKeT8xrLAQxfXGokmK/Goo8HZW9RZipqfJ9xbO3sgfEmabOxBnxF1PlA+YvURwQrvKA41af09/aJ9WlMXEtgGzfadEFn6N0=\", \"E\": \"AQABAA==\"}"
	f, err := FlagAgent.New("10.96.0.3", 8443, 3, privateKey, publlicKey)
	if err != nil {
		log.Panic(err)
		return
	}
	f.StartService()
}