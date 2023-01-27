import Head from 'next/head'
import Navbar from '../components/navbar'
import Image from 'next/image'
import { useFormik } from 'formik';
import barber_background from '../public/barber_bg.png'
import Footer from '../components/footer';

export default function Home() {
  const formik = useFormik({
    initialValues: {
      email: '',
      password: '',
      repeatPassword: '',
    },
    onSubmit: values => {
      // TODO: check values and yadda yadda
      alert(JSON.stringify(values, null, 3));
    },
  });
  return (
    <>
      <Head>
        <title>Home | Barber Shop</title>
        <link rel="icon" type="image/png" sizes="32x32" href="/barber-shop.png"></link>
      </Head>
      <Navbar/>
      <div className="w-full flex-col justify-center items-center bg-slate-800 h-full">
        <div className="relative w-full h-full">
          <Image className="w-full" src={barber_background} alt="barber salon"/>
          <div className="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 inset-0 flex justify-center items-center">
            <div className="shadow-md shadow-black/70 rounded-3xl bg-slate-800/90 text-white w-full">
              <h1 className="text-4xl text-center font-bold leading-tight tracking-tight text-slate-200 pt-5">
                Cut it out
              </h1>
              <h2 className="text-lg text-center font-bold leading-tight tracking-tight text-slate-300 pt-5 ">
                Find your future Favourite Barber
              </h2>
              <div className='flex justify-center'>
                <div className='w-1/2 m-5 flex items-center justify-center border-2 border-red-900/70 rounded-full'>
                  <input
                    type="search"
                    className="w-full text-base pl-5 bg-slate-300/0 focus:border-red-900 text-white bg-clip-padding rounded-full transition ease-in-out focus:outline-none" 
                    form-control id="barberSearch" placeholder="Search"
                  />
                  <button className="btn inline-block px-6 py-2.5 m-1 bg-red-900 text-white font-medium text-xs leading-tight uppercase rounded-full shadow-md focus:bg-red-800 hover:bg-red-800 focus:outline-none transition duration-150 ease-in-out flex items-center" type="button" id="search_button">
                    <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="search" className="w-4 " role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                      <path fill="currentColor" d="M505 442.7L405.3 343c-4.5-4.5-10.6-7-17-7H372c27.6-35.3 44-79.7 44-128C416 93.1 322.9 0 208 0S0 93.1 0 208s93.1 208 208 208c48.3 0 92.7-16.4 128-44v16.3c0 6.4 2.5 12.5 7 17l99.7 99.7c9.4 9.4 24.6 9.4 33.9 0l28.3-28.3c9.4-9.4 9.4-24.6.1-34zM208 336c-70.7 0-128-57.2-128-128 0-70.7 57.2-128 128-128 70.7 0 128 57.2 128 128 0 70.7-57.2 128-128 128z"></path>
                    </svg>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
        <Footer/>
      </div>
    </>
  )
}
