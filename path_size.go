package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, humanReadable, withHidden, recursive bool) (string, error) {
	size, err := getSize(path, withHidden, recursive)
	if err != nil {
		return "", err
	}
	return formatSize(size, humanReadable), nil
}

func getSize(path string, withHidden, recursive bool) (int, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	// FILE
	if !info.IsDir() {
		if isHidden(info.Name()) && !withHidden {
			return 0, nil
		}
		return int(info.Size()), nil
	}

	// DIR
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	total := 0

	for _, e := range entries {
		name := e.Name()

		if isHidden(name) && !withHidden {
			continue
		}

		fullPath := filepath.Join(path, name)

		if e.IsDir() {
			if !recursive {
				continue
			}

			s, err := getSize(fullPath, withHidden, recursive)
			if err != nil {
				return 0, err
			}
			total += s
		} else {
			s, err := getSize(fullPath, withHidden, recursive)
			if err != nil {
				return 0, err
			}
			total += s
		}
	}

	return total, nil
}

func isHidden(name string) bool {
	return strings.HasPrefix(name, ".")
}

func formatSize(sizeInBytes int, humanReadable bool) string {
	if !humanReadable {
		return fmt.Sprintf("%dB", sizeInBytes)
	}

	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
	)

	switch {
	case sizeInBytes >= EB:
		return fmt.Sprintf("%.1fEB", float64(sizeInBytes)/EB)
	case sizeInBytes >= PB:
		return fmt.Sprintf("%.1fPB", float64(sizeInBytes)/PB)
	case sizeInBytes >= TB:
		return fmt.Sprintf("%.1fTB", float64(sizeInBytes)/TB)
	case sizeInBytes >= GB:
		return fmt.Sprintf("%.1fGB", float64(sizeInBytes)/GB)
	case sizeInBytes >= MB:
		return fmt.Sprintf("%.1fMB", float64(sizeInBytes)/MB)
	case sizeInBytes >= KB:
		return fmt.Sprintf("%.1fKB", float64(sizeInBytes)/KB)
	default:
		return fmt.Sprintf("%dB", sizeInBytes)
	}
}
