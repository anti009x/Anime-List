/*bawaan react*/
import React from 'react';

/*bagian asset*/
import logo from './logo192.png';

/*bagian components*/
import Menu from './components/menu';
import Movielist from './components/movielist';
/*bagian style*/
import './App.css'




function App() {
  return (
    <>
      {/* Header */}
      <div className='container mb-3'>
        <div className='navbar navbar-light bg-dark w-100 h-100 d-inline-block'>
          <img src={logo} alt='anime jpg' height={100} width={100} />
          <span className='navbar-brand text-light h1'>OtakuNime</span>
        </div>
      </div>

      {/* Main Content */}
      <div className='container'>
        <p className='d-flex justify-content-end text-light bg-dark w-100 h-100' style={{ padding: '10px' }}>
          Daftar Semua Anime
        </p>

        <div className='d-flex justify-content-start'>
          {/* Navigation Menu */}
          <Menu></Menu>

          
          <div className='col-10 d-flex justify-content-start '>
            {/* Anime Cards */}
            <Movielist></Movielist>
          </div>
          
        </div>
      </div>
    </>
  );
}

export default App;
