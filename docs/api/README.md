# List of all available APIs divided in

- free-to-use : 	` `
- require-login : 	🔑
- require-barber : 	💈
- require-admin :	🛠️

# User

## `POST` /user/register/

request

```json
{
	"email" : "sus@kek.com",
	"password" : "super_secret",
}
```

response ✔️ -> status : `201`

```json
{
	"token" : "token",
}
```

response ❌ -> status : `400`

```json
{
	"error" : "errorMessage",
}
```

## `POST` /user/login/

request

```json
{
	"email" : "sus@kek.com",
	"password" : "super_secret",
}
```

response ✔️ -> status : `200`

```json
{
	"token" : "token",
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```

## 🔑 `GET` /user/self/

> also the current appointment

response ✔️ -> status : `200`

```json
{
	"user" : {},
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```

## 🔑 `DELETE` /user/self/

response ✔️ -> status : `202`

```json
{
	"message" : "Message",
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```


## 🔑 `GET` /user/logout/

> maybe only client side (just delete the token)


## 🛠️ `GET` /admin/user?email=""

> ordered by something

response ✔️ -> status : `200`

```json
{
	"users" : [
		{
			"id": "",
			"email": "",
		},
	],
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```

## 🛠️ `GET` /admin/user/:id

response ✔️ -> status : `200`

```json
{
	"user" : {
		"id": "",
		"email": "",
		"permissions" : [
			
		]
	},
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```

## 🛠️ `DELETE` /admin/user/:id

response ✔️ -> status : `202`

```json
{
	"message" : "Message",
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```

## 🛠️ `PUT` /admin/user/:id

> PARTIAL UPDATE

> TODO: check shopPermission

response ✔️ -> status : `200`

```json
{
	"user" : {
		"email": "",		
		"permissions" : [
			
		]
	},
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```

> maybe add an analytics listing

# Password

## 🔑 `GET` /user/password_recovery/

response ✔️ -> status : `200`

```json
{
	"memssage" : "errorMessage",
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```

## 🔑 `POST` /user/confirm_recovery?token=""

request

```json
{
	"password" : "super_secret",
}
```

response ✔️ -> status : `200`

```json
{
	"memssage" : "errorMessage",
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```

# BarberShop

## 🔑 `GET` /barber_shop?lat=""&lon=""&name=""&radius=""

response ✔️ -> status : `200`

```json
{
	"barberShops" : [
		{
			"TODO" : "TODO"		
		},
	],
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```


## 🔑 `GET` /barber_shop/:id

response ✔️ -> status : `200`

```json
{
	"barberShop" : {
		"TODO" : "TODO"
	},
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```

## 🛠️ `POST` /admin/barber_shop/:id

response ✔️ -> status : `201`

```json
{
	"barberShop" : {},
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```


## 💈 `PUT` /barber_shop/:id

> PARTIAL UPDATE

request 

```json
{
	"barberShop" : {},
}
```

response ✔️ -> status : `201`

```json
{
	"message" : "message",
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```


## 🛠️ `DELETE` /admin/barber_shop/:id

response ✔️ -> status : `202`

```json
{
	"message" : "message",
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```

# Calendar

## 🔑 | 💈 | 🛠️ `GET` /barber_shop/:id/calendar

# Appointment

## 🔑 `POST` /barber_shop/:id/appointment

request 

```json
{
	"date" : "01/12/2000",
	"slot" : 1,
}
```

response ✔️ -> status : `201`

```json
{
	"message" : "message",
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```

## 🔑 `DELETE` /user/self/appointment/

request

```json
{

}
```

response ✔️ -> status : `202`

```json
{
	"message" : "message",
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```


## 💈 `DELETE` /appointment/:id

request

```json
{

}
```

response ✔️ -> status : `202`

```json
{
	"message" : "message",
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```


# Holidays

## 💈 `POST` /barber_shop/:id/holidays

request 

```json
{
	"date" : "01/12/2000",
	"slot" : 1,
	"unavailableEmployees" : 7,
}
```

response ✔️ -> status : `201`

```json
{
	"message" : "message",
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```

> every time the number of workers is updated in the calendar


# Reviews

## 🔑 `POST` /barber_shop/:id/review/

request 

```json
{
	"title" : "kek",
	"body" : "kekkeroni",
	"rating" : 3
}
```

response ✔️ -> status : `201`

```json
{
	"message" : "message",
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```


## 🔑 `GET` /barber_shop/:id/review?page=""&sort_by=""

`sort_by` = `upvote` | `downvote` | `recent` | `oldest`

response ✔️ -> status : `200`

```json
{
	"reviews" : [
		{

		},
	],

	"len" : 2
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```

## 🛠️ `DELETE` /review/:id

response ✔️ -> status : `200`

```json
{
	"message" : "message",
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```

## 💈 `POST` /review/:id/report

request

```
TODO : 
```

response ✔️ -> status : `200`

```json
{
	"message" : "message",
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```

## 🔑 `POST` /review/:id/vote

request
```json
{
	"value" : +1,
}
```

response ✔️ -> status : `200`

```json
{
	"message" : "message",
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```


# Anal

## 🛠️ `GET` /admin/analytics/:id

response ✔️ -> status : `200`

```json
{
	"data" : {},
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```

## 🛠️ `GET` /admin/barber_shop/:id/analytics/:id

response ✔️ -> status : `200`

```json
{
	"data" : {},
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```
