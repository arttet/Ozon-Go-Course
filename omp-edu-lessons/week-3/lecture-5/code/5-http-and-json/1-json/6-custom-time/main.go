package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type myStruct struct {
	Time       time.Time  `json:"time"`
	CustomTime CustomTime `json:"custom_time"`
}

type CustomTime time.Time

const customTimeLayout = "2006-01-02 15:04:05"

func (t CustomTime) MarshalJSON() ([]byte, error) {
	s := `"` + time.Time(t).Format(customTimeLayout) + `"`
	return []byte(s), nil
}

func (t *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "null" {
		return nil
	}

	cTime, err := time.Parse(customTimeLayout, s)
	if err != nil {
		return fmt.Errorf("time.Parse(): %w", err)
	}

	*t = CustomTime(cTime)
	return nil
}

func main() {
	fmt.Println()

	var jsonBytes []byte

	{
		in := myStruct{
			Time:       time.Now(),
			CustomTime: CustomTime(time.Now()),
		}
		var err error
		jsonBytes, err = json.MarshalIndent(in, "", "\t")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonBytes))
	}

	fmt.Println()

	{
		out := myStruct{}
		err := json.Unmarshal(jsonBytes, &out)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Time", out.Time)
		fmt.Println("CustomTime", time.Time(out.CustomTime))
	}
}
