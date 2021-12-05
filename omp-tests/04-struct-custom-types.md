# Тест 4. Структуры. Собственные типы

1. Что выведет код?

   ```go
   package main

   import "fmt"

   type Point struct {
       X, Y int
   }

   func main() {
       var p Point

       if p == nil {
           fmt.Println("true")
       } else {
           fmt.Println("false")
       }
   }
   ```

   - [ ] true
   - [ ] false
   - [ ] Код не скомпилируется
   - [ ] Код запаникует

   <details>
   <summary>Ответ с пояснением</summary>

   - [ ] true
   - [ ] false
   - [x] **Код не скомпилируется**
   - [ ] Код запаникует

   **Объяснение:**

   `nil` - некорректное значение для Point. Поэтому попытка сравнения p с `nil` не пройдет компиляцию
   </details>

1. Что выведет код?

   ```go
   package main

   import "fmt"

   type Point struct {
       X, Y int
   }

   func main() {
       var p Point

       if p == (Point{}) {
           fmt.Println("true")
       } else {
           fmt.Println("false")
       }
   }
   ```

   - [ ] true
   - [ ] false
   - [ ] Код не скомпилируется
   - [ ] Код запаникует

   <details>
   <summary>Ответ с пояснением</summary>

   - [x] **true**
   - [ ] false
   - [ ] Код не скомпилируется
   - [ ] Код запаникует

   **Объяснение:**

   Структуры одинакового типа сравниваются на равенство по значению. При этом равенство достигается только в том случае, если все соответствующие элементы структур равны.
   </details>

1. Что выведет код?

   ```go
   package main

   import "fmt"

   type Point struct {
       X, Y int
   }

   func main() {
       var p Point

       if p == struct{X,Y int}{} {
           fmt.Println("true")
       } else {
           fmt.Println("false")
       }
   }
   ```

   - [ ] true
   - [ ] false
   - [ ] Код не скомпилируется
   - [ ] Код запаникует

   <details>
   <summary>Ответ с пояснением</summary>

   - [x] **true**
   - [ ] false
   - [ ] Код не скомпилируется
   - [ ] Код запаникует

   **Объяснение:**

   Мы объявили собственный тип `Point` на базе структуры из двух типов. При этом вполне допустимо сравнивать `Point` с анонимной структурой такой же сигнатуры
   </details>

1. Что выведет код?

   ```go
   package main

   import "fmt"

   type Point struct {
       X, Y int
   }

   type AnotherPoint struct {
       X, Y int
   }

   func main() {
       var p Point

       if p == (AnotherPoint{}) {
           fmt.Println("true")
       } else {
           fmt.Println("false")
       }
   }
   ```

   - [ ] true
   - [ ] false
   - [ ] Код не скомпилируется
   - [ ] Код запаникует

   <details>
   <summary>Ответ с пояснением</summary>

   - [ ] true
   - [ ] false
   - [x] **Код не скомпилируется**
   - [ ] Код запаникует

   **Объяснение:**

   Сравнение разных собственных типов недопустимо, даже если они под капотом имеет один и тот же тип (в данном случае - структура с полями X и Y). Однако данный код легко можно заставить работать, если явно преобразовать один тип к другому. Например, заменить `(AnotherPoint{})` на `Point(AnotherPoint{})`. Тогда функция выведет `true`.
   </details>

1. Что выведет код?

   ```go
   package main

   import "fmt"

   type Point struct {
       X, Y     int
       Comments []string
   }

   func main() {
       var p Point

       if p == (Point{}) {
           fmt.Println("true")
       } else {
           fmt.Println("false")
       }
   }
   ```

   - [ ] true
   - [ ] false
   - [ ] Код не скомпилируется
   - [ ] Код запаникует

   <details>
   <summary>Ответ с пояснением</summary>

   - [ ] true
   - [ ] false
   - [x] **Код не скомпилируется**
   - [ ] Код запаникует

   **Объяснение:**

   Объекты типа `Point` нельзя сравнивать ни с `nil`, ни друг с другом, поскольку в определении типа `Point` присутствует слайс `[]string`, который не поддерживает сравнение на равенство.
   </details>

1. Что выведет код?

   ```go
   package main

   import "fmt"

   type Point struct {
       X int
   }

   func (p Point) SetX(newX int) {
       p.X = newX
   }

   func main() {
       var p Point
       p.SetX(1)
       fmt.Println(p.X)
   }
   ```

   - [ ] 0
   - [ ] 1
   - [ ] Код не скомпилируется
   - [ ] Код запаникует

   <details>
   <summary>Ответ с пояснением</summary>

   - [x] **0**
   - [ ] 1
   - [ ] Код не скомпилируется
   - [ ] Код запаникует

   **Объяснение:**

   `p` - неявный аргумент функции `SetX`. А поскольку все аргументы передаются по значению, то при вызове `SetX` происходит копирование аргумента `p`. Здесь происходит модификация копии, которая будет уничтожена после выхода из функции
   </details>

1. Что выведет код?

   ```go
   package main

   import "fmt"

   type Point struct {
       X int
   }

   func (p Point) SetX(newX int) {
       p.X = newX
   }

   func main() {
       p := &Point{}
       p.SetX(1)
       fmt.Println(p.X)
   }
   ```

   - [ ] 0
   - [ ] 1
   - [ ] Код не скомпилируется
   - [ ] Код запаникует

   <details>
   <summary>Ответ с пояснением</summary>

   - [x] **0**
   - [ ] 1
   - [ ] Код не скомпилируется
   - [ ] Код запаникует

   **Объяснение:**

   `p` - неявный аргумент функции `SetX`. Несмотря на то, что `p` внутри функции `main` имеет тип `*Point`, перед вызовом происходит разыменовывание указателя и копирование полученного значения. Внутри `SetX` происходит модификация копии, которая будет уничтожена после выхода из функции
   </details>

1. Что выведет код?

   ```go
   package main

   import "fmt"

   type Point struct {
       X int
   }

   func (p *Point) SetX(newX int) {
       p.X = newX
   }

   func main() {
       var p Point
       p.SetX(1)
       fmt.Println(p.X)
   }
   ```

   - [ ] 0
   - [ ] 1
   - [ ] Код не скомпилируется
   - [ ] Код запаникует

   <details>
   <summary>Ответ с пояснением</summary>

   - [ ] 0
   - [x] **1**
   - [ ] Код не скомпилируется
   - [ ] Код запаникует

   **Объяснение:**

   Перед вызовом `SetX` неявно происходит взятие указателя. Внутри `SetX` происходит модификация значения по указателю (т.е. по адресу в памяти), который указывает на область памяти переменной `var p Point` из функции `main`. Поэтому эти изменения сохранятся даже после завершения вызова функции.
   </details>

1. Что выведет код?

   ```go
   package main

   import "fmt"

   type Point struct {
       X int
   }

   func (p *Point) Description() string {
       return "Hello"
   }

   func main() {
       var p *Point
       fmt.Println(p.Description())
   }
   ```

   - [ ] Hello
   - [ ] Код не скомпилируется
   - [ ] Код запаникует

   <details>
   <summary>Ответ с пояснением</summary>

   - [x] **Hello**
   - [ ] Код не скомпилируется
   - [ ] Код запаникует

   **Объяснение:**

   Zero-value для типа `*Point` - `nil`. Поскольку тип получателя совпадает с типом переменной `var p *Point`, то `nil` передается в качестве неявного аргумента функции `Description` "как есть". Поскольку он никак не используется, то никаких ошибок во время выполнения не происходит.
   </details>

1. Что выведет код?

   ```go
   package main

   import "fmt"

   type Point struct {
       X int
   }

   func (p Point) Description() string {
       return "Hello"
   }

   func main() {
       var p *Point
       fmt.Println(p.Description())
   }
   ```

   - [ ] Hello
   - [ ] Код не скомпилируется
   - [ ] Код запаникует

   <details>
   <summary>Ответ с пояснением</summary>

   - [ ] Hello
   - [ ] Код не скомпилируется
   - [x] **Код запаникует**

   **Объяснение:**

   Zero-value для типа `*Point` - `nil`. Поскольку тип получателя функции не совпадает с типом переменной `var p *Point`, то перед вызовом функции `Description` происходит разыменовывание nil-указателя, которое приводит к панике `panic: runtime error: invalid memory address or nil pointer dereference`
   </details>
