import { useEffect, useState } from "react";
import { Menu, Transition } from '@headlessui/react'
import { Fragment} from 'react'
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Chart from "../chart_components/chart";
import { faBarsStaggered } from '@fortawesome/free-solid-svg-icons' 
import {adminAnalytics} from "../../lib/admin"
import PaginatedList from "../chart_components/list_chart";
import { ColorRing } from "react-loader-spinner";

export default function AnalyticsAdmin(){
    const ChartKeys = ['AppointmentsByMonth', 'NewUsersByMonth', 'ReviewsByMonth', 'ViewsByMonth']
    const listKeys = ['AppointmentCancellationUserRanking', 'GetAppointmentCancellationShopRanking','GetEngagementShopRanking']
    const [selectedAnalytics,setSelectedAnalytics] = useState({shown:'',key:''})
    const shownElements = ["Appointments per month", "Views per month", "Reviews per month", "New users per month", 
                                "Appointment cancellation", "Shops with most cancellations",  "Shops by engagement" ]
    const [analytics,setAnalytics] = useState<any[]>([])
    const [dropDownElements, setdropDownElements] = useState<any[]>([])
    const [analyticsData, setAnalyticsData] = useState(undefined)
    const [loaded, setLoaded] = useState(false)
    useEffect(()=>{
        const fetchData = async () => {
            const analytics_response = await (await  adminAnalytics()).json()
            setAnalytics(analytics_response.adminAnalytics)
            const retrieved_dropdownElements = []
            var i = 0
            for(var key in analytics_response.adminAnalytics){
                var shown = shownElements[i]
                retrieved_dropdownElements.push({key, shown})
                i++
            }
            setdropDownElements(retrieved_dropdownElements)
            setSelectedAnalytics(retrieved_dropdownElements[0])
            const show_key:any = retrieved_dropdownElements[0].key
            setAnalyticsData(analytics[show_key])
            setLoaded(true)
        }
        fetchData()
    },[])

    useEffect(()=>{
        const show_key:any = selectedAnalytics.key
        setAnalyticsData(analytics[show_key])
    },[selectedAnalytics])
    if(loaded){
        return(<>
            <div className="flex flex-col text-xl justify-center w-full items-start px-3">
                {/* DROPDOWN */}
                <div className="inline-block leading-none px-2 mr-4 py-2 rounded-full bg-slate-700 bg-opacity-60 backdrop-blur-lg drop-shadow-lg hover:bg-slate-700 my-3 hover:text-slate-500 text-slate-200">
                    <Menu as="div" className="relative inline-block">
                        <Menu.Button className="inline-flex w-full justify-center items-center rounded-full bg-opacity-20 text-slate-200 focus:outline-none">
                            <FontAwesomeIcon  icon={faBarsStaggered} className="  pr-2"/>

                            <div className=' px-1 flex hover:text-slate-200'>
                                {selectedAnalytics.shown}
                            </div>
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
                        <Menu.Items className="absolute right-0 mt-3 w-56 origin-top-right divide-y divide-slate-600 rounded-md bg-slate-800 shadow-sm ring-1 ring-black ring-opacity-5 focus:outline-none shadow-md shadow-black/70">
                            <div  className="px-1 py-1">
                                {dropDownElements.map((element:any)=>
                                <div key={`container-`+element.key}  className="px-1 py-1">
                                    <Menu.Item key={`item-`+element.key} >
                                        {({ active }) => (
                                        <button key={`button-`+element.key}  className={`hover:bg-slate-500/80 text-slate-300 group flex w-full items-center text-left rcenterounded-md px-2 py-2 `}
                                        onClick={async (e) => {  
                                            setSelectedAnalytics(element)
                                        }}>
                                            {element.shown}
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
            </div>
            
            {(ChartKeys.includes(selectedAnalytics.key))?
                <Chart analyticsData={analyticsData} title={selectedAnalytics.shown}></Chart>:
                <PaginatedList analyticsData={analyticsData} title={selectedAnalytics.shown}/>}
        </>)
    }else{
        return(
            <div className="flex items-center justify-center w-full h-screen">
                <ColorRing
                visible={true}
                height="15%"
                width="15%"
                ariaLabel="blocks-loading"
                wrapperStyle={{}}
                wrapperClass="blocks-wrapper"
                colors={['#83beda', '#ac1c3e', '#f81629', '#0071ab', '#005873']}
                />
            </div>
        )
    }
}