import { headers, url } from "./request-utils";

export async function modifyShopDescription (shopid,values){
  const response = await fetch(url+`barbershop/`+shopid, {
    method: 'PUT',
    headers: headers(localStorage.getItem("token")),
    body: JSON.stringify({
      "description": values.description
    })
  })
  return response;
}
