# Pacote `cmd/all-in-one`

A ser documentado.

Mas por enquanto, aqui vão alguns exemplos: 

```terminal
$ nats -s localhost:4222 stream create payments --subjects "payment.process" --storage memory --replicas 1 --retention=limits --discard=old --max-msgs 1_000_000 --max-msgs-per-subject 100_000 --max-bytes 4GiB --max-age 1d --max-msg-size 10MiB --dupe-window 2m --allow-rollup --no-deny-delete --no-deny-purge
$ curl localhost:8080/payments
$ curl -X POST localhost:8080/payments -d '{"id": "some-uuid", "subscription_id":"jpkroehling", "amount":99, "status":"FAILED"}'
$ nats -s localhost:4222 stream view payments
```
