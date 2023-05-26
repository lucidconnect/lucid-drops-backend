package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/agnivade/levenshtein"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	space = regexp.MustCompile(`\s+`)
)

func StringToUUID(id string) uuid.UUID {
	newFormatedId, _ := uuid.FromString(id)
	return newFormatedId
}
func StringToFloat64(id string) float64 {
	newFormatedId, _ := strconv.ParseFloat(id, 64)
	return newFormatedId
}

// StrCompareDistance checks and returns the edit distance between [s1] and [s2]
func StrCompareDistance(s1, s2 string) int {
	reg, _ := regexp.Compile("[^a-zA-Z0-9 ]+")
	s1 = strings.ToLower(reg.ReplaceAllString(s1, ""))
	s2 = strings.ToLower(reg.ReplaceAllString(s2, ""))
	s1Arr := strings.SplitAfter(s1, " ")
	s2Arr := strings.SplitAfter(s2, " ")
	var nameFound int
	for i := 0; i < len(s1Arr); i++ {
		for j := 0; j < len(s2Arr); j++ {
			distance := levenshtein.ComputeDistance(s1Arr[i], s2Arr[j])
			if distance < 3 {
				nameFound++
				s2Arr = append(s2Arr[:j], s2Arr[j+1:]...)
				break
			}
		}
	}
	return nameFound
}

func GetStringPtr(ss ...string) *string {
	var acc string
	for _, s := range ss {
		acc += s
	}

	if acc == "" {
		return nil
	} else {
		return &acc
	}
}

func QualifyRef(qualifier interface{}, ref interface{}) string {
	return fmt.Sprintf("%v:%v", qualifier, ref)
}

func GetStrPtr(in string) *string {
	if in != "" {
		return &in
	}
	return nil
}

func GetIntPtr(in int) *int {
	return &in
}

func MaskedBVNPhoneNumber(phoneNumber string) (*string, error) {
	length := len(phoneNumber)

	if length < 11 {
		return nil, errors.New("the passed Phone Number should be greater than 11")
	}

	runeStr := []rune(phoneNumber)

	for i := 6; i <= len(runeStr)-4; i++ {
		runeStr[i] = 'x'
	}
	out := string(runeStr)
	return &out, nil
}

func CustomToTitleCase(in string) string {
	caser := cases.Title(language.English)
	return caser.String(in)
}

func RemoveWhiteSpace(in string) string {
	return space.ReplaceAllString(in, " ")
}

func ToFirstNamePlusInitials(in string) string {

	in = strings.Replace(in, ",", "", -1)
	names := strings.Split(in, " ")
	if len(names) < 2 {
		return strings.ToTitle(in)
	}

	if len(names[1]) < 1 {
		return strings.ToTitle(in)
	}

	return	fmt.Sprintf("%s %s.", names[0], strings.ToUpper(string(names[1][0:1])))
}
