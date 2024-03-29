import React from 'react';
import styles from './HomeCta.module.css';
import WatchList from './WatchList';
import { BestForYou } from './WatchList';



function HomeCta() {
    return (
        <div>
            <div className = {styles.HomeCta}>
                <div className = {styles.Title}> Near From You </div>
                <button className = {styles.SeeMore}>See more</button>
            </div>
            <WatchList/>
            <div className = {styles.HomeCta}>
                <div className = {styles.Title}> Best for you </div>
                <button className = {styles.SeeMore}>See more</button>
            </div>
            <BestForYou/>
        </div>
    )
}

export default HomeCta;