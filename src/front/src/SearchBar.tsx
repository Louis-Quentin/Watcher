import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import styles from './css/SearchBar.module.css';

import store1 from './css/img/store_1.jpg';
import store2 from './css/img/store_2.jpg';
import store3 from './css/img/store_3.jpg';

const images = [store1, store2, store3];

const SearchBar: React.FC = () => {
  const [value, setValue] = useState('');
  const [suggestions, setSuggestions] = useState([]);
  const [currentImageIndex, setCurrentImageIndex] = useState(0);
  const [rangeValue, setRangeValue] = useState(50); // Initialize range value
  const navigate = useNavigate();

  const navigateToContacts = () => {
    navigate('/search');
  };

  useEffect(() => {
    const fetchData = async () => {
      try {
        const { data } = await axios.get(
          `https://dummyjson.com/products/search?q=${value}`
        );
        setSuggestions(data.products);
      } catch (error) {
        console.error(error);
      }
    };

    if (value) {
      fetchData();
    } else {
      setSuggestions([]);
    }
  }, [value]);

  useEffect(() => {
    const intervalId = setInterval(() => {
      setCurrentImageIndex((prevIndex) => (prevIndex + 1) % images.length);
    }, 5000); // Change image every 5 seconds

    return () => clearInterval(intervalId);
  }, []);

  return (
    <div className={styles.container}>
      {images.map((image, index) => (
        <div
          key={index}
          className={`${styles.backgroundImage} ${index === currentImageIndex ? styles.visible : ''}`}
          style={{ backgroundImage: `url(${image})` }}
        />
      ))}
      <div className={styles.searchbarContainer}>
        <div className={styles.searchbar}>
        <input
            type="text"
            className={styles.textbox}
            placeholder="Find your store.."
            value={value}
            onChange={(e) => {
              setValue(e.target.value);
            }}
          />
          <div className={styles.rangeContainer}>
            <label htmlFor="rangeInput">{rangeValue} Km</label>
            <input
              type="range"
              id="rangeInput"
              className={styles.rangeSlider}
              value={rangeValue}
              min="0"
              max="100"
              onChange={(e) => setRangeValue(Number(e.target.value))}
            />
          </div>
          
          <button className={styles.button} onClick={navigateToContacts}>
            Search
          </button>
        </div>
      </div>
    </div>
  );
};

export default SearchBar;
