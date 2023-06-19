import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faArrowUp, faArrowDown, faStar, faTrash} from "@fortawesome/free-solid-svg-icons";
import { deleteVote, submitVote, deleteReview } from "../../lib/shops";
import { useEffect, useState } from "react";
import { getReviews } from "../../lib/shops";

export default function Reviews({shopid, userid}:any) {

    const [reviews, setreviewsData] = useState<any[]>([])
    const [reload, setReload] = useState(false)
    useEffect(()=>{
        const fetchData = async (shopid:any) => {
            const reviews_object = await (await getReviews(shopid)).json()
            if(reviews_object.reviews == null)
                setreviewsData([])
            else
                setreviewsData(reviews_object.reviews)
        }
        fetchData(shopid)
      },[reload])

    const handleVote = async (review:any,vote:boolean) =>{
        if(review.UpVotes.includes(userid) || review.DownVotes.includes(userid)){
            await deleteVote(shopid,review.ID)
        }else{
            await submitVote(shopid, review.ID, vote)
        }
        setReload(!reload)
    }

    return (
        <>
        <div>
        {reviews.map((review:any)=>
                <div key={review.ID} className="w-full text-slate-200 ">
                    <div key={review.ID+"container"} className="flex flex-col items-center justify-start w-full bg-slate-700 rounded-lg p-3 shadow-md shadow-black/30 mb-3 z-0">
                        <div key={review.ID+"name"} className="text-md w-full text-left font-normal">{review.Username}</div>
                        <div key={review.ID+"title"} className="text-xl mb-3 font-bold flex items-center justify-between w-full text-left">
                            <div className="flex items-center">
                                {[...Array(review.Rating)].map((_,index)=><FontAwesomeIcon key={index} icon={faStar} className="text-sm"/>)}
                            </div>
                            <div className="flex items-center">
                                <button key={review.ID+"arrowUp"} className="hover: text-white mr-3 text-sm" onClick={(e)=>handleVote(review,true)}>
                                    {<FontAwesomeIcon key={review.ID+"arrowUpIcon"} icon={faArrowUp} className={(review.UpVotes.includes(userid))?"text-orange-500":""}/>}
                                </button>
                                <button key={review.ID+"arrowDown"} className="hover: text-white mr-3 text-sm" onClick={(e)=>handleVote(review,false)}>
                                    <FontAwesomeIcon key={review.ID+"arrowDownIcon"}  icon={faArrowDown} className={(review.DownVotes.includes(userid))?"text-orange-500":""}/>
                                </button>
                                <div className={`text-sm ${(review.UpVotes.length - review.DownVotes.length) >= 0 ? "text-green-700" : "text-rose-600"}`}> {((review.UpVotes.length - review.DownVotes.length) >= 0) ? "+": "-"}{review.UpVotes.length - review.DownVotes.length}</div>
                            </div>
                        </div>
                        <div key={review.ID+"separator"} className="w-full border-b border-slate-500"></div>
                        <div className="text-justify top-0 w-full text-sm p-1 font-normal">
                            {review.Content}
                        </div>
                    </div>
                </div>
            )}
        </div>
        </>
    );
}