import { useEffect, useState, useRef } from "react";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faArrowAltCircleRight, faArrowAltCircleLeft } from "@fortawesome/free-solid-svg-icons";

export default function PaginatedList({analyticsData, title}:any) {    
    const [slice, setSlice] = useState(0)
    const page_size = 18
    const [shownData, setShownData]=useState<any[]>((Array.isArray(analyticsData))?analyticsData.slice(slice * page_size, (slice+1)*page_size):[])
    useEffect(()=>{
        setShownData((Array.isArray(analyticsData))?analyticsData.slice(slice * page_size, (slice+1)*page_size):[])
    },[analyticsData, slice])

    if(Array.isArray(analyticsData)){
        return (
        <>
        <div className="flex items-center justify-center w-full">
            <div className="w-2/3 h-full py-2 flex flex-col justify-start items-center bg-slate-700 rounded-lg  px-3 text-slate-300">
                <h1 className="text-xl font-bold pb-2">{title}</h1>
                {shownData.map((element:any,index)=>{
                    if(typeof(element) == 'object'){
                        var keys = []
                        for(var key in element){
                            keys.push(key)
                        }
                        return<>
                            <ul key={element[keys[1]]+''+index}>
                                {element[keys[1]]}: {element[keys[0]]} 
                            </ul>
                        </>
                    }else{
                        return<>
                            <ul key={element+''+index}>
                                {element}
                            </ul>
                        </>
                    }
                })}
                <div className="flex text-slate-200 items-center justify-center">
                    <button onClick={
                        (e)=>{
                            if(slice > 0){
                                setSlice(slice - 1)
                            }
                        }
                    }><FontAwesomeIcon className="px-2 text-xl py-2" icon={faArrowAltCircleLeft}/></button>
                    {slice+1}
                    <button onClick={
                        (e)=>{
                            if(slice < (analyticsData.length/page_size)-1){
                                setSlice(slice + 1)
                            }
                        }
                    }><FontAwesomeIcon className="px-2 text-xl py-2" icon={faArrowAltCircleRight}/></button>
                </div>
            </div>
        </div>
        </>
        )
    }else if(typeof(analyticsData) == 'number'){
        return <>
        <div className="w-full flex items-centr justify-center text-slate-300 text-xl">
            Weighted rating based on reviews and upvotes: {analyticsData}
        </div>
        </>}
    else
        return <></>
}