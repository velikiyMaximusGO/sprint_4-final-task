package daysteps

import (
	"errors"
	"fmt"
	"log"
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
	// TODO: реализовать функцию
	buf := strings.Split(data, ",")
	if len(buf) != 2 {
		return 0, 0, errors.New("the slice length should be equal to two")
	}
	steps, err := strconv.Atoi(buf[0])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, errors.New("the number of steps cannot be a negative number or zero")
	}
	periodWalk, err2 := time.ParseDuration(buf[1])
	if err2 != nil {
		return 0, 0, err2
	}
	if periodWalk <= 0 {
		return 0, 0, errors.New("the length of a walk cannot be a negative number or zero")
	}
	return steps, periodWalk, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, periodWalk, err := parsePackage(data)

	if err != nil {
		log.Println(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distance := float64(steps) * stepLength / mInKm
	sumCalories, err2 := spentcalories.WalkingSpentCalories(steps, weight, height, periodWalk)

	if err2 != nil {
		log.Println(err2)
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, sumCalories)
}
