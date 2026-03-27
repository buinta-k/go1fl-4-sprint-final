package daysteps

import (
	"time"
	"strings"
	"strconv"
	"fmt"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	slice:=strings.Split(data,",")
	
	if len(slice)!=2 {
		return 0, 0, fmt.Errorf("Длина слайса не равна 2")
	}
	
	steps, err:=strconv.Atoi(slice[0]) 
	if err!=nil {
			return 0, 0, err
			}
		
	if steps<=0 {
			return 0, 0 , fmt.Errorf("Количество шагов должно быть больше 0")
		}
	
	t, err2:=time.ParseDuration(slice[1])
	if err!=nil {
			return 0, 0, err2
	}
	return steps, t, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, t, err:=parsePackage(data)
	if err!=nil {
		fmt.Println(err)
		return ""
	}
	
	if steps<=0 {
		return ""
	}
	
	distance:=(float64(steps)*stepLength)/mInKm

	calories:=WalkingSpentCalories(steps, weight, height, distance)

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %f км.\nВы сожгли %f",steps,distance,calories)
}
