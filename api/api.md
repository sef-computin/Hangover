# Gateway:
1) Регистрация
POST /api/v1/auth/register
2) Авторизация
POST /api/v1/auth/authorize
3) Обновить токен
POST /api/v1/auth/refresh

3) Получить список мероприятий
GET /api/v1/events/
Body: {}
RESPONSE: Event[]

4) Создать мероприятие
POST /api/v1/events/create
Body: Event{}
RESPONSE: {}

5) Редактировать мероприятие
PATCH /api/v1/events/{event_uuid} 
Body: Event{}
RESPONSE: {}

6) Отменить мероприятие
DELETE /api/v1/events/{event_uuid} 
Body: {}
RESPONSE: {}

7) Записаться на мероприятие
POST /api/v1/events/{event_uuid}/enroll
Body: 
8) Отменить запись
DELETE /api/v1/enrollments/{enroll_uuid}

# Events
1) Получить список мероприятий
GET /api/v1/events
2) Создать мероприятие
POST /api/v1/events/create
3) Редактировать мероприятие
PATCH /api/v1/events/{event_uuid}

# Enrollments
1) Записать пользователя на мероприятие
2) Подтвердить запись пользователю
3) Отменить запись

# Logs
1) Сохранить лог
