import { useEffect, useState } from 'react'
import '../App.css'

type Review = {
  id: string,
  Name: string,
  Updated: Date,
  Rating: string, 
  Title: string,
  Content: string
}


// Represents Reviews Service on FE
function ReviewsInterface() {
  const apiUrl = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';
  const [data, setData] = useState<Review[] | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(()=> {
      const fetchAllReviews = async () => {
        try {
          const response = await fetch(
            `${apiUrl}/api/reviews/?time=64`,
            
          );
          if (!response.ok) {
            throw new Error(`HTTP error: Status ${response.status}`);
          }

          let reviewsData = await response.json();
          setData(reviewsData.data);
          setError(null);

        }catch (err: any){

          setError(err.message);
        setData(null);
        }finally {
          setLoading(false);
        }
      }

      fetchAllReviews()
  }, [])

  // TODO: Add Pagination when there is more than 10 reviews per Page

  return (
    <>
      <div>
        <div className="topnav" id="myTopnav">
          <a>Home</a>
        </div>
        <div >
        <div className="AppDiv">
          <img src="/sims-logo.png" />
          <h3>The Sims FreePlay Reviews</h3>
          <p>Customize your virtual town & story</p>
          <p>Electronic Arts</p>
        </div>
      {loading && (
          <div className="text-xl font-medium">Loading Reviews...</div>
        )}
        {error && <div className="text-red-700">{error}</div>}
        <ul>
          {data  &&
            data.map((response: Review) => (
              <li
                key={response.id}
                className="border-b border-gray-100 text-sm sm:text-base"
              >
                <div
                  className='inner'
                >
                  <h3>Rating: {response.Rating}</h3>
                  <h5>Author: {response.Name}</h5>
                  <h5>{new Date(response.Updated).toLocaleString()}</h5>
                  <p>{response.Title}</p>
                 
                  
                  {response.Content}
                </div>
              </li>
            ))}
        </ul>
        </div>
        
       </div>
    </>
  )
}

export default ReviewsInterface
