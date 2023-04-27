import { headers, url } from "./request-utils";

export async function getShopData(id) {
  const token = localStorage.getItem("token")
  const response = await fetch(url+'barber_shop/'+id, {
    method: 'GET',
    headers: headers(localStorage.getItem("token"))
  })
  return response;
}

export function getReviews(id){
}

export async function shopCalendar(id){
  const response = await fetch(url+`barbershop/`+id+`/calendar`, {
      method: 'GET',
      headers: headers(localStorage.getItem("token"))
  })
  return response;
}
export async function submitReview(id){
  const token = localStorage.getItem("token")
  const response = await fetch(url+'/barber_shop/'+id+'/review/', {
    method: 'POST',
    headers: headers(localStorage.getItem("token")),
    body: JSON.stringify({
      "title" : "kek",
      "body" : "kekkeroni",
      "rating" : 3
    })
  })
  return response;
}