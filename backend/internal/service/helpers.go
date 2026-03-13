package service

import (
	"strings"
	"time"
)

const dateLayout = "2006-01-02"

func cleanString(value string) string {
	return strings.TrimSpace(value)
}

func cleanOptionalString(value *string) *string {
	if value == nil {
		return nil
	}
	trimmed := strings.TrimSpace(*value)
	if trimmed == "" {
		return nil
	}
	return &trimmed
}

func parseRequiredDate(value string) (time.Time, error) {
	return time.Parse(dateLayout, strings.TrimSpace(value))
}

func parseOptionalDate(value *string) (*time.Time, error) {
	if value == nil || strings.TrimSpace(*value) == "" {
		return nil, nil
	}

	parsed, err := time.Parse(dateLayout, strings.TrimSpace(*value))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

func parseDateTime(value string) (time.Time, error) {
	trimmed := strings.TrimSpace(value)

	if parsed, err := time.Parse(time.RFC3339, trimmed); err == nil {
		return parsed, nil
	}

	parsed, err := time.Parse(dateLayout, trimmed)
	if err != nil {
		return time.Time{}, err
	}
	return parsed, nil
}
