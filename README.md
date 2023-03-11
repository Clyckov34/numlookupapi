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
    <p>С поддержкой модуля Go просто добавьте следующий импорт</p> 

```
import "github.com/Clyckov34/numlookupapi"
```

<p>к вашему коду, а затем go [build|run|test]автоматически получит необходимые зависимости</p>
<p>В противном случае выполните следующую команду Go, чтобы установить numlookupapi пакет</p>

```
$ go get github.com/Clyckov34/numlookupapi
```
</div>
<div>
    <h2>Пример</h2>
    <p>Сначала вам нужно импортировать пакет numlookupapi для использования numlookupapi, один из простейших примеров выглядит следующим образом example.go:</p>

```go
package main

import (
	"fmt"
	"log"

	"github.com/Clyckov34/numlookupapi"
)

func main() {
	var api = numlookupapi.Params{ApiKey: "API-KEY"}
	response, err := api.GetRequest("+79963567210")
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	result, err := response.Data()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(result)
	fmt.Println(result.PhoneLocalFormat())
}

```

</div>
