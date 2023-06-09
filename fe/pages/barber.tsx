import Head from 'next/head'
import Navbar from '../components/navbar'
import {useEffect, useState, useRef} from 'react';
import UserInfos from '../components/user_components/account_infos';
import Footer from '../components/footer';
import {getOwnedShops} from '../lib/barber'
import ModifyShop from '../components/barber_components/modify_shop';
import BarberReservations from '../components/barber_components/barber_reservations';
import { useRouter } from 'next/router';
import { getUserInfos } from '../lib/user';
import AnalytcsContainer from '../components/barber_components/analytics_container';
export default function User() {

  const [content, setContent] = useState("account_info");
  const [loaded,setLoaded] = useState(false)
  const [userData, setUserData] = useState<any[]>([])
  const [shopsData, setShopsData] = useState<any[]>([])
  const router = useRouter()
  let displayed_element;
  // check if logged in and barber
  useEffect(()=>{
    const token = localStorage.getItem('token')
    if(!token){
      router.push("/")
    }else{
      const fetchData = async () => {
        const retrievedData = await (await getUserInfos()).json() 
        // if anyone tries to access without being a barber -> unauthorized
        if(retrievedData.user.Type !== 'barber'){
          router.push("/401")
        }else{
          setUserData(retrievedData)
          const ownedShops = await (await getOwnedShops()).json()
          setShopsData(ownedShops.barbershops)
          setLoaded(true)
        }
      }
      fetchData()
    }
  },[])
  if (content == "account_info") {
    displayed_element = <UserInfos userdata={userData}/>;
  } else if (content == "modify_shop"){
    displayed_element = <ModifyShop shops={shopsData}/>;
  } else if (content == "reservations"){
    displayed_element = <BarberReservations shops={shopsData}/>;
  } else if (content == "analytics"){
    displayed_element = <AnalytcsContainer shops={shopsData}></AnalytcsContainer>;
  } 
  
  if(!loaded){
    return <div></div> 
  }

  return (
    <>
    <Head>
      <title>Account | Barber Shop</title>
      <link rel="icon" type="image/png" sizes="32x32" href="/barber-shop.png"></link>
    </Head>
    <Navbar/>
    <svg className='w-full bg-slate-800 h-full' viewBox='0 0 1442 100' preserveAspectRatio="xMidYMid">
        <path className='w-full fill-slate-900' d="M 0 90 C 480 0 600 0 720 10.7 C 840 21 960 43 1080 48 C 1200 53 1320 43 1380 37.3 L 1440 32 L 1440 0 L 1380 0 C 1320 0 1200 0 1080 0 C 960 0 840 0 720 0 C 600 0 480 0 360 0 C 240 0 120 0 60 0 L 0 0 Z"></path>
    </svg>
    <div className="flex flex-col lg:flex-row justify-center items-start w-full bg-slate-800 px-5 lg:pl-10 pb-10 h-full">
      {/* menu */}
        <div className='w-full lg:w-1/5 lg:h-screen border-b lg:border-r lg:border-b-0 border-slate-500 text-slate-300 mb-2.5 pb-2.5'>
            <ul className='flex flex-row lg:flex-col items-center lg:items-start justify-between'>
                <li className='mx-2 '>
                    <button className={`hover:text-white focus:outline-none ${content == "account_info" ? "font-bold" : ""}`} onClick={event => {setContent("account_info")}}>Account</button>
                </li>
                <li className='mx-2 '>
                    <button className={`hover:text-white focus:outline-none ${content == "modify_shop" ? "font-bold" : ""}`} onClick={event => {setContent("modify_shop")}}>Modify Shop</button>
                </li>
                <li className='mx-2 '>
                    <button className={`hover:text-white focus:outline-none ${content == "reservations" ? "font-bold" : ""}`} onClick={event => {setContent("reservations")}}>Reservations</button>
                </li>
                <li className='mx-2 '>
                    <button className={`hover:text-white focus:outline-none ${content == "analytics" ? "font-bold" : ""}`} onClick={event => {setContent("analytics")}}>Analytics</button>
                </li>
                {/* <li className='mx-2 '>
                    <button className={`hover:text-white focus:outline-none ${content == "calendar" ? "font-bold" : ""}`} onClick={event => {setContent("calendar")}}>Calendar</button>
                </li> */}
            </ul>
        </div>

        <div className='w-full h-full flex justify-center'>
            {displayed_element}
        </div>
    </div>
    <Footer/>
    </>
  )
}

