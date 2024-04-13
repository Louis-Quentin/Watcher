import React from 'react';
import ReactDOM from 'react-dom';
import { useState } from 'react';
import  styles from './css/HomeLanding.module.css';
import {Link} from 'react-router-dom';
import SearchBar from './SearchBar';

function HomeLanding() {
    const LandingImg = require('./css/img/HomeLanding.jpg');
    return (    
        <div className={styles.HomeVod}>
          <video autoPlay muted loop id="HomeVod">
            <source src="http://localhost:3000/videos/HomePage_vod.mp4" type="video/mp4"/>
          </video>
          <div className={styles.Header}>
            <div className={styles.TitleHeader}>WATCHER</div>
            <div className={styles.Title}>Your Go-To Platform for Ordering and Locating Watches</div>
            <div className={styles.Content}>Save time and effort by ordering watches online from the comfort of your home.  Explore a wide range of watch brands, styles, and models to find the perfect match for your preferences.</div>
            <div className={styles.Button}>
                <div className={styles.Line}></div>
                <div className={styles.Text}>Discover More</div>
            </div>
          </div>
        </div>   
      );}
export default HomeLanding;