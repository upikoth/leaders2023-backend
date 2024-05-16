# Leaders2023-backend

## Как собрать приложение:

docker build --platform=linux/amd64 -t cr.yandex/crpg6rg1bsgvdmff8sm1 .

docker push cr.yandex/crpg6rg1bsgvdmff8sm1

env переменная PORT задается в serverless container автоматически самим сервисом.
