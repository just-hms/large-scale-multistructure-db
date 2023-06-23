# UML Class diagram

```mermaid
classDiagram
direction LR

	Barber --|> User
	Barber "1,*"--"1" BarberShop : Owns
	Admin --|> User
	Calendar "1,*"--"1" Slot
	ShopView "1"--"0,*" BarberShop
	Appointment "1"--"0,*" BarberShop
	BarberShop "1"--"1" Calendar
	Review "1"--"0,*" BarberShop
	User "0,*"--"1" ShopView
	User "0,*"--"0,*" Review : HasUpvoted
	User "0,*"--"1" Appointment
	User "0,*"--"0,*" Review : HasDownvoted
	User "0,*"--"1" Review
	

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