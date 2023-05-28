import { getShopData,submitReview,shopCalendar,getReviews, getAppointment } from '../lib/shops';
import { getUserInfos } from '../lib/user';
import { useFormik } from 'formik';
import Image from 'next/image';
import Navbar from '../components/navbar';
import Head from 'next/head';
import barber_background from '../public/barber_profile.jpg'
import barber_propic from '../public/barber_bg.png'
import Footer from '../components/footer';
import Reviews from '../components/shop_component/reviews';
import Calendar from '../components/shop_component/calendar';
import ReactStars from 'react-stars'
import { useEffect, useState } from 'react';
import { useRouter } from 'next/router'

export default function Shop() {
  const [loaded,setLoaded] = useState(false)
  const router = useRouter()
  const [shopData, setshopData] = useState<any>('')
  const { shopid } = router.query
  const [userid, setUserid] = useState<string>('')
  const [shopAppointments, setShopAppointments] = useState<any[]>([])
  const formik = useFormik({
    initialValues: {
        content: '',
        rating: 0,
    },
    onSubmit: async (values:any) => {
      const response = await submitReview(shopid,values)
      if(response.ok){
          window.location.reload()
      }
    },
  });

  useEffect(()=>{
    if(shopid != undefined){
      const token = localStorage.getItem('token')
      if(!token){
        router.push("/")
      }else{
        const fetchData = async (shopid:any) => {
          // get basic shop data
          const response = await getShopData(shopid)
          if (response.ok){                                         
            const json_response = await response.json()
            setshopData(json_response.barbershop)
          }else{
            router.push("/404")
          }
          // get shop calendar
          const calendar = (await (await shopCalendar(shopid)).json()).calendar.Slots
          setShopAppointments(calendar)
          const user_object = await(await getUserInfos()).json()
          setUserid(user_object.user.ID)
          setLoaded(true)
        }
        fetchData(shopid)
      }
    }
  },[shopid])

  if(!loaded){
    return <div></div>
  }
  return (
    <>
      <Head>
        <title>{shopData.Name} | Barber Shop</title>
      <link rel="icon" type="image/png" sizes="32x32" href="/barber-shop.png"></link>
      </Head>
      <Navbar style="absolute top-0">
        <svg className='w-full bg-slate-800/0 h-full' viewBox='0 0 1440 100' preserveAspectRatio="xMidYMid">
            <path className='w-full fill-slate-900' d="M 0 90 C 480 0 600 0 720 10.7 C 840 21 960 43 1080 48 C 1200 53 1320 43 1380 37.3 L 1440 32 L 1440 0 L 1380 0 C 1320 0 1200 0 1080 0 C 960 0 840 0 720 0 C 600 0 480 0 360 0 C 240 0 120 0 60 0 L 0 0 Z"></path>
        </svg>
      </Navbar>
      <div className='h-full'>
        <div className='h-96 w-full'>
          <Image className="w-full h-full object-cover " src={barber_background} alt="barber salon"/>
        </div>
        <div className='flex flex-col lg:flex-row w-full bg-slate-800 h-full'>
          {/* REVIEWS */}
          <div className="w-full lg:w-1/2 lg:h-full bg-slate-800 mt-0 px-3 pb-3 lg:py-3 order-last lg:order-none">
            <div className="w-full top-0 transform lg:-translate-y-20 flex justify-center items-center">
              <div className='w-full relative max-h-128 overflow-auto rounded-3xl shadow-md shadow-black/70'>
                <div className="bg-slate-800 text-white w-full flex flex-col items-center justify-center ">
                  <div className="mx-5 text-center font-bold leading-tight tracking-tight text-slate-200 sticky top-0 bg-slate-800 w-full flex flex-col items-center">
                    <h1 className='text-2xl py-1 border-b border-slate-600 w-5/6'>Reviews</h1>
                  </div>
                  <div className="text-lg text-center w-3/4 font-bold leading-tight tracking-tight text-slate-300 break-words p-3 mx-5">
                    <Reviews shopid={shopid} userid={userid}></Reviews>
                  </div>
                  {/* leave a review */}
                  <div className='bg-slate-800 border-t border-slate-600 w-full sticky bottom-0'>
                    <form onSubmit={formik.handleSubmit} className="shadow-md shadow-black/70  bg-slate-800 text-white w-full flex flex-col items-center justify-start">
                      <h1 className="text-md text-center font-bold leading-tight tracking-tight text-slate-200 py-2">
                        Leave a review!
                      </h1>
                      <div className='w-3/4 border-b border-slate-700 pt-1'></div>
                      {/* <p className='text-sm'>Give the review a title:</p>
                      <input name="title" id="title" onChange={formik.handleChange} value={formik.values.title} className='bg-slate-700 w-3/4 focus:outline-none resize-none rounded-md p-1.5 text-sm break-words mt-1'/> */}
                      <div className="text-sm text-justify leading-tight tracking-tight text-slate-300 break-words w-3/4 py-2 flex flex-col">
                        <div className='flex items-center justify-between'>
                          How did we do?
                          <ReactStars
                            count={5}
                            size={20}
                            color2={'#ffffff'} 
                            half={false}
                            onChange={(value): void => {
                              formik.setFieldValue("rating", value);
                            }} 
                            value={formik.values.rating}/>
                        </div>
                        <textarea name="content" id="content" onChange={formik.handleChange} value={formik.values.content} className='bg-slate-700 focus:outline-none resize-none rounded-md p-1.5 text-sm break-words mt-1'/>                  
                        <button type="submit" className="w-full text-sm bg-slate-600 hover:bg-slate-500 focus:outline-none rounded-lg border-slate-700 text-sm py-2 text-center mt-3 z-10">What I thought</button>
                      </div>
                    </form>
                  </div>
                </div>
              </div>
            </div>

          </div>
          <div className='flex flex-col items-center w-full lg:w-4/6 justify-start'>
            {/* DESCRIPTION */}
            <div className='flex flex-col items-center order-first lg:order-none w-full px-3 lg:py-3'>
                <div className="w-full top-0 transform -translate-y-40 lg:-translate-y-20 inset-0 flex justify-center items-center">
                  <div className="w-full h-full flex flex-col items-center justify-start rounded-3xl bg-slate-700 bg-opacity-60 backdrop-blur-lg drop-shadow-lg">
                    <div className='w-1/2 transform -translate-y-1/2 h-20 w-20 shadow shadow-black/70 rounded-full'>
                      <Image className="w-full h-full object-cover rounded-full " src={barber_propic} alt="barber salon"/>
                    </div>
                    <h1 className="text-2xl text-center font-bold leading-tight tracking-tight text-slate-200 ">
                      {shopData.Name}
                    </h1>
                    <div className='w-3/4 border-b border-slate-600 pt-1'></div>
                    <p className="text-md text-justify leading-tight tracking-tight text-slate-300 break-words w-3/4 p-3 ">
                      {shopData.Description}
                    </p>
                  </div>
              </div>
            </div>
            {/* CALENDAR */}
            <Calendar shopid={shopid} calendar={shopAppointments} employees={shopData.Employees}></Calendar>
          </div>
        </div>
        <Footer/>
      </div>
    </>
  );
}
