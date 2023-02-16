import {headers} from "./request-utils"


const url = "http://127.0.0.1:7000/api/" 
export function getReservation(user) {
    // this function will retrieve all the shop reviews
    return {
      id:1111,
      name:"Barbiere di Siviglia",
      date:"27/02/1998",
    };
  }

export async function getUserInfos(){
  const token = localStorage.getItem("token")
  // req.Header.Add("Authorization", "Bearer "+tc.token)
  const response = await fetch(url+'user/self/', {
    method: 'GET',
    headers: headers(localStorage.getItem("token"))
  })
  return response;
}

export async function signup(values){
  const response = await fetch(url+'user/', {
    method: 'POST',
    headers: headers(),
    body: JSON.stringify({
            "email": values.email,
            "password":values.password
    })
  })
  return response;
}

export async function signin(values){
  const response = await fetch(url+'user/login/', {
    method: 'POST',
    headers: headers(),
    body: JSON.stringify({
            "email": values.email,
            "password":values.password
    })
  })
  return response;
}

export async function deleteAccount(){
  const response = await fetch(url+'user/self/', {
    method: 'DELETE',
    headers: headers(localStorage.getItem("token"))
  })
  return response;
}

// TODO: !!!
export async function changePassword(values){
  const response = await fetch(url+'user/lost_password/', {
    method: 'POST',
    headers: headers(),
    body: JSON.stringify({
            "username": values.email,
    })
  })
  if(response.ok){
    const response_json = await response.json()
    const reset_password = await fetch(url+'user/reset_password?reset_token='+encodeURIComponent(response_json.resetToken), {
      method: 'POST',
      headers: headers(),
      body: JSON.stringify({
              "new_password": values.password,
      })
    })
    return reset_password
  }
}
