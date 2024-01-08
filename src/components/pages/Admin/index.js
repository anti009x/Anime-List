import React from 'react'

import { FcFullTrash } from "react-icons/fc";

import { FcSupport } from "react-icons/fc";
import { FcOk } from "react-icons/fc";
import { FcAddDatabase } from "react-icons/fc";
import { FcDeleteDatabase } from "react-icons/fc";
function Admin() {
  return (
    <>
   
    <div className='d-flex justify-content-end mb-2 bg-transparent mr-3'>
   
    <button type="button" className="btn  shadow-lg p-3  bg-white rounded">
    <FcAddDatabase 

    size={30}
    
    />
    </button>
    <button type="button" className="btn  ml-2 shadow p-3  bg-white rounded">
    <FcDeleteDatabase 
      size={30}
    />

    </button>

    </div>
   <table className="table shadow p-3 mb-5 bg-white rounded  ">
  <thead className='thead-dark'>
    <tr>
      <th className="col-2">#</th>
      <th className="col-2">Judul</th>
      <th className="col-6 text-center">Informasi</th>
      <th className="col-2">Rating</th>
      <th className="col-2">Genre</th>
      <th className="col-2">Rilis</th>
      <th className="col-2">Link Download</th>
      <th className="col-4">Action</th>

      <th className='col-2'>Status</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th scope="row">1</th>
      <td>Date Alive</td>
      <td >Padp 30 tahun yang lalu terjadi sebuah fenomena aneh bernama "Spacequake" yang menghancurkan pusat Eurasia. Dari Peristiwa itu tercatat 150 juta orang menjadi korban dan sampai sekarang, masih terjadi Spacequake itu namun dalam skala yang lebih kecil dari pada yang awal secara random diberbagai tempat di dunia. Itsuka Shidou, seorang anak SMA biasa pada saat terjadi Spacequake bertemu dengan seorang gadis misterius yang tidak memiliki nama yang berpakain gaun dengan armor dan memegang pedang ditangannya.</td>
      <td>7.86</td>
      <td>Romance,
        Action,Drama</td>
      <td>2014</td>
      <td>drive.google.com</td>
      <td >
      <FcFullTrash
      
      size={40}

      

      />
     <FcSupport 

      size={35}
     />
      </td>

      <td >
      <div className="form-check d-flex justify-content-center">
  <input
    className="form-check-input"
    type="radio"

  />
</div>


      </td>


    </tr>

  </tbody>
</table>
   </>
  )
}

export default Admin
