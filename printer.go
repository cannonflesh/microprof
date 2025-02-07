package microprof

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/cpu"
)

type logger interface {
	Infof(string, ...any)
	Errorf(string, ...any)
}

type units string

const (
	UnitsBytes units = "bytes"
	UnitsKb    units = "Kb"
	UnitsMb    units = "Mb"
	UnitsGb    units = "Gb"
)

func PrintProfilingInfo(l logger, u units, byCPU bool) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	l.Infof("Allocated Memory: %s\n", formatMemoryUsage(memStats.Alloc, u))
	l.Infof("Total Allocated Memory: %s\n", formatMemoryUsage(memStats.TotalAlloc, u))
	l.Infof("Heap Memory: %s\n", formatMemoryUsage(memStats.HeapAlloc, u))
	l.Infof("Heap System Memory: %s\n", formatMemoryUsage(memStats.HeapSys, u))
	l.Infof("Garbage Collector Memory: %s\n", formatMemoryUsage(memStats.GCSys, u))
	percentage, err := cpu.Percent(0, byCPU)
	if err != nil {
		l.Errorf("Getting CPU load: %v\n", err)
	}
	if len(percentage) > 0 {
		l.Infof("CPU usage: %s\n", formatCpuUsage(percentage))
	}
}

func unitsFactor(u units) float32 {
	kb := float32(1024)
	mb := kb * kb
	gb := mb * kb

	values := map[units]float32{
		UnitsBytes: 1,
		UnitsKb:    kb,
		UnitsMb:    mb,
		UnitsGb:    gb,
	}

	return values[u]
}

func refineUnits(u units) units {
	if u != UnitsGb && u != UnitsMb && u != UnitsKb && u != UnitsBytes {
		return UnitsBytes
	}

	return u
}

func formatMemoryUsage(v uint64, u units) string {
	u = refineUnits(u)
	if u == UnitsBytes {
		return fmt.Sprintf("%df Bytes", v)
	}

	return fmt.Sprintf("%0.4f %s", float32(v)/unitsFactor(u), u)
}

func formatCpuUsage(perc []float64) string {
	if len(perc) == 1 {
		return strconv.FormatFloat(perc[0], 'f', 4, 64)
	}

	res := make([]string, 0, len(perc))
	for i, v := range perc {
		res = append(res, strconv.Itoa(i)+": "+strconv.FormatFloat(v, 'f', 4, 64))
	}

	return strings.Join(res, " | ")
}
