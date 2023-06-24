<style>
	.barber *, .user *, .admin *,  .barberuser *, .adminuser *, .adminbarber *,  .nil *{
		fill : none !important;
		stroke : none !important;
		background-size: 100% 100%;
		background-repeat: no-repeat;
		border-radius : 0.7rem;
	}

	.nil *{
		background-color : white;	
	}

	.barber * {
		background-color : #bbf7d0;		
	}

	.user *{
		background-color : #7dd3fc;		
	}

	.admin *{
		background-color : #fca5a5;		
	}

	.barberuser *{
		background-image: linear-gradient(90deg, #bbf7d0 50%, #bbf7d0 50%, #7dd3fc 50%, #7dd3fc 50%); 
	}

	.adminuser *{
		background-image: linear-gradient(90deg, #bbf7d0 33.33%, #fca5a5 33.33%, #fca5a5 66.66%, #7dd3fc 66.66%); 
	}

	.adminbarber *{
		background-image: linear-gradient(90deg, #bbf7d0 50%, #bbf7d0 50%, #fca5a5 50%, #fca5a5 50%); 
	}

	.nodeLabel, .edgeLabel{
		font-size: 3rem !important;
	}

	.nodeLabel{
		padding: 10px 10px;
	}


</style>

# Use-case diagram
```mermaid
flowchart LR

%% users
generic_user["<div style='width:200px;height:250px'><img src='stick.png' alt='kek'><h3 style='position:absolute;bottom:0px;text-align:center;width:100%;'>user</h3></div>"]

logged_user["<div style='width:200px;height:250px'><img src='stick.png' alt='kek'><h3 style='position:absolute;bottom:0px;text-align:center;width:100%;'>logged</h3></div>"]

barber_user["<div style='width:200px;height:250px'><img src='stick.png' alt='kek'><h3 style='position:absolute;bottom:0px;text-align:center;width:100%;'>barber</h3></div>"]

admin_user["<div class='admin' style='width:200px;height:250px'><img src='stick.png' alt='kek'><h3 style='position:absolute;bottom:0px;text-align:center;width:100%;'>admin</h3></div>"]

%% browse shops subgraph

logged_user ---- browse_shops
subgraph  
	%% entities
	view_shop_analytics -.extends.-> view_shop
	browse_shops([browse barber shops])
	find_shops([find shops])
	view_shop([view shop])
	review([review shop])
	view_reviews([view reviews])
	booking([book appointment])
	modify_shop([modify shop info])
	view_shop_analytics([view shop analytics])
	delete_shop([delete shop])

	%% relations
	browse_shops -.include.-> find_shops
	view_shop -.extends.-> find_shops
	review -.extends.-> view_shop
	modify_shop -.extends.->view_shop
	view_reviews-.extends.-> view_shop
	up_vote-.extends.->view_reviews
	down_vote-.extends.->view_reviews
	booking -.extends.-> view_shop
	delete_shop -.extends.->view_shop
end


%% user profile
logged_user ---- view_profile_info

subgraph  

	%% entities
	view_profile_info([view profile info])
	del_appointment([delete appointment])
	view_appointments([view appointments])
	browse_owned_shops([browse owned shops])
	curr_appointment([view current appointment])
	del_acc([delete account])
	delete_appointment([delete appointment])

	%% relations
	browse_owned_shops -.include.-> view_appointments
	delete_appointment -.extends.-> view_appointments
	
	browse_owned_shops -.extends.->view_profile_info
	curr_appointment -.extends.->view_profile_info
	del_acc -.extends.->view_profile_info
	del_appointment-.extends.->curr_appointment
end


%% admin subgraph
admin_user --- browse_users
admin_user --- user_analytics
admin_user --- create_shop

subgraph  
	%% entities
	browse_users([browse users])
	find_user([find user])
	view_user([view user])
	delete_user([delete user])
	modify_perm([edit barbershop ownership])
	user_analytics([view app analytics])
	create_shop([create shop])

	%% relations
	browse_users -.include.-> find_user
	find_user -.include.->view_user
	delete_user-.extends.->view_user
	modify_perm-.extends.->view_user
	user_analytics
	create_shop
end

%% login subgraph
generic_user ---- login
generic_user ---- signup

subgraph  
	login([login])
	signup([signup])
	pswd_rec([password recovery])

	pswd_rec -.extends.->login
end

%% user definitions
admin_user --> generic_user
barber_user --> generic_user
logged_user --> generic_user



%% colors
class browse_owned_shops barber
class login nil
class signup nil
class modify_shop barber
class up_vote user
class down_vote user
class dis_reviews barber
class barber_user barber
class view_appointments barber
class view_shop_analytics barber
class delete_appointment barber
class delete_shop admin
class delete_review admin
class browse_users admin
class find_user admin
class view_user admin
class delete_user admin
class modify_perm admin
class create_shop admin
class user_analytics admin
class curr_appointment user
class view_profile_info barberuser
class pswd_rec nil
class del_appointment user
class del_acc barberuser
class browse_shops user
class find_shops user
class view_shop user
class view_reviews user
class review user
class booking user

%% user colors
class generic_user nil
class logged_user user
class barber_user barber
class admin_user admin

```
