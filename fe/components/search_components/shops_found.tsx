import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {faStar, faMapLocationDot} from "@fortawesome/free-solid-svg-icons";
import Image from "next/image";
import Link from "next/link";
import barber_background from '../../public/barber_profile.jpg'
// import GeneralDropdown from "../general_dropdown";

export default function ShopsFound({shops}:any) {
    return (
        <>
        <div className="w-full flex-col text-xl justify-start items-center px-3 overflow-auto">
            <div className="w-full bg-slate-800 sticky top-0 flex flex-col lg:flex-row justify-between items-center pb-2 px-2 border-b border-slate-600 mb-3">
                <div className="py-2 lg:p-0">Found {(shops)?shops.length:0} Barber Shops</div>

            </div>
            {(shops !== undefined)?shops.map((shop:any)=>
            <Link href={"/shop?shopid="+shop.ID} key={shop.ID} className="w-full text-slate-200 px-2 flex flex-col items-center justify-start">
                <div key={shop.ID+"container"} className="flex flex-col items-center justify-start w-full rounded-lg pb-5">
                    <div className="flex w-full items-start justify-start">
                        <div key={shop.ID+"title"} className="text-sm flex items-center lg:items-start justify-start w-full text-left">
                            {/* TODO SOURCE */}
                            <div className='h-32 lg:w-1/2'>
                                <Image className="w-full h-full object-cover rounded-lg shadow-md shadow-black/30 hover:shadow-black/80" width="100" height="100" src={shop.ImageLink} alt="barber salon"/>
                            </div>
                            <div className="flex flex-col items-start justify-start w-full px-3">
                                <div key={shop.ID+"name"} className="text-xl text-left font-bold hover:underline">{shop.Name}</div>
                                <div className="flex justify-center items-center">
                                    <FontAwesomeIcon key={shop.ID+"starIcon"} icon={faStar} className="text-sm text-rose-700"/>
                                    <div className=" pl-1">{shop.Rating}/5</div>
                                </div>
                                <div className="flex w-full lg:w-2/3 flex-col lg:flex-row items-start lg:items-center justify-between">
                                    <div className="w-full">{shop.description}</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </Link>
            ):<></>}
        </div>
        </>
    );
}