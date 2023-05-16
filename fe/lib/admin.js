import { Coming_Soon } from "@next/font/google";
import { headers, url } from "./request-utils";


export async function getAccountInfos(){
  const response = await fetch(url+`admin/user`, {
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

// export async function getReviews(shop) {
//     // this function will retrieve all the shop reviews
//     return [{
//       id:1111,
//       name:"Pippo Baudo",
//       shop:"Barbiere di Siviglia",
//       title:"Gatti fritti",
//       review:"Distanza dal ristorante: 950m 4 ordini totali richiesti al momento della recensione: 2 consegne e 2 cancellazioni. Qualità, quantità e prezzo del ristorante sono eccellenti in loco, ma il servizio relativo alle consegne è del tutto inadeguato. Entrambe le volte che ho ricevuto la consegna il cibo è arrivato danneggiato in qualche modo. Particolarmente grave il caso del Mafè (composto da abbondante salsa di consistenza liquida, oleosa) spedito in contenitori di stagnola con tappo di carta. ",
//       upvotes:10,
//       vote:3,
//     },{
//       id:1112,
//       name:"Pippo Baudo",
//       shop:"Barbiere di Siviglia",
//       title:"Gatti fritti",
//       review:"Distanza dal ristorante: 950m 4 ordini totali richiesti al momento della recensione: 2 consegne e 2 cancellazioni. Qualità, quantità e prezzo del ristorante sono eccellenti in loco, ma il servizio relativo alle consegne è del tutto inadeguato. Entrambe le volte che ho ricevuto la consegna il cibo è arrivato danneggiato in qualche modo. Particolarmente grave il caso del Mafè (composto da abbondante salsa di consistenza liquida, oleosa) spedito in contenitori di stagnola con tappo di carta. ",
//       upvotes:-10,
//       vote:5,
//     },{
//       id:1113,
//       name:"Pippo Baudo",
//       shop:"Barbiere di Siviglia",
//       title:"Gatti fritti",
//       review:"Distanza dal ristorante: 950m 4 ordini totali richiesti al momento della recensione: 2 consegne e 2 cancellazioni. Qualità, quantità e prezzo del ristorante sono eccellenti in loco, ma il servizio relativo alle consegne è del tutto inadeguato. Entrambe le volte che ho ricevuto la consegna il cibo è arrivato danneggiato in qualche modo. Particolarmente grave il caso del Mafè (composto da abbondante salsa di consistenza liquida, oleosa) spedito in contenitori di stagnola con tappo di carta. ",
//       upvotes:10,
//       vote:2,
//     }];
// }
// export async function getAccount(type) {
//   // this function will retrieve all the shop reviews
//   return [{
//     id:1111,
//     name:"Pippo Baudo",
//   },{
//     id:1112,
//     name:"Pippo Baudo",
//   },{
//     id:1113,
//     name:"Pippo Baudo",
//   }];
// }
