package restserver

import (
	"strconv"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func BenchmarkBcrypt(b *testing.B) {

	for _, password := range []string{
		"$2y$05$z/OEmNQamd6m6LSegUErh.r/Owk9Xwmc5lxDheIuHY2Z7XiS6FtJm",
		"$2y$12$gSmDNRXxnRHbFt7UbMz4UeK29wvJJZ33dc9x32y5ATiCXCom0UtrC",
	} {
		cost, err := bcrypt.Cost([]byte(password))
		if err != nil {
			b.Fatal(err)
		}

		b.Run(strconv.FormatInt(int64(cost), 10), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				err := bcrypt.CompareHashAndPassword([]byte(password), []byte("test"))
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
