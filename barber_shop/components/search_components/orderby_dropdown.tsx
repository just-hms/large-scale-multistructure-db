import { Menu, Transition } from '@headlessui/react'
import { Fragment} from 'react'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faBarsStaggered } from '@fortawesome/free-solid-svg-icons' 

export default function OrderByDropdown({elements}) {
  return (
    <div className="inline-block text-sm leading-none rounded-full text-white border-slate-700 hover:text-slate-500">
    <Menu as="div" className="relative inline-block">
        <Menu.Button className="inline-flex w-full justify-center items-center rounded-full bg-slate-800 bg-opacity-20 text-sm font-medium text-white focus:outline-none">
            {/* <Image className="w-full h-full" src="https://flowbite.s3.amazonaws.com/blocks/marketing-ui/logo.svg" width="20" height="20" alt="logo" /> */}
            <p className='pr-3 '>Order by</p> 
            <FontAwesomeIcon  icon={faBarsStaggered} className=" text-xl pr-2"/>
        </Menu.Button>
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
                <div className="px-1 py-1">
                {elements.map((element)=>
                <div className="px-1 py-1">
                    <Menu.Item>
                        {({ active }) => (
                        <button className={`hover:bg-slate-500/80 text-white group flex w-full items-center rounded-md px-2 py-2 text-sm`}>
                            {element}
                        </button>
                        )}
                    </Menu.Item>
                </div>
                )}

                </div>
            </Menu.Items>
        </Transition>
    </Menu>
</div>
  )
}