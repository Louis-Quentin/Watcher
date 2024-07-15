import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import styles from './css/Navbar.module.css';

function Navbar() {
  const [isActive, setIsActive] = useState(false);
  const [showSearchBar, setShowSearchBar] = useState(false);

  const logo = require('./css/img/profile_picture_promptDreams.png');

  const toggleActiveClass = () => {
    setIsActive(!isActive);
  };

  const removeActive = () => {
    setIsActive(false);
  };

  const toggleSearchBar = () => {
    setShowSearchBar(!showSearchBar);
  };

  return (
    <>
      <nav className={`${styles.navbar} ${showSearchBar ? styles.expanded : ''}`}>
        <ul className={`${styles.navMenu} ${isActive ? styles.active : ''}`}>
          <div className={styles.NavLeft}>
            <img
              src={require('./css/img/burger_menu.png')}
              alt="Menu"
              onClick={toggleActiveClass}
              className={`${styles.BurgerMenu}`}
            />
            <li onClick={removeActive}>
              <a href='Map' className={`${styles.navLink} ${styles.Map}`}>Map</a>
            </li>
            <li onClick={removeActive}>
              <a href='/stores' className={`${styles.navLink}`}>Stores</a>
            </li>
            <li onClick={toggleSearchBar}>
              <a href='#home' className={`${styles.navLink}`}>Search</a>
            </li>
          </div>
          <img src={require('./css/img/watcher_logo.png')} className={`${styles.logo}`} alt="Logo" />
          <div className={styles.NavRight}>
            <li onClick={removeActive}>
              <a href='#home' className={`${styles.navLink}`}>Contact</a>
            </li>
            <li onClick={removeActive}>
              <Link to="/inscription" className={`${styles.navLink}`}>Signup</Link>
            </li>
            <li onClick={removeActive}>
              <Link to="/login" className={`${styles.navLink}`}>LogIn</Link>
            </li>
            <a href="/profile">
              <img src={require('./css/img/burger_menu.png')} alt="Menu" className={`${styles.profileLogo}`} />
            </a>
          </div>
        </ul>

        {showSearchBar && (
          <div className={styles.searchContainer}>
            <input type="text" className={styles.searchInput} placeholder="Search..." />
          </div>
        )}
      </nav>

      {showSearchBar && (
        <div className={`${styles.overlay} ${showSearchBar ? styles.visible : ''}`} onClick={toggleSearchBar}></div>
      )}
    </>
  );
}

export default Navbar;
