/*bawaan react*/
import React from 'react';
import { BrowserRouter as Router,Routes,Route, Link } from 'react-router-dom';

/*bagian asset*/
import logo from './logo192.png';

/*bagian components*/
import Menu from './components/menu';
/*bagian style*/
import './App.css'
import Home from './components/pages/Home';
import Genres from './components/pages/Genres';
import Movielist from './components/movielist';
import Admin from './components/pages/Admin';





function App() {
  return (
    <Router>
      {/* Header */}
      <div className='container mb-3'>
        <div className='navbar navbar-light bg-dark w-100 h-100 d-inline-block'>
          <img src={logo} alt='anime jpg' height={100} width={100} />
          <span className='navbar-brand text-light h1'>OtakuNime</span>
        </div>
      </div>

      {/* Main Content */}
      <div className='container'>
        <p className='d-flex justify-content-start text-light bg-dark w-100 h-100' style={{ padding: '10px' }}>
        <Link to='/movie' style={{textDecoration: 'none', color:'yellow',marginTop:2,marginRight:3}}>Daftar Semua Anime</Link> ||
          <Link to='/menu' style={{textDecoration: 'none', color:'red',marginLeft:4,marginTop:2}}>MENU</Link>
          
        </p>

        <div className='d-flex justify-content-center mr-4'>
          {/* Navigation Menu */}
          
      {/* <Menu></Menu> */}
          
          <div className='col-10 d-flex justify-content-center mr-4'>
            {/* Anime Cards */}
            <Routes>
              <Route path='/' element={<Home/>}/>
              <Route path='/genres' element={<Genres/>}/>
              <Route path='/movie' element={<Movielist/>}/>
              <Route path='/admin' element={<Admin/>}/>
            </Routes>
          </div>
          
        </div>
        <div className='col-4 d-flex justify-content-start '>
      <Routes>
              <Route path='/menu' element={<Menu/>}/>
            </Routes>
          
      </div>
      </div>
    
    </Router>
    
  );
}

export default App;
