import { useFormik } from 'formik';
import {useRouter} from 'next/router'
import {useState, useEffect} from 'react'
import Image from 'next/image';
import Link from 'next/link';
import barber_icon from '../public/barber-shop.png'

export default function LoginForm() {
    const router = useRouter()
    const [route] = useState()
    const formik = useFormik({
        initialValues: {
            email: '',
            password: ''
        },
            onSubmit: values => {
                fetch('http://127.0.0.1:7000/user/login/', {
                    method: 'POST',
                    headers: {
                        'Accept': 'application/json',
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                            "email": "email",
                            "password":"password"
                    })
                })
                .then(response => response.json())
                .then(response => console.log(JSON.stringify(response)))
                .catch((e) => {
                console.error(`An error occurred: ${e}`)
                });
                // router.push("/home")
        },
    });
    return (
        <>
        <section className="bg-slate-50 dark:bg-slate-900">
        <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto h-screen lg:py-0">
            <Link href="/" className="flex items-center mb-6 text-2xl font-semibold text-slate-900 dark:text-white">
                <Image width="40" src={barber_icon} alt="barber salon"/>
                Barber Shop    
            </Link>
            <div className="w-full bg-white rounded-lg shadow md:mt-0 sm:max-w-md xl:p-0 dark:bg-slate-800 dark:border-slate-700">
                <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
                    <h1 className="text-xl text-center font-bold leading-tight tracking-tight text-slate-900 md:text-2xl dark:text-white">
                        Sign in to your account
                    </h1>
                    <form className="space-y-4 md:space-y-6" onSubmit={formik.handleSubmit}>
                        <div>
                            <label htmlFor="email" className="block mb-2 text-sm font-medium text-slate-900 dark:text-white">Your email</label>
                            <input type="email" name="email" id="email" onChange={formik.handleChange} value={formik.values.email} className="bg-slate-50 border border-slate-300 text-slate-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-slate-700 dark:border-slate-600 dark:placeholder-slate-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="name@company.com" required/>
                        </div>
                        <div>
                            <label htmlFor="password" className="block mb-2 text-sm font-medium text-slate-900 dark:text-white">Password</label>
                            <input type="password" name="password" id="password" onChange={formik.handleChange} value={formik.values.password} placeholder="••••••••" className="bg-slate-50 border border-slate-300 text-slate-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-slate-700 dark:border-slate-600 dark:placeholder-slate-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" required/>
                        </div>
                        <div className="flex items-center justify-between">
                            <div className="flex items-start">
                                <div className="flex items-center h-5">
                                    <input id="remember" aria-describedby="remember" type="checkbox" className="w-4 h-4 border border-slate-300 rounded bg-slate-50 focus:ring-3 focus:ring-primary-300 dark:bg-slate-700 dark:border-slate-600 dark:focus:ring-primary-600 dark:ring-offset-slate-800"/>
                                </div>
                                <div className="ml-3 text-sm">
                                    <label htmlFor="remember" className="text-slate-500 dark:text-slate-300">Remember me</label>
                                </div>
                            </div>
                            <a href="#" className="text-sm font-medium text-primary-600 hover:underline dark:text-slate-300">Forgot password?</a>
                        </div>
                        <button type="submit" className="w-full text-white bg-rose-800 hover:bg-rose-700 focus:ring-4 focus:outline-none focus:ring-rose-300 font-medium rounded-lg border-slate-700 text-sm px-5 py-2.5 text-center dark:bg-white-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800">Sign in</button>
                        <p className="text-sm font-light text-slate-500 dark:text-slate-400">
                            Don’t have an account yet? <Link href="/signup" className="font-medium text-primary-600 hover:underline dark:text-primary-500">Sign up</Link>
                        </p>
                    </form>
                </div>
            </div>
        </div>
        </section>
        </>
    )
}