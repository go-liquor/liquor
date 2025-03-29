package lqstring

import (
	"math/rand"
	"net/mail"
	"net/url"
	"regexp"
	"time"

	pluralize "github.com/gertd/go-pluralize"
	"github.com/golang-cz/textcase"
	"github.com/google/uuid"
)

var (
	pl                = pluralize.NewClient()
	numericRegex      = regexp.MustCompile(`^[0-9]+$`)
	alphanumericRegex = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	letterRunes       = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

// ToSnakeCase converts a string to snake_case format.
//
// Parameters:
//   - str: The string to convert
//
// Returns:
//   - string: The converted string in snake_case
func ToSnakeCase(str string) string {
	return textcase.SnakeCase(str)
}

// ToCamelCase converts a string to camelCase format.
//
// Parameters:
//   - str: The string to convert
//
// Returns:
//   - string: The converted string in camelCase
func ToCamelCase(str string) string {
	return textcase.CamelCase(str)
}

// ToKebabCase converts a string to kebab-case format.
//
// Parameters:
//   - str: The string to convert
//
// Returns:
//   - string: The converted string in kebab-case
func ToKebabCase(str string) string {
	return textcase.KebabCase(str)
}

// ToPascalCase converts a string to PascalCase format.
//
// Parameters:
//   - str: The string to convert
//
// Returns:
//   - string: The converted string in PascalCase
func ToPascalCase(str string) string {
	return textcase.PascalCase(str)
}

// ToPlural converts a singular word to its plural form.
//
// Parameters:
//   - str: The singular word
//
// Returns:
//   - string: The plural form of the word
func ToPlural(str string) string {
	return pl.Plural(str)
}

// ToSingular converts a plural word to its singular form.
//
// Parameters:
//   - str: The plural word
//
// Returns:
//   - string: The singular form of the word
func ToSingular(str string) string {
	return pl.Singular(str)
}

// IsPlural checks if a word is in plural form.
//
// Parameters:
//   - str: The word to check
//
// Returns:
//   - bool: true if the word is plural, false otherwise
func IsPlural(str string) bool {
	return pl.IsPlural(str)
}

// IsSingular checks if a word is in singular form.
//
// Parameters:
//   - str: The word to check
//
// Returns:
//   - bool: true if the word is singular, false otherwise
func IsSingular(str string) bool {
	return pl.IsSingular(str)
}

// IsEmail validates if a string is a valid email address.
//
// Parameters:
//   - str: The string to validate
//
// Returns:
//   - bool: true if the string is a valid email address, false otherwise
func IsEmail(str string) bool {
	_, err := mail.ParseAddress(str)
	return err == nil
}

// IsURL validates if a string is a valid URL.
//
// Parameters:
//   - str: The string to validate
//
// Returns:
//   - bool: true if the string is a valid URL, false otherwise
func IsURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// IsNumeric checks if a string contains only numeric characters.
//
// Parameters:
//   - str: The string to check
//
// Returns:
//   - bool: true if the string contains only numbers, false otherwise
func IsNumeric(str string) bool {
	return numericRegex.MatchString(str)
}

// IsAlphanumeric checks if a string contains only letters and numbers.
//
// Parameters:
//   - str: The string to check
//
// Returns:
//   - bool: true if the string contains only letters and numbers, false otherwise
func IsAlphanumeric(str string) bool {
	return alphanumericRegex.MatchString(str)
}

// RandomString generates a random string of the specified length.
//
// Parameters:
//   - length: The desired length of the random string
//
// Returns:
//   - string: A random string containing letters
func RandomString(length int) string {
	b := make([]rune, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letterRunes[r.Intn(len(letterRunes))]
	}
	return string(b)
}

// UUID generates a new UUID v4.
//
// Returns:
//   - string: A new UUID string
func UUID() string {
	return uuid.New().String()
}
