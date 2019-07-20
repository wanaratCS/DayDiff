package calcualte

import (
	"testing"
	"time"
)

func Test_calculate_Input_9_5_7_4_Should_Be_40(t *testing.T) {
	expected, number1, number2, number3, number4 := 40, 9, 5, 7, 4
	actual := calculate(expected, number1, number2, number3, number4)
	if actual != expected {
		t.Errorf("expecred %d but it got %d", expected, actual)
	}
}

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