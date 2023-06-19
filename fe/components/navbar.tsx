
import Dropdown from './user_components/user_dropdown'
import Link from 'next/link'
import barber_icon from '../public/barber-shop.png'
import Image from 'next/image'
export default function Navbar({children, style}:any) {
  return (
    <>
        <div className={`z-50 w-full ` + style}>
            <div className="flex items-center justify-between flex-wrap bg-slate-900 pt-2 pl-6 pr-6 pb-1">
                <div className="flex items-center justify-between flex-shrink-0 text-slate-200 mr-6">
                   <Link href="/home" className="flex items-center font-semibold text-xl tracking-tight">
                        <Image width="40" src={barber_icon} alt="barber salon"/>
                        Barber Shop
                    </Link>
                </div>
                <Dropdown elements = {["Profile", "Log Out"]}/>
            </div>
            {children}
        </div>
    </>
  )}