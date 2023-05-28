import Image from 'next/image';
import barber_background from '../../public/barber_profile.jpg'
import barber_propic from '../../public/barber_bg.png'
import { useFormik } from 'formik';
import { useEffect, useState, useRef } from 'react';
import { modifyShopDescription, modifyShopEmployees } from '../../lib/barber';
import Holidays from './holiday';

export default function ModifiedShop({ shopData }:any) {
  const textAreaRef = useRef<any>();
  const formik = useFormik({
    initialValues: {
        description: shopData.Description,
        employees: shopData.Employees,
    },
    onSubmit: async (values:any) => {
      const response = await modifyShopDescription(shopData.ID,values)
      if(response.status == 202){
        alert("description changed")
        shopData.Description = values.description
      }else{
        alert("there was an error")
      }
    },
  });
  useEffect(()=>{
    formik.values.description = shopData.Description
    textAreaRef.current.value = shopData.Description
  },[shopData.Description])
  return (
    <>
      <div className='h-full w-full lg:px-5'>
        <div className='h-96 w-full'>
          <Image className="w-full h-full object-cover rounded-xl" src={barber_background} alt="barber salon"/>
        </div>
        <div className='flex flex-col lg:flex-row w-full bg-slate-800 h-full'>
          <div className='flex flex-col items-center w-full justify-start'>
            {/* DESCRIPTION */}
            <div className='flex flex-col items-center order-first lg:order-none w-full px-2'>
                <div className="w-full top-0transform -translate-y-40 lg:-translate-y-20 inset-0 flex justify-center items-center">
                  <form onSubmit={formik.handleSubmit} className="w-full px-5 h-full flex flex-col items-center justify-start rounded-3xl bg-slate-700 bg-opacity-60 backdrop-blur-lg drop-shadow-lg">
                    <div className='w-20 h-20 transform -translate-y-1/2 shadow shadow-black/70 rounded-full'>
                      <Image className="w-full h-full object-cover rounded-full " src={barber_propic} alt="barber salon"/>
                    </div>
                    <h1 className="text-2xl text-center font-bold leading-tight tracking-tight text-slate-200 border-b w-3/4 mb-2.5 pb-2.5 border-slate-500 ">
                      {shopData.Name}
                    </h1>
                    <div className='w-full h-40'>
                      <textarea ref={textAreaRef} name="description" id="description" onChange={formik.handleChange} value={formik.values.description} className="bg-slate-600 h-40 lg:h-full resize-y bg-opacity-60 backdrop-blur-lg drop-shadow-lg text-slate-200 focus:outline-none rounded-md p-1.5 text-sm break-words mt-1 w-full" />
                    </div>
                    <button type="submit" className="px-3 py-1 mx-2 my-4 bg-rose-800 bg-opacity-70 text-slate-300 text-xs rounded-full focus:bg-red-700 hover:bg-red-700 focus:outline-none transition duration-150 ease-in-out ">
                        Submit Changes
                    </button> 
                  </form>
              </div>
              <Holidays shopData={shopData}></Holidays>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
