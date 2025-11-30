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

	buf := strings.Split(data, ",")

	if len(buf) != 3 {
		return 0, "", 0, errors.New("Неправильное количество данных")
	}

	steps, err1 := strconv.Atoi(buf[0])

	if err1 != nil {
		return 0, "", 0, err1
	}

	if steps <= 0 {
		return 0, "", 0, err1
	}
	typeActivity := buf[1]
	durationActivity, err2 := time.ParseDuration(buf[2])

	if err2 != nil {
		return 0, "", 0, err2
	}

	return steps, typeActivity, durationActivity, nil
}

func distance(steps int, height float64) float64 {
	lenStep := height * stepLengthCoefficient
	return float64(steps) * lenStep / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	distance := distance(steps, height)

	speed := distance / duration.Hours()

	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, typeActivity, durationActivity, err := parseTraining(data)

	if err != nil || steps <= 0 || durationActivity <= 0 {
		log.Println(err)
	}

	switch typeActivity {
	case "Ходьба":
		distance := distance(steps, height)
		speed := meanSpeed(steps, height, durationActivity)
		calories, err := WalkingSpentCalories(steps, weight, height, durationActivity)
		if err != nil {
			return "", errors.New("Проблема")
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км.ч\nСожгли калорий: %.2f\n", typeActivity, durationActivity.Hours(), distance, speed, calories), nil
	case "Бег":
		distance := distance(steps, height)
		speed := meanSpeed(steps, height, durationActivity)
		calories, err := RunningSpentCalories(steps, weight, height, durationActivity)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", typeActivity, durationActivity.Hours(), distance, speed, calories), nil
	default:
		return "", errors.New("Неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("Входные параметры некорректны")
	}

	speed := meanSpeed(steps, height, duration)

	minute := duration.Minutes()

	return (weight * speed * minute) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("Входные параметры некорректны")
	}

	speed := meanSpeed(steps, height, duration)

	minute := duration.Minutes()

	return (weight * speed * minute) / minInH, nil
}
