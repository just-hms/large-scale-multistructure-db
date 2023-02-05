import Head from 'next/head'
import { Inter } from '@next/font/google'
import LoginForm from '../components/login_form'
import SignupForm from '../components/signup_form'

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  return (
    <>
    <Head>
      <title>Sign up</title>
    </Head>
    <SignupForm/>
    </>
  )
}
