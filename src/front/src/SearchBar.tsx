import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import styles from './css/SearchBar.module.css';

function SearchBar() {
  const [value, setValue] = useState('');
  const [suggestions, setSuggestions] = useState([]);
  const navigate = useNavigate(); // This hook should be called inside the component

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
        console.log(error);
      }
    };

    fetchData();
  }, [value]);

  return (
    <div className={styles.container}>
      <input
        type="text"
        className={styles.textbox}
        placeholder="Search data..."
        value={value}
        onChange={(e) => {
          setValue(e.target.value);
        }}
      />
      <button onClick={navigateToContacts}>Watch now</button>
    </div>
  );
}

export default SearchBar;
