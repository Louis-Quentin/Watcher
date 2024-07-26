import React from 'react';
import Navbar from './NavBar';
import Footer from './Footer';
import styles from './css/RetailerPage.module.css';
import { useLocation } from 'react-router-dom';
import StarRating from './StarRating';

interface Photo {
    photo_reference: string;
  }
  
  interface OpeningHours {
    open_now: boolean;
    weekday_text: string[];
  }
  

const RetailerPage: React.FC = () => {
  const location = useLocation();
  const state = location.state;
  const retailer = state?.retailer;

  const formatTime = (time: string) => {
    const hours = time.substring(0, 2);
    const minutes = time.substring(2);
    return `${hours}:${minutes}`;
  };

  if (!retailer) {
    return <div>No retailer data available</div>;
  }

  console.log(retailer.opening_hours);
  return (
    <div>
      <Navbar />
      <div className={styles.Retailer}>
        {retailer.photoUrl ? (
          <img className={styles.storeImg} src={retailer.photoUrl} alt={retailer.name} />
        ) : (
          <div className={styles.storeImg}>No image available</div>
        )}
        <div className={styles.storeData}>
            <div className={styles.Name}>{retailer.name}</div>
            <div className={styles.dataCol}>
                <div className={styles.childRow}>
                    <div className={styles.square}/>
                    <div className={`${styles.text}`}>{retailer.vicinity}</div>
                </div>
            </div>
            <div className={styles.dataCol}>
                <div className={styles.childRow}>
                    <div className={styles.square}/>
                    <div className={styles.line}>
                        <StarRating rating={retailer.rating}/>
                        <div className={styles.text}>by {retailer.user_ratings_total} users</div>
                    </div>
                </div>
          </div>
          <div>
            {retailer.reviews[0]}
          </div>
        </div>
      </div>
      <Footer />
    </div>
  );
};

export default RetailerPage;
