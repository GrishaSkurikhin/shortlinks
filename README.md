# Shortlinks
<div>
  <img src="https://github.com/devicons/devicon/blob/master/icons/go/go-original-wordmark.svg" title="Go" width="40" height="40"/>&nbsp;
  <img src="https://github.com/devicons/devicon/blob/master/icons/javascript/javascript-original.svg" title="Javascript" width="40" height="40"/>&nbsp; 
  <img src="https://github.com/devicons/devicon/blob/master/icons/mongodb/mongodb-original.svg" title="MongoDB" width="40" height="40"/>&nbsp; 
  <img src="https://github.com/devicons/devicon/blob/master/icons/docker/docker-original.svg" title="Docker" width="40" height="40"/>&nbsp; 
</div>

## Как запустить проект
1) Для сборки необходимы предустановленные **git** и **docker**.
2) Выполнить команды в терминале
```
    git clone https://github.com/GrishaSkurikhin/shortlinks
    cd shortlinks
    docker-compose up
```
При необходимости можно поменять хост и порт приложения в файле docker-compose.

## Краткое описание проекта
**Shortlinks** - веб-приложение для сокращения ссылок. Пользователь записывает в поля формы оригинальную ссылку и придуманное им короткое имя и получает укороченную ссылку на основе этого имени. Далее пользователь может использовать укороченную ссылку как оригинальную.
![image](https://user-images.githubusercontent.com/71190776/201477717-38ee05a2-7042-41ba-933b-be07670d4a1b.png)

## Архитектура проекта
Серверная часть написана на golang с использованием фреймворка gin. Бэкенд взаимодействует с фронтом, написанным на javascript, с помощью REST API. Методы api:
1) ```GET "/"``` - получение основной html-страницы
2) ```POST "/create"``` - создание короткой ссылки
3) ```GET "/re/:shortname"``` - переадресация на длинную ссылку

Для хранения информации о ссылках была использована нереляционная СУБД **MongoDB**. Записи в ней представляют собой отношение ```длинная_ссылка:короткая_ссылка```. Для обращения к БД на go был написан простой контроллер, реализующий интерфейс DataSource. При необходимости можно заменить MongoDB на любую другую СУБД, без изменения серверной части.

## Что можно улучшить
1) Добавить возможность регистрации пользователя
2) Добавить возможность удаления, изменения и просмотра своих записанных ссылок
3) Добавить выбор при генерации новой ссылки: по введенному имени или с использованием автоматически сгенерированного токена 
4) Добавить возможность управления доступом к новой ссылке
