package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagAgent"
)

func main() {
	privateKey := "{\"D\": \"AgKnZmeTw7qcv2/zMeaAIbQyw5YOtOzKcpkX6nbtXpskIqZsDqE5HP1XhMGS9b2ZHzQOS3vDPKzAkaVvd88fWIxRgwUG82jPz6jOSPo4Othn15zskHl80NSFY3r90C3ESBuFJooSwlHcSBn/lE3G/S3qhL7Yya8+L30AKLajLmulg6zlB6NtchmucpjFcPXqYJ4BNik2vbm+o5xsLqwm+AYaV3WVEff+ze3/ifEx2JOdEH++ExOmmekIB91sQpCfjWpeocJ46vfY/x8sJzEvMEfUapoRI/vB8p9r1Jzkvq/lecpphrTk5GsO60PoQZi3oFPXIzYT0avxUbUoG5Ge1Mk=\", \"E\": \"AQABAA==\", \"N\": \"AsM04hzHxhHzX014V9TpXFLhxQxwVg0JTk5MTKsxkmQzFSJk3Qcb7lEracJutSuh/XcXtwmwQFsmU+XjDRlUFuOSoQ1+1R8X+5SG+hyBBjdczxa/wWEyXzLrBreuKdPiqE+pP0Yw9Pj9Zzo72VCIdqKI9AmifLRl66dno7ZqnInmjHrI7OBgZg7hrIjvYvYTzqnUmox4tLExrxe0Nl0OmzWm5J9VWSSEzXOJOXNAunHWCVrm1oNfygs+9ziI0L5ZN61t5GGsZHBVMiTlH+5nDT+yd6UPRCwk56QlKMaqq73Kw50fU2iQJBukk/5FunmHHq+asVvgM7zwImiNn6UqNO8=\", \"Primes1\": \"As9LGtMnfqBndj8PKaqu3Du2Nt3xFXyK/WWPWG0HU+2LR3fC554NQBtibySE92Hgl3kwyJFxqvAfW/Ee3TH5Bp8zyicAcqv30ewPjqlnJjDj1Tl4ywFeLOMsJDR2XvQw3GhdWgK6uENfMG+vmspzXzEIb0lthDmMtSbk/6vyHVHd\", \"Primes2\": \"AvESvuYpMu5qo1/oqoBvobp93XlAn5LNUX3Ohf6uCe1t7lkO2CrWr+L3ru3uonrp12I/yiM+ntqaBKt3A0FtFRk3QXCi05ilWC9WMUtVcCJ1hcY4+44cbSI+aPQbCgqwOW+uTAhxDBR3sBkm3jkCpOYMdYTjbrmcQ7PVx8ga8cM7\"}"
	publlicKey := "{\"N\": \"AqRzRA81LnCcUJdJCASXw4XhAmQczECCcY4zDiYKggmsTGAebcOXaq9HelkpYp1cN7fNhYHO+AoIjAoiBrkv9JwugE2ro9Htl6bmuWJN7LMVIEyrLxGM0XLoC1mEjNNqidg54tVwhkW3wWGFky+WOeCElKo+NBnHdAe1ekSSeQLLsKPbg1WRPZtky7aseexr+NZqKbAHPLQdsLhIddlOMRNnNllWM0q1BFUl1IeRQCAivPtrGrAy1XqVRKeT8xrLAQxfXGokmK/Goo8HZW9RZipqfJ9xbO3sgfEmabOxBnxF1PlA+YvURwQrvKA41af09/aJ9WlMXEtgGzfadEFn6N0=\", \"E\": \"AQABAA==\"}"
	f, err := FlagAgent.New("10.96.0.2", 8443, 3, privateKey, publlicKey)
	if err != nil {
		log.Panic(err)
		return
	}
	f.StartService()
}