# Class diagram

``` mermaid
classDiagram
direction LR

	ShopPermission --|> Permission
	User "0,*"-- Visualization
	Reservation --"0,*" BarberShop
	Review --"0,*" BarberShop
	User "0,*"-- Vote
	User "0,*"-- Reservation
	User "1,*"--"0,*" Permission
	GlobalPermission --|> Permission
	Visualization --"0,*" BarberShop
	ShopPermission --"0,*" BarberShop
	BarberShop -- Calendar
	Calendar "1,*"-- Slot
	Vote --"0,*" Review
	User "0,*"-- Review

	Permission : string Name

	User : string Email
	User : string Password

	BarberShop : string Name
	BarberShop : float AverageRating 
	BarberShop : float Latitude
	BarberShop : float Longitude
	BarberShop : int EmployeesNumber

	Reservation : datetime CreatedAt
	Reservation : datetime Start
	Reservation : time Duration

	Visualization : datetime CreatedAt

	Slot : datetime Start
	Slot : int BookedAppointments
	Slot : int UnavailableEmployees

	Review : string Content
	Review : datetime CreatedAt
	Review : int Rating

	Vote : bool Up
```

