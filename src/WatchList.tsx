import React from 'react';
import styles from './WatchList.module.css';
import test from './watchs_sample.json';

const WatchList = () => {
    return(
        <div className= {styles.comp}>
            <div className= {styles.card}>
                <div className= {styles.textbox}>{test.distance}</div>
                <div className= {styles.textbox}>{test.name}</div>
                <div className= {styles.textbox}>{test.img}</div>
                <div className= {styles.textbox}>{test.shipping}</div>
                <div className= {styles.textbox}>{test.shop}</div>
            </div>
        </div>
    );
}

export const BestForYou = () => {
    return(
        <div className= {styles.comp}>
            <div className= {styles.card}>
                <div className= {styles.textbox}>{test.distance}</div>
                <div className= {styles.textbox}>{test.name}</div>
                <div className= {styles.textbox}>{test.img}</div>
                <div className= {styles.textbox}>{test.shipping}</div>
                <div className= {styles.textbox}>{test.shop}</div>
            </div>
        </div>
    );
}

export default WatchList;