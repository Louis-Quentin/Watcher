import React from 'react';
import { Link } from 'react-router-dom';
import styles from './css/UserProfile.module.css';

const UserProfile: React.FC = () => {
  return (
    <div className={styles.container}>
      <Link to="/" className={styles.backButton}>Return</Link>
      <div className={styles.column}>
        <Link to="/orders" className={styles.link}>
          <div className={styles.item}>Orders</div>
          <hr className={styles.separator} />
          <div className={styles.arrow}></div>
        </Link>
      </div>
      <div className={styles.column}>
        <Link to="/messages" className={styles.link}>
          <div className={styles.item}>Messages</div>
          <hr className={styles.separator} />
          <div className={styles.arrow}></div>
        </Link>
      </div>
      <div className={styles.column}>
        <Link to="/waiting-lists" className={styles.link}>
          <div className={styles.item}>Waiting lists</div>
          <hr className={styles.separator} />
          <div className={styles.arrow}></div>
        </Link>
      </div>
      <div className={styles.column}>
        <Link to="/settings" className={styles.link}>
          <div className={styles.item}>Settings</div>
          <hr className={styles.separator} />
          <div className={styles.arrow}></div>
        </Link>
      </div>
      <div className={styles.column}>
        <Link to="/help" className={styles.link}>
          <div className={styles.item}>Help</div>
          <hr className={styles.separator} />
          <div className={styles.arrow}></div>
        </Link>
      </div>
      <div className={styles.column}>
        <Link to="/info" className={styles.link}>
          <div className={styles.item}>Info</div>
          <hr className={styles.separator} />
          <div className={styles.arrow}></div>
        </Link>
      </div>
    </div>
  );
}

export default UserProfile;
