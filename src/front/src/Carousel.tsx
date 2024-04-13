import React, { useState, useEffect } from 'react';
import Slider from 'react-slick';
import 'slick-carousel/slick/slick.css';
import 'slick-carousel/slick/slick-theme.css';
import { Watch } from './DataStructure.js';
import styles from './Carousel.module.css';
import HomeStyles from './css/HomeLanding.module.css'

const Carousel: React.FC<{ watchlist: Watch[] }> = ({ watchlist }) => {
  const [slidesToShow, setSlidesToShow] = useState<number>(3);
  const hostname = "http://localhost:3000/";

  useEffect(() => {
    const handleResize = () => {
      // Adjust the number of slides to show based on screen width
      const newSlidesToShow = window.innerWidth < 1100 ? 1 : 3;
      setSlidesToShow(newSlidesToShow);
    };

    // Initial call to set the number of slides to show
    handleResize();

    // Add event listener for window resize
    window.addEventListener('resize', handleResize);

    // Cleanup function to remove event listener
    return () => {
      window.removeEventListener('resize', handleResize);
    };
  }, []);

  const settings = {
    dots: false,
    infinite: true,
    speed: 9000,
    slidesToShow: slidesToShow, // Set the number of slides to show dynamically
    slidesToScroll: 1, // Change to scroll one slide at a time
    autoplay: true,
    autoplaySpeed: 0,
    cssEase: 'linear',
    pauseOnHover: false,
    pauseOnFocus: false,
  };

  return (
    <Slider {...settings}>
      {watchlist.map((watch, index) => (
        <div key={index} className={styles.slide}>
         <div className={styles.slideContent} style={{ backgroundImage: `url(${watch.Url})` }}>
            <div className={styles.slideHeader}>
              <div className={`${styles.Name} ${styles.Tag} ${HomeStyles.ColorDa}`}>{watch.Name}</div>
              <div className={`${styles.Price} ${styles.Tag} ${HomeStyles.ColorDa2}`}>{watch.Price} $</div>
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
