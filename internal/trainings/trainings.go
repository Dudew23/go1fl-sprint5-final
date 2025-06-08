package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

var (
	ErrWrongInfo     = errors.New("incorrect data")
	ErrWrongTrain    = errors.New("неизвестный тип тренировки")
	ErrStepsLessZero = errors.New("steps cannot be less than zero")
	ErrDurLessZero   = errors.New("duration cannot be less than zero")
)

// Метод парсит строку с данными формата "3456,Ходьба,3h00m" и записывает данные в соответствующие поля структуры Training.
func (t *Training) Parse(datastring string) (err error) {

	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return ErrWrongTrain
	}

	stepsStr := parts[0]
	if stepsStr == "" {
		return ErrWrongTrain
	}
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return ErrWrongTrain
	}
	if steps <= 0 {
		return ErrWrongTrain
	}
	t.Steps = steps

	trainingType := parts[1]
	if trainingType == "" {
		return ErrWrongTrain
	}
	t.TrainingType = trainingType

	durationStr := parts[2]
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return ErrWrongTrain
	}
	if duration <= 0 {
		return ErrWrongTrain
	}
	t.Duration = duration

	return nil
}

// Метод формирует и возвращает строку с данными о тренировке, исходя из того, какой тип тренировки был передан.
func (t Training) ActionInfo() (string, error) {
	var info string

	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	if t.TrainingType == "Плавание" {
		return "", ErrWrongTrain
	}

	switch t.TrainingType {
	case "Бег":
		calories, _ := spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		info = fmt.Sprintf(
			"Тип тренировки: Бег\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.Duration.Hours(), distance, speed, calories,
		)
	case "Ходьба":
		calories, _ := spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		info = fmt.Sprintf(
			"Тип тренировки: Ходьба\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.Duration.Hours(), distance, speed, calories,
		)
	case "Плавание":
		return "", ErrWrongTrain
	default:
		return "", ErrWrongTrain
	}

	return info, nil
}
