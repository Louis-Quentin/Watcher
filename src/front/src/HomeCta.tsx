import React, { useEffect, useState } from 'react';
import styles from './HomeCta.module.css';
import WatchList from './WatchList'; // Import the WatchList function
import { Watch } from './DataStructure.js'; // Import the Watch type from your types file
import Carousel from './Carousel';

function HomeCta() {
    const [watchList, setWatchList] = useState<Watch[]>([]);

    useEffect(() => {
        // Fetch watch list data and update state
        const fetchWatchData = async () => {
            try {
                const data = await WatchList(); // Assuming WatchList is a function that fetches the watch list data
                setWatchList(data); // Assuming data is an array of Watch objects
            } catch (error) {
                console.error('Error fetching watch list:', error);
            }
        };

        fetchWatchData();
    }, []);

    return (
        <div>
            <div className={styles.HomeCta}>
                <div className={styles.Title}> Near From You </div>
                <button className={styles.SeeMore}>See more</button>
            </div>
            <Carousel watchlist={watchList}/>
            <div className={styles.HomeCta}>
                <div className={styles.Title}> Best for you </div>
                <button className={styles.SeeMore}>See more</button>
            </div>
        </div>
    );
}

export default HomeCta;
