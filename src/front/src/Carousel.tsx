import React from 'react';
import Slider from 'react-slick';
import 'slick-carousel/slick/slick.css';
import 'slick-carousel/slick/slick-theme.css';
import { Watch } from './DataStructure.js';
import styles from './Carousel.module.css';
import { hostname } from 'os';

const Carousel: React.FC<{ watchlist: Watch[] }> = ({ watchlist }) => {
  const hostname = "http://localhost:3000/";
  const settings = {
    dots: false,
    infinite: true,
    speed: 9000,
    slidesToShow: 3,
    slidesToScroll: 1, // Change to scroll one slide at a time
    autoplay: true,
    autoplaySpeed: 0,
    cssEase: 'linear',
    pauseOnHover: true,
    pauseOnFocus: true,
  };

  watchlist.forEach((element) => {
    console.log(element.Url);
  })
  return (
    <Slider {...settings}>
      {watchlist.map((watch, index) => (
        <div key={index} className={styles.slide}>
         <div className={styles.slideContent} style={{ backgroundImage: `url(${watch.Url})` }}>
            <div className={styles.slideHeader}>
              <div className={`${styles.Name} ${styles.Tag}`}>{watch.Name}</div>
              <div className={`${styles.Price} ${styles.Tag}`}>{watch.Price} $</div>
              <div className={`${styles.Range} ${styles.Tag}`}>{watch.Range} km</div>
            </div>
            <div className={styles.slideBottom}>
              <div className={`${styles.Brand} ${styles.Tag}`}>{watch.Brand}</div>
              <div className={`${styles.Available} ${styles.Tag}`} style={{ backgroundImage: watch.Available ? `url('${hostname}images/in-stock.png')` : `url('${hostname}images/out-of-stock.png')` }}>
            </div>
          </div>
        </div>
      </div>
      ))}
    </Slider>
  );
};

export default Carousel;
