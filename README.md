<div>
    <center>
        <h1>Numlookupapi API</h1>
    </center>
    <p><a href="https://app.numlookupapi.com/" target="_blank">numlookupapi.com</a> - Автоматизируйте процесс проверки номера телефона, проверяя информацию об операторе связи с помощью API для поиска телефонных номеров по всему миру.</p>
<div>
<div>
    <h2>Тарифный планы</h2>
    <ul>
        <li>Бесплатный: 100 запросов / в месяц</li>
        <li>Платный: от $9.99 /мес.</li>
    </ul>
    <p><a href="https://numlookupapi.com/pricing/">Подробные тарифные планы</a></p>
</div>
<div>
    <h2>Получить numlookupapi</h2>
    <p>Импортный пакет</p> 

```
import "github.com/Clyckov34/numlookupapi"
```

<p>Установка пакета</p>

```
go get github.com/Clyckov34/numlookupapi
```
</div>
<div>
    <h2>Пример: Общий вывод</h2>
	<p>example.go:</p>

```go
package main

import (
	"fmt"
	"log"

	"github.com/Clyckov34/numlookupapi"
)

func main() {
	var api = numlookupapi.Client{
		ApiKey: "API-KEY",
	}

	res, err := api.GetResponse("+79963567210")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)
}	

```

```
$ go run example.go
&{true 79963567212 9963567212 +79963567212 +7 RU Russian Federation Volgograd Oblast LLC Skartel (YOTA) mobile}
```

</div>
<div>
    <h2>Пример: Вывод по конкретным данным</h2>
    <p>example.go:</p>

```go
package main

import (
	"fmt"
	"log"

	"github.com/Clyckov34/numlookupapi"
)

func main() {
	var api = numlookupapi.Client{
		ApiKey: "API-KEY",
	}

	res, err := api.GetResponse("+79963567210")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Valid:", res.Valid)
	fmt.Println("Number:", res.Number)
	fmt.Println("Local Format:", res.LocalFormat)
	fmt.Println("International Format:", res.InternationalFormat)
	fmt.Println("Country Prefix:", res.CountryPrefix)
	fmt.Println("Country Code:", res.CountryCode)
	fmt.Println("Country Name:", res.CountryName)
	fmt.Println("Location:", res.Location)
	fmt.Println("Carrier:", res.Carrier)
	fmt.Println("Line Type:", res.LineType)

}


```

```
$ go run example.go
Valid: true
Number: 79963567210
Local Format: 9963567210
International Format: +79963567210
Country Prefix: +7
Country Code: RU
Country Name: Russian Federation
Location: Volgograd Oblast
Carrier: LLC Skartel (YOTA)
Line Type: mobile
```

</div>
<div>
	<h2>Заметки</h2>
	<ul>
		<li>Все номира должны начинаться со знака "+". Например: +79963567210</li>
		<li>Чтобы получить API-KEY, зарегистрируйтесь в <a href="https://app.numlookupapi.com/dashboard">Личном кабинете</a>, и скопируйте ключ</li>
	</ul>
</div>