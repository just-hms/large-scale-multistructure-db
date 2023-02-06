import Link from 'next/link'
export default function UserInfos() {
  return (
    <div className='flex flex-col items-center justify-center text-left text-slate-300 text-lg w-full lg:w-5/6 rounded-3xl bg-slate-700 py-3 px-3 lg:px-0 shadow-md shadow-black/70'>
      <h1 className='text-2xl'>My account</h1>
      <p className='pt-2 pb-2 w-full lg:w-1/2'>Email</p>
      <div className='border border-slate-500 rounded-md p-2 w-full lg:w-1/2 overflow-scroll'>email@email.com</div>
      <p className='pt-2 pb-2 w-full lg:w-1/2'>Password</p>
      <div className='flex flex-col lg:flex-row w-full lg:w-1/2 justify-start items-center pb-5'>
        <div className='border border-slate-500 rounded-md p-2 mr-2 w-full lg:w-2/3 overflow-scroll'>●●●●●●●●●●</div>
        <Link href="#" className='w-full lg:w-1/3 pt-1 lg:p-0 text-center lg:text-right hover:text-white'>Change Password</Link>
      </div>
      <button type="submit" className="w-full lg:w-1/2 lg:mt-5 text-white bg-slate-600 hover:bg-slate-500 focus:outline-none font-medium rounded-2xl text-sm px-5 py-2.5 text-center ">Log Out</button>

      <div className='w-full lg:w-1/2 pb-5 border-b border-slate-500'></div>
      <button type="submit" className="w-full lg:w-1/2 mt-5 text-white bg-rose-800 hover:bg-rose-700 focus:outline-none font-medium rounded-2xl border-slate-700 text-sm px-5 py-2.5 text-center ">Delete Account</button>
    </div>
  )
}