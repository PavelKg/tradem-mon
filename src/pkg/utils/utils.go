package utils

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
	"unicode"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func ParseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func FormatIntToString(n int64, thousandsSeparator byte) string {
	in := strconv.FormatInt(n, 10)
	numOfDigits := len(in)
	if n < 0 {
		numOfDigits-- // First character is the - sign (not a digit)
	}
	numOfCommas := (numOfDigits - 1) / 3

	out := make([]byte, len(in)+numOfCommas)
	if n < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = thousandsSeparator
		}
	}
}

func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func CreateUniqueCSVFilename(path string) string {
	return path + GenerateRandomString(10) + ".csv"
}

func RemoveFile(file string) {
	err := os.Remove(file)
	if err != nil {
		log.Println(err)
	}
}

func ParseDelimiter(name string) rune {
	var delimiter rune
	switch name {
	case "comma":
		delimiter = ','
	default:
		delimiter = ';'
	}
	return delimiter
}

func Round(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(f*shift+.5) / shift
}

func SameStringSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	// create a map of string -> int
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		// 0 value for int is 0, so just increment a counter for the string
		diff[_x]++
	}
	for _, _y := range y {
		// If the string _y is not in diff bail out early
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y] -= 1
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	return len(diff) == 0
}

func JwtGenSignedToken(subject string, secret string) (string, error) {
	const jwtLifeTime = 72
	// Create the Claims
	claims := jwt.MapClaims{
		"sub": subject,
		"exp": time.Now().Add(time.Hour * jwtLifeTime).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("token not signed")
	}
	return t, nil

}

func GetUserIdFromJwt(ctx *fiber.Ctx) (string, error) {
	token := ctx.Locals(JwtContextKey).(*jwt.Token)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if val, ok := claims["sub"]; ok {
			return val.(string), nil
		}
	}
	return "", fmt.Errorf("jwt error format")
}

func RemoveSpace(s string) string {
	rr := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsSpace(r) {
			rr = append(rr, r)
		}
	}
	return string(rr)
}
