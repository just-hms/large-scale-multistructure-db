import { useRouter } from "next/router"
import { useState, useEffect } from "react";
import Head from "next/head"
import Navbar from "../components/navbar"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {faLocationDot} from "@fortawesome/free-solid-svg-icons";
import {findShops} from "../lib/search"
import ShopsFound from "../components/search_components/shops_found"
import Link from "next/link";
import Footer from "../components/footer";

export default function Search() {
    const router = useRouter()
    // query parameter
    const { area } = router.query 
    // fetch latitude and longitude given a request
    const [shops,setShops] = useState<any[]>([])
    const [loaded,setLoaded] = useState(false)

    // SEARCH BAR FILTER
    const [query, setQuery] = useState('');
    const searchFilter = (shops:any) => {
      if(shops)
        return shops.filter(
          (el:any) => el.Name.toLowerCase().includes(query.toLocaleLowerCase())
        )
    }
    //Handling the input on our search bar
    const handleChange = (e:any) => {
      setQuery(e.target.value)
      searchFilter(shops)
    }

    const filteredShops = searchFilter(shops)

    useEffect(() => {
      const token = localStorage.getItem('token')
      const fetchData = async () =>{
        const response = await findShops(area)
        if(response.status == 200){
          const json = await response.json()
          console.log(json.barberShops)
          setShops(json.barberShops)
        }
        setLoaded(true)
      }
      if(!token){
        router.push("/")
      }else{
        if(area){
          fetchData()
        }
      }
    }, [area]);

    if(!loaded){
      return <div></div> //show nothing or a loader
    }
    return (
    <>
    <Head>
        <title>Search | Barber Shop</title>
        <link rel="icon" type="image/png" sizes="32x32" href="/barber-shop.png"></link>
    </Head>
    <Navbar/>
    <svg className='w-full bg-slate-800 h-full' viewBox='0 0 1440 100' preserveAspectRatio="xMidYMid">
        <path className='w-full fill-slate-900' d="M 0 90 C 480 0 600 0 720 10.7 C 840 21 960 43 1080 48 C 1200 53 1320 43 1380 37.3 L 1440 32 L 1440 0 L 1380 0 C 1320 0 1200 0 1080 0 C 960 0 840 0 720 0 C 600 0 480 0 360 0 C 240 0 120 0 60 0 L 0 0 Z"></path>
    </svg>
    <div className="w-full flex flex-col text-slate-200 justify-start p-5 items-center bg-slate-800 h-full">
        <div className="w-full flex mb-5 flex-col lg:flex-row items-center">
            {/* LAST SEARCH PARAMS */}
            <div className="w-full lg:w-1/3 flex items-center justify-center">
                <div className='w-full text-xl font-bold py-1.5 flex flex-col items-center justify-center rounded-full bg-slate-700'>
                  <div className="flex w-full items-center justify-center">
                    <FontAwesomeIcon icon={faLocationDot} className="text-xl pr-1 text-rose-600"/>
                      {area}
                  </div>
                  <Link href="/home" className="underline text-sm">Change Location</Link>
                </div>
            </div>
            {/* SEARCH BAR */}
            <div className="w-full lg:w-2/3 flex items-center justify-center">
                <div className='w-full py-4 my-5 lg:m-5 flex items-center justify-center rounded-full bg-red-900 bg-opacity-60 backdrop-blur-lg drop-shadow-lg'>
                  <div className="btn inline-block ml-3 text-slate-200 font-medium text-sm leading-tight uppercase rounded-full flex items-center" >
                    <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="search" className="w-4 " role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                      <path className='fill-slate-100/90 drop-shadow-lg backdrop-blur-lg '  d="M505 442.7L405.3 343c-4.5-4.5-10.6-7-17-7H372c27.6-35.3 44-79.7 44-128C416 93.1 322.9 0 208 0S0 93.1 0 208s93.1 208 208 208c48.3 0 92.7-16.4 128-44v16.3c0 6.4 2.5 12.5 7 17l99.7 99.7c9.4 9.4 24.6 9.4 33.9 0l28.3-28.3c9.4-9.4 9.4-24.6.1-34zM208 336c-70.7 0-128-57.2-128-128 0-70.7 57.2-128 128-128 70.7 0 128 57.2 128 128 0 70.7-57.2 128-128 128z"></path>
                    </svg>
                  </div>
                  <input
                    type="text"
                    className="w-full font-bold text-slate-100 pl-5 bg-slate-700/0 bg-clip-padding rounded-full transition ease-in-out focus:outline-none " 
                     id="barberSearch" placeholder="Search"
                    onChange={handleChange}
                  />
                </div>
            </div>
        </div>
        <div className="w-full flex flex-col lg:flex-row items-center lg:items-start">
            {/* SHOPS */}
            <div className="w-full flex h-screen flex-col items-center justify-start">
                <ShopsFound shops={filteredShops}/>
            </div>
        </div>
    </div>
    <Footer/>
    </>
    )
}
