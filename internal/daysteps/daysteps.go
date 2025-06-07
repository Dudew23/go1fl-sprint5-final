package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

var (
	ErrWrongInfo     = errors.New("incorrect data")
	ErrStepsLessZero = errors.New("steps cannot be less than zero")
	ErrDurLessZero   = errors.New("duration cannot be less than zero")
)

// Метод парсит строку с данными формата "678,0h50m" и записывает данные в соответствующие поля структуры DaySteps.
func (ds *DaySteps) Parse(datastring string) (err error) {
	dataSlice := strings.Split(datastring, ",")

	if len(dataSlice) != 2 || datastring == "" {
		return ErrWrongInfo
	}

	ds.Steps, err = strconv.Atoi(dataSlice[0])
	if err != nil {
		return err
	}
	if ds.Steps <= 0 {
		return ErrStepsLessZero
	}

	ds.Duration, err = time.ParseDuration(dataSlice[1])
	if err != nil {
		return err
	}

	if ds.Duration <= 0 {
		return ErrDurLessZero
	}

	return nil
}

// Метод формирует и возвращает строку с данными о прогулке.
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories,
	), nil
}
