import { useState, useEffect } from "react";
import React from "react"
import { deleteUser } from "../../lib/admin";
import { faTrash } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { assignShop } from "../../lib/admin";
import { getShopData } from "../../lib/shops";

export default function SingleAccountManagement({account,userkey}:any) {

  const [shopid, setShopid] = useState('');
  const [ownedShopsNames, setOwnedShopsNames] = useState<any[]>([])
  const handleChangeShopid = (e:any) => {
    setShopid(e.target.value)
  }  
  const handlePasteShopid = (e:any) => {
    setShopid(e.clipboardData.getData('Text'))
  }
  const addShop = async (e:any) => {
    let shops:any = []
    // iterate through already owned shops
    if(account.OwnedShops){
      account.OwnedShops.forEach((shop:any) => {
        shops.push(shop)
    });
      shops.push(shopid)
    }else{
      shops = [shopid]
    }
    // add shop
    const response = await assignShop(account.ID, shops)
    if(response.status == 202)
        window.location.reload()
  }
  useEffect(()=>{
    const fetchData = async()=>{
      let shops = []
      for(var i in account.OwnedShops){
        var infos = await(await getShopData(account.OwnedShops[i])).json()
        console.log(infos)
        shops.push(infos.barbershop.Name)
      }
      setOwnedShopsNames(shops)
    }
    fetchData()
  },[])

  return (
    <div key={userkey} className="w-full text-slate-200 my-4 px-2 flex flex-col items-center justify-start">
    <div key={userkey+"container"} className=" flex flex-col bg-slate-700 items-center justify-start w-full rounded-lg">
        <div key={userkey+"main_container"} className="flex w-full items-start justify-start">
            <div key={userkey+"title"} className="flex flex-col px-3 items-center lg:items-start justify-start w-full text-left">
                  <div key={userkey+"container"} className="flex w-full items-center py-1 justify-between">
                    <p key={userkey+"email"} className=" text-xl font-bold">{(account.Username)?<>{account.Username} |</>:<></>} {account.Email}</p>
                    <button key={userkey+"delete_button"} type="button" id="delete_button"
                    onClick={async ()  =>{
                      const response = await deleteUser(account.ID)
                      if(response.status === 202)
                        window.location.reload()
                    }}>        
                      <FontAwesomeIcon key={userkey+"icon"} icon={faTrash} className="text-xl pr-1 text-slate-400 hover:text-slate-100"/>
                    </button>
                  </div>

                  <p key={userkey+"account_type"} className="w-full pb-1 border-b border-slate-500">Account type: {account.Type}</p>
                    {/* GET SHOPS NAMES AND SHOW */}
                  {(account.OwnedShops)?<><div key={userkey+"owned_shop"} className="w-full pb-1 border-b border-slate-500">Owned Shops:{
                    ownedShopsNames.map( (shop:any)=>{
                    return <div key={userkey+"-"+shop}>- {shop}</div>
                    })
                  }</div></>:<></>}
                  <label key={userkey+"shop_to_assign"} htmlFor="shoptoassign">Shop To Assign:</label>
                  <input key={userkey+"shop_input"} id={"shoptoassign-"+account.ID} type="text" placeholder="shopID" className="w-1/2 my-2 font-bold text-slate-100 px-3 py-1 bg-slate-600 bg-clip-padding rounded-full transition ease-in-out focus:outline-none"
                  onChange={handleChangeShopid}
                  onPaste={handlePasteShopid}/>

                  <button key={userkey+"assign_shop_button"} className="px-4 py-2 my-2 bg-rose-800 text-slate-300 text-xs rounded-full focus:bg-rose-700 hover:bg-rose-700 focus:outline-none transition duration-150 ease-in-out " type="button" id="assign_button"
                  onClick={addShop}>
                      Assign Shop
                  </button>

            </div>
        </div>
    </div>
</div>
  )
}