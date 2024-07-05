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