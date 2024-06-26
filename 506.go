/*
506. Относительные ранги
Легкий
Темы
Компании
Вам дан целочисленный массив scoreразмером n, где score[i]— результат спортсмена на соревновании. Все результаты гарантированно уникальны .ith

Спортсмены распределяются в зависимости от их результатов: спортсмен, занявший место, имеет наивысший балл, спортсмен, занявший место, имеет наивысший балл и так далее. Место каждого спортсмена определяет его ранг:1st2nd2nd

Место в разряде спортсменов .1st"Gold Medal"
Место в разряде спортсменов .2nd"Silver Medal"
Место в разряде спортсменов .3rd"Bronze Medal"
Для спортсменов, занявших место на место, их рейтинг равен номеру места (т. е. ранг спортсмена, занимающего место, равен ).4thnthxth"x"
Возвращает массив answerразмера nгде answer[i]— звание спортсмена .ith

Пример 1:

Ввод: счет = [5,4,3,2,1]

	Выход: ["Золотая медаль","Серебряная медаль","Бронзовая медаль","4","5"]
	Объяснение: Места: [1- е место , 2 -й , 3- й , 4- й , 5 -й ].

Пример 2:

Ввод: счет = [10,3,8,9,4]

	Выход: ["Золотая медаль","5","Бронзовая медаль","Серебряная медаль","4"]
	Объяснение: Места: [1 место , 5- е , 3- е , 2- е , 4 -е ].

Ограничения:

n == score.length
1 <= n <= 104
0 <= score[i] <= 106
Все значения scoreуникальны .
*/
package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	if len(score) == 0 {
		return []string{}
	}

	if len(score) == 1 {
		return []string{"Gold Medal"}
	}

	sorted := make([]int, len(score))

	for i, v := range score {
		sorted[i] = v
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))

	mRes := make(map[int]int, len(sorted))
	for i, v := range sorted {
		mRes[v] = i + 1
	}
	fmt.Println(mRes)

	res := make([]string, len(sorted))
	for i, v := range score {
		switch mRes[v] {
		case 1:
			res[i] = "Gold Medal"
		case 2:
			res[i] = "Silver Medal"
		case 3:
			res[i] = "Bronze Medal"
		default:
			res[i] = strconv.Itoa(mRes[v])
		}
	}

	return res

}

func main() {

	fmt.Println(findRelativeRanks([]int{1, 2}))
}
