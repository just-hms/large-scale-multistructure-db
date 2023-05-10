import { useState } from "react";
import GeneralDropdown from "../general_dropdown"
import React from "react"
import { deleteUser } from "../../lib/admin";
import { faTrash } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

export default function ManageUsers({accounts}:any) {

  // change to useEffect triggered by onsubmit which changes the users shown
  const [query, setQuery] = useState('');
  const searchFilter = (accounts:any) => {
    if(accounts){
      return accounts.filter(
        (el:any) => {
          return el.Email.toLowerCase().includes(query.toLocaleLowerCase())
        }
      )
    }
  }
  const filtered = searchFilter(accounts)

  //Handling the input on our search bar
  const handleChange = (e:any) => {
    setQuery(e.target.value)
  }

  return (
  <div className='flex flex-col items-start justify-start text-left text-slate-300 text-lg w-full'>
      <div className="w-full h-full bg-slate-800 ">
          <div className="w-full flex justify-center items-start">
            <div className='relative w-full lg:w-3/4 max-h-128 overflow-auto rounded-3xl shadow-md shadow-black/70'>
              <div className="sticky top-0 bg-slate-700 w-full flex flex-col items-center justify-center border-b border-slate-600 px-5 pt-5">
                <p className="text-2xl font-bold">Accounts</p> 
                <div className="flex flex-col lg:p-0 lg:flex-row items-center justify-between ">
                  {/* <GeneralDropdown elements={["Barber", "User"]} placeholder="Type" classname="px-1 py-2 hover:text-slate-500 rounded-full text-slate-200 bg-slate-800 bg-opacity-60 backdrop-blur-lg drop-shadow-lg"><></></GeneralDropdown> */}
                  {/* SEARCH BAR */}
                  <div className=" text-lg text-center font-bold leading-tight tracking-tight text-slate-300 break-words p-5">
                    <div className='w-full py-2 my-5 lg:m-5 flex items-center justify-center rounded-full bg-slate-800 bg-opacity-60 backdrop-blur-lg drop-shadow-lg'>
                      <div className="btn inline-block ml-3 text-slate-200 font-medium text-sm leading-tight uppercase rounded-full flex items-center" >
                        <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="search" className="w-4 " role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                          <path className='fill-slate-100/90 drop-shadow-lg backdrop-blur-lg '  d="M505 442.7L405.3 343c-4.5-4.5-10.6-7-17-7H372c27.6-35.3 44-79.7 44-128C416 93.1 322.9 0 208 0S0 93.1 0 208s93.1 208 208 208c48.3 0 92.7-16.4 128-44v16.3c0 6.4 2.5 12.5 7 17l99.7 99.7c9.4 9.4 24.6 9.4 33.9 0l28.3-28.3c9.4-9.4 9.4-24.6.1-34zM208 336c-70.7 0-128-57.2-128-128 0-70.7 57.2-128 128-128 70.7 0 128 57.2 128 128 0 70.7-57.2 128-128 128z"></path>
                        </svg>
                      </div>
                      <input
                        type="text"
                        className="w-full font-bold text-slate-100 pl-5 bg-slate-700/0 bg-clip-padding rounded-full transition ease-in-out focus:outline-none" 
                        id="barberSearch" placeholder="Search"
                        onChange={handleChange}
                      />
                    </div>
                  </div>
                </div>
              </div>
              <div className="p-3 bg-slate-800/80">
                {
                filtered.map((account:any)=>
                <div key={account.ID} className="w-full text-slate-200 my-4 px-2 flex flex-col items-center justify-start">
                    <div key={account.ID+"container"} className=" flex flex-col bg-slate-700 items-center justify-start w-full rounded-lg">
                        <div className="flex w-full items-start justify-start">
                            <div key={account.ID+"title"} className="flex flex-col px-3 items-center lg:items-start justify-start w-full text-left">
                                  <div className="flex w-full items-center py-1 justify-between">
                                    <p className=" text-xl font-bold">{account.Email}</p>
                                    <button className="" type="button" id="search_button"
                                    onClick={async ()  =>{
                                      const response = await deleteUser(account.ID)
                                      if(response.status === 202)
                                        window.location.reload()
                                    }}>        
                                      <FontAwesomeIcon icon={faTrash} className="text-xl pr-1 text-slate-400 hover:text-slate-100"/>
                                    </button>
                                  </div>
                                  <p className="w-full pb-1 border-b border-slate-500">Account type: {account.Type}</p>
                                  <button className="px-4 py-2 my-2 bg-rose-800 text-slate-300 text-xs rounded-full focus:bg-rose-700 hover:bg-rose-700 focus:outline-none transition duration-150 ease-in-out " type="button" id="search_button">
                                      Modify Permission
                                  </button>
                            </div>
                        </div>
                    </div>
                </div>
                )}
              </div>
            </div>
        </div>
      </div>
    </div>
  )
}