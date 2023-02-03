
export default function ReportedReviews({reported_reviews}) {
    console.log(reported_reviews.reported_reviews)
    return (
    <div className='flex flex-col items-start justify-start text-left text-slate-300 text-lg w-full'>
        <div className="w-full h-full bg-slate-800 ">
            <div className="w-full flex justify-center items-start">
                <div className='relative w-3/4 max-h-96 overflow-auto rounded-3xl shadow-md shadow-black/70'>
                    <div className="sticky top-0 bg-slate-700 w-full flex flex-col items-center justify-center border-b border-slate-600  px-5 pt-0">
                        <h1 className="text-xl text-center font-bold leading-tight tracking-tight text-slate-200 sticky top-0 w-full pt-3 pb-3">
                            Reported Reviews
                        </h1>
                    </div>
                    <div>
                    {reported_reviews.map((review)=>
                            <div key={review.id} className="w-full text-slate-200 p-3 pb-0">
                                <div key={review.id+"container"} className="flex flex-col items-start justify-start w-full bg-slate-700 rounded-3xl p-3 shadow-md shadow-black/30 mb-3">
                                    <div key={review.id+"name"} className="text-md w-full text-left font-normal">{review.name}</div>
                                    <div key={review.id+"title"} className="text-xl flex items-center justify-between w-full text-left">
                                        <div className="flex items-center">
                                            <p className="pr-2 font-bold text-md">{review.title}</p>
                                            <p className="pr-2 text-md">- {review.shop}</p>
                                        </div>
                                        <div className="flex items-center">
                                        <div className={`text-sm ${review.upvotes > 0 ? "text-green-700" : "text-rose-600"}`}> { (review.upvotes > 0) ? "+": ""} {review.upvotes}</div>
                                        </div>
                                    </div>
                                    <div key={review.id+"separator"} className="w-full border-b border-slate-500"></div>
                                    <div className="text-justify top-0 w-full text-sm p-1 font-normal">
                                        {review.review}
                                    </div>
                                    <div className="flex items-center justify-between w-full">
                                        <button className="px-6 py-2.5 m-1 bg-green-900 bg-opacity-70 text-slate-300 text-xs rounded-full focus:bg-green-800 hover:bg-green-800 focus:outline-none transition duration-150 ease-in-out " type="button" id="search_button">
                                            Remove from reported reviews
                                        </button>
                                        <button className="px-6 py-2.5 m-1 bg-rose-900 bg-opacity-70 text-slate-300 text-xs rounded-full focus:bg-red-800 hover:bg-red-800 focus:outline-none transition duration-150 ease-in-out " type="button" id="search_button">
                                            Delete Review
                                        </button>
                                    </div>
                                </div>
                            </div>
                        )}
                    </div>
                </div>
            </div>
        </div>
    </div>
    )
  }
  
  // TODO: For each account we gotta return the mail and the actions, i.e: delete account, change permissions, on click over button show results if any