// src/Carousel.tsx
import React, { useState, useEffect } from 'react';
import styles from './css/BasicCarousel.module.css';
import 'slick-carousel/slick/slick.css'; 
import 'slick-carousel/slick/slick-theme.css';
import Slider from 'react-slick';
import { CustomArrowProps } from 'react-slick';
import { NextArrow, PrevArrow } from './CustomArrows';


// Import images
import watcherLogo from './css/img/watcher_logo.png';
import SearchBar from './SearchBar';

interface CarouselProps {
  images: string[];
  interval?: number;
}

const BasicCarousel: React.FC<CarouselProps> = ({ images, interval = 3000 }) => {
  const [currentIndex, setCurrentIndex] = useState(0);
  const [slidesToShow, setSlidesToShow] = useState<number>(1);
  const [placeholder, setPlaceholder] = useState<string>("Search..");

  const showSlide = (index: number) => {
    const totalSlides = images.length;
    if (index >= totalSlides) {
      setCurrentIndex(0);
    } else if (index < 0) {
      setCurrentIndex(totalSlides - 1);
    } else {
      setCurrentIndex(index);
    }
  };

  const nextSlide = () => {
    showSlide(currentIndex + 1);
  };

  const prevSlide = () => {
    showSlide(currentIndex - 1);
  };

  useEffect(() => {
    const slideInterval = setInterval(nextSlide, interval);
    return () => clearInterval(slideInterval);
  }, [currentIndex, interval]);

  const settings = {
    dots: true,
    infinite: true,
    speed: 1000,
    slidesToShow: slidesToShow, // Set the number of slides to show dynamically
    slidesToScroll: 1, // Change to scroll one slide at a time
    autoplay: false,
    autoplaySpeed: 0,
    cssEase: 'linear',
    pauseOnHover: false,
    pauseOnFocus: false,
    nextArrow: <NextArrow className={styles.slickArrow} />,
    prevArrow: <PrevArrow className={styles.slickArrow} />, 
  };

  return (
    <div>
      <Slider {...settings}>
        {images.map((image, index) => (
          
          <div key={index} className={styles.slide}>
            <div
              className={styles.slideContent}
              style={{ backgroundImage: `url(${image})` }}>
                <SearchBar/>
            </div>
          </div>
        ))}
      </Slider>
    </div>
  );
};

export default BasicCarousel;