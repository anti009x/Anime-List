import React from 'react'
import anime from '../../img/anime.jpg'
const Moviesdetail = () => {
  return (  


    <>              
               
                        <div className=" d-flex justify-content-center card text-center" >
                        <div className="card-header  " style={{fontSize:20,fontWeight:'bold'}}>
                            Date Alive
                        </div>
                        <div className="card-body  card bg-dark text-white  ">
              
                           <div className='mb-3'>
                            <img src={anime} alt='anime' style={{width:'40%'}}></img>
                            <p className="card-text">
                            Padp 30 tahun yang lalu terjadi sebuah fenomena aneh bernama "Spacequake" yang menghancurkan pusat Eurasia. Dari Peristiwa itu tercatat 150 juta orang menjadi korban dan sampai sekarang, masih terjadi Spacequake itu namun dalam skala yang lebih kecil dari pada yang awal secara random diberbagai tempat di dunia. Itsuka Shidou, seorang anak SMA biasa pada saat terjadi Spacequake bertemu dengan seorang gadis misterius yang tidak memiliki nama yang berpakain gaun dengan armor dan memegang pedang ditangannya.
                            </p>
                            </div>

                    <div className='  d-flex justify-content-center '>  
                    <div className="list-group">
                    <p className="list-group-item list-group-item-action active">
                        Infromasi
                    </p>
                    <p className="list-group-item list-group-item-action">Rating : 7.86</p>
                    <p className="list-group-item list-group-item-action">Genre : Romance,Action,Drama</p>
                    <p className="list-group-item list-group-item-action">Tahun Liris : 2014</p>
                    </div>
                            </div>
                            <p className="btn btn-primary">Link Downlaod</p>
                        </div>
                        <div className="card-footer text-muted">
                            Link Masih Tersedia
                        </div>
                        </div>

           
              

    </>
  )
}

export default Moviesdetail
