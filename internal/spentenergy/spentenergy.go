package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

var ErrWrongInfo = errors.New("incorrect data")

// Возвращает два значения:
// float64 — количество калорий, потраченных при ходьбе.
// error — ошибку, если входные параметры некорректны (подумайте, какие значения параметров имеют смысл).
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, ErrWrongInfo
	}

	speed := MeanSpeed(steps, height, duration)

	minutes := duration.Minutes()

	return weight * speed * minutes / minInH * walkingCaloriesCoefficient, nil
}

// Возвращает два значения:
// float64 — количество калорий, потраченных при беге.
// error — ошибку, если входные параметры некорректны (подумайте, какие значения параметров имеют смысл).
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, ErrWrongInfo
	}

	speed := MeanSpeed(steps, height, duration)

	minutes := duration.Minutes()

	return weight * speed * minutes / minInH, nil
}

// Функция принимает количество шагов steps, рост пользователя height и
// продолжительность активности duration  и возвращает среднюю скорость.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	dist := Distance(steps, height)

	return dist / duration.Hours()
}

// Функция принимает количество шагов и рост пользователя в метрах, а возвращает дистанцию в километрах.
func Distance(steps int, height float64) float64 {
	return height * stepLengthCoefficient * float64(steps) / mInKm
}
