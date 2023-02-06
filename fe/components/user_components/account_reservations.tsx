import Image from 'next/image'
import barber_background from '../../public/barber_profile.jpg'
import Link from 'next/link'

export default function AccountReservation({reservationData}) {
    console.log(reservationData)
  return (
    <>    
    <div className='flex flex-col items-center justify-start text-left text-slate-300 text-lg w-full'>
        <div className='flex flex-col items-center justify-center w-full lg:w-2/3 bg-slate-700 rounded-lg shadow-md shadow-black/20'>
            <div className='h-40 lg:h-32 w-full'>
                <Image className="w-full h-full object-cover rounded-t-lg shadow-md shadow-black/10" src={barber_background} alt="barber salon"/>
            </div>
            <div className='w-full pl-5 pr-5 pt-2'>
                {/* TODO ADD LINK TO THE SHOP */}
                <Link href="" className='pt-2 pb-2 w-full font-bold text-xl'>{reservationData.name}</Link>
                <p className='w-full text-md'>Date:</p>
                <p className='pb-2 w-full border-b border-slate-500 text-md'>{reservationData.date}</p>
                <div className='flex w-full justify-start items-center pb-5'>
                    <div className='rounded-md pt-2 mr-2 w-full flex items-center justify-end'>
                        <button className='w-1/2 pt-2 text-slate-200 bg-rose-800 hover:bg-rose-700 focus:outline-none font-medium rounded-lg border-slate-700 text-sm px-5 py-2.5 text-center'>Cancel Reservation</button>
                    </div>
                </div>
            </div>
        </div>
    </div>  
    </>

  )}
