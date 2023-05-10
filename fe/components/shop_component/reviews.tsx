import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faArrowUp, faArrowDown, faStar, faTrash} from "@fortawesome/free-solid-svg-icons";
import { deleteVote, submitVote, deleteReview } from "../../lib/shops";

export default function Reviews({reviews,shopid, userid}:any) {
    // console.log(reviews)
    // console.log(userid)
    const handleVote = async (review:any,vote:boolean) =>{
        if(review.UpVotes.includes(userid) || review.DownVotes.includes(userid)){
            await deleteVote(shopid,review.ID)
        }else{
            await submitVote(shopid, review.ID, vote)
        }
        window.location.reload()
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
                                {/* <p className="pr-2">{review.title}</p> */}
                                {[...Array(review.Rating)].map((_,index)=><FontAwesomeIcon key={index} icon={faStar} className="text-sm"/>)}
                            </div>
                            <div className="flex items-center">
                                <button key={review.ID+"arrowUp"} className="hover: text-white mr-3 text-sm" onClick={(e)=>handleVote(review,true)}>
                                    {<FontAwesomeIcon key={review.ID+"arrowUpIcon"} icon={faArrowUp} className={(review.UpVotes.includes(userid))?"text-orange-500":""}/>}
                                </button>
                                <button key={review.ID+"arrowDown"} className="hover: text-white mr-3 text-sm" onClick={(e)=>handleVote(review,false)}>
                                    <FontAwesomeIcon key={review.ID+"arrowDownIcon"}  icon={faArrowDown} className={(review.DownVotes.includes(userid))?"text-orange-500":""}/>
                                </button>
                                {(review.UserID === userid)?<button key={review.ID+"trashbin"} className="hover: text-white mr-3 text-sm" onClick={(e)=>deleteReview(shopid, review.ID)}>
                                    <FontAwesomeIcon key={review.ID+"trashbin"}  icon={faTrash}/>
                                </button>:<></>}
                                <div className={`text-sm ${review.UpVotes.length >= 0 ? "text-green-700" : "text-rose-600"}`}> {((review.UpVotes.length - review.DownVotes.length) >= 0) ? "+": "-"}{review.UpVotes.length - review.DownVotes.length}</div>
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