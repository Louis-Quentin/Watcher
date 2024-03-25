import { useState } from 'react';
import  styles from './HomeLanding.module.css';
import {Link} from 'react-router-dom';
import 'react-slideshow-image/dist/styles.css'
import Slideshow from './SlideShow';

function HomeLanding() {
    const LandingImg = require('./HomeLanding.jpg');

      return (
        <div className="App">
            <header className="App-header">
                <Slideshow></Slideshow>
            </header>
        </div>)
}
export default HomeLanding;