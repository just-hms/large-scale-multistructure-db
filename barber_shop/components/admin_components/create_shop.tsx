import { useFormik } from 'formik';
import GeneralDropdown from '../search_components/general_dropdown';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faBarsStaggered } from '@fortawesome/free-solid-svg-icons' 

export default function CreateShop() {
  const formik = useFormik({
    initialValues: {
        name: '',
        description: ''
    },
        onSubmit: values => {
            // TODO: check values and yadda yadda
            console.log(values)
    },
});
return (
    <>
     <div className='flex flex-col items-center justify-start text-left text-slate-300 text-lg w-full'>
        <div className='relative max-h-96 rounded-3xl bg-slate-700 shadow-md shadow-black/70 p-5 w-3/4'>
          <form className="space-y-4 md:space-y-6 w-full flex flex-col justify-center items-center" onSubmit={formik.handleSubmit}>
              <h1 className='text-xl font-bold text-center w-3/4 border-b border-slate-500'>Create Shop</h1>
              <div className='w-3/4'>
                  <label htmlFor="name" className="block mb-2 text-sm font-medium text-slate-200">Shop's Name</label>
                  <input type="text" name="name" id="name" onChange={formik.handleChange} value={formik.values.name} className="border border-slate-600 text-slate-300 sm:text-sm rounded-lg focus:ring-slate-700 block w-full p-2.5 bg-slate-700" placeholder="Name" required=""/>
                  <label htmlFor="description" className="mt-2 block mb-2 text-sm font-medium text-slate-200">Shop's Description</label>
                  <input type="text" name="description" id="description" onChange={formik.handleChange} value={formik.values.description} placeholder="Description" className="border border-slate-600 text-slate-300 sm:text-sm rounded-lg focus:ring-slate-700 block w-full p-2.5 bg-slate-700" required=""/>
              </div>
              <div className='flex w-3/4 items-center justify-end'>
                <div className="text-lg text-slate-200">

                  <div className="px-3 py-2 rounded-full bg-slate-700 shadow-sm shadow-black/70">
                    <GeneralDropdown elements={['pippo','pluto','topolino']} placeholder="Select Barber"><></></GeneralDropdown>
                  </div>
                </div>
              </div>
              <button type="submit" className="w-3/4 text-white bg-rose-800 hover:bg-rose-700 focus:ring-4 focus:outline-none focus:ring-rose-300/0 font-medium rounded-lg border-slate-700 text-sm px-5 py-2.5 text-center dark:bg-white-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800">Create it!</button>
          
          </form>
        </div>
      </div>
    </>
)}
  
  // TODO: For each account we gotta return the mail and the actions, i.e: delete account, change permissions, on click over button show results if any