import { getAllShops, getReviews, getShopData } from '../../lib/shops';
import Image from 'next/image';
import Navbar from '../../components/navbar';
import Head from 'next/head';

import barber_background from '../../public/barber_profile.jpg'
import barber_propic from '../../public/barber_bg.png'
import Footer from '../../components/footer';
import Reviews from '../../components/shop_component/reviews';
import ReactStars from 'react-stars'

export default function Shop({ shopData, reviewsData }) {
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
        <div className='flex w-full items-start justify-between bg-slate-800 h-full'>
          {/* left */}
          <div className="w-2/6 h-full bg-slate-800 m-10 ml-15 mt-0">
            <div className="w-full top-0 transform -translate-y-20 flex justify-center items-center">
              <div className='relative max-h-96 overflow-auto rounded-3xl shadow-md shadow-black/70'>
                <div className="bg-slate-800 text-white w-full flex flex-col items-center justify-center p-5 pt-0">
                  <h1 className="text-2xl text-center font-bold leading-tight tracking-tight text-slate-200 sticky top-0 bg-slate-800 w-full border-b border-slate-600 pt-3 pb-3 ">
                    Reviews
                  </h1>
                  <div className="text-lg text-center font-bold leading-tight tracking-tight text-slate-300 break-words p-3">
                    <Reviews>{reviewsData}</Reviews>
                  </div>
                </div>
              </div>
            </div>
          </div>
          {/* center */}
          <div className='flex flex-col items-center w-3/6 pb-3'>
              <div className="w-full top-0 transform -translate-y-20 inset-0 flex justify-center items-center">
                <div className="shadow-md shadow-black/70 rounded-3xl bg-slate-800 text-white w-full flex flex-col items-center justify-start">
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
            {/* leave a review */}
            <div className="w-full bg-slate-800">
                <div className="shadow-md shadow-black/70 rounded-3xl bg-slate-700 text-white w-full flex flex-col items-center justify-start">
                  <h1 className="text-xl text-center font-bold leading-tight tracking-tight text-slate-200 p-2">
                    Leave a review!
                  </h1>
                  <div className='w-3/4 border-b border-slate-600 pt-1'></div>
                  <div className="text-md text-justify leading-tight tracking-tight text-slate-300 break-words w-3/4 p-3 flex flex-col">
                    How did we do?
                    <ReactStars
                      count={5}
                      size={20}
                      color2={'#ffffff'} />
                    <textarea className='bg-slate-600 focus:outline-none rounded-md p-3 text-sm break-words mt-1'  name="" id="" />
                    <button type="submit" className="w-full bg-slate-600 hover:bg-slate-500 focus:outline-none rounded-lg border-slate-700 text-sm px-5 py-2.5 text-center mt-3">That's what I thought</button>
                  </div>
                </div>
            </div>
          </div>
          {/* right */}
          <div className="sticky top-0 w-1/6 h-1/3 bg-slate-800 m-10 mr-15 mt-0">
            <div className="flex justify-center items-center">
              <div className="bg-slate-800 text-white w-full rounded-lg bg-slate-700 mt-3 shadow-md shadow-black/70">
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

