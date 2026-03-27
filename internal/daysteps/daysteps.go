package daysteps

import (
	"time"
	"strings"
	"strconv"
	"fmt"
	"log"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
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
	
	t, err2:=time.ParseDuration(strings.TrimSpace(slice[1]))
	if err2!=nil {
			return 0, 0, err2
	}
	return steps, t, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, t, err:=parsePackage(data)
	if err!=nil {
		log.Println(err)
		return ""
	}
	
	if steps<=0 {
		return ""
	}
	
	distance:=(float64(steps)*stepLength)/mInKm

	calories, err2:=spentcalories.WalkingSpentCalories(steps, weight, height, t)
	
	if err2 != nil {
		log.Println(err)
   		 return "" 
	}
	
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)

}
