import Dropdown from './user_dropdown'
import Link from 'next/link'
import barber_icon from '../public/barber-shop.png'
import Image from 'next/image'
export default function Navbar() {
  return (
    <>
        <div className='z-50'>
            <div className="flex items-center justify-between flex-wrap bg-slate-900 pt-2 pl-6 pr-6 pb-1">
                <div className="flex items-center flex-shrink-0 text-slate-200 mr-6">
                   <Link href="/home" className="flex items-center font-semibold text-xl tracking-tight">
                        {/* TODO:change with svg */}
                        <Image width="40" src={barber_icon} alt="barber salon"/>
                        Barber Shop
                    </Link>
                </div>
                <div className="w-full block flex-grow flex items-center w-auto justify-between">
                    <div className="text-sm lg:flex-grow collapse lg:visible">
                        <Link href="#responsive-header" className="block mt-4 lg:inline-block lg:mt-0 text-slate-200 hover:text-white mr-4">
                        Option One
                        </Link>
                        <Link href="#responsive-header" className="block mt-4 lg:inline-block lg:mt-0 text-slate-200 hover:text-white mr-4">
                        Option Two
                        </Link>
                        <Link href="#responsive-header" className="block mt-4 lg:inline-block lg:mt-0 text-slate-200 hover:text-white">
                        Option Three
                        </Link>
                    </div>
                    <Dropdown elements = {["Profile", "Log Out"]}/>
                </div>
            </div>
        </div>
    </>
  )}