import { useState } from "react";
import React from "react"
import SingleAccountManagement from "./single_account_management";
import { getAccountInfos } from "../../lib/admin";

export default function ManageUsers() {
  // change to useEffect triggered by onsubmit which changes the users shown
  const [users, setUsers] = useState([])
  const [query, setQuery] = useState('');
  //Handling the input on our search bar
  const handleChange = (e:any) => {
    setQuery(e.target.value)
  }
  const handleSearch = async (e:any) => {
    const response = await (await getAccountInfos(query)).json()
    setUsers(response.users)
  }

  return (
  <div className='flex flex-col items-start justify-start text-left text-slate-300 text-lg w-full'>
      <div className="w-full h-full bg-slate-800 ">
          <div className="w-full flex justify-center items-start">
            <div className='relative w-full lg:w-3/4 max-h-128 overflow-auto rounded-3xl shadow-md shadow-black/70'>
              <div className="sticky top-0 bg-slate-700 w-full flex flex-col items-center justify-center border-b border-slate-600 px-5 pt-5">
                <p className="text-2xl font-bold">Accounts</p> 
                <div className="flex flex-col lg:p-0 lg:flex-row items-center justify-between ">
                  {/* SEARCH BAR */}
                  <div className=" text-lg text-center font-bold leading-tight tracking-tight text-slate-300 break-words p-5">
                    <div className='w-full py-2 my-5 lg:m-5 flex items-center justify-center rounded-full bg-slate-800 bg-opacity-60 backdrop-blur-lg drop-shadow-lg'>
                      <button className="btn inline-block ml-3 text-slate-200 font-medium text-sm leading-tight uppercase rounded-full flex items-center" 
                      onClick={handleSearch}>
                        <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="search" className="w-4 " role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                          <path className='fill-slate-100/90 drop-shadow-lg backdrop-blur-lg '  d="M505 442.7L405.3 343c-4.5-4.5-10.6-7-17-7H372c27.6-35.3 44-79.7 44-128C416 93.1 322.9 0 208 0S0 93.1 0 208s93.1 208 208 208c48.3 0 92.7-16.4 128-44v16.3c0 6.4 2.5 12.5 7 17l99.7 99.7c9.4 9.4 24.6 9.4 33.9 0l28.3-28.3c9.4-9.4 9.4-24.6.1-34zM208 336c-70.7 0-128-57.2-128-128 0-70.7 57.2-128 128-128 70.7 0 128 57.2 128 128 0 70.7-57.2 128-128 128z"></path>
                        </svg>
                      </button>
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
                users.map((account:any)=>
                  <SingleAccountManagement key={account.ID} userkey={account.ID} account={account}/>
                )}
              </div>
            </div>
        </div>
      </div>
    </div>
  )
}