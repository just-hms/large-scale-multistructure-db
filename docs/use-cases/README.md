<style>
	.barber *, .user *, .admin *, .all *, .barberuser *, .adminuser *, .adminbarber *,  .nil *{
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
	
	.all *{
		background-image: linear-gradient(90deg, #bbf7d0 33.33%, #fca5a5 33.33%, #fca5a5 66.66%, #7dd3fc 66.66%); 
	}

	.nodeLabel, .edgeLabel{
		font-size: 3rem !important;
	}

	.nodeLabel{
		padding: 10px 10px;
	}


</style>

# Use-case diagram
<!-- 
```mermaid
flowchart TB

admin[admin]
logged_user[logged user]
barber[barber]

class admin admin
class barber barber
class logged_user user
``` -->

```mermaid
flowchart TB

%% main user
logged_user["<div class='nil' style='width:200px;height:250px'><img src='stick.png' alt='kek'></div>"]

%% admin subgraph

admin["<div class='admin' style='width:200px;height:250px'><img src='stick.png' alt='kek'></div>"]

admin --- browse_users
admin --- user_analytics
admin --- create_shop

subgraph  

	%% entities
	browse_users([browse users])
	find_user([find user])
	view_user([view user])
	delete_user([delete user])
	modify_perm([modify permissions])
	user_analytics([view app analytics])

	%% relations
	browse_users -.include.-> find_user
	find_user -.include.->view_user
	delete_user-.extends.->view_user
	modify_perm-.extends.->view_user
	user_analytics
	create_shop
end

%% generic shop entities

logged_user --- browse_shops
logged_user ---- view_profile_info

subgraph  

	%% entities
	browse_shops(["browse barber shops"])
	find_shops([find shops])
	view_shop([view shop])
	comment([comment])
	view_comments([view comments])
	booking([book an appointment])
	modify_shop([modify shop info])
	add_holidays([add holidays])
	rep_comments([report comment])
	view_appointments([view appointments])
	delete_appointment([delete an appointment])
	view_shop_analytics([view shop analytics])
	delete_shop([delete shop])
	delete_comment([delete comments])

	%% relations
	browse_shops -.include.-> find_shops
	view_shop -.extends.-> find_shops
	view_shop -.include.-> view_appointments
	delete_appointment -.extends.-> view_appointments
	delete_comment -.extends.->view_comments
	view_shop_analytics -.extends.-> view_shop
	comment -.extends.-> view_shop
	rep_comments -.extends.->view_comments
	modify_shop -.extends.->view_shop
	add_holidays -.extends.->modify_shop
	view_comments-.extends.-> view_shop
	booking -.extends.-> view_shop
	delete_shop -.extends.->view_shop
end


%% user profile

subgraph  

	%% entities
	view_profile_info([view profile info])
	pswd_rec([password recovery])
	del_acc([delete account])
	curr_appointment([view current appointment])
	del_appointment([delete appointment])

	%% relations
	pswd_rec -.extends.->view_profile_info
	curr_appointment -.extends.->view_profile_info
	del_acc -.extends.->view_profile_info
	del_appointment-.extends.->curr_appointment
end

class modify_shop barber
class add_holidays barber
class dis_comments barber
class barber_user barber
class view_appointments barber
class view_shop_analytics adminbarber
class delete_appointment barber
class admin admin
class delete_shop admin
class delete_comment admin
class browse_shops all
class browse_users admin
class find_user admin
class view_user admin
class delete_user admin
class modify_perm admin
class create_shop admin
class user_analytics admin
class rep_comments barber
class curr_appointment user
class view_profile_info all
class pswd_rec all
class del_appointment user
class del_acc barberuser
class find_shops all
class view_shop all
class view_comments all
class comment user
class booking user
class logged_user nil

```

