import { useFormik } from 'formik';
import { modifyShopEmployees } from '../../lib/barber';
import { useEffect, useRef } from 'react';
export default function Holidays({ shopData }:any) {
    const inputRef = useRef<any>();
    const formik = useFormik({
        initialValues: {
            employeesNumber: shopData.Employees,
        },
        onSubmit: async (values) => {
            const response = await modifyShopEmployees(shopData.ID,values)
            if(response.status == 202){
              alert("done")
              shopData.Employees = values.employeesNumber
            }else{
              alert("there was an error")
            }
        },
        });
    useEffect(()=>{
        formik.values.employeesNumber = shopData.Employees
        inputRef.current.value = shopData.Employees
        },[shopData.Employees])
  return (
    <>
        <form className="w-full p-5 text-slate-300 h-full flex flex-col items-center justify-start rounded-3xl bg-slate-700 bg-opacity-60 backdrop-blur-lg drop-shadow-lg" onSubmit={formik.handleSubmit}>
        <h1 className='text-xl font-bold py-2'>Modify Number of available employees</h1>
        <input ref={inputRef} type="number" name="employeesNumber" id="employeesNumber" onChange={formik.handleChange} value={formik.values.employeesNumber} placeholder="0" className="border border-slate-500 bg-slate-600  text-slate-300 sm:text-sm rounded-lg focus:ring-slate-700 block p-2.5 bg-slate-700 " required/>
        <button type="submit" className="px-3 py-1 mx-2 my-4 bg-rose-800 bg-opacity-70 text-slate-300 text-xs rounded-full focus:bg-red-700 hover:bg-red-700 focus:outline-none transition duration-150 ease-in-out ">Submit Changes</button>
        </form>
    </>
  );
}
