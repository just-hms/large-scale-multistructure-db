import { Menu, Transition } from '@headlessui/react'
import { Fragment} from 'react'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faBarsStaggered } from '@fortawesome/free-solid-svg-icons' 
import {useEffect, useState, useRef, createContext, useContext } from 'react';
import React from 'react';

export default function GeneralDropdown({placeholder,elements,children,classname}:{placeholder:any,elements:any,children:any,classname:any}) {
    if(placeholder == undefined)
        placeholder = elements[0]
    const [selected_value,setSelectedValue] = useState(placeholder)
  return (
    <div className={`inline-block leading-none `+` ` + classname}>
        <Menu as="div" className="relative inline-block">
            <Menu.Button className="inline-flex w-full justify-center items-center rounded-full bg-opacity-20 text-slate-200 focus:outline-none">
                
                <div className=' px-1 flex hover:text-white'>
                    <div className='pr-1'>{children}</div>
                    {selected_value}
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
                        {elements.map((element:any)=>
                        <div key={`container-`+element}  className="px-1 py-1">
                            <Menu.Item key={`item-`+element} >
                                {({ active }) => (
                                <button key={`button-`+element}  className={`hover:bg-slate-500/80 text-white group flex w-full items-center rounded-md px-2 py-2 `}
                                onClick={event => {setSelectedValue(element);}}>
                                    {element}
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
  )
}