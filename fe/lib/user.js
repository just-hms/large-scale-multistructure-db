import { headers, url } from "./request-utils";

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
  const response = await fetch(url+'user/self', {
    method: 'GET',
    headers: headers(localStorage.getItem("token"))
  })
  return response;
}

export async function signup(values){
  const response = await fetch(url+'user', {
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
  const response = await fetch(url+'user/login', {
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
  const response = await fetch(url+'user/self', {
    method: 'DELETE',
    headers: headers(localStorage.getItem("token"))
  })
  return response;
}

// TODO: !!!
export async function changePassword(values){
  const response = await fetch(url+'user/lostpassword', {
    method: 'POST',
    headers: headers(),
    body: JSON.stringify({
      "email": values.email,
    })
  })
  if(response.ok){  
    const response_json = await response.json()
    const reset_password = await fetch(url+'user/resetpassword/'+encodeURIComponent(response_json.resetToken), {
      method: 'POST',
      headers: headers(),
      body: JSON.stringify({
              "NewPassword": values.password,
      })
    })
    return reset_password
  }
  return response
}
