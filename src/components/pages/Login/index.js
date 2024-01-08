import React from 'react'

import gmail from '../../img/gmail.png'

import welcome from '../../img/welcome.jpg'

import style from './style.css'




const Login = () => {
  return (
    <>

<div className='.body'>
 <div class="night">
   <div class="shooting_star"></div>
   <div class="shooting_star"></div>
   <div class="shooting_star"></div>
   <div class="shooting_star"></div>
   <div class="shooting_star"></div>
   <div class="shooting_star"></div>
   <div class="shooting_star"></div>
   <div class="shooting_star"></div>
   <div class="shooting_star"></div>
   <div class="shooting_star"></div>
   <div class="shooting_star"></div>
   <div class="shooting_star"></div>
   <div class="shooting_star"></div>
   <div class="shooting_star"></div>
   <div class="shooting_star"></div>
   <div class="shooting_star"></div>
   <div class="shooting_star"></div>
   <div class="shooting_star"></div>
   <div class="shooting_star"></div>
   <div class="shooting_star"></div>
   </div>

    <section className="vh-100">
  <div className="container py-5 h-100">
    <div className="row d-flex align-items-center justify-content-center h-100">
      <div className="col-md-8 col-lg-7 col-xl-6 shadow p-3 bg-white rounded mb-md-3 mb-sm-4 mb-2 mt-2">
        <img
          src={welcome}
          className="img-fluid"
          alt="Phone image"
          
        />
      </div>
      <div className="col-md-7 col-lg-5 col-xl-5 offset-xl-1">
        <form>
          {/* Email input */}
          <div className="form-outline mb-4">
            <input
              type="email"
              id="form1Example13"
              className="form-control form-control-lg"
            />
            <label className="form-label" htmlFor="form1Example13">
              Email address
            </label>
          </div>
          {/* Password input */}
          <div className="form-outline mb-4">
            <input
              type="password"
              id="form1Example23"
              className="form-control form-control-lg"
            />
            <label className="form-label" htmlFor="form1Example23">
              Password
            </label>
          </div>
          <div className="d-flex justify-content-around align-items-center mb-4">
            {/* Checkbox */}
            <div className="form-check">
              <input
                className="form-check-input"
                type="checkbox"
                defaultValue=""
                id="form1Example3"
                defaultChecked=""
              />
              <label className="form-check-label" htmlFor="form1Example3">
                {" "}
                Remember me{" "}
              </label>
            </div>
            <a href="#!">Forgot password?</a>
          </div>
          {/* Submit button */}
          <button type="submit" className="btn btn-primary btn-lg btn-block">
            Sign in
          </button>
          <div className="divider d-flex align-items-center my-4">
            <p className="text-center fw-bold mx-3 mb-0 text-white" >OR</p>
          </div>

          <div
                 className="btn btn-primary btn-lg btn-block mt-2"
                 style={{ backgroundColor: "#C0C2C9" }}

                 role="button"
          >
          <img src={gmail}  style={{width:20,height:20}} >

          </img>
            <i className="ml-2 " />
            Continue with Email

          </div>
        </form>
      </div>
    </div>
  </div>
</section>


</div>

    </>
  )
}

export default Login
