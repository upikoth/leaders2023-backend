# Leaders2023-backend

## Как собрать приложение:

cat secrets/authorized_key.json | docker login \
  --username json_key \
  --password-stdin \
  cr.yandex

docker build --platform=linux/amd64 -t cr.yandex/crpo5i4epql8ladb2336 .

docker push cr.yandex/crpo5i4epql8ladb2336

env переменная PORT задается в serverless container автоматически самим сервисом.
