import { getAccount } from "../../lib/admin"
import GeneralDropdown from "../general_dropdown"
import React from "react"
// import { useDropdownContext, DropdownContext} from "../search_components/general_dropdown"

export default function ManageUsers() {
  return (
      <div className='flex flex-col items-center justify-center text-left text-slate-300 text-lg w-full'>
        <div className="w-full h-full m-10 ml-15 mt-0">
              <div className="w-full flex justify-center items-center">
                <div className='relative max-h-96 rounded-3xl bg-slate-700 shadow-md shadow-black/70'>
                  <div className="w-full flex flex-col items-center justify-center p-5 pt-0">
                    <h1 className="text-2xl text-center font-bold leading-tight tracking-tight text-slate-200 sticky top-0 bg-slate-700 w-full border-b border-slate-600 pt-3 pb-3">
                      Accounts
                    </h1>
                    <div className="flex flex-col py-2 lg:p-0 lg:flex-row items-center justify-between ">
                      {/* <div className="px-3 py-2 rounded-full bg-slate-800 bg-opacity-60 backdrop-blur-lg drop-shadow-lg"> */}
                        <GeneralDropdown elements={["Barber", "User"]} placeholder="Type" classname="px-1 py-2 hover:text-slate-500 rounded-full text-slate-200 bg-slate-700 shadow-sm shadow-slate-900/60"><></></GeneralDropdown>
                      {/* </div> */}
                      {/* SEARCH BAR */}
                      <div className=" text-lg text-center font-bold leading-tight tracking-tight text-slate-300 break-words p-3">
                        <div className='flex justify-center'>
                          <div className='w-full lg:m-5 flex items-center justify-center rounded-full bg-slate-800 bg-opacity-60 backdrop-blur-lg drop-shadow-lg'>
                            <input
                              type="search"
                              className="w-full font-bold text-slate-300 pl-5 bg-slate-700/0 bg-clip-padding rounded-full transition ease-in-out focus:outline-none" 
                              form-control id="barberSearch" placeholder="Search"
                            />
                            <button className="btn inline-block px-6 py-2.5 m-1 bg-slate-700 bg-opacity-60 backdrop-blur-lg drop-shadow-lg text-slate-300 font-medium text-xs leading-tight uppercase rounded-full focus:bg-slate-600 hover:bg-slate-600 focus:outline-none transition duration-150 ease-in-out flex items-center" type="button" id="search_button">
                              <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="search" className="w-4 " role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                                <path className='fill-slate-100/90 drop-shadow-lg backdrop-blur-lg'  d="M505 442.7L405.3 343c-4.5-4.5-10.6-7-17-7H372c27.6-35.3 44-79.7 44-128C416 93.1 322.9 0 208 0S0 93.1 0 208s93.1 208 208 208c48.3 0 92.7-16.4 128-44v16.3c0 6.4 2.5 12.5 7 17l99.7 99.7c9.4 9.4 24.6 9.4 33.9 0l28.3-28.3c9.4-9.4 9.4-24.6.1-34zM208 336c-70.7 0-128-57.2-128-128 0-70.7 57.2-128 128-128 70.7 0 128 57.2 128 128 0 70.7-57.2 128-128 128z"></path>
                              </svg>
                            </button>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
        </div>
  )
}


// TODO: For each account we gotta return the mail and the actions, i.e: delete account, change permissions, on click over button show results if any