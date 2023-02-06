import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {faStar, faMapLocationDot} from "@fortawesome/free-solid-svg-icons";
import Image from "next/image";
import Link from "next/link";
import barber_background from '../../public/barber_profile.jpg'
import GeneralDropdown from "../general_dropdown";

export default function ShopsFound({shops}) {
    return (
        <>
        <div className="w-full flex-col text-xl justify-start items-center px-3">
            <div className="w-full flex flex-col lg:flex-row justify-between items-center pb-2 px-2">
                <div className="py-2 lg:p-0">Found {shops.length} Barber Shops</div>
                <div className="flex items-center justify-end">
                    <GeneralDropdown elements={[1,2,3]}><div>Ordered By</div></GeneralDropdown>
                </div>
            </div>
            <div className="w-full border-b border-slate-600 mb-3"/>
            {shops.map((shop)=>
            //TODO: ADD HREF TO SHOP'S PAGE
            <Link href="" key={shop.id} className="w-full text-slate-200 px-2 flex flex-col items-center justify-start">
                <div key={shop.id+"container"} className="flex flex-col items-center justify-start w-full rounded-lg pb-5">
                    <div className="flex w-full items-start justify-start">
                        <div key={shop.id+"title"} className="text-sm flex items-center lg:items-start justify-start w-full text-left">
                            <div className='h-32 lg:w-1/2'>
                                <Image className="w-full h-full object-cover rounded-lg shadow-md shadow-black/30" src={barber_background} alt="barber salon"/>
                            </div>
                            <div className="flex flex-col items-start justify-start w-full px-3">
                                <div key={shop.id+"name"} className="text-xl text-left font-bold hover:underline">{shop.name}</div>
                                <div className="flex justify-center items-center">
                                    <FontAwesomeIcon key={shop.id+"starIcon"} icon={faStar} className="text-sm text-rose-700"/>
                                    <div className=" pl-1">{shop.meanRating}/5 ({shop.reviewNumber})</div>
                                </div>
                                <div className="flex w-full lg:w-2/3 flex-col lg:flex-row items-start lg:items-center justify-between">
                                    <div className="w-full">{shop.description}</div>
                                    <div className="flex items-center justify-start">
                                        <FontAwesomeIcon key={shop.id+"locationIcon"} icon={faMapLocationDot} className="text-sm pr-2"/>
                                        <div className="font-bold ">{shop.distance} km</div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </Link>
            )}
        </div>
        </>
    );
}