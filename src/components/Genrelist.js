import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';

const Genrelist = () => {
  const [genrelist, setGenrelist] = useState([]);

  useEffect(() => {
    // Fetch genre list data from an API or perform any other asynchronous operation here
    // For now, I'll set the genre list statically
    setGenrelist([
      { id: 1, genre_name: 'Action' },
      { id: 2, genre_name: 'Comedy' },
      { id: 3, genre_name: 'Fantasy' },
      { id: 4, genre_name: 'Romance' },
      

    ]);

  }, []);

  return (
    <>
      <div className='row d-flex justify-content-center  ' >
        {genrelist.map((genre, index) => (
          <div className='col-4 mb-3 ' key={index}>
            <div className='card ' >
              {/* <img className="card-img-top" src="..." alt="Card image cap" /> */}
              <div className='card-body '  >
                <h5 className='card-title d-flex justify-content-center '>{genre.genre_name}</h5>
              </div>
              <div className='card-footer d-flex justify-content-center'>
                <a href="#" className="btn btn-primary">
                <Link to={`/genres/${genre.id}`} style={{textDecoration : 'none',color:'white'}}>click here !</Link>
                </a>
              </div>
            </div>
          </div>
        ))}
      </div>
    </>
  );
};

export default Genrelist;
