package stat

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	BASE_TEN         int = 10
	INTEGER_BIT_SIZE int = 64
)

type Stat struct {
	Name       string
	Value      int
	UserAgent  string
	Date       time.Time
	Referer    string
	AppVersion string
}

func FromRequest(request *http.Request) (s Stat) {
	s = Stat{}
	s.UserAgent = request.Header["User-Agent"][0]
	s.Referer = request.Header["Referer"][0]

	queryParameters, _ := url.ParseQuery(request.URL.RawQuery)
	for name, _ := range queryParameters {
		rawValue := queryParameters[name][0]

		if name == "date" {
			parsedInt := convertStringToInt(rawValue)
			s.Date = time.Unix(parsedInt, 0)
		} else if name == "appversion" {
			s.AppVersion = rawValue
		} else {
			s.Name = name
			parsedInt := convertStringToInt(rawValue)
			s.Value = int(parsedInt)
		}
	}
	return s
}

func convertStringToInt(raw string) int64 {
	parsedInt, err := strconv.ParseInt(raw, BASE_TEN, INTEGER_BIT_SIZE)
	if err != nil {
		log.Print(err)
	}
	return parsedInt
}

func (s *Stat) String() string {
	return fmt.Sprintf("Name: %s | Value: %d | User Agent: %s | Date: %s | Referer: %s", s.Name, s.Value, s.UserAgent, s.Date.Format(time.UnixDate), s.Referer)
}
