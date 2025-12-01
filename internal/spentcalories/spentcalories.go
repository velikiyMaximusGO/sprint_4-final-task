package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	buf := strings.Split(data, ",")

	if len(buf) != 3 {
		return 0, "", 0, errors.New("incorrect amount of data")
	}

	steps, err1 := strconv.Atoi(buf[0])

	if err1 != nil {
		return 0, "", 0, err1
	}

	if steps <= 0 {
		return 0, "", 0, errors.New("the number of steps cannot be a negative number or zero")
	}
	typeActivity := buf[1]
	durationActivity, err2 := time.ParseDuration(buf[2])

	if err2 != nil {
		return 0, "", 0, err2
	}

	if durationActivity <= 0 {
		return 0, "", 0, errors.New("the duration of activity cannot be a negative number or zero")
	}

	return steps, typeActivity, durationActivity, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	lenStep := height * stepLengthCoefficient
	return float64(steps) * lenStep / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}

	distance := distance(steps, height)

	speed := distance / duration.Hours()

	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, typeActivity, durationActivity, err := parseTraining(data)

	if err != nil {
		log.Println(err)
		return "", err
	}

	if steps <= 0 {
		return "", errors.New("the number of steps cannot be a negative number or zero")
	}

	if durationActivity <= 0 {
		return "", errors.New("the duration of activity cannot be a negative number or zero")
	}

	switch typeActivity {
	case "Ходьба":
		distance := distance(steps, height)
		speed := meanSpeed(steps, height, durationActivity)
		calories, err := WalkingSpentCalories(steps, weight, height, durationActivity)
		if err != nil {
			return "", errors.New("problem")
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", typeActivity, durationActivity.Hours(), distance, speed, calories), nil
	case "Бег":
		distance := distance(steps, height)
		speed := meanSpeed(steps, height, durationActivity)
		calories, err := RunningSpentCalories(steps, weight, height, durationActivity)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", typeActivity, durationActivity.Hours(), distance, speed, calories), nil
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("the number of steps cannot be a negative number or zero")
	}

	if weight <= 0 {
		return 0, errors.New("the weight cannot be a negative number or zero")
	}

	if height <= 0 {
		return 0, errors.New("height cannot be a negative number or zero")
	}

	if duration <= 0 {
		return 0, errors.New("the running time cannot be a negative number or zero")
	}

	speed := meanSpeed(steps, height, duration)

	minute := duration.Minutes()

	return (weight * speed * minute) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("the number of steps cannot be a negative number or zero")
	}

	if weight <= 0 {
		return 0, errors.New("the weight cannot be a negative number or zero")
	}

	if height <= 0 {
		return 0, errors.New("height cannot be a negative number or zero")
	}

	if duration <= 0 {
		return 0, errors.New("the running time cannot be a negative number or zero")
	}

	speed := meanSpeed(steps, height, duration)

	minute := duration.Minutes()

	return (weight * speed * minute) / minInH * walkingCaloriesCoefficient, nil
}
