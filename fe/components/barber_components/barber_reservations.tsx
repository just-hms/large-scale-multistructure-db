import { useEffect, useState } from "react";

export default function BarberReservations({shops}:any) {
    return (
        <>
        <div className="w-full flex flex-col text-xl justify-center items-center px-3">
            {shops.map((shop:any)=>{
                if(shop.Appointments){
                    return shop.Appointments.map((appointment:any)=>{
                    if(appointment.Status === 'pending'){
                        return <>
                        <div key={appointment.ID} className="w-full lg:w-3/4 rounded-2xl bg-slate-700 shadow-sm shadow-black/70 text-slate-200 px-2 my-2 flex flex-col items-center justify-center">
                            <p className="font-bold py-2" key={appointment.ID+"shopname"}>Barber shop: {shop.Name}</p>
                            <div key={appointment.ID+"container"} className="flex flex-col items-start justify-center text-justify py-2">
                                <p>User: {appointment.Username}</p>
                                <p key={appointment.ID+"name"} className="text-left pr-3">Date: {new Date(appointment.StartDate).toLocaleDateString()} {new Date(appointment.StartDate).toLocaleTimeString()}</p>
                            </div>
                        </div>
                        </>
                    }else{
                        return <></>
                    }
                    })
                }
                else{
                    return<></>
                }  
            }
            )}
        </div>
        </>
    );
}