import React from 'react';
import { CustomArrowProps } from 'react-slick';
import styles from './css/BasicCarousel.module.css';


const NextArrow: React.FC<CustomArrowProps> = (props) => {
  const { className, style, onClick } = props;
  return (
    <div
      className={`${className}`}
      style={{ ...style, display: 'block', right: '10px', zIndex: 1 }}
      onClick={onClick}
    >
      &#10095;
    </div>
  );
};

const PrevArrow: React.FC<CustomArrowProps> = (props) => {
  const { className, style, onClick } = props;
  return (
    <div
      className={`${className}`}
      style={{ ...style, display: 'block', left: '10px', zIndex: 1 }}
      onClick={onClick}
    >
      &#10094;
    </div>
  );
};

export { NextArrow, PrevArrow };
