```bash
docker compose build
docker compose up -d clickhouse rabbit
...
docker compose up migrations
docker compose up producer step1 step2 step3
docker compose up producer mono 
```
Need to wait several seconds after up ``clickhouse`` and ``rabbit``

Reports with work time are stored in ``reports/``

### Описание работы
* producer - генерит ивенты в очередь `enter`. Конфиги:
  * svc.total - общее количество генерящихся событий
  * svc.eps - можно выставить отдельное количество event per second
* step1, step2, step3 - получают и обрабатывают сообщение, перекидывая в следующие очереди (step1, step2, step3)
  * svc.total - сколько всего ожидает сервис получить сообщений
* step4 - получает сообщение из очереди `step3` и кладет результат в clickhouse
  * svc.total - сколько всего ожидает сервис получить сообщений
* mono - получает сообщение из очереди `enter`, процессит через ту же логику (step1, step2, step3) без прокидывания через rabbit между сервисами, и складывает в clickhouse
  * svc.total - сколько всего ожидает сервис получить сообщений

Для эмуляции логики обработки сообщений, дается CPU bound задача на каждом из этапов

Параллелизм обработки отсутствует

Можно запускать `producer` параллельно с настроенным eps или отдельно, заранее (что бы не давать дополнительную нагрузку и оценить производительность конкретного блока, предположив что внутренний и внешний рэббит - отдельные системы)