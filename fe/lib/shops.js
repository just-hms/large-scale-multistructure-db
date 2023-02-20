import { headers } from "./request-utils";

const url = "http://127.0.0.1:7000/api/" 
export async function getShopData(shopid) {
  const token = localStorage.getItem("token")
  const response = await fetch(url+'barber_shop/'+shopid, {
    method: 'GET',
    headers: headers(localStorage.getItem("token"))
  })
  return response;
}

export function getReviews(shopid){
}
  
export async function submitReview(shopid){
  const token = localStorage.getItem("token")
  const response = await fetch(url+'/barber_shop/'+shopid+'/review/', {
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