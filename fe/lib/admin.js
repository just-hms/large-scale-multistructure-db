import {headers} from "./request-utils"
const url = "http://127.0.0.1:7000/api/" 

export async function getAccountInfos(){
  const response = await fetch(url+`admin/user?email=""`, {
    method: 'GET',
    headers: headers(localStorage.getItem("token"))
  })
  return response;
}

export async function getReviews(shop) {
    // this function will retrieve all the shop reviews
    return [{
      id:1111,
      name:"Pippo Baudo",
      shop:"Barbiere di Siviglia",
      title:"Gatti fritti",
      review:"Distanza dal ristorante: 950m 4 ordini totali richiesti al momento della recensione: 2 consegne e 2 cancellazioni. Qualità, quantità e prezzo del ristorante sono eccellenti in loco, ma il servizio relativo alle consegne è del tutto inadeguato. Entrambe le volte che ho ricevuto la consegna il cibo è arrivato danneggiato in qualche modo. Particolarmente grave il caso del Mafè (composto da abbondante salsa di consistenza liquida, oleosa) spedito in contenitori di stagnola con tappo di carta. ",
      upvotes:10,
      vote:3,
    },{
      id:1112,
      name:"Pippo Baudo",
      shop:"Barbiere di Siviglia",
      title:"Gatti fritti",
      review:"Distanza dal ristorante: 950m 4 ordini totali richiesti al momento della recensione: 2 consegne e 2 cancellazioni. Qualità, quantità e prezzo del ristorante sono eccellenti in loco, ma il servizio relativo alle consegne è del tutto inadeguato. Entrambe le volte che ho ricevuto la consegna il cibo è arrivato danneggiato in qualche modo. Particolarmente grave il caso del Mafè (composto da abbondante salsa di consistenza liquida, oleosa) spedito in contenitori di stagnola con tappo di carta. ",
      upvotes:-10,
      vote:5,
    },{
      id:1113,
      name:"Pippo Baudo",
      shop:"Barbiere di Siviglia",
      title:"Gatti fritti",
      review:"Distanza dal ristorante: 950m 4 ordini totali richiesti al momento della recensione: 2 consegne e 2 cancellazioni. Qualità, quantità e prezzo del ristorante sono eccellenti in loco, ma il servizio relativo alle consegne è del tutto inadeguato. Entrambe le volte che ho ricevuto la consegna il cibo è arrivato danneggiato in qualche modo. Particolarmente grave il caso del Mafè (composto da abbondante salsa di consistenza liquida, oleosa) spedito in contenitori di stagnola con tappo di carta. ",
      upvotes:10,
      vote:2,
    }];
}
export async function getAccount(type) {
  // this function will retrieve all the shop reviews
  return [{
    id:1111,
    name:"Pippo Baudo",
  },{
    id:1112,
    name:"Pippo Baudo",
  },{
    id:1113,
    name:"Pippo Baudo",
  }];
}
  