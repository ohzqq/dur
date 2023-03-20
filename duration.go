package dur

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type Duration time.Duration

type Stamp int

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
	return format.Parse(dur)
}

func (format Stamp) Parse(dur string) Duration {
	ds := "0s"
	switch format {
	case HH:
		fallthrough
	case MM:
		fallthrough
	case SS:
		ds = fmt.Sprintf(format.DurFmt(), dur)
	case MMSS:
		var mm, ss int64
		_, err := fmt.Sscanf(dur, format.ScanFmt(), &mm, &ss)
		if err != nil {
			log.Fatal(err)
		}
		ds = fmt.Sprintf(format.DurFmt(), mm, ss)
	case HHMMSS:
		var hh, mm, ss int64
		_, err := fmt.Sscanf(dur, format.ScanFmt(), &hh, &mm, &ss)
		if err != nil {
			log.Fatal(err)
		}
		ds = fmt.Sprintf(format.DurFmt(), hh, mm, ss)
	case HHMMSSsss:
		var hh, mm, ss, ms int64
		_, err := fmt.Sscanf(dur, format.ScanFmt(), &hh, &mm, &ss, &ms)
		if err != nil {
			log.Fatal(err)
		}
		ds = fmt.Sprintf(format.DurFmt(), hh, mm, ss, ms)
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

const (
	DurHH        = "%vh"
	DurMM        = "%vm"
	DurSS        = "%vs"
	DurMMSS      = "%vm%vs"
	DurHHMMSS    = "%vh%vm%vs"
	DurHHMMSSsss = "%vh%vm%vs%vms"
	DurCuestamp  = "%vm%vs"
	DurTimestamp = "%vh%vm%vs"
	DurFullStamp = "%vh%vm%vs%vms"
)

func (ts Stamp) DurFmt() string {
	switch ts {
	case HH:
		return DurHH
	case MM:
		return DurMM
	case SS:
		return DurSS
	case MMSS:
		return DurMMSS
	case HHMMSS:
		return DurHHMMSS
	case HHMMSSsss:
		return DurHHMMSSsss
	case Cuestamp:
		return DurCuestamp
	case Timestamp:
		return DurTimestamp
	case FullStamp:
		return DurFullStamp
	}
	return ""
}

const (
	ScanHH        = "%d"
	ScanMM        = "%d"
	ScanSS        = "%d"
	ScanMMSS      = "%d:%d"
	ScanHHMMSS    = "%d:%d:%d"
	ScanHHMMSSsss = "%d:%d:%d.%d"
	ScanCuestamp  = "%d:%d"
	ScanTimestamp = "%d:%d:%d"
	ScanFullStamp = "%d:%d:%d.%d"
)

func (ts Stamp) ScanFmt() string {
	switch ts {
	case HH:
		return ScanHH
	case MM:
		return ScanMM
	case SS:
		return ScanSS
	case MMSS:
		return ScanMMSS
	case HHMMSS:
		return ScanHHMMSS
	case HHMMSSsss:
		return ScanHHMMSSsss
	case Cuestamp:
		return ScanCuestamp
	case Timestamp:
		return ScanTimestamp
	case FullStamp:
		return ScanFullStamp
	}
	return ""
}

const (
	StampHH        = "%02v"
	StampMM        = "%02v"
	StampSS        = "%02v"
	StampMMSS      = "%02v:%02v"
	StampHHMMSS    = "%02v:%02v:%02v"
	StampHHMMSSsss = "%02v:%02v:%02v.%03d"
	StampCuestamp  = "%02v:%02v"
	StampTimestamp = "%02v:%02v:%02v"
	StampFullStamp = "%02v:%02v:%02v.%03d"
)

func (ts Stamp) StampFmt() string {
	switch ts {
	case HH:
		return StampHH
	case MM:
		return StampMM
	case SS:
		return StampSS
	case MMSS:
		return StampMMSS
	case HHMMSS:
		return StampHHMMSS
	case HHMMSSsss:
		return StampHHMMSSsss
	case Cuestamp:
		return StampCuestamp
	case Timestamp:
		return StampTimestamp
	case FullStamp:
		return StampFullStamp
	}
	return ""
}
