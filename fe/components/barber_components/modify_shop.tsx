import { useState } from 'react';
import ModifiedShop from './modified_shop';
import { Menu, Transition } from '@headlessui/react'
import { Fragment} from 'react'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faBarsStaggered } from '@fortawesome/free-solid-svg-icons' 
import React from 'react';

export default function ModifyShop({shops}:any) {
  // dropdown shenanigans
  let dropdownElements = []
  for(var element in shops){
    dropdownElements.push(shops[element])
  }
  const [shopData,setShopData] = useState(dropdownElements[0])
  const [selectedShop,setSelectedShop] = useState(dropdownElements[0].Name)
  return (
    <>
    <div className='flex flex-col items-end justify-start h-full w-full px-5'>
        {/* DROPDOWN */}
        <div className="inline-block leading-none px-1 mr-4 py-2 rounded-full bg-slate-700 bg-opacity-60 backdrop-blur-lg drop-shadow-lg hover:bg-slate-700 my-3 hover:text-slate-500 text-slate-200">
          <Menu as="div" className="relative inline-block">
              <Menu.Button className="inline-flex w-full justify-center items-center rounded-full bg-opacity-20 text-slate-200 focus:outline-none">
                  <div className=' px-1 flex hover:text-white'>
                      {selectedShop}
                  </div>
                  <FontAwesomeIcon  icon={faBarsStaggered} className="  pr-2"/>
              </Menu.Button>
              <Transition
              as={Fragment}
              enter="transition ease-out duration-100"
              enterFrom="transform opacity-0 scale-95"
              enterTo="transform opacity-100 scale-100"
              leave="transition ease-in duration-75"
              leaveFrom="transform opacity-100 scale-100"
              leaveTo="transform opacity-0 scale-95"
              >
                  <Menu.Items className="absolute right-0 mt-3 w-56 origin-top-right divide-y divide-slate-600 rounded-md bg-slate-800 shadow-sm ring-1 ring-black ring-opacity-5 focus:outline-none z-10 shadow-md shadow-black/70">
                      <div  className="px-1 py-1">
                          {dropdownElements.map((element:any)=>
                          <div key={`container-`+element.ID}  className="px-1 py-1">
                              <Menu.Item key={`item-`+element.ID} >
                                  {({ active }) => (
                                  <button key={`button-`+element.ID}  className={`hover:bg-slate-500/80 text-white group flex w-full items-center rounded-md px-2 py-2 `}
                                  onClick={async (e) => {  
                                    setShopData(element)
                                    setSelectedShop(element.Name)
                                  }}>
                                      {element.Name}
                                  </button>
                                  )}
                              </Menu.Item>
                          </div>
                          )}
                      </div>
                  </Menu.Items>
              </Transition>
          </Menu>
      </div>
      <ModifiedShop shopData={shopData}></ModifiedShop>
    </div>
    </>
  );
}
