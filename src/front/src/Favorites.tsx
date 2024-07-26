// src/Favorites.tsx
import React, { useRef } from 'react';
import Slider from 'react-slick';
import styles from './css/Favorites.module.css';
import 'slick-carousel/slick/slick.css';
import 'slick-carousel/slick/slick-theme.css';

import image1 from './css/img/rolex_gmt.jpg';
import image2 from './css/img/rolex_hulk.webp';
import image3 from './css/img/rolex_pepsi.jpg';
import image4 from './css/img/rolex_smurf.webp';
import image5 from './css/img/rolex_coke.jpg';
import image6 from './css/img/rolex_starbucks.jpeg';
import profilePic from './css/img/profile.jpg';

const NextArrow = (props: any) => {
  const { className, onClick } = props;
  return (
    <div
      className={`${className} ${styles.arrowButton} ${styles.nextArrow}`}
      onClick={onClick}
    >
      &gt;
    </div>
  );
};

const PrevArrow = (props: any) => {
  const { className, onClick } = props;
  return (
    <div
      className={`${className} ${styles.arrowButton} ${styles.prevArrow}`}
      onClick={onClick}
    >
      &lt;
    </div>
  );
};

const Favorites: React.FC = () => {
  const sliderRef = useRef<Slider>(null);

  const sliderSettings = {
    dots: true,
    infinite: true,
    speed: 500,
    slidesToShow: 4,
    slidesToScroll: 1,
    nextArrow: <NextArrow />,
    prevArrow: <PrevArrow />
  };

  const images = [
    {
      src: image1,
      description: 'Rolex GMT Master 2 "Bruce Wayne"',
      price: '$14 400'
    },
    {
      src: image2,
      description: 'Rolex Submariner Date "Hulk"',
      price: '$21 000'
    },
    {
      src: image3,
      description: 'Rolex GMT Master 2 "Pepsi"',
      price: '$25 700'
    },
    {
      src: image4,
      description: 'Rolex Submariner Date "Smurf"',
      price: '$35 000'
    },
    {
      src: image5,
      description: 'Rolex GMT Master 2 "Coke"',
      price: '$13 150'
    },
    {
      src: image6,
      description: 'Rolex Submariner Date "Starbucks"',
      price: '$15 600'
    }
  ];

  const handleArrowClick = () => {
    if (sliderRef.current) {
      sliderRef.current.slickNext();
    }
  };

  return (
    <div className={styles.container}>
      <div className={styles.profileContainer}>
        <img src={profilePic} alt="Profile" className={styles.profilePic} />
        <div className={styles.profileInfo}>
          <h2 className={styles.profileName}>Redouane El Shahawi</h2>
          <p className={styles.profileDate}>Member since July 26, 2024</p>
        </div>
      </div>
      <div className={styles.header}>
        <h1 className={styles.title}>Favoris</h1>
        <a href="/favorites" className={styles.viewAllLink}>View All</a>
      </div>
      <div className={styles.sliderContainer}>
        <Slider ref={sliderRef} {...sliderSettings}>
          {images.map((image, index) => (
            <div key={index} className={styles.slide}>
              <div className={styles.imageContainer}>
                <img src={image.src} alt={`Slide ${index}`} className={styles.sliderImage} />
              </div>
              <p className={styles.description}>{image.description}</p>
              <p className={styles.price}>{image.price}</p>
            </div>
          ))}
        </Slider>
        <button className={`${styles.arrowButton} ${styles.prevArrow}`} onClick={() => sliderRef.current?.slickPrev()}>
          &lt;
        </button>
        <button className={`${styles.arrowButton} ${styles.nextArrow}`} onClick={() => sliderRef.current?.slickNext()}>
          &gt;
        </button>
      </div>
    </div>
  );
};

export default Favorites;