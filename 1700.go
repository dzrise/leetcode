/*
1700. Число студентов, не имеющих возможности пообедать
Легкий
Темы
Компании
Намекать
В школьной столовой во время обеденного перерыва предлагаются круглые и квадратные бутерброды, обозначенные цифрами 0и 1соответственно. Все студенты стоят в очереди. Каждый студент предпочитает квадратные или круглые бутерброды.

Количество бутербродов в столовой равно количеству студентов. Бутерброды складываются в стопку . На каждом этапе:

Если студент, стоящий в начале очереди, предпочитает бутерброд, лежащий наверху стопки, он возьмет его и покинет очередь.
В противном случае они покинут ее и перейдут в конец очереди.
Это продолжается до тех пор, пока ни один из студентов в очереди не захочет взять верхний бутерброд и, следовательно, не сможет есть.

Вам даны два целочисленных массива students, sandwichesгде sandwiches[i]— тип сэндвича в стеке ( это вершина стека) и — предпочтение студента в начальной очереди ( это начало очереди). Возвращает количество студентов, которые не могут есть.i​​​​​​thi = 0students[j]j​​​​​​thj = 0

Пример 1:

Входные данные: студенты = [1,1,0,0], бутерброды = [0,1,0,1]

	Выходные данные: 0

Объяснение:
- Передний ученик покидает верхний сэндвич и возвращается в конец очереди, образуя учеников = [1,0,0,1].
- Передний ученик покидает верхний сэндвич и возвращается в конец строки, делая учеников = [0,0,1,1].
- Передний ученик берет верхний сэндвич и покидает линию, образуя учеников = [0,1,1] и сэндвичей = [1,0,1].
- Передний ученик покидает верхний сэндвич и возвращается в конец строки, делая учеников = [1,1,0].
- Передний ученик берет верхний сэндвич и покидает линию, образуя учеников = [1,0] и сэндвичей = [0,1].
- Передний ученик покидает верхний сэндвич и возвращается в конец строки, делая учеников = [0,1].
- Передний ученик берет верхний сэндвич и выходит из очереди, образуя учеников = [1] и сэндвичей = [1].
- Передний ученик берет верхний сэндвич и выходит из очереди, образуя учеников = [] и сэндвичи = [].
Следовательно, все студенты могут есть.
Пример 2:

Входные данные: студенты = [1,1,1,0,0,1], бутерброды = [1,0,0,0,1,1]

	Выходные данные: 3

Ограничения:

1 <= students.length, sandwiches.length <= 100
students.length == sandwiches.length
sandwiches[i]есть 0или 1.
students[i]есть 0или 1.
*/
package main

import "fmt"

func countStudents(students []int, sandwiches []int) int {
	adds := len(students) * len(students)
	for i := 0; i < len(sandwiches); i++ {
		if students[i] == sandwiches[0] {
			students = append(students[:i], students[i+1:]...)
			sandwiches = sandwiches[1:]
			i--
		} else {
			cur := students[i]
			students = append(students[:i], students[i+1:]...)
			students = append(students, cur)
			if adds > 0 {
				adds--
				i--
			}
		}
	}

	return len(students)
}

func main() {
	fmt.Println(countStudents([]int{1, 1, 0, 0}, []int{0, 1, 0, 1}))
	fmt.Println(countStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}))
	//fmt.Println(countStudents([]int{0, 0, 0, 0, 0, 1}, []int{1, 1, 1, 0, 0, 0}))
}
