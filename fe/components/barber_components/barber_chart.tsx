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
    // var slice = 1
    const chartRef = useRef<ChartJS<"line", number[], string>>(null)
    useEffect(()=>{
        chartRef.current?.clear()
        if(chartRef.current?.data.labels){
            while(chartRef.current?.data.labels.length > 0){
                chartRef.current?.data.labels.pop()
            }
        }
        chartRef.current?.data.datasets.pop()
        if(analyticsData != undefined){
            var labels:any = Object.keys(analyticsData).slice(-30);
            labels.map((label:any)=>{chartRef.current?.data.labels?.push(label)})
            chartRef.current?.data.datasets.push({
                label: title,
                data: (title == "Review votes per month")? labels.map((label:any) => analyticsData[label].upCount - analyticsData[label].downCount) :labels.map((label:any) => analyticsData[label]),
                borderColor: 'rgb(255, 99, 132)',
                backgroundColor: 'rgba(255, 99, 132, 0.5)',
            })
            chartRef.current?.update()
        }else{
        }
    },[analyticsData])

    return (
    <>
    {( title != "Inactive users" && title != "Weighted Rating")?
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
    </div>:<></>}
    </>
    )
}