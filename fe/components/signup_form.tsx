import { useFormik } from 'formik';
import { useRouter } from 'next/router'
import { useState } from 'react'
import Image from 'next/image';
import Link from 'next/link';
import barber_icon from '../public/barber-shop.png'
import { signup } from '../lib/user';

export default function SignupForm() {
    const router = useRouter()
    const [route] = useState()
    const [response_state,setResponseState] = useState("")
    const formik = useFormik({
        initialValues: {
            email: '',
            password: '',
            repeatPassword: '',
        },
        onSubmit: async (values) => {
            if(values.password != values.repeatPassword){
                alert("asd")
                setResponseState("error")
            }else{
                const response = await signup(values)
                if(response.ok){
                    return router.push("/");
                }
            }
        },
    });
    return (
        <>
            <section className="bg-cian-50 dark:bg-slate-900">
                <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto h-screen lg:py-0">
                    <Link href="/" className="flex items-center mb-6 text-2xl font-semibold text-slate-900 dark:text-white">
                        <Image width="40" src={barber_icon} alt="barber salon"/>
                        Barber Shop
                    </Link>
                    <div className="w-full bg-white rounded-lg shadow  md:mt-0 sm:max-w-md xl:p-0 dark:bg-slate-800 dark:border-slate-700">
                        <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
                            <h1 className="text-xl text-center font-bold leading-tight tracking-tight text-slate-900 md:text-2xl dark:text-white">
                                Join Barber Shop
                            </h1>
                            <form className="space-y-4 md:space-y-6" onSubmit={formik.handleSubmit}>
                                <div>
                                    <label htmlFor="email" className="block mb-2 text-sm font-medium text-slate-900 dark:text-white">Your email</label>
                                    <input type="email" name="email" id="email" onChange={formik.handleChange} value={formik.values.email} className="bg-slate-50 border border-slate-300 text-slate-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-slate-700 dark:border-slate-600 dark:placeholder-slate-400 dark:text-white dark:focus:ring-slate-500 dark:focus:border-slate-500" placeholder="name@company.com" required />
                                </div>
                                <div>
                                    <label htmlFor="password" className="block mb-2 text-sm font-medium text-slate-900 dark:text-white">Password</label>
                                    <input type="password" name="password" id="password" onChange={formik.handleChange} value={formik.values.password} placeholder="••••••••" className="bg-slate-50 border border-slate-300 text-slate-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-slate-700 dark:border-slate-600 dark:placeholder-slate-400 dark:text-white dark:focus:ring-slate-500 dark:focus:border-slate-500" required />
                                </div>
                                <div>
                                    <label htmlFor="repeatPassword" className="block mb-2 text-sm font-medium text-slate-900 dark:text-white">Repeat Password</label>
                                    <input type="password" name="repeatPassword" id="repeatPassword" onChange={formik.handleChange} value={formik.values.repeatPassword} placeholder="••••••••" className="bg-slate-50 border border-slate-300 text-slate-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-slate-700 dark:border-slate-600 dark:placeholder-slate-400 dark:text-white dark:focus:ring-slate-500 dark:focus:border-slate-500" required />
                                </div>
                                <div className="flex items-center justify-center    ">
                                    <div className="flex items-start">
                                        <div className="flex items-center h-5">
                                            <input id="remember" aria-describedby="remember" type="checkbox" className="w-4 h-4 border border-slate-300 rounded bg-slate-50 focus:ring-3 focus:ring-primary-300 dark:bg-slate-700 dark:border-slate-600 dark:focus:ring-primary-600 dark:ring-offset-slate-800" />
                                        </div>
                                        <div className="ml-3 text-sm">
                                            <label htmlFor="remember" className="text-slate-500 dark:text-slate-300">Remember me</label>
                                        </div>
                                    </div>
                                </div>
                                <button type="submit" className="w-full text-white bg-rose-800 hover:bg-rose-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg border-slate-700 text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800">Sign up</button>
                                <p className="text-sm font-light text-slate-500 dark:text-slate-400"></p>
                            </form>
                        </div>
                    </div>
                </div>
            </section>
        </>
    )
}