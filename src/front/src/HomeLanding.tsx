import React from 'react';
import ReactDOM from 'react-dom';
import { useState } from 'react';
import  styles from './css/HomeLanding.module.css';
import {Link} from 'react-router-dom';
import SearchBar from './SearchBar';

function HomeLanding() {
    const LandingImg = require('./css/img/HomeLanding.jpg');
      return (    
        <div className = {styles.LandingPage}>
                <SearchBar/>
        </div>   
                
)}
export default HomeLanding;