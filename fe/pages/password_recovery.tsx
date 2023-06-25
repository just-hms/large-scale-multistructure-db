import { useFormik } from 'formik';
import { useRouter } from 'next/router'
import { useState, useEffect } from 'react'
import Image from 'next/image';
import Link from 'next/link';
import barber_icon from '../public/barber-shop.png'
import { changePassword } from '../lib/user';

export default function PasswordRecovery() {
    const router = useRouter()
    const [loaded,setLoaded] = useState(false)
    const [error,setError] = useState("")
    const formik = useFormik({
        initialValues: {
            email: '',
            password: '',
            repeatPassword: '',
        },
        onSubmit: async (values) => {
            if(values.password != values.repeatPassword){
                setError("Passwords not matching")
            }else{
                const response = await changePassword(values)
                if(response.ok){
                    return router.push("/");
                }else if(response.status == 400){
                    setError("Email not valid")
                }
            }
        },
    });
    useEffect(()=>{
        if(!localStorage.getItem('token')){
          router.push("/")
        }else{
          setLoaded(true)
        }
      },[])
    if(!loaded){
        return <div></div> //show nothing or a loader
    }
    return (
        <>
            <section className="bg-slate-900 h-screen overflow-auto">
                <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto h-screen lg:py-0">
                    <Link href="/home" className="flex items-center mb-6 text-2xl font-semibold text-slate-300">
                        <Image width="40" src={barber_icon} alt="barber salon"/>
                        Barber Shop
                    </Link>
                    <div className="w-full rounded-lg shadow  md:mt-0 sm:max-w-md xl:p-0 bg-slate-800 border-slate-700">
                        <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
                            <form className="space-y-4 md:space-y-6" onSubmit={formik.handleSubmit}>
                                <div>
                                    <label htmlFor="email" className="block mb-2 text-sm font-medium text-slate-300">Your email</label>
                                    <input type="email" name="email" id="email" onChange={formik.handleChange} value={formik.values.email} className="border sm:text-sm rounded-lg block w-full p-2.5 bg-slate-700 border-slate-600 placeholder-slate-400 text-slate-300 focus:ring-slate-500 focus:border-slate-500" placeholder="name@company.com" required />
                                </div>
                                <div>
                                    <label htmlFor="password" className="block mb-2 text-sm font-medium text-slate-300">New Password</label>
                                    <input type="password" name="password" id="password" onChange={formik.handleChange} value={formik.values.password} placeholder="••••••••" className="border sm:text-sm rounded-lg block w-full p-2.5 bg-slate-700 border-slate-600 placeholder-slate-400 text-slate-300 focus:ring-slate-500 focus:border-slate-500" required />
                                </div>
                                <div>
                                    <label htmlFor="repeatPassword" className="block mb-2 text-sm font-medium text-slate-300">Repeat New Password</label>
                                    <input type="password" name="repeatPassword" id="repeatPassword" onChange={formik.handleChange} value={formik.values.repeatPassword} placeholder="••••••••" className="border sm:text-sm rounded-lg block w-full p-2.5 bg-slate-700 border-slate-600 placeholder-slate-400 text-slate-300 focus:ring-slate-500 focus:border-slate-500" required />
                                </div>
                                <p className='text-rose-600 text-sm'>{error}</p>
                                <button type="submit" className="w-full text-slate-300 bg-rose-800 hover:bg-rose-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg border-slate-700 text-sm px-5 py-2.5 text-center">Change Password</button>
                                <p className="text-sm font-light text-slate-500"></p>
                            </form>
                        </div>
                    </div>
                </div>
            </section>
        </>
    )
}