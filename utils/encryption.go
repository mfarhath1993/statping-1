// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/statping/statping
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package utils

import (
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

// HashPassword returns the bcrypt hash of a password string
func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

// NewSHA1Hash returns a random SHA1 hash based on a specific length
func NewSHA256Hash() string {
	d := make([]byte, 10)
	rand.Seed(Now().UnixNano())
	rand.Read(d)
	return fmt.Sprintf("%x", sha256.Sum256(d))
}

var characterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandomString generates a random string of n length
func RandomString(n int) string {
	b := make([]rune, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = characterRunes[rand.Intn(len(characterRunes))]
	}
	return string(b)
}
