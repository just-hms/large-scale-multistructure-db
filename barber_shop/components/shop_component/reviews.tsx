import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faArrowUp, faArrowDown, faStar} from "@fortawesome/free-solid-svg-icons";
export default function Reviews(reviews) {
    return (
        <>
            <div>
            {reviews.children.map((review)=>
                    <div key={review.id} className="w-full text-slate-200">
                        <div key={review.id+"container"} className="flex flex-col items-center justify-start w-full bg-slate-700 rounded-lg p-3 shadow-md shadow-black/30 mb-3">
                            <div key={review.id+"name"} className="text-md w-full text-left font-normal">{review.name}</div>
                            <div key={review.id+"title"} className="text-xl mb-3 font-bold flex items-center justify-between w-full text-left">
                                <div className="flex items-center">
                                    <p className="pr-2">{review.title}</p>
                                    {[...Array(review.vote)].map(_=><FontAwesomeIcon key={review.id+"arrowUpIcon"} icon={faStar} className="text-sm"/>)}
                                </div>
                                <div className="flex items-center">
                                    <button key={review.id+"arrowUp"} className="hover: text-white mr-3 text-sm">
                                        <FontAwesomeIcon key={review.id+"arrowUpIcon"} icon={faArrowUp}/>
                                    </button>
                                    <button key={review.id+"arrowDown"} className="hover: text-white mr-3 text-sm">
                                        <FontAwesomeIcon key={review.id+"arrowDownIcon"}  icon={faArrowDown}/>
                                    </button>
                                    <div className={`text-sm ${review.upvotes > 0 ? "text-green-700" : "text-rose-600"}`}> { (review.upvotes > 0) ? "+": ""} {review.upvotes}</div>
                                </div>
                            </div>
                            <div key={review.id+"separator"} className="w-full border-b border-slate-500"></div>
                            <div className="text-justify top-0 w-full text-sm p-1 font-normal">
                                {review.review}
                            </div>
                        </div>
                    </div>
                )}
            </div>
        </>
    );
}