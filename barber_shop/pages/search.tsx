import { useRouter} from "next/router"
import Head from "next/head"
import Navbar from "../components/navbar"
import Image from "next/image"
import {findShops} from "../lib/search"
import ShopsFound from "../components/search_components/shops_found"

export default function Search({shopData}) {
    const router = useRouter()
    const { prop_name } = router.query
    console.log(shopData)
    return (
    <>
    <Head>
        <title>Search | Barber Shop</title>
        <link rel="icon" type="image/png" sizes="32x32" href="/barber-shop.png"></link>
    </Head>
    <Navbar/>
    <svg className='w-full bg-slate-800 h-full' viewBox='0 0 1442 100' preserveAspectRatio="xMidYMid">
        <path className='w-full fill-slate-900' d="M 0 90 C 480 0 600 0 720 10.7 C 840 21 960 43 1080 48 C 1200 53 1320 43 1380 37.3 L 1440 32 L 1440 0 L 1380 0 C 1320 0 1200 0 1080 0 C 960 0 840 0 720 0 C 600 0 480 0 360 0 C 240 0 120 0 60 0 L 0 0 Z"></path>
    </svg>
    <div className="w-full flex flex-col text-slate-200 justify-start p-5 items-center bg-slate-800 h-full">
        <div className="w-full flex mb-5">
            {/* LAST SEARCH PARAMS   */}
            <div className="w-1/3 flex items-center justify-center">
                <div className='w-full p-5 flex items-center justify-center rounded-full bg-slate-700'>
                    asdl,mjasolkd
                </div>
            </div>
            {/* SEARCH BAR */}
            <div className=" w-2/3 flex items-center justify-center">
                <div className='w-full m-5 flex items-center justify-center rounded-full bg-red-900 bg-opacity-60 backdrop-blur-lg drop-shadow-lg'>
                  <input
                    type="search"
                    className="w-full font-bold text-slate-100 pl-5 bg-slate-700/0 bg-clip-padding rounded-full transition ease-in-out focus:outline-none" 
                    form-control id="barberSearch" placeholder="Search"
                  />
                  <button className="btn inline-block px-6 py-5 m-1 bg-red-800 bg-opacity-60 backdrop-blur-lg drop-shadow-lg text-white font-medium text-xs leading-tight uppercase rounded-full focus:bg-red-800 hover:bg-red-800 focus:outline-none transition duration-150 ease-in-out flex items-center" type="button" id="search_button">
                    <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="search" className="w-4 " role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                      <path className='fill-slate-100/90 drop-shadow-lg backdrop-blur-lg'  d="M505 442.7L405.3 343c-4.5-4.5-10.6-7-17-7H372c27.6-35.3 44-79.7 44-128C416 93.1 322.9 0 208 0S0 93.1 0 208s93.1 208 208 208c48.3 0 92.7-16.4 128-44v16.3c0 6.4 2.5 12.5 7 17l99.7 99.7c9.4 9.4 24.6 9.4 33.9 0l28.3-28.3c9.4-9.4 9.4-24.6.1-34zM208 336c-70.7 0-128-57.2-128-128 0-70.7 57.2-128 128-128 70.7 0 128 57.2 128 128 0 70.7-57.2 128-128 128z"></path>
                    </svg>
                  </button>
                </div>
            </div>
        </div>
        <div className="w-full flex">
            {/* IDK */}
            <div className="border w-1/3 flex flex-col items-center justify-center">
                left menu 
                <p>stuff</p>
            </div>
            {/* SHOPS */}
            <div className=" w-2/3 flex flex-col items-center justify-center">
                <ShopsFound shops={shopData}/>
            </div>
        </div>
    </div>
    </>
    )
}

export async function getServerSideProps() {
    // TODO: actually retrieve datas
    const shopData = findShops()
    return {
      props: {
        shopData
      },
    }
  }