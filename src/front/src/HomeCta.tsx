import React, { useEffect, useState, useRef } from 'react';
import styles from './HomeCta.module.css';
import HomeStyles from './css/HomeLanding.module.css';
import WatchList from './WatchList'; // Import the WatchList function
import { Watch } from './DataStructure.js'; // Import the Watch type from your types file
import Carousel from './Carousel';

interface CardProps {
    watch: Watch;
}

const Card: React.FC<CardProps> = ({ watch }) => {
    return (
        <div className={`${styles.Card}`}>
            <img src={watch.Url} alt="Watch" className={`${styles.CardImage} `} />
            <div className={styles.CardName}>
                <div className={` ${styles.CardTag} ${HomeStyles.TextDa2} `}>{watch.Name}</div>
            </div>
            <div className={styles.CardContent}>
                <div className={` ${styles.CardTag} ${HomeStyles.TextDa2} `}>{watch.Price} $</div>
                <div className={` ${styles.CardTag} ${HomeStyles.TextDa2} `}>{watch.Range} km</div>
                <div className={` ${styles.CardTag} ${HomeStyles.TextDa2} `}>{watch.Brand}</div>
            </div>
            <div className={styles.CardName}>
                <div className={` ${styles.CardTag} ${HomeStyles.TextDa2} `}>{watch.Available ? "Available" : "Order"}</div>
            </div>
        </div>
    );
}

const HomeCards: React.FC<{ watchlist: Watch[] }> = ({ watchlist }) => {
    const [isVisible, setIsVisible] = useState(false);
    const homeCardsRef = useRef<HTMLDivElement>(null);

    useEffect(() => {
        const observer = new IntersectionObserver(
            (entries) => {
                entries.forEach((entry) => {
                    if (entry.isIntersecting) {
                        setIsVisible(true);
                        observer.unobserve(entry.target);
                    }
                });
            },
            { threshold: 0.8 } // Trigger when 10% of the element is visible
        );

        if (homeCardsRef.current) {
            observer.observe(homeCardsRef.current);
        }

        return () => {
            if (homeCardsRef.current) {
                observer.unobserve(homeCardsRef.current);
            }
        };
    }, []);

    return (
        <div ref={homeCardsRef} className={`${styles.CardContainer} ${isVisible ? styles.Visible : ''}`}>
            {isVisible &&
                watchlist.map((watch, index) => (<Card key={index} watch={watch}/>))
            }
        </div>
    );
}

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
        <div className={styles.HomeCtaContainer}>
            <div className={styles.HomeCta}>
                <div className={`${styles.Title} ${styles.Visible}`}>Close To You</div>
            </div>
            <HomeCards watchlist={watchList.slice(0, 3)} />
            <HomeCards watchlist={watchList.slice(3, 6)} />
        </div>
    );
}

export default HomeCta;