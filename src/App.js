import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import logo from './logo192.png';
import './App.css';
import Home from './components/pages/Home';
import Genres from './components/pages/Genres';
import Movielist from './components/movielist';
import Admin from './components/pages/Admin';
import Moviesdetail from './components/pages/Movies';
import Login from './components/pages/Login';
import Menu from './components/menu';

// Layout component with common header and main content
const Layout = ({ children }) => (
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
      <p className='text-light bg-dark w-100 h-100' style={{ padding: '10px' }}>
        <Link to='/movie' style={{ textDecoration: 'none', color: 'yellow', marginTop: 2, marginRight: 3 }}>
          Daftar Semua Anime
        </Link>{' '}
        ||
        <Link to='/menu' style={{ textDecoration: 'none', color: 'red', marginRight: 3, marginTop: 2, marginLeft: 4 }}>
          MENU
        </Link>{' '}
        ||
        <Link to='/' style={{ textDecoration: 'none', color: 'yellow', marginRight: 3, marginTop: 2, marginLeft: 4 }}>
          HOME
        </Link>{' '}
        ||
        <div className='d-flex justify-content-end'>
          <Link to='/login' style={{ textDecoration: 'none', color: 'yellow' }}>
            LOGIN
          </Link>
        </div>
      </p>

      {children}
    </div>
  </>
);

function App() {
  return (
    <Router>
      <Routes>
        <Route
          path='/login'
          element={
            <Login />
          }
        />
        <Route
          path='/*'
          element={
            <Layout>
              {/* Your Routes */}
              <Routes>
                <Route path='/menu' element={<Menu />} />
                <Route path='/moviedetail' element={<Moviesdetail />} />
                <Route exact path='/movie/:id' element={<Moviesdetail />} />
                <Route path='/admin' element={<Admin />} />
                <Route path='/movie' element={<Movielist />} />
                <Route path='/' element={<Home />} />
                <Route path='/genres' element={<Genres />} />
              </Routes>
            </Layout>
          }
        />
      </Routes>
    </Router>
  );
}

export default App;