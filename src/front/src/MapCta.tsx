import React, { useEffect, useState, useRef } from 'react';
import styles from './css/MapCta.module.css';
import HomeStyles from './css/HomeLanding.module.css';
import WatchList from './WatchList'; // Import the WatchList function
import { Watch } from './DataStructure.js'; // Import the Watch type from your types file

function MapCta() {



    return (
        <div className={styles.MapCta}>
            <div className={styles.Background}>
                <div className={styles.Header}>
                <div className={styles.TitleHeader}>WATCHER</div>
                <div className={styles.Title}>Your Go-To Platform for Locating Watches</div>
                <div className={styles.Content}>Easily find nearby stores with our location-based feature. Effortlessly navigate to your closest destination and enjoy seamless shopping experiences on-the-go!</div>
                <div className={styles.Button}>
                <div className={styles.Line}></div>
                <div className={styles.Text}>Discover More</div>
            </div>
          </div>
        </div>
    </div>
    );
    
}

export default MapCta;