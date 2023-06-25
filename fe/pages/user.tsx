import Head from 'next/head'
import Navbar from '../components/navbar'
import {useEffect, useState, useRef} from 'react';
import UserInfos from '../components/user_components/account_infos';
import AccountReservation from '../components/user_components/account_reservations';
import Footer from '../components/footer';
import {getUserInfos}  from '../lib/user';
import { useRouter } from 'next/router';


export default function User() {
  const router = useRouter()
  const [loaded,setLoaded] = useState(false)
  const [content, setContent] = useState("account_info");
  const [userData, setUserData] = useState<any[]>([])
  const [reservationData, setReservationData] = useState<any[]>([])
  let displayed_element;

  useEffect(()=>{
    if(!localStorage.getItem('token')){
      router.push("/")
    }else{
      const fetchData = async () => {
        const userInfos = await (await getUserInfos()).json()
        setUserData(userInfos)
        setReservationData(userInfos.user.CurrentAppointment)
        setLoaded(true)
      }
      fetchData()
    }
  },[])

  if(!loaded){
    return <div></div> //show nothing or a loader
  }
  if (content == "account_info") {
    displayed_element = <UserInfos userdata={userData}/>;
  } else if (content == "account_reservation"){
    displayed_element = <><AccountReservation reservationData={reservationData}/></>;
  } else if(content == ""){
    }

  return (
    <>
    <Head>
      <title>Account | Barber Shop</title>
      <link rel="icon" type="image/png" sizes="32x32" href="/barber-shop.png"></link>
    </Head>
    <Navbar/>
    <svg className='w-full bg-slate-800 h-full' viewBox='0 0 1440 100' preserveAspectRatio="xMidYMid">
        <path className='w-full fill-slate-900' d="M 0 90 C 480 0 600 0 720 10.7 C 840 21 960 43 1080 48 C 1200 53 1320 43 1380 37.3 L 1440 32 L 1440 0 L 1380 0 C 1320 0 1200 0 1080 0 C 960 0 840 0 720 0 C 600 0 480 0 360 0 C 240 0 120 0 60 0 L 0 0 Z"></path>
    </svg>
    <div className="flex flex-col lg:flex-row justify-center items-start w-full bg-slate-800 px-5 lg:pl-10 pb-10 h-full">
        <div className='w-full lg:w-1/5 lg:h-screen border-b lg:border-r lg:border-b-0 border-slate-500 text-slate-300 mb-2.5 pb-2.5'>
            <ul className='flex flex-row lg:flex-col items-center lg:items-start justify-between '>
                <li>
                    <button className={`hover:text-white focus:outline-none ${content == "account_info" ? "font-bold" : ""}`} onClick={event => {setContent("account_info")}}>Account</button>
                </li>
                <li>
                    <button className={`hover:text-white focus:outline-none ${content == "account_reservation" ? "font-bold" : ""}`} onClick={event => {setContent("account_reservation")}}>Last booked session</button>
                </li>
            </ul>
        </div>

        <div className='w-full lg:w-4/5 h-full flex justify-center'>
            {displayed_element}
        </div>
    </div>
    <Footer/>
    </>
  )
}