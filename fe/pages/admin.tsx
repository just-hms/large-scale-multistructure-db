import Head from 'next/head'
import Navbar from '../components/navbar'
import {useEffect, useState, useRef} from 'react';
import Footer from '../components/footer';
import ManageUsers from '../components/admin_components/manage_accounts';
import ReportedReviews from '../components/admin_components/reported_reviews';
import { getReviews } from '../lib/admin';
import CreateShop from '../components/admin_components/create_shop';

export default function Admin({reviewsData}) {
  const [content, setContent] = useState("manage_accounts");
  let displayed_element;
  if (content == "manage_accounts") {
    displayed_element = <ManageUsers/>;
  } else if (content == "reported_reviews"){
    displayed_element = <ReportedReviews reported_reviews={reviewsData}/>
  } else if(content == "view_analytics"){
    displayed_element = <></>;
  } else if(content == "create_shop"){
    displayed_element = <CreateShop/>;
  }
  return (
    <>
    <Head>
      <title>Admin | Barber Shop</title>
      <link rel="icon" type="image/png" sizes="32x32" href="/barber-shop.png"></link>
    </Head>
    <Navbar/>
    <svg className='w-full bg-slate-800 h-full' viewBox='0 0 1442 100' preserveAspectRatio="xMidYMid">
        <path className='w-full fill-slate-900' d="M 0 90 C 480 0 600 0 720 10.7 C 840 21 960 43 1080 48 C 1200 53 1320 43 1380 37.3 L 1440 32 L 1440 0 L 1380 0 C 1320 0 1200 0 1080 0 C 960 0 840 0 720 0 C 600 0 480 0 360 0 C 240 0 120 0 60 0 L 0 0 Z"></path>
    </svg>
    <div className="flex flex-col lg:flex-row justify-center items-start w-full bg-slate-800 px-5 lg:pl-10 pb-10 h-full">
        <div className='w-full flex-col lg:w-1/5 lg:h-screen lg:border-r lg:border-b-0 border-b border-slate-500 text-slate-300 mb-2.5 pb-2.5'>
            <ul className='flex flex-row lg:flex-col justify-between'>
                <li>
                    <button className={`hover:text-white focus:outline-none ${content == "manage_accounts" ? "font-bold" : ""}`} onClick={event => {setContent("manage_accounts")}}>Manage Accounts</button>
                </li>
                <li>
                    <button className={`hover:text-white focus:outline-none ${content == "reported_reviews" ? "font-bold" : ""}`} onClick={event => {setContent("reported_reviews")}}>Reported Reviews</button>
                </li>
                <li>
                    <button className={`hover:text-white focus:outline-none ${content == "view_analytics" ? "font-bold" : ""}`} onClick={event => {setContent("view_analytics")}}>View Analytics</button>
                </li>
                <li>
                    <button className={`hover:text-white focus:outline-none ${content == "create_shop" ? "font-bold" : ""}`} onClick={event => {setContent("create_shop")}}>Create Shop</button>
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

export async function getStaticProps() {
  // TODO: actually retrieve datas
  // const postData = getShopData(params.shop)
  const reviewsData =  await getReviews("shopname")
  return {
    props: {
      reviewsData
    },
  } 
}