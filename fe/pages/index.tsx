import Head from 'next/head'
import { Inter } from '@next/font/google'
import LoginForm from '../components/login_form'

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  return (
    <>
    <Head>
      <title>Barber Shop</title>
    </Head>
    <LoginForm/>
    </>
  )
}
