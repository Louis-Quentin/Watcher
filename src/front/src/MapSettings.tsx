import React from 'react';
import { useState } from 'react';
import styles from './css/MapSettings.module.css';  // Adjust the path as necessary
import buttonImage from './css/img/map_settings.png';  // Adjust the path to your image as necessary

const MapSettings = () => {
    const [isSettingsVisible, setIsSettingsVisible] = useState(false);
    const [rangeValue, setRangeValue] = useState(50); // Initialize range value

    const toggleSettingsVisibility = () => {
        setIsSettingsVisible(!isSettingsVisible);
      };
      return (
        <div className={styles.container}>
          {/* Other content */}
          <button
            className={styles.absoluteButton}
            style={{ backgroundImage: `url(${buttonImage})` }}
            onClick={toggleSettingsVisibility}
          ></button>
          
          <div className={`${styles.settingsScreen} ${isSettingsVisible ? styles.visible : ''}`}>
            <h3>Filter</h3>
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
          </div>
        </div>
      );
    };
export default MapSettings;
