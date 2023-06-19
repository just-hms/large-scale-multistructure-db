import Image from 'next/image'
import barber_background from '../../public/barber_profile.jpg'
import Link from 'next/link'
import { useEffect, useState } from 'react'
import { getShopData } from '../../lib/shops';
import { deleteAppointment } from '../../lib/user';


export default function AccountReservation({reservationData}:any) {
    const [shopData,setshopData] = useState()
    const date = (reservationData)?(new Date(Date.parse(reservationData.StartDate))).toLocaleDateString()+" "+(new Date(Date.parse(reservationData.StartDate))).toLocaleTimeString():''
    useEffect(()=>{
        if(reservationData){
            const fetchData = async () => {
                const shopInfo = (await (await getShopData(reservationData.BarbershopID)).json()).barbershop
                setshopData(shopInfo)
            }
            fetchData()
        }
    },[])
  return (
    <>    
    <div className='flex flex-col items-center justify-start text-left text-slate-300 text-lg w-full'>
        <div className='flex flex-col items-center justify-center w-full lg:w-2/3 bg-slate-700 rounded-lg shadow-md shadow-black/20'>
            {(reservationData != null)?
                <>
                <div className='h-40 lg:h-32 w-full'>
                    <Image className="w-full h-full object-cover rounded-t-lg shadow-md shadow-black/10" src={barber_background} alt="barber salon"/>
                </div>
                <div className='w-full pl-5 pr-5 pt-2'>
                    <Link href={"/shop?shopid="+reservationData.BarbershopID} className='pt-2 pb-2 w-full hover:underline font-bold text-xl'>{(shopData)?shopData.Name:''}</Link>
                    <p className='w-full text-md'>Date:</p>
                    <p className='pb-2 w-full border-b border-slate-500 text-md'>{date}</p>
                    <div className='flex w-full justify-start items-center pb-5'>
                        <div className='rounded-md pt-2 mr-2 w-full flex items-center justify-end'>
                            <button className='w-1/2 pt-2 text-slate-200 bg-rose-800 hover:bg-rose-700 focus:outline-none font-medium rounded-lg border-slate-700 text-sm px-5 py-2.5 text-center'
                            onClick={async (e)=>{
                                const response = await deleteAppointment()
                                if (response.status == 202)
                                    window.location.reload()
                            }}>Cancel Reservation</button>
                        </div>
                    </div>
                </div>
                </>:
                <>
                <div className='w-full px-5 py-10 text-center'>
                    <p className='w-full text-lg'>No reservations</p>
                </div>
                </>
            }
        </div>
    </div>  
    </>

  )}
