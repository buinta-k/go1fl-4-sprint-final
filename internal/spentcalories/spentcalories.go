package spentcalories

import (
	"time"
	"strings"
	"strconv"
	"log"
	"fmt"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	slice:=strings.Split(data,",")
	if len(slice)!=3 {
		return 0, "", 0, fmt.Errorf("Длина слайса не равна 3")
	}
	
	steps, err:=strconv.Atoi(slice[0])
	if err!=nil {
		return 0, "", 0, err
	}

	t, err2:=time.ParseDuration(slice[2])
	if err2!=nil {
		return 0, "", 0, err2
	}

	return steps, slice[1], t, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	length:=heigth*stepLengthCoefficient
	return ( float64(steps)*length)/mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration<0 {
		return 0
	}

	distance1:=distance(steps,height)

	return distance1/duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, sport, duration, err:=parseTraining(data)
	if err!=nil {
		log.Println(err)
	}
	distance1:=distance(steps, height)
	speed:=meanSpeed(steps, height, duration)
	
	switch sport{
		case "Ходьба":
		calories:=WalkingSpentCalories(steps, weight, height, duration)
		fmt.Printf("Тип тренировки: %s\nДлительность: %.2f\nДистанция: %d км.\nСкорость: %d км/ч.\nСожгли калорий: %d", sport, duration, distance1, calories)
		case "Бег":
		calories:=RunningSpentCalories(steps, weigh, height, duration)
		fmt.Printf("Тип тренировки: %s\nДлительность: %.2f\nДистанция: %d км.\nСкорость: %d км/ч.\nСожгли калорий: %d", sport, duration, distance1, calories)
	}
	
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps<=0 || weight<=0 || heigth<=0 || duration<=0 {
		return 0, fmt.Errorf("Некорректный параметр")
	}

	averageSpeed:=meanSpeed(steps, height, duration)

	durationInMinutes:=duration.Minutes()

	return (weight*averageSpeed*durationInMinutes)/minInH
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps<=0 || weigth<=0 || height<=0 || duration<=0 {
		return 0, fmt.Errorf("Некорректный параметр")
	}
	
	averageSpeed:=meanSpeed(steps, height, duration)

	durationInMinutes:=duration.Minutes()

	count:=(weight*averageSpeed*durationInMinutes)/minInH

	return count*walkingCaloriesCoefficient
}
