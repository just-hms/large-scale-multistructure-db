import Dropdown from './dropdown'
import Link from 'next/link'
import barber_icon from '../public/barber-shop.png'
import Image from 'next/image'
export default function Navbar() {
  return (
    <nav className="flex items-center justify-between flex-wrap bg-slate-900 pt-2 pl-6 pr-6 pb-2">
        <div className="flex items-center flex-shrink-0 text-white mr-6 ">
            {/* <svg className="fill-current h-8 w-8 mr-2" width="54" height="54" viewBox="0 0 54 54" xmlns="http://www.w3.org/2000/svg"><path d="M13.5 22.1c1.8-7.2 6.3-10.8 13.5-10.8 10.8 0 12.15 8.1 17.55 9.45 3.6.9 6.75-.45 9.45-4.05-1.8 7.2-6.3 10.8-13.5 10.8-10.8 0-12.15-8.1-17.55-9.45-3.6-.9-6.75.45-9.45 4.05zM0 38.3c1.8-7.2 6.3-10.8 13.5-10.8 10.8 0 12.15 8.1 17.55 9.45 3.6.9 6.75-.45 9.45-4.05-1.8 7.2-6.3 10.8-13.5 10.8-10.8 0-12.15-8.1-17.55-9.45-3.6-.9-6.75.45-9.45 4.05z" /></svg> */}

            <Link href="/home" className="flex items-center font-semibold text-xl tracking-tight">
                {/* TODO:change with svg */}
                <Image width="40" src={barber_icon} alt="barber salon"/>
                Barber Shop
            </Link>
        </div>
        <div className="w-full block flex-grow flex items-center w-auto justify-between">
            <div className="text-sm lg:flex-grow">
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
            <Dropdown/>
        </div>
    </nav>
  )}