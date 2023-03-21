package dur

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

type Stamp int

const (
	HH Stamp = iota
	MM
	SS
	MMSS
	HHMMSS
	HHMMSSsss
	TimeStamp
	CueStamp
	FullStamp
)

func Parse(format Stamp, dur string) (time.Duration, error) {
	times, err := SplitTimestamp(format, dur)
	if err != nil {
		return time.Duration(0), fmt.Errorf("%w\n", err)
	}

	ds, err := FormatDuration(format, times...)
	if err != nil {
		return time.Duration(0), fmt.Errorf("%w\n", err)
	}

	def, err := time.ParseDuration(ds)
	if err != nil {
		return time.Duration(0), fmt.Errorf("%w\n", err)
	}

	return def, nil
}

func (format Stamp) Parse(dur string) (time.Duration, error) {
	return Parse(format, dur)
}

func Format(format Stamp, d time.Duration) (string, error) {
	times, err := SplitDurString(format, d.String())
	if err != nil {
		return "", fmt.Errorf("%w\n", err)
	}

	ds, err := FormatDuration(format, times...)
	if err != nil {
		return "", fmt.Errorf("%w\n", err)
	}

	return ds, nil
}

func (ts Stamp) Format(d time.Duration) (string, error) {
	return Format(ts, d)
}

func FormatStamp(format Stamp, times ...int) (string, error) {
	var ts []any
	for _, t := range times {
		ts = append(ts, t)
	}

	var buf bytes.Buffer
	_, err := fmt.Fprintf(&buf, format.StampFmt(), ts...)
	if err != nil {
		return "", fmt.Errorf("%w\n", err)
	}
	return buf.String(), nil
}

func FormatDuration(format Stamp, times ...int) (string, error) {
	var ts []any
	for _, t := range times {
		ts = append(ts, t)
	}

	var buf bytes.Buffer
	_, err := fmt.Fprintf(&buf, format.DurFmt(), ts...)
	if err != nil {
		return "", fmt.Errorf("%w\n", err)
	}
	return buf.String(), nil
}

func SplitTimestamp(format Stamp, dur string) ([]int, error) {
	var ts []int
	switch format {
	case HH:
		fallthrough
	case MM:
		fallthrough
	case SS:
		t, err := strconv.Atoi(dur)
		if err != nil {
			return ts, fmt.Errorf("%w\n", err)
		}
		ts = append(ts, t)
	case MMSS:
		var mm, ss int
		_, err := fmt.Sscanf(dur, format.ScanFmt(), &mm, &ss)
		if err != nil {
			return ts, fmt.Errorf("%w\n", err)
		}
		ts = append(ts, mm, ss)
	case HHMMSS:
		var hh, mm, ss int
		_, err := fmt.Sscanf(dur, format.ScanFmt(), &hh, &mm, &ss)
		if err != nil {
			return ts, fmt.Errorf("%w\n", err)
		}
		ts = append(ts, hh, mm, ss)
	case HHMMSSsss:
		var hh, mm, ss, ms int
		_, err := fmt.Sscanf(dur, format.ScanFmt(), &hh, &mm, &ss, &ms)
		if err != nil {
			return ts, fmt.Errorf("%w\n", err)
		}
		ts = append(ts, hh, mm, ss, ms)
	}
	return ts, nil
}

func SplitDurString(format Stamp, dur string) ([]int, error) {
	var ts []int
	switch format {
	case HH:
		var hh int
		_, err := fmt.Sscanf(dur, format.DurFmt(), &hh)
		if err != nil {
			return ts, fmt.Errorf("%w\n", err)
		}
		ts = append(ts, hh)
	case MM:
		var mm int
		_, err := fmt.Sscanf(dur, format.DurFmt(), &mm)
		if err != nil {
			return ts, fmt.Errorf("%w\n", err)
		}
		ts = append(ts, mm)
	case SS:
		var ss int
		_, err := fmt.Sscanf(dur, format.DurFmt(), &ss)
		if err != nil {
			return ts, fmt.Errorf("%w\n", err)
		}
		ts = append(ts, ss)
	case MMSS:
		var mm, ss int
		_, err := fmt.Sscanf(dur, format.DurFmt(), &mm, &ss)
		if err != nil {
			return ts, fmt.Errorf("%w\n", err)
		}
		ts = append(ts, mm, ss)
	case HHMMSS:
		var hh, mm, ss int
		_, err := fmt.Sscanf(dur, format.DurFmt(), &hh, &mm, &ss)
		if err != nil {
			return ts, fmt.Errorf("%w\n", err)
		}
		ts = append(ts, hh, mm, ss)
	case HHMMSSsss:
		var hh, mm, ss, ms int
		_, err := fmt.Sscanf(dur, format.DurFmt(), &hh, &mm, &ss, &ms)
		if err != nil {
			return ts, fmt.Errorf("%w\n", err)
		}
		ts = append(ts, hh, mm, ss, ms)
	}
	return ts, nil
}

const (
	DurHH        = "%vh"
	DurMM        = "%vm"
	DurSS        = "%vs"
	DurMMSS      = "%vm%vs"
	DurHHMMSS    = "%vh%vm%vs"
	DurHHMMSSsss = "%vh%vm%v.%vs"
	DurCuestamp  = "%vm%vs"
	DurTimestamp = "%vh%vm%vs"
	DurFullStamp = "%vh%vm%v.%vs"
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
	case CueStamp:
		return DurCuestamp
	case TimeStamp:
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
	case CueStamp:
		return ScanCuestamp
	case TimeStamp:
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
	StampHHMMSSsss = "%02v:%02v:%02v.%03v"
	StampCuestamp  = "%02v:%02v"
	StampTimestamp = "%02v:%02v:%02v"
	StampFullStamp = "%02v:%02v:%02v.%03v"
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
	case CueStamp:
		return StampCuestamp
	case TimeStamp:
		return StampTimestamp
	case FullStamp:
		return StampFullStamp
	}
	return ""
}
