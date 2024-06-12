import React from 'react';
import ReactDOM from 'react-dom';
import { useState } from 'react';
import  styles from './css/Navbar.module.css';
import {Link} from 'react-router-dom';
import { url } from 'inspector';

function ToggleMenu () {

}

function Navbar() {

  const [isActive, setIsActive] = useState(false);
  const logo = require('./css/img/profile_picture_promptDreams.png');

  const toggleActiveClass = () => {
    setIsActive(!isActive);
  };

  const removeActive = () => {
    setIsActive(false)
  }
  return (
    
        <nav className={`${styles.navbar}`}>
          <ul className={`${styles.navMenu} ${isActive ? styles.active : ''}`}>
            <div className={styles.NavLeft}>
            <img  src={require("./css/img/burger_menu.png")} alt="Menu" onClick={ToggleMenu} className={`${styles.BurgerMenu}`} />
            <li onClick={removeActive}>
              <a href='Map' className={`${styles.navLink} ${styles.Map}`}>Map</a>
            </li>
            <li onClick={removeActive}>
              <a href='#home' className={`${styles.navLink}`}>Stores</a>
            </li>
            <li onClick={removeActive}>
              <a href='#home' className={`${styles.navLink}`}>Order</a>
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
                <img src={require("./css/img/burger_menu.png")} alt="Menu" onClick={ToggleMenu} className={`${styles.profileLogo}`} />
              </a>
            </div>
          </ul>

          <div className={`${styles.hamburger} ${isActive ? styles.active : ''}`}  onClick={toggleActiveClass}>
            <span className={`${styles.bar}`}></span>
            <span className={`${styles.bar}`}></span>
            <span className={`${styles.bar}`}></span>
          </div>
        </nav>
  );
}

export default Navbar;