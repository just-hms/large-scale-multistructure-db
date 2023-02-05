import Head from 'next/head'
import Navbar from '../components/navbar'
import Image from 'next/image'
import barber_background from '../public/barber_bg_1.png'
import barber_background_vertical from '../public/barber_bg_vertical.png'
import React from 'react';
import { useState } from 'react';
import { useRouter } from 'next/router'

export default function Home() {

  const [query, setQuery] = useState('');
  const router = useRouter()
  const handleChange=(event)=>{
    setQuery(event.target.value);
  }
  function handleClick() {
    if(query.length > 0)
      router.push("/search?area="+query);
  }

  return (
    <>
      <Head>
        <title>Home | Barber Shop</title>
        <link rel="icon" type="image/png" sizes="32x32" href="/barber-shop.png"></link>
      </Head>
      <Navbar/>
      <div className="w-full flex-col justify-center items-center bg-slate-900 h-screen">
        <div className="w-full h-full">
          {/* large screen image */}
          <Image className="top-0 hidden lg:inline lg:w-full h-full object-cover z-0" src={barber_background} alt="barber salon"/>
          {/* small screen image */}
          <Image className="top-0 lg:hidden display w-full  h-full object-cover z-0" src={barber_background_vertical} alt="barber salon"/>
          <div className="absolute w-full px-3 lg:p-0 lg:w-auto top-2/3 lg:top-2/3 left-1/2 transform -translate-x-1/2 -translate-y-1/2 inset-0 flex justify-center items-center">
            <div className=" rounded-3xl w-full bg-slate-700 bg-opacity-60 backdrop-blur-lg drop-shadow-lg">
              <h1 className="text-4xl text-center font-bold leading-tight tracking-tight text-slate-200 pt-5">
                Cut it out
              </h1>
              <h2 className="text-lg text-center font-bold leading-tight tracking-tight text-slate-300 pt-5 ">
                Find your future Favourite Barber
              </h2>
              <div className='flex justify-center'>
                <div className='w-full lg:w-1/2 m-5 flex items-center justify-center rounded-full bg-red-900 bg-opacity-60 backdrop-blur-lg drop-shadow-lg'>
                  <input
                    type="search"
                    className="w-full font-bold text-slate-100 pl-5 bg-slate-700/0 bg-clip-padding rounded-full transition ease-in-out focus:outline-none" 
                    form-control id="barberSearch" placeholder="Desired area"
                    onChange={handleChange}
                  />
                  <button className="btn inline-block px-6 py-2.5 m-1 bg-red-800 bg-opacity-60 backdrop-blur-lg drop-shadow-lg text-white font-medium text-xs leading-tight uppercase rounded-full focus:bg-red-800 hover:bg-red-800 focus:outline-none transition duration-150 ease-in-out flex items-center" type="button" id="search_button"
                    onClick={handleClick}>
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
    </>
  )/*  */
}