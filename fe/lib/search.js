import { headers, url } from "./request-utils";

// export async function findShops(lat,lon){
// const response = await fetch(url+`barbershop`, {
//     method: 'POST',
//     headers: headers(localStorage.getItem("token")),
//     body: JSON.stringify({
//       "latitude": lat,
//       "longitude": lon,
//       "radius": 100000,
//     })
// })
// return response;
// }

export async function findShops(area){
  const geocoding = await fetch(url+`geocoding/search`,{
    method: 'POST',
    headers: headers(localStorage.getItem("token")),
    body: JSON.stringify({
      "area": area,
    })
  }).then((response)=>response.json())

  const data = await fetch(url+`barbershop`, {
    method: 'POST',
    headers: headers(localStorage.getItem("token")),
    body: JSON.stringify({
      "latitude": geocoding.geocodes.Latitude,
      "longitude": geocoding.geocodes.Longitude,
      "radius": 100000,
    })
  })
  return data
}