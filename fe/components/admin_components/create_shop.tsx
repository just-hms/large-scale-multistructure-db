import { useFormik } from 'formik';
import { createShop } from '../../lib/admin';
import { useState } from 'react';

export default function CreateShop({accounts}:any) {
  const [message,setMessage] = useState("")
  const formik = useFormik({
    initialValues: {
        name: '',
        address: '',
        description: '',
        employeesNumber: '',
    },
        onSubmit: async (values, {resetForm}) => {
            setMessage("") 
            const response = await createShop(values)
            if(response.ok){
              setMessage("Barber Shop Created") 
              resetForm();
            }
    },
});
return (
    <>
     <div className='flex flex-col items-center justify-start text-left text-slate-300 text-lg w-full'>
        <div className='relative max-h-96 overflow-auto rounded-3xl bg-slate-700 shadow-md shadow-black/70 px-5 w-full lg:w-3/4'>
          <h1 className='sticky py-5 top-0 bg-slate-700 w-full text-xl font-bold text-center w-full border-b border-slate-500'>Create Shop</h1>
          <form className="space-y-4 md:space-y-6 py-5 w-full flex flex-col justify-center items-center" onSubmit={formik.handleSubmit}>
              <div className='w-full lg:w-3/4'>
                  <label htmlFor="name" className="block mb-2 text-sm font-medium text-slate-200">Shop Name</label>
                  <input type="text" name="name" id="name" onChange={formik.handleChange} value={formik.values.name} className="border border-slate-500  bg-slate-600 text-slate-300 sm:text-sm rounded-lg focus:ring-slate-700 block w-full p-2.5 bg-slate-700" placeholder="Name" required/>
                  <label htmlFor="description" className="mt-2 block mb-2 text-sm font-medium text-slate-200">Shop Description</label>
                  <input type="text" name="description" id="description" onChange={formik.handleChange} value={formik.values.description} placeholder="Description" className="border border-slate-500 bg-slate-600  text-slate-300 sm:text-sm rounded-lg focus:ring-slate-700 block w-full p-2.5 bg-slate-700" required/>
                  <label htmlFor="address" className="mt-2 block mb-2 text-sm font-medium text-slate-200">Address</label>
                  <input type="text" name="address" id="address" onChange={formik.handleChange} value={formik.values.address} placeholder="Address" className="border border-slate-500 bg-slate-600  text-slate-300 sm:text-sm rounded-lg focus:ring-slate-700 block w-full p-2.5 bg-slate-700" required/>
                  <label htmlFor="employeesNumber" className="mt-2 block mb-2 text-sm font-medium text-slate-200">Employees Number</label>
                  <input type="number" name="employeesNumber" id="employeesNumber" onChange={formik.handleChange} value={formik.values.employeesNumber} placeholder="0" className="border border-slate-500 bg-slate-600  text-slate-300 sm:text-sm rounded-lg focus:ring-slate-700 block w-full p-2.5 bg-slate-700 " required/>
              </div>
              <button type="submit" className="w-full lg:w-3/4 text-slate-200 bg-rose-800 hover:bg-rose-700 focus:ring-4 focus:outline-none focus:ring-rose-300/0 font-medium rounded-2xl border-slate-700 text-sm px-5 py-2.5 text-center dark:bg-white-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800">Create it!</button>
              <p>{message}</p>
          </form>
        </div>
      </div>
    </>
)}
  
