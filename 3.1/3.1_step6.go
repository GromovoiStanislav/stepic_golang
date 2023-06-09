package main

import "fmt"

/*
В ходе анализа результатов переписи населения информация была сохранена в
объекте типа map:
	groupCity := map[int][]string{
		10:   []string{...}, // города с населением 10-99 тыс. человек
		100:  []string{...}, // города с населением 100-999 тыс. человек
		1000: []string{...}, // города с населением 1000 тыс. человек и более
	}

При подготовке важного отчета о городах с населением 100-999 тыс. человек был
подготовлен другой объект типа map:
	cityPopulation := map[string]int{
		город: население города в тысячах человек,
	}

Однако база данных с информацией о точной численности населения содержала ошибки,
поэтому в cityPopulation в т.ч. была сохранена информация о городах, которые входят
в другие группы из groupCity.

Ваша программа имеет доступ к обоим указанным отображениям, требуется исправить
cityPopulation, чтобы в ней была сохранена информация только о городах из
группы groupCity[100].

Функция main() уже объявлена, доступ к отображениям осуществляется по
указанным именам. По результатам выполнения ничего больше делать не требуется,
проверка будет осуществлена автоматически.
*/

func main() {
	/*
	 * Группировка городов по численности населения в тысячах человек
	 * от 10 до 100, от 100 до 1000 и более 1000:
	 * groupCity map[int][]string{
	 *	 10: []string{...},
	 *	 100: []string{...},
	 *	 1000: []string{...},
	 * }
	 *
	 * Население городов в тысячах человек:
	 * cityPopulation map[string]int{...}
	 */

	groupCity := map[int][]string{
		10:   {"A_10", "B_10", "C_10"},
		100:  {"A_100", "B_100", "C_100"},
		1000: {"A_1000", "B_1000", "C_1000"},
	}
	fmt.Println(groupCity)

	cityPopulation := map[string]int{
		"A_10": 10, "B_10": 15, "C_10": 20,
		"A_100": 100, "B_100": 200, "C_100": 400,
		"A_1000": 1001, "B_1000": 2000, "C_1000": 3000,
	}

	// for number, cities := range groupCity {
	// 	if number == 100 {
	// 		continue
	// 	}

	// 	for _, city := range cities {
	// 		delete(cityPopulation, city)
	// 	}
	// }
	// fmt.Println(cityPopulation)


	// for _, city := range groupCity[1000] {
	// 	delete(cityPopulation, city)
	// }
	// for _, city := range groupCity[10] {
	// 	delete(cityPopulation, city)
	// }


	for _, city := range append(groupCity[10], groupCity[1000]...) {
		delete(cityPopulation, city)
	}
}