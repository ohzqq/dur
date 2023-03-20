package dur

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type Duration time.Duration

type Stamp int
type DurFmt int
type ScanFmt int
type StampFmt int

const (
	HH Stamp = iota
	MM
	SS
	MMSS
	HHMMSS
	HHMMSSsss
	Timestamp
	Cuestamp
	FullStamp
)

const (
	nano   = 1000000000
	ssNano = nano
	mmNano = 60 * nano
	hhNano = 60 * 60 * nano
)

func Parse(format Stamp, dur string) Duration {
	ds := "0s"
	switch format {
	case HH:
		h, _ := strconv.Atoi(dur)
		hh := int64(h * hhNano)
		return Duration(time.Duration(hh))
	case MM:
		m, _ := strconv.Atoi(dur)
		mm := int64(m * mmNano)
		return Duration(time.Duration(mm))
	case SS:
		s, _ := strconv.Atoi(dur)
		ss := int64(s * ssNano)
		return Duration(time.Duration(ss))
	case MMSS:
		var mm, ss int64
		_, err := fmt.Sscanf(dur, MMSS, &mm, &ss)
		if err != nil {
			log.Fatal(err)
		}
		s := ss * ssNano
		m := mm * mmNano
		return Duration(time.Duration(m + s))
	case HHMMSS:
		var hh, mm, ss int64
		_, err := fmt.Sscanf(dur, HHMMSS, &hh, &mm, &ss)
		if err != nil {
			log.Fatal(err)
		}
		s := ss * ssNano
		m := mm * mmNano
		h := hh * hhNano
		return Duration(time.Duration(h + m + s))
	case HHMMSSsss:
		var hh, mm int64
		var ss float64
		_, err := fmt.Sscanf(dur, HHMMSSsss, &hh, &mm, &ss)
		if err != nil {
			log.Fatal(err)
		}
		println(0.033 * ssNano)
		s := int64(ss * ssNano)
		m := mm * mmNano
		h := hh * hhNano
		return Duration(time.Duration(h + m + s))
	}
	def, err := time.ParseDuration(ds)
	if err != nil {
		log.Fatal(err)
	}
	return Duration(def)
}

func formatHH(hh int) string {
	return strconv.Itoa(hh) + "h"
}

func formatMM(dur int) string {
	return strconv.Itoa(dur) + "m"
}

func formatSS(dur int) string {
	return strconv.Itoa(dur) + "s"
}

func formatSSsss(dur float64) string {
	return strconv.FormatFloat(dur, 'f', 3, 64) + "s"
}
