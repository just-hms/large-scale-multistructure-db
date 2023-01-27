import Image from 'next/image'
import barber_background from '../../public/barber_profile.jpg'

export default function AccountReservation({reservationData}) {
    console.log(reservationData)
  return (
    <>    
    <div className='flex flex-col items-center justify-start text-left text-slate-300 text-lg w-full'>
        <div className='flex flex-col items-center justify-center w-2/3 bg-slate-700 rounded-lg shadow-md shadow-black/20'>
            <div className='h-32 w-full'>
                <Image className="w-full h-full object-cover rounded-t-lg shadow-md shadow-black/10" src={barber_background} alt="barber salon"/>
            </div>
            <div className='w-full pl-5 pr-5'>
                <p className='pt-2 pb-2 w-full font-bold text-xl'>{reservationData.name}</p>
                <p className='pt-2 pb-2 w-full border-b border-slate-500 text-md'>Date:</p>
                <div className='flex w-full justify-start items-center pb-5'>
                    <div className='rounded-md pt-2 mr-2 w-full flex items-center justify-between'>
                        {reservationData.date}
                        <button className='w-1/2 pt-2text-white bg-rose-800 hover:bg-rose-700 focus:outline-none font-medium rounded-lg border-slate-700 text-sm px-5 py-2.5 text-center'>Cancel Reservation</button>
                    </div>
                </div>
            </div>
        </div>
    </div>  
    </>

  )}
