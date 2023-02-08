import { Menu, Transition } from '@headlessui/react'
import { Fragment} from 'react'
import { useRouter } from 'next/router'

export default function UserDropdown({elements}:any){
  const router = useRouter()
  return (
    <div className="inline-block text-sm leading-none rounded-full text-white border-slate-700 hover:text-slate-500">
      <Menu as="div" className="relative inline-block">
        <div>
          <Menu.Button className="inline-flex w-full justify-center rounded-full bg-slate-800 bg-opacity-20 text-sm font-medium text-white focus:outline-none">
            {/* <Image className="w-full h-full" src="https://flowbite.s3.amazonaws.com/blocks/marketing-ui/logo.svg" width="20" height="20" alt="logo" /> */}
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} className="w-10 h-10 stroke-slate-500 hover:stroke-slate-400">
             <path strokeLinecap="round" strokeLinejoin="round" d="M17.982 18.725A7.488 7.488 0 0012 15.75a7.488 7.488 0 00-5.982 2.975m11.963 0a9 9 0 10-11.963 0m11.963 0A8.966 8.966 0 0112 21a8.966 8.966 0 01-5.982-2.275M15 9.75a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
          </Menu.Button>
        </div>
        <Transition
          as={Fragment}
          enter="transition ease-out duration-100"
          enterFrom="transform opacity-0 scale-95"
          enterTo="transform opacity-100 scale-100"
          leave="transition ease-in duration-75"
          leaveFrom="transform opacity-100 scale-100"
          leaveTo="transform opacity-0 scale-95"
        >
          <Menu.Items className="absolute right-0 mt-3 w-56 origin-top-right divide-y divide-slate-600 rounded-md bg-slate-800 shadow-sm ring-1 ring-black ring-opacity-5 focus:outline-none z-10 shadow-md shadow-black/70">
            {elements.map((element:any)=>
              <div key={`container-`+element} className="px-1 py-1">
                <Menu.Item key={`item-`+element}>
                {({ active }) => (
                  <button key={`link-`+element} className={`hover:bg-slate-500/80 text-white group flex w-full items-center rounded-md px-2 py-2 text-sm`}
                  onClick={ async (event) => {
                    if(element === "Profile"){
                      return router.push("/user");
                    }else{
                      const response = await fetch("/api/logout", {
                        method: "POST",
                        headers: { "Content-Type": "application/json" }
                      });
                      if (response.ok) {
                        return router.push("/");
                      }
                    }
                    }}>
                    {element}
                    
                  </button>
                  )}
                  </Menu.Item>
              </div>
            )}
          </Menu.Items>
        </Transition>
      </Menu>
    </div>
  )
}