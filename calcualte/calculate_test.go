package calcualte

import (
	"testing"
	"time"
)


func Test_calculateSecond_Input_Start_5_2_1996_End_20_7_2019_Should_Be_740188800(t *testing.T){
	expected := 740188800
	StartDate := time.Date(1996, 2, 5, 0, 0, 0, 0, time.UTC)
	EndDate := time.Date(2019, 7, 20, 0, 0, 0, 0, time.UTC)
	actual := calculateSecond(StartDate,EndDate)
	if actual != expected {
		t.Errorf("expecred %d but it got %d", expected, actual)
	}
}

func Test_calculateMinutes_Input_Start_5_2_1996_End_20_7_2019_Should_Be_12336480(t *testing.T){
	expected := 12336480
	StartDate := time.Date(1996, 2, 5, 0, 0, 0, 0, time.UTC)
	EndDate := time.Date(2019, 7, 20, 0, 0, 0, 0, time.UTC)
	actual := calculateMinutes(StartDate,EndDate)
	if actual != expected {
		t.Errorf("expecred %d but it got %d", expected, actual)
	}
}

func Test_calculateHour_Input_Start_5_2_1996_End_20_7_2019_Should_Be_205608(t *testing.T){
	expected := 205608
	StartDate := time.Date(1996, 2, 5, 0, 0, 0, 0, time.UTC)
	EndDate := time.Date(2019, 7, 20, 0, 0, 0, 0, time.UTC)
	actual := calculateHour(StartDate,EndDate)
	if actual != expected {
		t.Errorf("expecred %d but it got %d", expected, actual)
	}
}
func Test_calculateSecond_Input_Start_2_8_1995_End_20_7_2019_Should_Be_756345600(t *testing.T){
	expected := 756345600
	StartDate := time.Date(1995, 8, 2, 0, 0, 0, 0, time.UTC)
	EndDate := time.Date(2019, 7, 20, 0, 0, 0, 0, time.UTC)
	actual := calculateSecond(StartDate,EndDate)
	if actual != expected {
		t.Errorf("expecred %d but it got %d", expected, actual)
	}
}