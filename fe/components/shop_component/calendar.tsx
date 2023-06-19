import { getAppointment } from "../../lib/shops"
import styled from "@emotion/styled";
// Calendar
import FullCalendar from '@fullcalendar/react' // must go before plugins
import dayGridPlugin from '@fullcalendar/daygrid' // a plugin!
import timeGridPlugin from '@fullcalendar/timegrid'
// Confirm
import { confirmAlert } from "react-confirm-alert";
import "react-confirm-alert/src/react-confirm-alert.css";
import moment from 'moment-timezone';
import { DateEnv } from "@fullcalendar/core/internal";
// confirm pop-up

const submit = ({shopid, info}:any) => {
confirmAlert({
    customUI: ({ onClose }) => {
      return (
        <div className='custom-ui flex border bg-slate-800 text-slate-100 p-5 rounded-lg'>
            <div className="flex-col items-center justify-center">
                <h1 className="text-xl">Do you want to book this appointment?</h1>
                <div className="flex items-center justify-between w-full">
                    <button className="m-2 w-1/3 bg-slate-700 hover:bg-slate-600 focus:border-rose-600 font-medium rounded-lg border-slate-800 text-sm px-5 py-2.5 text-center"
                        onClick={onClose}>No</button>
                    <button className="m-2 w-2/3 bg-slate-600 hover:bg-slate-500 focus:border-rose-600 font-medium rounded-lg border-slate-800 text-sm px-5 py-2.5 text-center"
                        onClick={async (e) => {
                            // alert((new Date(info.event.startStr)).toISOString())
                            // alert(info.event.startStr)
                            const response = await getAppointment(shopid,info.event.startStr)
                            if(response.status == 201){
                                alert("Slot booked correctly!")
                            }else{
                                alert((await response.json()).error)
                            }   
                            // alert(parseInt((new Date(info.event.startStr).getTime() / 1000).toFixed(0)))
                            onClose()
                        }}
                        >
                        Yes, Book it!
                    </button>
                </div>
            </div>
        </div>
      );
    }
  });
}
// calendar styling
export const StyleWrapper = styled.div`
  .fc{
    color: white;
  }
  .fc-day-today {
    background: none !important;
} `

// this creates the calendar slots based on the already booked appointments
// must return an array of objects with: title, start, end
export const craftEventObject = (index:any,calendar:any) =>{
    // console.log(calendar)
    const date = moment().add(30*index, 'm')

    if(date.hours() > 20|| date.hours() < 7 ){
        return null
    }
    if (date.minutes() >= 30){
        date.set('hour',date.hours()+1)
        date.set('minute',0)
    }else{
        date.set('minute',30)
    }
    date.set('second',0)
    date.set('millisecond',0)
    // TODO TAKE THIS OUTSIDE
    for(var i in calendar){
        if((new Date(calendar[i].Start)).getTime() === date.toDate().getTime()){
            if(calendar[i].BookedAppointments == calendar[i].Employees){
                return null
            }
        }
    }
    return {
        'title' : "Slot Disponibile",
        'start' : date.toDate().toISOString(),
        'end'   : moment(date).add(30, 'm').toDate().toISOString(),
    }
}

export const fillCalendar = (calendar:any, employees:any)=>{
    let events = []
    if(employees != 0){
        for(let index = 0; index < 48*30; index+=1){
            var event = craftEventObject(index,calendar)
            if(event){
                events.push(event)
            }
        }
    }
    return events
}

export default function Calendar({shopid,calendar, employees}:any) {
    return (
        <>
        <div className="w-full lg:w-5/6 h-1/3 mt-0 px-3 lg:py-3 transform -translate-y-20">
            <div className="flex justify-center items-center">
                <div className="w-full rounded-lg bg-slate-700 shadow-md shadow-black/70 p-3 ">
                    <h1 className="text-2xl text-center font-bold leading-tight tracking-tight text-slate-200 pt-5 ml-3 mr-3 break-words">
                    Free time Slots
                    </h1>
                    <StyleWrapper>
                    <FullCalendar
                        plugins={[ dayGridPlugin , timeGridPlugin]}
                        initialView="timeGridDay"
                        slotMinTime="08:00:00"
                        slotMaxTime="20:00:00"
                        events={fillCalendar(calendar, employees)}
                        allDaySlot={false}
                        eventClick={
                            async (info) => {
                                submit({shopid,info})
                            }
                        }
                    />
                    </StyleWrapper>
                </div>
            </div>
        </div>
        </>
    );
}