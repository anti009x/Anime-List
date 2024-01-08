import React from 'react'
import { Link } from 'react-router-dom'

const Menu = () => {
  return (
  <>

 
     <div className='position-fixed 'style={{ marginLeft: '20rem' }}>
      <h1>WELCOME MY YOUTUBE CHANNEL</h1>
    </div>
    <div className='col-2' >
            <div className='card' style={{ width: '12rem' }}>
              <div className='card-header'>
                <Link to='/' style={{textDecoration: 'none' , color:'black'}}>HOME</Link>
              </div>
              <ul className='list-group list-group-flush'>
                <li className='list-group-item'>
                  <Link to='/movie' style={{textDecoration: 'none', color:'black'}}>MOVIE</Link>
                </li>
                <li className='list-group-item'>
                  <Link to='/genres' style={{textDecoration: 'none', color:'black'}}>GENRES</Link>
                </li>
                <li className='list-group-item'>
                  <Link to='/admin' style={{textDecoration: 'none', color:'black'}}>ADMIN</Link>
                </li>
              </ul>
            </div>
          </div>
        

  </>
  )
}

export default Menu
