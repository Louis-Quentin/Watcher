import React from 'react';
import ReactDOM from 'react-dom';
import { useState } from 'react';
import  styles from './Navcta.module.css';
import {Link} from 'react-router-dom';

function NavbarCta() {
    const [isActive, setIsActive] = useState(false);
    const toggleActiveClass = () => {
        setIsActive(!isActive);
      };
    
      const removeActive = () => {
        setIsActive(false)
      }
      return (
        <div className="App">
            <header className="App-header">
                    <div className={`${styles.NavCta}`}>Join our NewsLetter !</div>
            </header>
        </div>)
}
export default NavbarCta;