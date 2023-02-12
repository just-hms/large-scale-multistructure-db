import { useFormik } from 'formik';
import {useRouter} from 'next/router'
import {useState, useEffect} from 'react'
import Image from 'next/image';
import Link from 'next/link';
import barber_icon from '../public/barber-shop.png'
import { signin } from '../lib/user';
import { getUserInfos } from '../lib/user';

export default function LoginForm() {
    const router = useRouter()
    const [error,setError] = useState("")
    const formik = useFormik({
        initialValues: {
            email: '',
            password: ''
        },
        onSubmit: async (values) =>{
            const response = await signin(values)
            if (response.ok){
                const response_json = await response.json()
                localStorage.setItem("token", response_json.token)
                const fetchData = async () => {
                    const userData = (await (await getUserInfos()).json())
                    localStorage.setItem("isAdmin", userData.user.IsAdmin)
                    return router.push("/home");
                  }
                fetchData()
            }else{
                setError("Username or Password not valid")
            }
        },
    });
    return (
        <>
        <section className="bg-slate-900 h-screen overflow-auto">
        <div className="flex h-full flex-col items-center justify-center px-6 py-8 mx-auto lg:py-0">
            <Link href="/" className="flex items-center mb-6 text-2xl font-semibold text-slate-300">
                <Image width="40" src={barber_icon} alt="barber salon"/>
                Barber Shop    
            </Link>
            <div className="w-full rounded-lg shadow md:mt-0 sm:max-w-md xl:p-0 bg-slate-800 border-slate-700">
                <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
                    <h1 className="text-xl text-center font-bold leading-tight tracking-tight md:text-2xl text-slate-300">
                        Sign in to your account
                    </h1>
                    <form className="space-y-4 md:space-y-6" onSubmit={formik.handleSubmit}>
                        <div>
                            <label htmlFor="email" className="block mb-2 text-sm font-medium text-slate-300">Your email</label>
                            <input type="email" name="email" id="email" onChange={formik.handleChange} value={formik.values.email} className="bg-slate-50 border border-slate-300 text-slate-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-slate-700 dark:border-slate-600 dark:placeholder-slate-400 dark:text-slate-300 dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="name@company.com" required/>
                        </div>
                        <div>
                            <label htmlFor="password" className="block mb-2 text-sm font-medium text-slate-300">Password</label>
                            <input type="password" name="password" id="password" onChange={formik.handleChange} value={formik.values.password} placeholder="••••••••" className="bg-slate-50 border border-slate-300 text-slate-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-slate-700 dark:border-slate-600 dark:placeholder-slate-400 dark:text-slate-300 dark:focus:ring-blue-500 dark:focus:border-blue-500" required/>
                        </div>
                        <p className='text-rose-600 text-sm'>{error}</p>
                        {/* <Link href="/password_recovery" className="text-sm font-medium hover:underline text-slate-300">Forgot Password?</Link> */}
                        <button type="submit" className="w-full text-slate-300 bg-rose-800 hover:bg-rose-700 focus:ring-4 focus:outline-none focus:ring-rose-300 font-medium rounded-lg border-slate-700 text-sm px-5 py-2.5 text-center dark:bg-white-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800">Sign in</button>
                        <p className="text-sm font-light text-slate-400 ">
                            Don’t have an account yet? <Link href="/signup" className="font-medium text-slate-400 hover:underline">Sign up</Link>
                        </p>
                    </form>
                </div>
            </div>
        </div>
        </section>
        </>
    )
}