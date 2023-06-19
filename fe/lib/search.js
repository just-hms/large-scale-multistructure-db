import { headers, url } from "./request-utils";

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
      "latitude": geocoding.geocodes[0].Latitude,
      "longitude": geocoding.geocodes[0].Longitude,
      "radius": 3000,
    })
  })
  return data
}

export async function findShopByName(name){
  const data = await fetch(url+`barbershop`, {
    method: 'POST',
    headers: headers(localStorage.getItem("token")),
    body: JSON.stringify({
      "name": name
    })
  })
  return data
}