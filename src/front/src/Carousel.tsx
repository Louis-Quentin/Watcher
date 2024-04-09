import React from 'react';
import Slider from 'react-slick';
import 'slick-carousel/slick/slick.css';
import 'slick-carousel/slick/slick-theme.css';
import { Watch } from './DataStructure.js';

const Carousel = ({ watchlist }: {watchlist:Watch[]}) => {
  const settings = {
    dots: true,
    infinite: true,
    speed: 500,
    slidesToShow: watchlist.length, // Show all slides
    slidesToScroll: 1
  };

  return (
    <Slider {...settings}>
      {watchlist.map((watch, index) => (
        <div key={index}>
          {/* Render your watch element here */}
          <h3>{watch.Name}</h3>
          <h4>{watch.Brand}</h4>
          <h5>{watch.Price}</h5>
          <h5>{watch.Range}</h5>
          {/* Add other watch information as needed */}
        </div>
      ))}
    </Slider>
  );
};

export default Carousel;
