// src/UserProfile.tsx
import React from 'react';
import { Link } from 'react-router-dom';
import styles from './css/UserProfile.module.css';
import Navbar from './NavBar';
import HomeLanding from './HomeLanding';
import AllBrands from './AllBrands';
import Favorites from './Favorites';
import Footer from './Footer'

const UserProfile: React.FC = () => {
  const handleProfileClick = () => {
    <HomeLanding />;
  };

  return (
    <div className={styles.container}>
      <Navbar />
      <div className={styles.column}>
        <div className={styles.link} onClick={handleProfileClick}>
          <div className={styles.item}>Profile</div>
        </div>
      </div>
      <div className={styles.column}>
        <Link to="/orders" className={styles.link}>
          <div className={styles.item}>Orders</div>
        </Link>
      </div>
      <div className={styles.column}>
        <Link to="/favorites" className={styles.link}>
          <div className={styles.item}>Favorites</div>
        </Link>
      </div>
      <div className={styles.column}>
        <Link to="/messages" className={styles.link}>
          <div className={styles.item}>Messages</div>
        </Link>
      </div>
      <div className={styles.column}>
        <Link to="/waiting-lists" className={styles.link}>
          <div className={styles.item}>Waiting lists</div>
        </Link>
      </div>
      <div className={styles.column}>
        <Link to="/settings" className={styles.link}>
          <div className={styles.item}>Settings</div>
        </Link>
      </div>
      <div className={styles.userContainer}>
      <div className={styles.userContent}>
        <div className={styles.section}>
          <Favorites />
        </div>
        <div className={styles.section}>
          <AllBrands />
          </div>
        <div className={styles.section}>
          <Footer/>
        </div>
      </div>
    </div>
    </div>
  );
};

export default UserProfile;