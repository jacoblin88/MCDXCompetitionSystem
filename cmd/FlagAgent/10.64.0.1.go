package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagAgent"
)

func main() {
	privateKey := "{\"D\": \"Ap8TFlwI9/gzDhqNqvLrKf9ltCyCF9HPZrdeEQjN4v3LkW2ZuUlSwGLf9hlx6NiLK1S5pGVa1b4B+tuJcl8BJC/YapVPz/77y6hzBV4KZGURFAvQ5EHB/P4NDUmiaDCZ5FKDOCuDDdosI7M3lTYZPaBDcxRU5vCOuqfEi+uG2tDWP6QRjhAYDcJNijRDu/VetQ1zxBfsdO5xjGaVMxAIEpeqRj2NcD89hwPXkXJ9RPxaZXGhITZjywHTjXWdHOYHk1B33vfyqVdv51QH47NkzZmX3teYypEz8M+vo+tAm9xMej9MusS6Em2ri6cCCS0sp03Ph5+1fgIIacjXxD1avgE=\", \"E\": \"AQABAA==\", \"N\": \"AqB8zj8telRo8m8VtlrimsbPBV5Ka8ZlnShfn43eH5K4rsJggAMS9FtuEJ7NUD8Bx+Ggooy74dT25F9a2lOvzsIxieP7Wts5Cvz3237tWHaAGVXmNgxB0tp1k+NvbIcyRgn+BjzD8j5GG519oxyu/XPFj9t1xD8Dcvkv32a56KMhrjyaA1GO5x31GG2QiKkiti6zYrzt2eUFNG2AlYCzzUMcfiUa+JC3V9ICKY9072N2DXtTeVMi4u0FZleMP2CXcHmS+ja0ExCUC5sSCVKxMyyPDv4xCSYXuCLOG3ad9T8Urfr4LiHfltMpKs8UchykpIGzdE5Y3s9NWHzj6EXSDds=\", \"Primes1\": \"AsUuURT4lFDPp/JC4JAi9s+96rTRu7f3pOFsNAWQIxBZB2tu9fElHTPldwnMK5SL81Cc1Y/67QoBpo/o3R8qdvC49ojs1DiNFOnKYyXzXVrAn48qSDjAvmKuh4aWSCAtGr04X8enh1brk7cJ7CR+VMkPmdWn/3VvAXmmlKNRv7RN\", \"Primes2\": \"AtBcZ4EBTzbdM1+qQQ6/EUAOQi2/Hmixs5N6Ow7pBJ8HGX7BarZopbe3taowsMwcvyL4MSnRwlOZyLVbGHq+r0eqXzgZYRwM2vGgUtyUw5LQ7UJaHyiBmmXVgOR4h75BvRE/62cqb51SBvngV8hDDK/VgnUZn6m5jkhyrYv+y37H\"}"
	publlicKey := "{\"N\": \"AqRzRA81LnCcUJdJCASXw4XhAmQczECCcY4zDiYKggmsTGAebcOXaq9HelkpYp1cN7fNhYHO+AoIjAoiBrkv9JwugE2ro9Htl6bmuWJN7LMVIEyrLxGM0XLoC1mEjNNqidg54tVwhkW3wWGFky+WOeCElKo+NBnHdAe1ekSSeQLLsKPbg1WRPZtky7aseexr+NZqKbAHPLQdsLhIddlOMRNnNllWM0q1BFUl1IeRQCAivPtrGrAy1XqVRKeT8xrLAQxfXGokmK/Goo8HZW9RZipqfJ9xbO3sgfEmabOxBnxF1PlA+YvURwQrvKA41af09/aJ9WlMXEtgGzfadEFn6N0=\", \"E\": \"AQABAA==\"}"
	f, err := FlagAgent.New("10.64.0.1", 8443, 3, privateKey, publlicKey)
	if err != nil {
		log.Panic(err)
		return
	}
	f.StartService()
}