package logparse

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type EquipStatus struct {
	Status      bool
	Errcount    int
	Databases   int
	UploadCount int
	AvgProctime float64
}

func ParseJSONLog(filePath string) (st EquipStatus, err error) {

	f, err := os.Open(filePath)

	if err != nil {

		return
	}

	defer f.Close()

	var (
		dbTotal, dbProc, errs, total, athletes int
		avg                                    float64
	)

	s := bufio.NewScanner(f)
	for s.Scan() {
		var m map[string]any

		line := s.Text()

		err = json.Unmarshal([]byte(line), &m)

		if err != nil {

			continue
		}

		total++

		if m["level"] == "info" {
			switch m["msg"] {
			case "No data":
				err = fmt.Errorf("no log data")
				return
			case "Dados enviados com sucesso":
				dbProc++
				avg += m["duration"].(float64)
				athletes += m["athlete_count"].(int)
			case "Arquivos encontrados, iniciando MADB":
				dbTotal = m["databases"].(int)
			}
		}

		if m["level"] == "error" {
			errs++
		}
	}

	st.Status = true

	st.AvgProctime = avg / float64(dbProc)
	st.Databases = dbTotal
	st.UploadCount = athletes
	st.Errcount = errs

	if dbTotal == 0 {
		st.Status = false
	}

	if st.UploadCount == 0 {
		st.Status = false
	}

	if dbProc == 0 {
		st.Status = false
	}

	if errs > (total / 2) {
		st.Status = false
	}

	return
}
