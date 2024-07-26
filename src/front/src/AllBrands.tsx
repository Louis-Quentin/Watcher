// src/AllBrands.tsx
import React, { useRef } from 'react';
import Slider from 'react-slick';
import styles from './css/AllBrands.module.css';
import 'slick-carousel/slick/slick.css';
import 'slick-carousel/slick/slick-theme.css';


import image1 from './css/img/logo_rolex.png';
import image2 from './css/img/logo_AP.jpg';
import image3 from './css/img/logo_jc.jpg';
import image4 from './css/img/logo_omega.jpg';
import image5 from './css/img/logo_rm1.webp';
import image6 from './css/img/logo_pp.svg';

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

const AllBrands: React.FC = () => {
  const sliderRef = useRef<Slider>(null); 

  const sliderSettings = {
    dots: true,
    infinite: true,
    speed: 500,
    slidesToShow: 2,
    slidesToScroll: 1,
    nextArrow: <NextArrow />,
    prevArrow: <PrevArrow />
  };

  const images = [
    {
      src: image1,
      description: 'Rolex',
      price: 'Since 1905',
      link: '/rolex'
    },
    {
      src: image2,
      description: 'Audemars Piguet',
      price: 'Since 1875',
      link: '/ap'
    },
    {
      src: image3,
      description: 'Jacob & Co',
      price: 'Since 1986',
      link: '/jacobandco'
    },
    {
      src: image4,
      description: 'Omega',
      price: 'Since 1894',
      link: '/omega'
    },
    {
      src: image5,
      description: 'Richard Mille',
      price: 'Since 2001',
      link: '/richard_mille'
    },
    {
      src: image6,
      description: 'Patek Philippe',
      price: 'Since 1839',
      link: '/patek'
    }
  ];

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <div className={styles.profileContainer}>
          {/* You can add profile picture and details here */}
        </div>
        <h1 className={styles.title}>All Brands</h1>
        <a href="/favorites" className={styles.viewAllLink}>View All</a>
      </div>
      <div className={styles.sliderContainer}>
        <Slider ref={sliderRef} {...sliderSettings}>
          {images.map((image, index) => (
            <a href={image.link} key={index} className={styles.slide}>
              <div className={styles.imageContainer}>
                <img src={image.src} alt={`Slide ${index}`} className={styles.sliderImage} />
                <div className={styles.descriptionContainer}>
                  <p className={styles.description}>{image.description}</p>
                  <p className={styles.price}>{image.price}</p>
                </div>
              </div>
            </a>
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

export default AllBrands;