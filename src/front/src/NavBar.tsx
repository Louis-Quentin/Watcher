import React from 'react';
import ReactDOM from 'react-dom';
import { useState } from 'react';
import  styles from './css/Navbar.module.css';
import {Link} from 'react-router-dom';

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
    <div className="App">
      <header className="App-header">

        <nav className={`${styles.navbar}`}>

          {/* logo */}
          <Link to="/">
            <img
                alt='logo' style={{ width: 100 }} src={String(logo)} />
          </Link>
          <ul className={`${styles.navMenu} ${isActive ? styles.active : ''}`}>
            <li onClick={removeActive}>
              <a href='#home' className={`${styles.navLink}`}>Home</a>
            </li>
            <li onClick={removeActive}>
              <a href='#home' className={`${styles.navLink}`}>Catalog</a>
            </li>
            <li onClick={removeActive}>
              <a href='#home' className={`${styles.navLink}`}>All products</a>
            </li>
            <li onClick={removeActive}>
              <a href='#home' className={`${styles.navLink}`}>Contact</a>
            </li>
            <li onClick={removeActive}>
              <Link to="/inscription" className={`${styles.navLink}`}>Inscription</Link>
            </li>
            <li onClick={removeActive}>
              <Link to="/login" className={`${styles.navLink}`}>Log in</Link>
            </li>
          </ul>

          <div className={`${styles.hamburger} ${isActive ? styles.active : ''}`}  onClick={toggleActiveClass}>
            <span className={`${styles.bar}`}></span>
            <span className={`${styles.bar}`}></span>
            <span className={`${styles.bar}`}></span>
          </div>
        </nav>

      </header>
    </div>
  );
}

export default Navbar;