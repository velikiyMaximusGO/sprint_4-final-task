package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	buf := strings.Split(data, ",")
	steps, err := strconv.Atoi(buf[0])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, errors.New("Количество шагов должно быть больше нуля")
	}
	periodWalk, err2 := time.ParseDuration(buf[1])
	if err != nil {
		return 0, 0, err2
	}
	return steps, periodWalk, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, periodWalk, err := parsePackage(data)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distance := float64(steps) * stepLength / mInKm
	sumCalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, periodWalk)

	if err != nil {
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, distance, sumCalories)
}
