const url = "http://127.0.0.1:7000/api/" 
import {headers} from "./request-utils"

export async function findShops(lat,lon){
const response = await fetch(url+`barber_shop?lat=""&lon=""&name=""&radius=""`, {
    method: 'GET',
    headers: headers(localStorage.getItem("token"))
})
console.log(response)
return response;
}