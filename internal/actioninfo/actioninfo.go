package actioninfo

import (
	"fmt"
	"log/slog"
)

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, i := range dataset {
		err := dp.Parse(i)
		if err != nil {
			slog.Warn(":)")
		}
	}

	info, err := dp.ActionInfo()
	if err != nil {
		slog.Warn(":)")
	}

	fmt.Println(info)
}
