: <<'COMMENT'
195. Десятая линия
Легкий
Темы
Компании
Учитывая текстовый файл  file.txt, выведите только 10-ю строку файла.

Пример:

Предположим, что он file.txtимеет следующее содержимое:

Линия 1
Линия 2
Линия 3
Линия 4
Линия 5
Линия 6
Линия 7
Линия 8
Линия 9
Строка 10
Ваш скрипт должен вывести десятую строку:

Строка 10
Примечание:
1. Если файл содержит менее 10 строк, что следует вывести?
2. Есть как минимум три разных решения. Постарайтесь изучить все возможности.
COMMENT

#!/bin/bash
filename="file.txt"
if [ $(wc -l < "$filename") -ge 10 ]; then
    head -n 10 "$filename" | tail -n 1
fi

awk 'NR == 10' file.txt

sed -n '10p' file.txt


n=1;
while read line ;
do
    if [  $n -eq 10  ]; then
        echo "$line"
    fi
    n=$((n+1));
done < file.txt