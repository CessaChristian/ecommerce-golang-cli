package utils

import "time"

var wibLocation, _ = time.LoadLocation("Asia/Jakarta")

// Convert UTC → WIB
func ToWIB(t time.Time) time.Time {
    return t.In(wibLocation)
}

// Format WIB dengan format default
func FormatWIB(t time.Time) string {
    return t.In(wibLocation).Format("02 Jan 2006 15:04")
}

func FormatWIBPtr(t *time.Time) string {
    if t == nil {
        return "-"
    }
    return t.In(wibLocation).Format("02 Jan 2006 15:04")
}
