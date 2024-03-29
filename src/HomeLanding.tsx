import React from 'react';
import ReactDOM from 'react-dom';
import { useState } from 'react';
import  styles from './HomeLanding.module.css';
import {Link} from 'react-router-dom';
import SlideShow from './SlideShow';
import SearchBar from './SearchBar';

function HomeLanding() {
    const LandingImg = require('./HomeLanding.jpg');

      return (
        <div className={styles.LandingPage}>
            <SearchBar/>
        </div>)
}
export default HomeLanding;