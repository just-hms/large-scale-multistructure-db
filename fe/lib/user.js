export function getReservation(user) {
    // this function will retrieve all the shop reviews
    return {
      id:1111,
      name:"Barbiere di Siviglia",
      date:"27/02/1998",
    };
  }
export async function signup(values){
  const response = await fetch('http://127.0.0.1:7000/user/', {
    method: 'POST',
    headers: new Headers({
        'Content-Type': 'application/json',
    }),
    body: JSON.stringify({
            "email": values.email,
            "password":values.password
    })
  })
  return response;
}