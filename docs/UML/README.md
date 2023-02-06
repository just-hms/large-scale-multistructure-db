# UML Class diagram

```mermaid
classDiagram
direction LR

	Barber --|> User
	Admin --|> User
	Barber "1,*"--"1" BarberShop : Owns
	User "1,*"--"0,*" BarberShop : HasPermissions
	ShopView "1"--"0,*" BarberShop
	Calendar "1,*"--"1" Slot
	User "0,*"--"1" ShopView
	Appointment "1"--"0,*" BarberShop
	Review "1"--"0,*" BarberShop
	User "0,*"--"1" Appointment
	BarberShop "1"--"1" Calendar
	User "0,*"--"0,*" Review : HasDownvoted
	User "0,*"--"1" Review
	User "0,*"--"0,*" Review : HasUpvoted
	

	class User {
		Email : String
		Password : String
	}

	class Barber

	class Admin

	class BarberShop {
		Name : String
		AverageRating : Float
		Latitude : Float
		Longitude : Float
		EmployeesNumber : Int
	}

	class Appointment {
		CreatedAt : DateTime
		Start : DateTime
		Duration : Time
	}

	class ShopView {
		CreatedAt : DateTime
	}
	
	class Slot{
		Start : DateTime
		BookedAppoIntments : Int 
		UnavailableEmployees : Int 
	}

	class Review {
 		Content : String
		CreatedAt : DateTime
		Rating : Int
		Reported : Bool
	}
```