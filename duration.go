package dur

import (
	"log"
	"time"
)

type Duration time.Duration

type DurFmt string

const (
	HH         DurFmt = "15"
	MM                = "04"
	SS                = "05"
	MMSS              = MM + ":" + SS
	HHMMSS            = HH + ":" + MMSS
	HHMMSSsss         = HHMMSS + ".000"
	Timestamp         = HHMMSS
	Cuestamp          = MMSS
	MilliStamp        = HHMMSSsss
)

func Parse(format, dur string) Duration {
	switch DurFmt(format) {
	case HH:
		d, err := time.ParseDuration(dur + "h")
		if err != nil {
			log.Fatal(err)
		}
		return Duration(d)
	case MM:
		d, err := time.ParseDuration(dur + "m")
		if err != nil {
			log.Fatal(err)
		}
		return Duration(d)
	case SS:
		//ss, err := strconv.Atoi(dur)
		//if err != nil {
		//  log.Fatal(err)
		//}
		d, err := time.ParseDuration(dur + "s")
		if err != nil {
			log.Fatal(err)
		}
		return Duration(d)
	case MMSS:
	case HHMMSS:
	case HHMMSSsss:
	}
	def, _ := time.ParseDuration("0s")
	return Duration(def)
}

func formatHH(hh string) string {
	//hh, err := strconv.Atoi(hh)
	//if err != nil {
	//  log.Fatal(err)
	//}
	return hh + "h"
}

func formatMM(dur string) string {
	//mm, err := strconv.Atoi(dur)
	//if err != nil {
	//  log.Fatal(err)
	//}
	return dur + "m"
}
