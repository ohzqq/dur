package dur

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type Timestamp struct {
	HH  int
	MM  int
	SS  int
	MS  int
	Dur time.Duration
}

func Parse(e string) (Timestamp, error) {
	var hh string
	var mm string
	var ss string

	switch split := strings.Split(e, ":"); len(split) {
	case 3:
		hh = split[0] + "h"
		mm = split[1] + "m"
		ss = split[2] + "s"
	case 2:
		mm = split[0] + "m"
		ss = split[1] + "s"
	case 1:
		s := split[0]
		if s == "" {
			s = "0"
		}
		ss = s + "s"
	}

	ts := Timestamp{}

	stamp := fmt.Sprintf("%s%s%s", hh, mm, ss)
	d, err := time.ParseDuration(stamp)
	if err != nil {
		return ts, err
	}
	ts.Dur = d

	s := strings.Split(fmt.Sprintf("%06.3f", d.Seconds()), ".")
	secs, err := strconv.Atoi(s[0])
	if err != nil {
		return ts, err
	}
	ts.MS, err = strconv.Atoi(s[1])
	if err != nil {
		return ts, err
	}
	ts.HH = secs / 3600
	ts.MM = secs % 3600 / 60
	ts.SS = secs % 60

	return ts, err
}

func (d Timestamp) Secs() string {
	return fmt.Sprintf("%02d", d.RoundSecs())
}

func (d Timestamp) Min() string {
	return fmt.Sprintf("%02d", d.MM)
}

func (d Timestamp) Hours() string {
	return fmt.Sprintf("%02d", d.HH)
}

func (d Timestamp) MMSS() string {
	ts := fmt.Sprintf("%02d:%02d", d.HH*60+d.MM, d.RoundSecs())
	return ts
}

func (d Timestamp) Cuestamp() string {
	return d.MMSS()
}

func (d Timestamp) HHMMSS() string {
	ts := fmt.Sprintf("%02d:%02d:%02d", d.HH, d.MM, d.RoundSecs())
	return ts
}

func (d Timestamp) Timestamp() string {
	return d.HHMMSS()
}

func (d Timestamp) String() string {
	ts := fmt.Sprintf("%02d:%02d:%02d.%03d", d.HH, d.MM, d.SS, d.MS)
	return ts
}

func (d Timestamp) RoundSecs() int {
	secs := fmt.Sprintf("%02d.%03d", d.SS, d.MS)
	fs, err := strconv.ParseFloat(secs, 64)
	if err != nil {
		return 0
	}
	return int(math.Round(fs))
}
