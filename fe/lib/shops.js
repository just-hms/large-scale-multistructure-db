import { headers, url } from "./request-utils";

export async function getShopData(id) {
  const token = localStorage.getItem("token")
  const response = await fetch(url+'barbershop/'+id, {
    method: 'GET',
    headers: headers(localStorage.getItem("token"))
  })
  return response;
}

export async function getReviews(id){
  const response = await fetch(url+`barbershop/`+id+`/review`, {
    method: 'GET',
    headers: headers(localStorage.getItem("token"))
  })
  return response;
}

export async function shopCalendar(id){
  const response = await fetch(url+`barbershop/`+id+`/calendar`, {
      method: 'GET',
      headers: headers(localStorage.getItem("token"))
  })
  return response;
}

export async function shopAnalytics(id){
  const response = await fetch(url+`barbershop/`+id+`/analytics`, {
      method: 'GET',
      headers: headers(localStorage.getItem("token"))
  })
  return response;
}

export async function submitReview(id,values){
  const response = await fetch(url+'barbershop/'+id+'/review', {
    method: 'POST',
    headers: headers(localStorage.getItem("token")),
    body: JSON.stringify({
      "content" : values.content,
      "rating" : values.rating
    })
  })
  return response;
}

export async function deleteReview(shopid, reviewid){
  await fetch(url+'barbershop/'+shopid+'/review/'+reviewid, {
    method: 'DELETE',
    headers: headers(localStorage.getItem("token"))
  })
}

export async function submitVote(shopid, reviewid, vote){
  const response = await fetch(url+'barbershop/'+shopid+'/review/'+reviewid+'/vote', {
    method: 'POST',
    headers: headers(localStorage.getItem("token")),
    body: JSON.stringify({
      "upvote" : vote,
    })
  })
  return response;
}

export async function deleteVote(shopid, reviewid){
  await fetch(url+'barbershop/'+shopid+'/review/'+reviewid+'/vote', {
    method: 'DELETE',
    headers: headers(localStorage.getItem("token"))
  })
}


export async function getAppointment(shopid,dateTime){
  const response = await fetch(url+'barbershop/'+shopid+'/appointment', {
    method: 'POST',
    headers: headers(localStorage.getItem("token")),
    body: JSON.stringify({
      "dateTime" : dateTime,
    })
  })
  return response
}

