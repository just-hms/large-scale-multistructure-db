const url = "http://127.0.0.1:7000/api/" 
import {headers} from "./request-utils"

// TODO
export async function findShops(lat,lon){
const response = await fetch(url+`barbershop`, {
    method: 'POST',
    headers: headers(localStorage.getItem("token")),
    body: JSON.stringify({
      "latitude": lat,
      "longitude": lon,
      "radius": 10,
    })
})
return response;
}