import Image from "next/image";
import barber_background from '../../public/barber_profile.jpg'

export default function BarberReservations({reservations}) {
    return (
        <>
        <div className="w-full flex flex-col text-xl justify-center items-center px-3">
            {reservations.map((reservation)=>
            <div key={reservation.id} className="w-full lg:w-3/4 rounded-2xl bg-slate-700 shadow-sm shadow-black/70 text-slate-200 px-2 my-2 flex flex-col items-center justify-center">
                <div key={reservation.id+"container"} className="flex flex-col items-start justify-start w-full rounded-lg py-2.5">
                    <p>{reservation.user}</p>
                    <div key={reservation.id+"title"} className="text-lg flex items-center lg:items-start justify-center w-full text-left">
                        <h1 key={reservation.id+"name"} className="text-left pr-3">{reservation.date}</h1>
                        <p className="w-full">{reservation.time}</p>
                    </div>
                </div>
            </div>
            )}
        </div>
        </>
    );
}