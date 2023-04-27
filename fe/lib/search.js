import { headers, url } from "./request-utils";

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
