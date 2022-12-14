package repo

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"

	"blynker/internal/iface"
	"blynker/internal/model"
)

var _ iface.Repository = &Repo{}

type Repo struct {
	Data model.Sensor
}

func (r *Repo) Save(data *model.Sensor) error {
	r.Data = *data

	updAt := r.Data.UpdatedAt.Format(time.RFC3339)
	temp := strconv.FormatFloat(r.Data.Temperature, 'G', 2, 64)
	light := strconv.FormatInt(int64(r.Data.Light), 10)
	movement := strconv.FormatBool(r.Data.Movement)

	filename := "data.csv"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	defer file.Close()
	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)
	err = writer.WriteAll([][]string{{updAt, temp, light, movement}})
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) Get() *model.Sensor {
	return &r.Data
}
