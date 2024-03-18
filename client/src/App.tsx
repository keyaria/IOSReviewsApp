import { useEffect, useState } from 'react'
import './App.css'

type Review = {
  id: string,
  Name: string,
  Updated: Date,
  Rating: string, 
  Title: string,
  Content: string
}

function App() {
  const [data, setData] = useState<Review[] | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(()=> {
      const fetchAllReviews = async () => {
        try {
          const response = await fetch(
            'http://localhost:8888/api/reviews/?time=148',
            
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
      {loading && (
          <div className="text-xl font-medium">Loading Reviews...</div>
        )}
        {error && <div className="text-red-700">{error}</div>}
        <ul>
          {data  &&
            data.map((response: any) => (
              <li
                key={response.id}
                className="border-b border-gray-100 text-sm sm:text-base"
              >
                <div
                  className='inner'
                >
                  <h3>Rating: {response.Rating}</h3>
                  <h5>Author{response.Author}</h5>
                  <h5>{response.Updated}</h5>
                  <p>{response.Title}</p>
                 
                  
                  {response.Content}
                </div>
              </li>
            ))}
        </ul>
       </div>
    </>
  )
}

export default App
