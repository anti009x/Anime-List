import React, { useEffect, useState } from 'react'
import { Link } from 'react-router-dom';

function Movielist() {
  const [movie, setmovie] = useState([]);
  useEffect(() => {

    setmovie([
      {id:1, title:'Batman', runtime:145},
      {id:2, title:'superman',runtime:145},
      {id:3, title:'iron-man',runtime:145},
      
   
    ]);
    
  }, [])
  
  return (
    <>

      <div className='row d-flex justify-content-center'>
              {movie.map((movie,index) => (
                <div className='col-4 mb-3' key={index}>
                  <div className='card'>
                    {/* <img className="card-img-top" src="..." alt="Card image cap" /> */}
                    <div className='card-body'>
                      <h5 className='card-title d-flex justify-content-center'>{movie.title}</h5>
                      <p className='card-text'>
                      sdkjfhsdkfhsdkjhfksdfdsfgsdfdsfsdsdfgsdfgfdsdfsdf
                      </p>
                    </div>
                    <div className='card-footer d-flex justify-content-center'>
                    <a href="#" className="btn btn-primary">
                <Link to={`/movie/${movie.id}`} style={{textDecoration : 'none',color:'white'}}>Check Detail!</Link>
                </a>
                    </div>
                  </div>
                </div>
              ))}
            </div>
 
            </>
  )
}

export default Movielist