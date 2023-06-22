import { useEffect, useState, useRef } from "react";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faArrowAltCircleRight, faArrowAltCircleLeft } from "@fortawesome/free-solid-svg-icons";
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend,
  } from 'chart.js';
import { Line } from 'react-chartjs-2';
import React from "react";
ChartJS.register(
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend
  );


export default function Chart({analyticsData, title}:any) {    
    const [slice, setSlice] = useState(0)
    const page_size = 30
    const chartRef = useRef<ChartJS<"line", number[], string>>(null)
    // handle the analytic change
    useEffect(()=>{
        chartRef.current?.clear()
        if(chartRef.current?.data.labels){
            while(chartRef.current?.data.labels.length > 0){
                chartRef.current?.data.labels.pop()
            }
        }
        chartRef.current?.data.datasets.pop()
        if(analyticsData != undefined){
            var labels:any = Object.keys(analyticsData);
            labels.slice(slice * page_size, (slice+1)*page_size).map((label:any)=>{chartRef.current?.data.labels?.push(label)})
            chartRef.current?.data.datasets.push({
                label: title,
                data: (title == "Review votes per month")? labels.slice(slice * page_size, (slice+1)*page_size).map((label:any) => analyticsData[label].upCount - analyticsData[label].downCount) :labels.slice(slice * page_size, (slice+1)*page_size).map((label:any) => analyticsData[label]),
                borderColor: 'rgb(255, 99, 132)',
                backgroundColor: 'rgba(255, 99, 132, 0.5)',
            })
            chartRef.current?.update()
        }
    },[analyticsData, slice])

    return (
    <>
    <div className="w-full flex flex-col justify-center items-center px-3">
        <Line
        ref={chartRef}
        options={{
            responsive: true,
            plugins: {
                legend: {
                    position: 'top' as const,  
                    labels:{
                        color: "#cbd5e1",
                        font:{
                            size:14
                        },
                    }                  
                },
                title: {
                    display: true,
                    text: title,
                    font:{
                        size:20
                    },
                    color:"#cbd5e1"
                },  
            },
            scales:{
                x: {  
                    grid: {
                        color: '#94a3b8',
                    },
                    ticks: {
                        color: "#cbd5e1",
                    }
                },
                y: {  
                    grid: {
                        color: '#94a3b8',
                    },
                    ticks: {
                        color: "#cbd5e1",
                    }
                }
            }
        }} 
        data={{
            datasets: [],
        }} />
        {/* paginated view handling */}
        <div className="flex text-slate-200 items-center justify-center">
            <button onClick={
                (e)=>{
                    if(slice > 0){
                        setSlice(slice - 1)
                    }
                }
            }><FontAwesomeIcon className="text-slate-200 px-2 text-xl py-2" icon={faArrowAltCircleLeft}/></button>
            {slice+1}
            <button onClick={
                (e)=>{
                    if(slice < (Object.keys(analyticsData).length/page_size)-1){
                        setSlice(slice + 1)
                    }
                }
            }><FontAwesomeIcon className="text-slate-200 px-2 text-xl py-2" icon={faArrowAltCircleRight}/></button>
        </div>
    </div>
    </>
    )
}