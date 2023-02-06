import { getAllShops, getReviews, getShopData } from '../../lib/shops';
import Image from 'next/image';
import Navbar from '../../components/navbar';
import Head from 'next/head';

import barber_background from '../../public/barber_profile.jpg'
import barber_propic from '../../public/barber_bg.png'
import Footer from '../../components/footer';
import Reviews from '../../components/shop_component/reviews';
import ReactStars from 'react-stars'
import { useRouter } from 'next/router';

export default function Shop({ shopData, reviewsData }:{ shopData:any, reviewsData:any }) {
  return (
    <>
      <Head>
        <title>{shopData.title} | Barber Shop</title>
      <link rel="icon" type="image/png" sizes="32x32" href="/barber-shop.png"></link>
      </Head>
      <Navbar></Navbar>
      <div className='h-full'>
        <div className='h-96 w-full'>
          <Image className="w-full h-full object-cover " src={barber_background} alt="barber salon"/>
        </div>
        <div className='flex flex-col lg:flex-row w-full bg-slate-800 h-full'>
          {/* REVIEWS */}
          <div className="w-full lg:w-1/2 lg:h-full bg-slate-800 mt-0 px-3 pb-3 lg:py-3 order-last lg:order-none">
            <div className="w-full top-0 transform lg:-translate-y-20 flex justify-center items-center">
              <div className='relative max-h-128 overflow-auto rounded-3xl shadow-md shadow-black/70'>
                <div className="bg-slate-800 text-white w-full flex flex-col items-center justify-center ">
                  <div className="mx-5 text-center font-bold leading-tight tracking-tight text-slate-200 sticky top-0 bg-slate-800 w-full flex flex-col items-center">
                    <h1 className='text-2xl py-1 border-b border-slate-600 w-5/6'>Reviews</h1>
                  </div>
                  <div className="text-lg text-center font-bold leading-tight tracking-tight text-slate-300 break-words p-3 mx-5">
                    <Reviews>{reviewsData}</Reviews>
                  </div>
                  {/* leave a review */}
                  <div className='bg-slate-800 border-t border-slate-600 w-full sticky bottom-0'>
                    <div className="shadow-md shadow-black/70 rounded-b-3xl bg-slate-800 text-white w-full flex flex-col items-center justify-start">
                      <h1 className="text-md text-center font-bold leading-tight tracking-tight text-slate-200 py-2">
                        Leave a review!
                      </h1>
                      <div className='w-3/4 border-b border-slate-700 pt-1'></div>
                      <div className="text-sm text-justify leading-tight tracking-tight text-slate-300 break-words w-3/4 py-2 flex flex-col">
                        <div className='flex items-center justify-between'>
                          How did we do?
                          <ReactStars
                            count={5}
                            size={20}
                            color2={'#ffffff'} />
                        </div>
                        <textarea className='bg-slate-700 focus:outline-none resize-none rounded-md p-1.5 text-sm break-words mt-1'  name="" id="" />
                        <button type="submit" className="w-full text-sm bg-slate-700 hover:bg-slate-600 focus:outline-none rounded-lg border-slate-700 text-sm py-2 text-center mt-3 z-10">That`&apos`s what I thought</button>
                      </div>
                    </div>
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
                    <div className='w-full transform -translate-y-1/2 h-20 w-20 shadow shadow-black/70 rounded-full'>
                      <Image className="w-full h-full object-cover rounded-full " src={barber_propic} alt="barber salon"/>
                    </div>
                    <h1 className="text-2xl text-center font-bold leading-tight tracking-tight text-slate-200 ">
                      {shopData.name}
                    </h1>
                    <div className='w-3/4 border-b border-slate-600 pt-1'></div>
                    <p className="text-md text-justify leading-tight tracking-tight text-slate-300 break-words w-3/4 p-3 ">
                      {shopData.description}
                    </p>
                  </div>
              </div>
            </div>
            {/* CALENDAR */}
            <div className="w-full lg:w-5/6 h-1/3 mt-0 px-3 lg:py-3 transform -translate-y-20">
              <div className="flex justify-center items-center">
                <div className="w-full rounded-lg bg-slate-700 shadow-md shadow-black/70 mt-3 ">
                  <h1 className="text-2xl text-center font-bold leading-tight tracking-tight text-slate-200 pt-5 ml-3 mr-3 break-words">
                    Calendar
                  </h1>
                  <div className="text-lg text-justify font-bold leading-tight tracking-tight text-slate-300 break-words p-3">
                    <div className='h-20 w-50 bg-white'>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <Footer/>
      </div>
    </>
  );  
}

export async function getStaticPaths() {
  const paths = await getAllShops();
  return {
    paths: [{
      params: {
        shop: 'shop1'
      },
    },{
      params: {
        shop: 'shop2'
      },
    }],
    fallback: false,
  };
}

export async function getStaticProps() {
  // TODO: actually retrieve datas
  // const postData = getShopData(params.shop)
  const reviewsData =  getReviews("shopname")
  const shopData = {
    name:"Barbiere di Siviglia",
    title:"Barbiere di Siviglia",
    description:"occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
  }
  return {
    props: {
      shopData,
      reviewsData
    },
  }
}

