import { Coming_Soon } from "@next/font/google";
import { headers, url } from "./request-utils";


export async function getAccountInfos(email){
  const response = await fetch(url+`admin/user?email=`+email, {
    method: 'GET',
    headers: headers(localStorage.getItem("token"))
  })
  return response;
}

export async function deleteUser (id){
  const response = await fetch(url+`admin/user/`+id, {
    method: 'DELETE',
    headers: headers(localStorage.getItem("token"))
  })
  return response;
}

export async function createShop(values){
  const geocoding = await fetch(url+`geocoding/search`,{
    method: 'POST',
    headers: headers(localStorage.getItem("token")),
    body: JSON.stringify({
      "area": values.address,
    })
  }).then((response)=>response.json())
  const lat = geocoding.geocodes.Latitude
  const lon = geocoding.geocodes.Longitude
  const response = await fetch(url+`admin/barbershop`, {
    method: 'POST',
    headers: headers(localStorage.getItem("token")),
    body: JSON.stringify({
      "employees_number": values.employeesNumber,
      "name": values.name,
      "description":values.shopDescription,
      "Latitude": lat,
      "Longitude": lon
    })
  })
  return response
}

export async function modifyShop(values, id){
  const geocoding = await fetch(url+`geocoding/search`,{
    method: 'POST',
    headers: headers(localStorage.getItem("token")),
    body: JSON.stringify({
      "area": values.address,
    })
  }).then((response)=>response.json())
  const lat = geocoding.geocodes.Latitude
  const lon = geocoding.geocodes.Longitude
  const response = await fetch(url+`admin/barbershop/`+id, {
    method: 'POST',
    headers: headers(localStorage.getItem("token")),
    body: JSON.stringify({
      "employees_number":values.employeesNumber,
      "name": (values.name)?values.name:'',
      "Latitude": lat,
      "Longitude": lon
    })
  })
  return response
}

export async function deleteShop (id){
  const response = await fetch(url+`admin/barbershop/`+id, {
    method: 'DELETE',
    headers: headers(localStorage.getItem("token"))
  })
  return response;
}

export async function assignShop (id,barbershop_to_add){
  const response = await fetch(url+`admin/user/`+id, {
    method: 'PUT',
    headers: headers(localStorage.getItem("token")),
    body: JSON.stringify({
      "barbershopsId": barbershop_to_add
    })
  })
  return response;
}

export async function modifyUserEmail (id,email){
  const response = await fetch(url+`admin/user/`+id, {
    method: 'PUT',
    headers: headers(localStorage.getItem("token")),
    body: JSON.stringify({
      "email": email
    })
  })
  return response;
}
