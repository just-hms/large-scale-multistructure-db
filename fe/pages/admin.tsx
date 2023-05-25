import Head from 'next/head'
import Navbar from '../components/navbar'
import {useEffect, useState, useRef} from 'react';
import Footer from '../components/footer';
import ManageUsers from '../components/admin_components/manage_accounts';
import ReportedReviews from '../components/admin_components/reported_reviews';
import { getAccountInfos } from '../lib/admin';
import CreateShop from '../components/admin_components/create_shop';
import { useRouter } from 'next/router';
import { getUserInfos } from '../lib/user';
import ManageShops from '../components/admin_components/manage_shops';

export default function Admin() {
  const [content, setContent] = useState("manage_accounts");
  const router = useRouter()
  let displayed_element;
  const [loaded,setLoaded] = useState(false)
  const [users, setUsers] = useState([])
  // fetch account datas and check if user is logged and admin
  useEffect(()=>{
    const fetchAccountsData = async () => {
      const response = await (await getAccountInfos()).json()
      const myself = await (await getUserInfos()).json()
      // if anyone tries to access without being an admin -> unauthorized
      if(myself.user.Type !== 'admin'){
        router.push("/401")
      }else{
        setUsers(response.users)
        setLoaded(true)
      }
    }
    const token = localStorage.getItem('token')
    const type = localStorage.getItem("type")
    if(!token ){
      router.push("/home")
    }else{
      fetchAccountsData()
    }
  },[])
  // dynamic displaying content based on the menu selected output
  if(!loaded){
    return <div></div> 
  }else{
    if (content == "manage_accounts") {
      displayed_element = <ManageUsers accounts={users}/>;
    // } else if (content == "reported_reviews"){
    //   displayed_element = <ReportedReviews reported_reviews={reviewsData}/>
    } else if(content == "view_analytics"){
      displayed_element = <></>;
    } else if(content == "create_shop"){
      displayed_element = <CreateShop accounts={users}/>;
    } else if(content == "manage_shops"){
      displayed_element = <ManageShops/>;
    }
  }

  return (
    <>
    <Head>
      <title>Admin | Barber Shop</title>
      <link rel="icon" type="image/png" sizes="32x32" href="/barber-shop.png"></link>
    </Head>
    <Navbar/>
    <svg className='w-full bg-slate-800 h-full' viewBox='0 0 1440 100' preserveAspectRatio="xMidYMid">
        <path className='w-full fill-slate-900' d="M 0 90 C 480 0 600 0 720 10.7 C 840 21 960 43 1080 48 C 1200 53 1320 43 1380 37.3 L 1440 32 L 1440 0 L 1380 0 C 1320 0 1200 0 1080 0 C 960 0 840 0 720 0 C 600 0 480 0 360 0 C 240 0 120 0 60 0 L 0 0 Z"></path>
    </svg>
    <div className="flex flex-col lg:flex-row justify-center items-start w-full bg-slate-800 px-5 lg:pl-10 pb-10 h-full">
        <div className='w-full flex-col lg:w-1/5 lg:h-screen lg:border-r lg:border-b-0 border-b border-slate-500 text-slate-300 mb-2.5 pb-2.5'>
          {/* menu */}
            <ul className='flex flex-row lg:flex-col justify-between'>
                <li>
                    <button className={`hover:text-white focus:outline-none ${content == "manage_accounts" ? "font-bold" : ""}`} onClick={event => {setContent("manage_accounts")}}>Manage Accounts</button>
                </li>
                {/* <li>
                    <button className={`hover:text-white focus:outline-none ${content == "reported_reviews" ? "font-bold" : ""}`} onClick={event => {setContent("reported_reviews")}}>Reported Reviews</button>
                </li> */}
                <li>
                    <button className={`hover:text-white focus:outline-none ${content == "view_analytics" ? "font-bold" : ""}`} onClick={event => {setContent("view_analytics")}}>View Analytics</button>
                </li>
                <li>
                    <button className={`hover:text-white focus:outline-none ${content == "create_shop" ? "font-bold" : ""}`} onClick={event => {setContent("create_shop")}}>Create Shop</button>
                </li>
                <li>
                    <button className={`hover:text-white focus:outline-none ${content == "manage_shops" ? "font-bold" : ""}`} onClick={event => {setContent("manage_shops")}}>Manage Shops</button>
                </li>
            </ul>
        </div>
        <div className='w-full lg:w-4/5 h-full'>
            {displayed_element}
        </div>
    </div>
    <Footer/>
    </>
  )
}
