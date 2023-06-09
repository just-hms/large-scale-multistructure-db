import { useEffect, useState } from "react";
import { shopCalendar } from "../../lib/shops";

export default function BarberReservations({shops}:any) {
    const [reservations,setReservations] = useState<any[]>([])
    useEffect(()=>{
        const fetchData = async (shops:any) => {
            var reservation_array:any = []
            for (var i in shops){
                const calendar = (await (await shopCalendar(shops[i].ID)).json()).calendar.Slots
                var name = shops[i].Name
                reservation_array.push({ name , calendar})
            }
            setReservations(reservation_array)
        }

        fetchData(shops)
    },[])

    return (
        <>
        <div className="w-full flex flex-col text-xl justify-center items-center px-3">
            {
            reservations.map((shop:any)=>{
                if(shop.calendar.length > 0){
                    return <>
                        <div key={shop.name} className="w-full lg:w-3/4 rounded-2xl bg-slate-700 shadow-sm shadow-black/70 text-slate-200 px-2 my-2 flex flex-col items-center justify-center">
                            <div className="font-bold py-2 text-center" key={shop.name+"shopname"}>{shop.name}</div>
                            <div key={shop.name+"container"} className="flex flex-col items-start justify-center text-justify py-2">
                                <div className="font-bold pb-2 text-center" key={shop.name+"shopname"}>Booked Slots:
                                {shop.calendar.map((appointment:any)=>{
                                    return <p key={shop.name+"name"} className="font-normal pr-3">Date: {new Date(appointment.Start).toLocaleDateString()} {new Date(appointment.Start).toLocaleTimeString()}</p>
                                })}
                                </div>
                            </div>
                        </div>                        
                    </>
                }
                else{
                    return<></>
                }  
                }
                )
            }   
        </div>
        </>
    );
}