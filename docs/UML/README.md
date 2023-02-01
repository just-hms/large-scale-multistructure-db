# Class diagram

```mermaid
classDiagram
direction LR

	User "1,*"--"0,*" GlobalPermission
	ShopView "0,*"--"1" BarberShop
	ShopPermission --|> Permission
	User "1,*"--"0,*" ShopPermission
	Calendar "1,*"--"1" Slot
	User "1"--"0,*" ShopView
	Appointment "1"--"0,*" BarberShop
	Review "1"--"0,*" BarberShop
	ShopPermission "1"--"0,*" BarberShop
	User "0,*"--"1" Appointment
	BarberShop "1"--"1" Calendar
	GlobalPermission --|> Permission
	User "0,*"--"0,*" Review : DownVotes
	User "1"--"0,*" Review
	User "0,*"--"0,*" Review : UpVotes

	class Permission {
		<<abstract>>
		Name : String
	}

	class User {
		Email : String
		Password : String
	}

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

# For later

Holidays are done only by editing the `UnavailableEmployees` field in the slots.

Only barbers can report reviews, so there is no need for a counter.

Admin will have a view of the reported reviews
