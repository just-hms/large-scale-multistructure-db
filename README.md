## Needed json format for the APIs:
- Login:
```
{
      token:,
      <!-- or any other way to identify type of user -->
      isAdmin:"...",
      isBarber:"...",
      <!--  -->
}
```
- Reviews (search by shop):
```
{
      id:1111,
      name/user(qualcosa che identifichi):"...",
      title:"...",
      review:"...",
      upvotes:#upvotes,
      vote:number of stars[1-5],
}
```
- Accounts:
```
{
    id:1111,
    name/user:"...",
}
```
- Reservations:need to be able to retrieve by user and by shop
- search by shop (all the upcoming ones):
```
{
    id:1111,
    date:"27/02/1998",
    time:time_slot,
    user:user_reserved
}
```
- search by user (only the last one):
```
{
    id:1111,
    shop:"shop_name",
    shop_id:shop_id
    date:"27/02/1998",
    time:time_slot
}
```
- Shops (search by geographic area):
```
{
    shop:'shop_name',
    shop_id:shop_id,
    meanRating:mean_stars_reviews,
    reviewNumber:#reviews,
    description:"small description of the shop",
    distance:calculated via coordinates,
    image:""
}
```
- Shop infos (search by shop):
```
{
    title:"shop_title",
    name:"shop",
    shop_id:shop_id
    description:""
}
```